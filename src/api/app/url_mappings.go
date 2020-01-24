package app

import "github.com/tv2169145/golang-testing/src/api/controllers"

func mapUrls() {
	router.GET("locations/countries/:country_id", controllers.GetCountry)
}
