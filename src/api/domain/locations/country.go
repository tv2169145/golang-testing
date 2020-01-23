package locations

type Country struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Locale         string         `json:"locale"`
	TimeZone       string         `json:"time_zone"`
	GeoInformation GeoInformation `json:"geo_information"`
	States         []States       `json:"states"`
}

type GeoInformation struct {
	Location Location `json:"location"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type States struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
