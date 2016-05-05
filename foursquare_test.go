package foursquare

import (
	"os"
	"testing"
)

func TestUserless(t *testing.T) {
	got := Userless()
	if got.credentials.clientID != os.Getenv("FOURSQUARE_CLIENT_ID") {
		t.Errorf("Userless returns bad client id")
	}
	if got.credentials.secret != os.Getenv("FOURSQUARE_CLIENT_SECRET") {
		t.Errorf("Userless returns bad client secret")
	}
}
