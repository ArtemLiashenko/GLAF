/*
Project Name: Google location address formatter
Author: Artem Liashenko
Description: Address formatter for locations.
Version: 1.0.1
*/

package glaf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type AddressComponent struct {
	Long_name  string
	Short_name string
	Types      []string
}

type Target struct {
	Lat float64
	Lng float64
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
func (gData *GeoData) GetFormated() (string, error) {
	if gData.Status != "OK" {
		return "", errors.New("location not found")
	}

	return gData.Results[0].Formatted_address, nil
}

//get Сoordinates from geocoding api response
func (gData *GeoData) GetСoordinates() (string, error) {

	if gData.Status != "OK" {
		return "", errors.New("location not found")
	}

	return strconv.FormatFloat(gData.Results[0].Geometry.Location.Lat, 'f', 10, 64) + ", " + strconv.FormatFloat(gData.Results[0].Geometry.Location.Lng, 'f', 10, 64), nil
}

//get Street number from geocoding api response (long version)
func (gData *GeoData) GetStreetNumLong() (string, error) {

	if gData.Status != "OK" {
		return "", errors.New("location not found")
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "street_number" {
			return gData.Results[0].Address_components[i].Long_name, nil
		}
	}

	return "", errors.New("Street Number (long) not found")
}

//get Street number from geocoding api response (short version)
func (gData *GeoData) GetStreetNumShort() (string, error) {

	if gData.Status != "OK" {
		return "", errors.New("location not found")
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "street_number" {
			return gData.Results[0].Address_components[i].Short_name, nil
		}
	}

	return "", errors.New("Street Number (short) not found")
}

//get Street from geocoding api response (long version)
func (gData *GeoData) GetStreetLong() (string, error) {

	if gData.Status != "OK" {
		return "", errors.New("location not found")
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "route" {
			return gData.Results[0].Address_components[i].Long_name, nil
		}
	}

	return "", errors.New("Street (long) not found")
}

//get Street from geocoding api response (short version)
func (gData *GeoData) GetStreetShort() (string, error) {

	if gData.Status != "OK" {
		return "", errors.New("location not found")
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "route" {
			return gData.Results[0].Address_components[i].Short_name, nil
		}
	}

	return "", errors.New("Street (short) not found")
}

//get City from geocoding api response (long version)
func (gData *GeoData) GetCityLong() (string, error) {

	if gData.Status != "OK" {
		return "", errors.New("location not found")
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "locality" {
			return gData.Results[0].Address_components[i].Long_name, nil
		}
	}

	return "", errors.New("city (long) not found")
}

//get City from geocoding api response (short version)
func (gData *GeoData) GetCityShort() (string, error) {

	if gData.Status != "OK" {
		return "", errors.New("location not found")
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "locality" {
			return gData.Results[0].Address_components[i].Short_name, nil
		}
	}

	return "", errors.New("city (short) not found")
}

//get State from geocoding api response (long version)
func (gData *GeoData) GetStateLong() (string, error) {

	if gData.Status != "OK" {
		return "", errors.New("location not found")
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "administrative_area_level_1" {
			return gData.Results[0].Address_components[i].Long_name, nil
		}
	}

	return "", errors.New("state (long) not found")
}

//get State from geocoding api response (short version)
func (gData *GeoData) GetStateShort() (string, error) {

	if gData.Status != "OK" {
		return "", errors.New("location not found")
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "administrative_area_level_1" {
			return gData.Results[0].Address_components[i].Short_name, nil
		}
	}

	return "", errors.New("state (long) not found")
}

//get Country from geocoding api response (long version)
func (gData *GeoData) GetCountryLong() (string, error) {

	if gData.Status != "OK" {
		return "", errors.New("location not found")
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "country" {
			return gData.Results[0].Address_components[i].Long_name, nil
		}
	}

	return "", errors.New("country (long) not found")
}

//get Country from geocoding api response (short version)
func (gData *GeoData) GetCountryShort() (string, error) {

	if gData.Status != "OK" {
		return "", errors.New("location not found")
	}

	for i := 0; i < len(gData.Results[0].Address_components); i++ {
		if gData.Results[0].Address_components[i].Types[0] == "country" {
			return gData.Results[0].Address_components[i].Short_name, nil
		}
	}

	return "", errors.New("country (short) not found")
}

//prepare and send request to geocoding api then get response and make srtuct from json
func Unify(locStr string, apiKey string) GeoData {

	spaces, _ := regexp.Compile(" ")
	prLoc := spaces.ReplaceAllString(strings.TrimSpace(locStr), "+")
	link := "https://maps.googleapis.com/maps/api/geocode/json?address=" + prLoc + "&key=" + apiKey
	var geoResult GeoData
	resp, httpGetErr := http.Get(link)

	if httpGetErr == nil {
		bytes, readAllErr := ioutil.ReadAll(resp.Body)
		if readAllErr == nil {
			err := json.Unmarshal([]byte(bytes), &geoResult)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return geoResult
}
