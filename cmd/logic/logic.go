package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Artists struct {
	Id            int
	Image         string
	Name          string
	Members       []string
	CreationDate  int
	FirstAlbum    string
	Locations     string
	ConcertDates  string
	Relations     string
	RelationsData Relations
	LocationsData Locations
}

type Relations struct {
	Id             int
	DatesLocations map[string][]string
}

type Locations struct {
	Id        int
	Locations []string
}

const (
	getArtists = "https://groupietrackers.herokuapp.com/api/artists"
)

func AllArtists(w http.ResponseWriter, r *http.Request) ([]Artists, bool) {
	var result []Artists
	resp, err := http.Get(getArtists)
	if err != nil {
		log.Fatal(err)
		return result, true
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return result, true
	}

	jsonErr := json.Unmarshal(body, &result)

	if jsonErr != nil {
		log.Fatal(err)
		return result, true
	}

	return result, false
}

func DetailArtist(w http.ResponseWriter, r *http.Request) (Artists, int) {
	var result Artists

	id := r.FormValue("id")
	if id == "" {
		fmt.Println("id is empty")
		return result, 404
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return result, 500
	}

	artistById, err := http.Get(getArtists + "/" + id)
	if err != nil {
		fmt.Println(err)
		return result, 404
	}

	artistByIdBody, err := ioutil.ReadAll(artistById.Body)
	if err != nil {
		fmt.Println(err)
		return result, 500
	}

	jsonErrArtist := json.Unmarshal(artistByIdBody, &result)
	if jsonErrArtist != nil {
		fmt.Println(err)
		return result, 500
	}

	relationById, err := http.Get(result.Relations)
	if err != nil {
		fmt.Println(err)
		return result, 404
	}

	relationByIdBody, err := ioutil.ReadAll(relationById.Body)
	if err != nil {
		fmt.Println(err)
		return result, 500
	}

	jsonErrRelation := json.Unmarshal(relationByIdBody, &result.RelationsData)
	if jsonErrRelation != nil {
		fmt.Println(err)
		return result, 500
	}

	locationsById, err := http.Get(result.Locations)
	if err != nil {
		fmt.Println(err)
		return result, 404
	}

	locationsByIdBody, err := ioutil.ReadAll(locationsById.Body)
	if err != nil {
		fmt.Println(err)
		return result, 500
	}

	jsonErrLocations := json.Unmarshal(locationsByIdBody, &result.LocationsData)
	if jsonErrLocations != nil {
		fmt.Println(err)
		return result, 500
	}

	return result, 0
}
