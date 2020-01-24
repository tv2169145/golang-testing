package test

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"github.com/tv2169145/golang-testing/src/api/domain/locations"
	"github.com/tv2169145/golang-testing/src/api/utils/errors"
	"net/http"
	"testing"
)

func TestGetCountryNotFound(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		URL: "https://api.mercadolibre.com/countries/AR",
		HTTPMethod: http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"message": "Country not found", "error": "not_found", "status": 404, "cause": []}`,
	})
	fmt.Println("about to functional test get countries")
	response, err := http.Get("http://localhost:8080/locations/countries/AR")
	assert.Nil(t, err)
	assert.NotNil(t, response)

	//bytes, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(bytes))

	var apiErr errors.ApiError
	err = json.NewDecoder(response.Body).Decode(&apiErr)
	//err = json.Unmarshal(bytes, &apiErr)
	fmt.Println(apiErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
}

func TestGetCountryNoError(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		URL: "https://api.mercadolibre.com/countries/AR",
		HTTPMethod: http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id":"AR","name":"Argentina","locale":"es_AR","currency_id":"ARS","decimal_separator":",","thousands_separator":".","time_zone":"GMT-03:00","geo_information":{"location":{"latitude":-38.416096,"longitude":-63.616673}},"states":[{"id":"AR-B","name":"Buenos Aires"},{"id":"AR-C","name":"Capital Federal"},{"id":"AR-K","name":"Catamarca"},{"id":"AR-H","name":"Chaco"},{"id":"AR-U","name":"Chubut"},{"id":"AR-W","name":"Corrientes"},{"id":"AR-X","name":"Córdoba"},{"id":"AR-E","name":"Entre Ríos"},{"id":"AR-P","name":"Formosa"},{"id":"AR-Y","name":"Jujuy"},{"id":"AR-L","name":"La Pampa"},{"id":"AR-F","name":"La Rioja"},{"id":"AR-M","name":"Mendoza"},{"id":"AR-N","name":"Misiones"},{"id":"AR-Q","name":"Neuquén"},{"id":"AR-R","name":"Río Negro"},{"id":"AR-A","name":"Salta"},{"id":"AR-J","name":"San Juan"},{"id":"AR-D","name":"San Luis"},{"id":"AR-Z","name":"Santa Cruz"},{"id":"AR-S","name":"Santa Fe"},{"id":"AR-G","name":"Santiago del Estero"},{"id":"AR-V","name":"Tierra del Fuego"},{"id":"AR-T","name":"Tucumán"}]}`,
	})
	response, err := http.Get("http://localhost:8080/locations/countries/AR")
	assert.Nil(t, err)
	var country locations.Country
	err = json.NewDecoder(response.Body).Decode(&country)
	assert.Nil(t, err)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}
