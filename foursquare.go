package foursquare

import "os"

const (
	baseURL = "https://api.foursquare.com/v2/"
	apiDate = "20160415"
	method  = "foursquare"
)

type appSecret struct {
	clientID string
	secret   string
}

// ResponseMeta is the metadata for the Foursquare response
type ResponseMeta struct {
	Code      int64  `json:"code"`
	RequestID string `json:"requestId"`
}

// UserlessAPI returns a container for doing userless API requests
type UserlessAPI struct {
	credentials appSecret
}

func (api *UserlessAPI) setCredentials() {
	api.credentials.clientID = os.Getenv("FOURSQUARE_CLIENT_ID")
	api.credentials.secret = os.Getenv("FOURSQUARE_CLIENT_SECRET")
}

// Userless returns a userless API container
func Userless() UserlessAPI {
	var uAPI UserlessAPI
	uAPI.setCredentials()
	return uAPI
}
