package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/movieofthenight/go-streaming-availability/v2"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Movie struct {
	Title         string `json:"title"`
	Poster        string `json:"poster"`
	StreamingLink string `json:"streamingLink"`
	StreamingLogo string `json:"streamingLogo"`
}

// Recommend recommends movies based on the query parameters.
func Recommend(writer http.ResponseWriter, request *http.Request) {
	searchRequest, country, services := consumeQueryParameters(
		request.URL.Query())
	searchResponse, _, err := searchRequest.Execute()
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	movies := readSearchResponse(searchResponse, country, services)
	response, err := json.Marshal(movies)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

// consumeQueryParameters consumes the query parameters
// and returns the search request, the country code and the services.
func consumeQueryParameters(query url.Values) (
	searchRequest streaming.ApiSearchByFiltersRequest,
	country string, services map[string]bool) {

	country = query.Get("country")

	var servicesParameter []string
	services = map[string]bool{}
	for i, service := range strings.Split(query.Get("services"), ",") {
		// Limit the number of services to 16
		// as it is the maximum allowed by the Streaming Availability API.
		if i >= 16 {
			break
		}
		services[service] = true
		// We want to follow the format of
		// "services=netflix.subscription,hulu.subscription"
		// so we append ".subscription" to each service.
		// Otherwise, for services like Prime Video and Apple TV
		// the results will also include movies
		// that are only available for rent&purchase;
		// or only available with extra channel subscriptions.
		servicesParameter = append(servicesParameter,
			fmt.Sprintf("%s.subscription", service))
	}

	// Create a new search request.
	searchRequest = streamingAvailabilityClient.SearchByFilters(
		context.Background())
	// We only want to show movies and not series.
	searchRequest = searchRequest.ShowType("movie")
	searchRequest = searchRequest.Country(country)
	// Pass the comma separated list of services.
	searchRequest = searchRequest.Services(
		strings.Join(servicesParameter, ","))
	if query.Has("genre") {
		searchRequest = searchRequest.Genres(query.Get("genre"))
	}
	if query.Has("keyword") {
		searchRequest = searchRequest.Keyword(query.Get("keyword"))
	}
	switch query.Get("movieType") {
	case TrendingNow:
		// We want to show the most popular movies of the last week.
		searchRequest = searchRequest.
			OrderBy("popularity_1week").Desc(true)
	case AllTimeClassics:
		// We want to show the most popular movies of all time.
		searchRequest = searchRequest.
			OrderBy("popularity_alltime").Desc(true)
	case OldiesButGoldies:
		// We want to show the most popular movies
		// that released at least 25 years ago.
		searchRequest = searchRequest.YearMax(int32(time.Now().Year() - 25)).
			OrderBy("popularity_alltime").Desc(true)
	case BestOfRecentYears:
		// We want to show the most popular movies for the past year.
		searchRequest = searchRequest.
			OrderBy("popularity_1year").Desc(true)
	}
	return
}

// readSearchResponse reads the search response and returns the movies found.
func readSearchResponse(searchResponse *streaming.SearchFiltersResponseSchema,
	country string, services map[string]bool) (movies []Movie) {

	// We want to get the poster URLs in parallel from TMDB API.
	// We use a WaitGroup to wait for all the goroutines to finish.
	wg := sync.WaitGroup{}
	posters := map[int]string{}
	postersMu := &sync.Mutex{}
	for _, movie := range searchResponse.Result {
		movie := movie
		wg.Add(1)
		go func() {
			defer wg.Done()
			poster, err := getPosterUrl(int(movie.TmdbId))
			if err != nil {
				log.Println(err)
				return
			}
			postersMu.Lock()
			posters[int(movie.TmdbId)] = poster
			postersMu.Unlock()
		}()
	}
	wg.Wait()

	// Once we have the poster URLs, we can create the response.
	for _, movie := range searchResponse.Result {
		// We only want to show movies that have a poster.
		if posters[int(movie.TmdbId)] == "" {
			continue
		}
		for _, streamingOption := range movie.StreamingInfo[country] {
			// Since streaming options include all the possible ways
			// to watch the movie in the selected country,
			// we want to make sure we use the streaming option that is
			// available via a subscription
			// on a service that the user requested.
			if services[streamingOption.Service] &&
				streamingOption.StreamingType == streaming.SUBSCRIPTION {
				movies = append(movies, Movie{
					Title:         movie.Title,
					Poster:        posters[int(movie.TmdbId)],
					StreamingLink: streamingOption.Link,
					StreamingLogo: countries[country].
						Services[streamingOption.Service].DarkThemeLogo,
				})
				break
			}
		}
	}
	return
}

// getPosterUrl fetches the poster url from TDMB API by TMDB Id.
// Returns an empty string if the movie does not have a poster.
func getPosterUrl(tmdbId int) (string, error) {
	tmdbInfo, err := tmdbClient.GetMovieDetails(tmdbId, nil)
	if err != nil {
		return "", err
	}
	if len(tmdbInfo.PosterPath) > 0 {
		return fmt.Sprintf("https://image.tmdb.org/t/p/w500/%s",
			tmdbInfo.PosterPath), nil
	}
	return "", nil
}
