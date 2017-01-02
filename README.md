# README #

GLAF (Google location address formatter) - easily unify free-format locations with google.

## Required: ##

#### Golang packages: ####
* encoding/json
* io/ioutil
* errors
* net/http
* regexp
* strings
* fmt (for debug only)

#### Other: ####
* Google API key ([How to get](https://developers.google.com/maps/documentation/geocoding/get-api-key))

## How to set it up: ##
* Run with terminal: *$ go get github.com/ArtemLiashenko/GLAF*
* Import "glaf" in your program

## Example: ##


```go

package main

import "fmt"
import "glaf"

func main() {
	myVar := glaf.Unify("Myru St, 96, Kharkiv", "[API key]")
	// or you can paste coordinates: 49.9371940, 36.4142605
	// or paste it in another order: 96 myru str Kharkiv
	// or paste it in another order: myru str Kharkiv 96
	// or paste it in another order: myru 96 Kharkiv
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
	fmt.Println(myVar.GetСoordinates())
	
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


## Author ##
Artem Liashenko <artem.liashenko@gmail.com>

## License ##
See [LICENSE](https://github.com/ArtemLiashenko/GLAF/blob/master/LICENSE)
