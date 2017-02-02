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
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// AddressComponent - geocodding API returned JSON object with administrative area info
type AddressComponent struct {
	Long_name  string
	Short_name string
	Types      []string
}

// Target - geocodding API returned JSON object with coordinates Lat & Lng
type Target struct {
	Lat float64
	Lng float64
}

// DoubleTarget - geocodding API returned JSON object with coordinates
type DoubleTarget struct {
	Northeast Target
	Southwest Target
}

// Geometry - geocodding API returned JSON object with coordinates (main)
type Geometry struct {
	Bounds        DoubleTarget
	Location      Target
	Location_type string
	Viewport      DoubleTarget
}

// Result - geocodding API returned JSON object with geo data
type Result struct {
	Address_components []AddressComponent
	Formatted_address  string
	Geometry           Geometry
	Place_id           string
	Types              []string
}

// GeoData - geocodding API returned JSON main object
type GeoData struct {
	Results []Result
	Status  string
}

// Error reports an error and the operation that caused it.
type Error struct {
	Op  string
	Err error
}

func (e *Error) Error() string { return e.Op + ": " + e.Err.Error() }

//validateResponce - validate API responce
func validateResponce(gData *GeoData) error {

	if len(gData.Results) > 1 {
		return errors.New("TOO_MANY_RESULTS")
	}

	if gData.Status != "OK" {
		return errors.New(gData.Status)
	}

	return nil
}

//GetFormated - get Formated address
func (gData *GeoData) GetFormated() (string, error) {

	if validateResponce(gData) != nil {
		return "", &Error{"validateResponce", validateResponce(gData)}
	}

	return gData.Results[0].Formatted_address, nil
}

//GetСoordinates - get Сoordinates from geocoding api response
func (gData *GeoData) GetСoordinates() (string, error) {

	if validateResponce(gData) != nil {
		return "", &Error{"validateResponce", validateResponce(gData)}
	}

	return strconv.FormatFloat(gData.Results[0].Geometry.Location.Lat, 'f', 10, 64) + ", " + strconv.FormatFloat(gData.Results[0].Geometry.Location.Lng, 'f', 10, 64), nil
}

//GetPostCode - get Postal code from geocoding api response
func (gData *GeoData) GetPostCode() (string, error) {

	if validateResponce(gData) != nil {
		return "", &Error{"validateResponce", validateResponce(gData)}
	}

	for _, addrComp := range gData.Results[0].Address_components {
		if addrComp.Types[0] == "postal_code" {
			return addrComp.Long_name, nil
		}
	}

	return "", &Error{"GetPostCode", errors.New("NO_RESULT")}
}

//GetStreetNumLong - get Street number from geocoding api response (long version)
func (gData *GeoData) GetStreetNumLong() (string, error) {

	if validateResponce(gData) != nil {
		return "", &Error{"validateResponce", validateResponce(gData)}
	}

	for _, addrComp := range gData.Results[0].Address_components {
		if addrComp.Types[0] == "street_number" {
			return addrComp.Long_name, nil
		}
	}

	return "", &Error{"GetStreetNumLong", errors.New("NO_RESULT")}
}

//GetStreetNumShort - get Street number from geocoding api response (short version)
func (gData *GeoData) GetStreetNumShort() (string, error) {

	if validateResponce(gData) != nil {
		return "", &Error{"validateResponce", validateResponce(gData)}
	}

	for _, addrComp := range gData.Results[0].Address_components {
		if addrComp.Types[0] == "street_number" {
			return addrComp.Short_name, nil
		}
	}

	return "", &Error{"GetStreetNumShort", errors.New("NO_RESULT")}
}

//GetStreetLong - get Street from geocoding api response (long version)
func (gData *GeoData) GetStreetLong() (string, error) {

	if validateResponce(gData) != nil {
		return "", &Error{"validateResponce", validateResponce(gData)}
	}

	for _, addrComp := range gData.Results[0].Address_components {
		if addrComp.Types[0] == "route" {
			return addrComp.Long_name, nil
		}
	}

	return "", &Error{"GetStreetLong", errors.New("NO_RESULT")}
}

