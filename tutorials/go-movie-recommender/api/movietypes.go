package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type MovieType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

const (
	TrendingNow       = "trendingNow"
	BestOfRecentYears = "bestOfRecentYears"
	AllTimeClassics   = "allTimeClassics"
	OldiesButGoldies  = "oldiesButGoldies"
)

var movieTypes = []MovieType{
	{Id: TrendingNow, Name: "Trending Now"},
	{Id: BestOfRecentYears, Name: "Best of Recent Years"},
	{Id: AllTimeClassics, Name: "All-Time Classics"},
	{Id: OldiesButGoldies, Name: "Oldies But Goldies"},
}

func MovieTypes(writer http.ResponseWriter, request *http.Request) {
	response, err := json.Marshal(movieTypes)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}
