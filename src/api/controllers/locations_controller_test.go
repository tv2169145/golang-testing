package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tv2169145/golang-testing/src/api/domain/locations"
	"github.com/tv2169145/golang-testing/src/api/services"
	"github.com/tv2169145/golang-testing/src/api/utils/errors"
	"github.com/tv2169145/golang-testing/src/api/utils/test_utils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	getCountryFunc func(countryId string) (*locations.Country, *errors.ApiError)
)

type locationsServiceMock struct {}

func (s *locationsServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return getCountryFunc(countryId)
}

func TestMain(m *testing.M) {
	//rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetCountryNotFound(t *testing.T) {
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return nil, &errors.ApiError{
			Status: http.StatusNotFound,
			Message: "Country not found",
		}
	}
	services.LocationService = new(locationsServiceMock)
	//rest.FlushMockups()
	//_ = rest.AddMockups(&rest.Mock{
	//	URL: "https://api.mercadolibre.com/countries/AR",
	//	HTTPMethod: http.MethodGet,
	//	RespHTTPCode: http.StatusNotFound,
	//	RespBody: `{"message": "Country not found", "error": "not_found", "status": 404, "cause": []}`,
	//})
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "", nil)
	c := test_utils.GetMockedContext(request, response)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	GetCountry(c)
	assert.EqualValues(t, http.StatusNotFound, response.Code)

	var apiErr errors.ApiError
	err := json.NewDecoder(response.Body).Decode(&apiErr)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}

func TestGetCountryNoError(t *testing.T) {
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return &locations.Country{
			Id: "AR",
			Name: "Argentina",
		}, nil
	}
	services.LocationService = new(locationsServiceMock)
	//rest.FlushMockups()
	//_ = rest.AddMockups(&rest.Mock{
	//	URL: "https://api.mercadolibre.com/countries/AR",
	//	HTTPMethod: http.MethodGet,
	//	RespHTTPCode: http.StatusOK,
	//	RespBody: `{"id":"AR","name":"Argentina","locale":"es_AR","currency_id":"ARS","decimal_separator":",","thousands_separator":".","time_zone":"GMT-03:00","geo_information":{"location":{"latitude":-38.416096,"longitude":-63.616673}},"states":[{"id":"AR-B","name":"Buenos Aires"},{"id":"AR-C","name":"Capital Federal"},{"id":"AR-K","name":"Catamarca"},{"id":"AR-H","name":"Chaco"},{"id":"AR-U","name":"Chubut"},{"id":"AR-W","name":"Corrientes"},{"id":"AR-X","name":"Córdoba"},{"id":"AR-E","name":"Entre Ríos"},{"id":"AR-P","name":"Formosa"},{"id":"AR-Y","name":"Jujuy"},{"id":"AR-L","name":"La Pampa"},{"id":"AR-F","name":"La Rioja"},{"id":"AR-M","name":"Mendoza"},{"id":"AR-N","name":"Misiones"},{"id":"AR-Q","name":"Neuquén"},{"id":"AR-R","name":"Río Negro"},{"id":"AR-A","name":"Salta"},{"id":"AR-J","name":"San Juan"},{"id":"AR-D","name":"San Luis"},{"id":"AR-Z","name":"Santa Cruz"},{"id":"AR-S","name":"Santa Fe"},{"id":"AR-G","name":"Santiago del Estero"},{"id":"AR-V","name":"Tierra del Fuego"},{"id":"AR-T","name":"Tucumán"}]}`,
	//})
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "", nil)
	c := test_utils.GetMockedContext(request, response)
	c.Params = gin.Params{
		{Key:"country_id", Value:"AR"},
	}
	GetCountry(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	var country locations.Country
	err := json.NewDecoder(response.Body).Decode(&country)
	assert.Nil(t, err)
	assert.EqualValues(t, "Argentina", country.Name)
	assert.EqualValues(t, "AR", country.Id)
}