//GetStreetShort - get Street from geocoding api response (short version)
func (gData *GeoData) GetStreetShort() (string, error) {

	if validateResponce(gData) != nil {
		return "", &Error{"validateResponce", validateResponce(gData)}
	}

	for _, addrComp := range gData.Results[0].Address_components {
		if addrComp.Types[0] == "route" {
			return addrComp.Short_name, nil
		}
	}

	return "", &Error{"GetStreetShort", errors.New("NO_RESULT")}
}

//GetCityLong - get City from geocoding api response (long version)
func (gData *GeoData) GetCityLong() (string, error) {

	if validateResponce(gData) != nil {
		return "", &Error{"validateResponce", validateResponce(gData)}
	}

	for _, addrComp := range gData.Results[0].Address_components {
		if addrComp.Types[0] == "locality" {
			return addrComp.Long_name, nil
		}
	}

	return "", &Error{"GetCityLong", errors.New("NO_RESULT")}
}

//GetCityShort - get City from geocoding api response (short version)
func (gData *GeoData) GetCityShort() (string, error) {

	if validateResponce(gData) != nil {
		return "", &Error{"validateResponce", validateResponce(gData)}
	}

	for _, addrComp := range gData.Results[0].Address_components {
		if addrComp.Types[0] == "locality" {
			return addrComp.Short_name, nil
		}
	}

	return "", &Error{"GetCityShort", errors.New("NO_RESULT")}
}

//GetStateLong - get State from geocoding api response (long version)
func (gData *GeoData) GetStateLong() (string, error) {

	if validateResponce(gData) != nil {
		return "", &Error{"validateResponce", validateResponce(gData)}
	}

	for _, addrComp := range gData.Results[0].Address_components {
		if addrComp.Types[0] == "administrative_area_level_1" {
			return addrComp.Long_name, nil
		}
	}

	return "", &Error{"GetStateLong", errors.New("NO_RESULT")}
}

//GetStateShort - get State from geocoding api response (short version)
func (gData *GeoData) GetStateShort() (string, error) {

	if validateResponce(gData) != nil {
		return "", &Error{"validateResponce", validateResponce(gData)}
	}

	for _, addrComp := range gData.Results[0].Address_components {
		if addrComp.Types[0] == "administrative_area_level_1" {
			return addrComp.Short_name, nil
		}
	}

	return "", &Error{"GetStateShort", errors.New("NO_RESULT")}
}

//GetCountryLong - get Country from geocoding api response (long version)
func (gData *GeoData) GetCountryLong() (string, error) {

	if validateResponce(gData) != nil {
		return "", &Error{"validateResponce", validateResponce(gData)}
	}

	for _, addrComp := range gData.Results[0].Address_components {
		if addrComp.Types[0] == "country" {
			return addrComp.Long_name, nil
		}
	}

	return "", &Error{"GetCountryLong", errors.New("NO_RESULT")}
}

//GetCountryShort - get Country from geocoding api response (short version)
func (gData *GeoData) GetCountryShort() (string, error) {

	if validateResponce(gData) != nil {
		return "", &Error{"validateResponce", validateResponce(gData)}
	}

	for _, addrComp := range gData.Results[0].Address_components {
		if addrComp.Types[0] == "country" {
			return addrComp.Short_name, nil
		}
	}

	return "", &Error{"GetCountryShort", errors.New("NO_RESULT")}
}

//Unify - prepare and send request to geocoding api then get response and make srtuct from json
func Unify(locStr string, apiKey string) GeoData {

	var geoResult GeoData
	link, urlErr := url.Parse("https://maps.googleapis.com/maps/api/geocode/json")

	if urlErr == nil {
		linkQ := link.Query()
		linkQ.Set("address", strings.TrimSpace(locStr))
		linkQ.Set("key", apiKey)
		link.RawQuery = linkQ.Encode()

		resp, httpGetErr := http.Get(link.String())
		defer resp.Body.Close()

		if httpGetErr != nil {
			return geoResult
		}

		bytes, readAllErr := ioutil.ReadAll(resp.Body)

		if readAllErr != nil {
			return geoResult
		}

		unmErr := json.Unmarshal([]byte(bytes), &geoResult)
		if unmErr != nil {
			return geoResult
		}

	}

	return geoResult
}
