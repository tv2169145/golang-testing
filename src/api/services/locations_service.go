package services

import (
	"github.com/tv2169145/golang-testing/src/api/domain/locations"
	"github.com/tv2169145/golang-testing/src/api/providers/locations_provider"
	"github.com/tv2169145/golang-testing/src/api/utils/errors"
)

type LocationsServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

type locationService struct {}

var (
	LocationService LocationsServiceInterface
)

func init() {
	LocationService = new(locationService)
}

func (s *locationService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return locations_provider.GetCountry(countryId)
}
