package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string     `json:"routeId"`
	ClientID  string     `json:"clientId"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func (route *Route) LoadPositions() error {

	// Validations: If the ID is not filled, i cannot know what file i am suposed to read!
	if route.ID == "" {
		return errors.New("route ID cannot be null or blank!")
	}

	// Trying to read the file. Returns the Error if it is not null, and ensures that the file is closed after everything in the function is runned
	file, fileError := os.Open("destinations/" + route.ID + ".txt")
	if fileError != nil {
		return fileError
	}
	defer file.Close()

	// Initalize the scanner of the file. In the for loop, it repeats with every line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		// Getting the Latitude and the Longitude. The Latitude is always before the comma, and the Longitude is always after it
		data := strings.Split(scanner.Text(), ",")

		// Reading The Latitude
		latitude, latitudeError := strconv.ParseFloat(data[0], 64)
		if latitudeError != nil {
			return nil
		}

		// Reading The Longitude
		longitude, longitudeError := strconv.ParseFloat(data[1], 64)
		if longitudeError != nil {
			return nil
		}

		// Making a new Position from the Latitude and Longitude
		route.Positions = append(route.Positions, Position{
			Latitude:  latitude,
			Longitude: longitude,
		})

	}

	// Returning null, or void
	return nil

}

func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)

	for key, value := range r.Positions {

		// Populates the PartialRoutePosition Type with the received Route Type
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{value.Latitude, value.Longitude}
		route.Finished = false

		// if total-1 is equal to the key, we are at the end of the route, therefore Finished is true
		if total-1 == key {
			route.Finished = true
		}

		// Parsing the route to a JSON
		jsonRoute, jsonError := json.Marshal(route)
		if jsonError != nil {
			return nil, jsonError
		}

		result = append(result, string(jsonRoute))
	}
	return result, nil
}
