package aqicn

// BaseFeed abstract feed data
type BaseFeed struct {
	StationID int          `json:"idx"`
	Time      ObsLocalTime `json:"time"`
	City      struct {
		Name           string     `json:"name"`
		GeoLocation    [2]float64 `json:"geo"`
		AttributionURL string     `json:"url"`
	} `json:"city"`
}

// SearchFeed feed data provided to SearchCity
type SearchFeed struct {
	BaseFeed
	AQI string `json:"aqi"`
}

// DetailFeed detailed feed data
type DetailFeed struct {
	BaseFeed
	AQI          int `json:"aqi"`
	Attributions []struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"attributions"`
	DominentPol string `json:"dominentpol"`
	IAQI        map[string]struct {
		Value Number `json:"v"`
	} `json:"iaqi"`
}
