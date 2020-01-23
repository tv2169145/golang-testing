package main

import (
	"fmt"
	"github.com/tv2169145/golang-testing/src/api/providers/locations_provider"
)

func main() {
	country, err := locations_provider.GetCountry("ARS")
	fmt.Println(country)
	fmt.Println(err)
}
