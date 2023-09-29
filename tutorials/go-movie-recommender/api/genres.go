package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sort"
)

type Genre struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var genres []Genre

func init() {
	response, _, err := streamingAvailabilityClient.
		Genres(context.Background()).Execute()
	if err != nil {
		panic(err)
	}
	for genreId, genreName := range response.Result {
		genres = append(genres, Genre{Id: genreId, Name: genreName})
	}
	sort.Slice(genres, func(i, j int) bool {
		return genres[i].Name < genres[j].Name
	})
}

func Genres(writer http.ResponseWriter, request *http.Request) {
	response, err := json.Marshal(genres)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}
