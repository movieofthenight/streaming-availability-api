package api

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/movieofthenight/go-streaming-availability/v2"
	"os"
)

var streamingAvailabilityClient = createStreamingAvailabilityClient()
var tmdbClient = createTmdbClient()

func createStreamingAvailabilityClient() *streaming.DefaultAPIService {
	apiKey := os.Getenv("STREAMING_AVAILABILITY_API_KEY")
	configuration := streaming.NewConfiguration()
	configuration.AddDefaultHeader("X-RapidAPI-Key", apiKey)
	client := streaming.NewAPIClient(configuration).DefaultAPI
	return client
}

func createTmdbClient() *tmdb.Client {
	apiKey := os.Getenv("TMDB_API_KEY")
	client, err := tmdb.Init(apiKey)
	if err != nil {
		panic(err)
	}
	return client
}
