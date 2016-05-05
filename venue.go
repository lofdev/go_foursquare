package foursquare

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// VenueContact contains venue contact info
type VenueContact struct {
	Twitter        string `json:"twitter"`
	Phone          string `json:"phone"`
	FormattedPhone string `json:"formattedPhone"`
}

// VenueLocation contains venue location info
type VenueLocation struct {
	Address     string  `json:"address"`
	CrossStreet string  `json:"crossStreet"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	PostalCode  string  `json:"postalCode"`
	Country     string  `json:"country"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	Distance    float64 `json:"distance"`
	IsFuzzed    string  `json:"isFuzzed"`
}

// VenueStats contains venue stats
type VenueStats struct {
	CheckinsCount int64 `json:"checkinsCount"`
	UsersCount    int64 `json:"usersCount"`
	TipCount      int64 `json:"tipCount"`
}

// Venue contains a venue
type Venue struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Contact  VenueContact  `json:"contact"`
	Location VenueLocation `json:"location"`
	//categories
	Verified bool       `json:"verified"`
	Stats    VenueStats `json:"stats"`
	URL      string     `json:"url"`
	// There is a bunch more to add
}

// ResponseMeta is the metadata for the Foursquare response
type ResponseMeta struct {
	Code      int64  `json:"code"`
	RequestID string `json:"requestId"`
}

// VenuesResponseBody is a bullshit wrapper around the venues data
type VenuesResponseBody struct {
	Venues []Venue `json:"venues"`
}

// VenuesResponse is the outer wrapper for the Foursquare Response
type VenuesResponse struct {
	Meta     ResponseMeta       `json:"meta"`
	Response VenuesResponseBody `json:"response"`
}

// VenuesSearch requests the /venues/search GET endpoint of Foursquare's API
func (api *UserlessAPI) VenuesSearch(query string, lat float64, lng float64) (result []Venue) {
	requestURL := api.venuesSearchURL(query, lat, lng)
	res, _ := http.Get(requestURL)
	var v VenuesResponse
	json.NewDecoder(res.Body).Decode(&v)
	return v.Response.Venues
}

//  Create a URL for requesting /venues/search
func (api *UserlessAPI) venuesSearchURL(query string, lat float64, lng float64) (result string) {
	return baseURL + "venues/search?limit=50&intent=browse&ll=" +
		strconv.FormatFloat(lat, 'f', -1, 64) + "," +
		strconv.FormatFloat(lng, 'f', -1, 64) + "&" +
		"radius=1000&" +
		"client_id=" + api.credentials.clientID + "&" +
		"client_secret=" + api.credentials.secret + "&" +
		"v=" + apiDate + "&" +
		"m=" + method + "&" +
		"query=" + query
}
