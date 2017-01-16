# GLAF #
[![Go Report Card](https://goreportcard.com/badge/github.com/ArtemLiashenko/GLAF)](https://goreportcard.com/report/github.com/ArtemLiashenko/GLAF) 
[![GoDoc](https://godoc.org/github.com/ArtemLiashenko/GLAF?status.svg)](https://godoc.org/github.com/ArtemLiashenko/GLAF)
[![License](https://img.shields.io/aur/license/yaourt.svg)](https://github.com/ArtemLiashenko/GLAF/blob/master/LICENSE)

## Package for easy unify locations with different formats. ##
## Good for unification of addresses when you parse various sources. ##
## Good for using Google Maps Geolocation API in Golang. ##

### Required: ###

#### Golang packages: ####
* encoding/json
* io/ioutil
* errors
* net/http
* net/url
* strings

#### Other: ####
* Google API key ([How to get](https://developers.google.com/maps/documentation/geocoding/get-api-key))

### How to set it up: ###
* Run with terminal: *$ go get github.com/ArtemLiashenko/GLAF*
* Import "glaf" in your program

### Example: ###


```go

package main

import "fmt"
import "glaf"

func main() {
	myVar := glaf.Unify("Myru St, 96, Kharkiv", "[API key]")
	// or you can paste coordinates: 49.9371940, 36.4142605
	// or paste it in another order: 96 myru str Kharkiv
	// or paste it in another order: myru str Kharkiv 96
	// or paste it in another order: mira 96 Kharkiv
	fmt.Println(myVar.GetStreetNumLong())
	fmt.Println(myVar.GetStreetNumShort())
	fmt.Println(myVar.GetStreetLong())
	fmt.Println(myVar.GetStreetShort())
	fmt.Println(myVar.GetCityLong())
	fmt.Println(myVar.GetCityShort())
	fmt.Println(myVar.GetStateLong())
	fmt.Println(myVar.GetStateShort())
	fmt.Println(myVar.GetCountryLong())
	fmt.Println(myVar.GetCountryShort())
	fmt.Println(myVar.GetFormated())
	fmt.Println(myVar.GetCoordinates())
	
/*
Output:

96 <nil>
96 <nil>
Myru Street <nil>
Myru St <nil>
Kharkiv <nil>
Kharkiv <nil>
Kharkiv Oblast <nil>
Kharkiv Oblast <nil>
Ukraine <nil>
UA <nil>
Myru St, 96, Kharkiv, Kharkiv Oblast, Ukraine <nil>
49.9371940000, 36.4142605000 <nil>
*/	

}
```


### Author ###
Artem Liashenko <artem.liashenko@gmail.com>

### License ###
See [LICENSE](https://github.com/ArtemLiashenko/GLAF/blob/master/LICENSE)
