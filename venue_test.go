package foursquare

import (
	"os"
	"testing"
)

func TestVenueSearchURL(t *testing.T) {
	it := Userless()
	// Is it wrong to use Galvanize Boulder as my location?
	properURL := "https://api.foursquare.com/v2/venues/search?limit=50&" +
		"intent=browse&ll=40.017786,-105.281948&" +
		"radius=1000&" +
		"client_id=" + os.Getenv("FOURSQUARE_CLIENT_ID") + "&" +
		"client_secret=" + os.Getenv("FOURSQUARE_CLIENT_SECRET") + "&" +
		"v=20160415&m=foursquare" +
		"&query=food"
	if it.venuesSearchURL("food", 40.017786, -105.281948) != properURL {
		t.Errorf("Improperly formed request url")
	}
}

func TestVenueSearch(t *testing.T) {
	it := Userless()
	if len(it.venuesSearchURL("food", 40.017786, -105.281948)) == 0 {
		t.Errorf("Foursquare returned nothing.")
	}
}
