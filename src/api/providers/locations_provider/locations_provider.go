package locations_provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/tv2169145/golang-testing/src/api/domain/locations"
	"github.com/tv2169145/golang-testing/src/api/utils/errors"
	"net/http"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	response := rest.Get(fmt.Sprintf(urlGetCountry, countryId))
	if response == nil || response.Response == nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient error to getting country %s", countryId),
		}
	}

	if response.StatusCode > 299 {
		var apiErr errors.ApiError
		if err := json.NewDecoder(bytes.NewReader(response.Bytes())).Decode(&apiErr); err != nil {
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error response when getting country %s", countryId),
			}
		}
		return nil, &apiErr
	}
	var result locations.Country
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryId),
		}
	}

	return &result, nil
}
