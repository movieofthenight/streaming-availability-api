package api

import (
	"encoding/json"
	"github.com/cevatbarisyilmaz/ip2country"
	"log"
	"net"
	"net/http"
	"strings"
)

type UserCountryResponse struct {
	Country  string `json:"country,omitempty"`
	Detected bool   `json:"detected"`
}

func UserCountry(writer http.ResponseWriter, request *http.Request) {
	var country string
	host, _, err := net.SplitHostPort(request.RemoteAddr)
	if err == nil {
		country, err = ip2country.Country(net.ParseIP(host))
		if err == nil {
			country = strings.ToLower(country)
		}
	}
	data := &UserCountryResponse{Country: country, Detected: country != ""}
	response, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}
