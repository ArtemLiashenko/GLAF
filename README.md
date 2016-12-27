# README #

GLAF (Google location address formatter) - easily unify free-format locations with google.

## Required: ##

#### Golang packages: ####
* encoding/json
* io/ioutil
* net/http
* regexp
* strings
* fmt (for debug only)

#### Other: ####
* Google API key ([How to get](https://developers.google.com/maps/documentation/geocoding/get-api-key))

## How to set it up: ##
Rut with terminal: *$ go get bitbucket.org/Zilibuka/uniloc*

## Example: ##


```
#!Go

package main

import "fmt"
import "glaf"

func main() {
	myVar := glaf.Unify("CA san francisco USA", "[API key]")
	fmt.Println(myVar.GetCityLong())
	fmt.Println(myVar.GetCityShort())
	fmt.Println(myVar.GetStateLong())
	fmt.Println(myVar.GetStateShort())
	fmt.Println(myVar.GetCountryLong())
	fmt.Println(myVar.GetCountryShort())
	fmt.Println(myVar.GetFormated())
}
```


####Output:####
```
#!go
San Francisco
SF
California
CA
United States
US
San Francisco, CA, USA
```


## Who do I talk to? ##
* Repo owner or admin (Artem Liashenko)