/*
Project Name: Google location address formatter
Author: Artem Liashenko
Description: Address formatter for locations.
Version: 1.0.1
*/

package glaf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type AddressComponent struct {
	Long_name  string
	Short_name string
	Types      []string
}

type Target struct {
	Lat float32
	Lng float32
}

type DoubleTarget struct {
	Northeast Target
	Southwest Target
}

type Geometry struct {
	Bounds        DoubleTarget
	Location      Target
	Location_type string
	Viewport      DoubleTarget
}

type Result struct {
	Address_components []AddressComponent
	Formatted_address  string
	Geometry           Geometry
	Place_id           string
	Types              []string
}

type GeoData struct {
	Results []Result
	Status  string
}

//get Formated address
func (gData GeoData) GetFormated() string {
	return gData.Results[0].Formatted_address
}

//get City from geocoding api response (long version)
func (gData GeoData) GetCityLong() string {

	if gData.Status != "OK" {
		return "location not found"
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "locality" {
			return gData.Results[0].Address_components[i].Long_name
		}
	}

	return "city (long) not found"
}

//get City from geocoding api response (short version)
func (gData GeoData) GetCityShort() string {

	if gData.Status != "OK" {
		return "location not found"
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "locality" {
			return gData.Results[0].Address_components[i].Short_name
		}
	}

	return "city (short) not found"
}

//get State from geocoding api response (long version)
func (gData GeoData) GetStateLong() string {

	if gData.Status != "OK" {
		return "location not found"
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "administrative_area_level_1" {
			return gData.Results[0].Address_components[i].Long_name
		}
	}

	return "state (long) not found"
}

//get State from geocoding api response (short version)
func (gData GeoData) GetStateShort() string {

	if gData.Status != "OK" {
		return "location not found"
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "administrative_area_level_1" {
			return gData.Results[0].Address_components[i].Short_name
		}
	}

	return "state (long) not found"
}

//get Country from geocoding api response (long version)
func (gData GeoData) GetCountryLong() string {

	if gData.Status != "OK" {
		return "location not found"
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "country" {
			return gData.Results[0].Address_components[i].Long_name
		}
	}

	return "country (long) not found"
}

//get Country from geocoding api response (short version)
func (gData GeoData) GetCountryShort() string {

	if gData.Status != "OK" {
		return "location not found"
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "country" {
			return gData.Results[0].Address_components[i].Short_name
		}
	}

	return "country (short) not found"
}

//prepare and send request to geocoding api then get response and make srtuct from json
func Unify(locStr string, apiKey string) GeoData {

	spaces, _ := regexp.Compile(" ")
	prLoc := spaces.ReplaceAllString(strings.TrimSpace(locStr), "+")
	link := "https://maps.googleapis.com/maps/api/geocode/json?address=" + prLoc + "&key=" + apiKey
	resp, _ := http.Get(link)
	bytes, _ := ioutil.ReadAll(resp.Body)

	var geoResult GeoData

	err := json.Unmarshal([]byte(bytes), &geoResult)
	if err != nil {
		fmt.Println(err)
	}

	return geoResult
}
