package authentication

import (
	"github.com/go-webauthn/webauthn/webauthn"
)

var (
	wconfig = &webauthn.Config{
		RPDisplayName: "Go Webauthn",                                          // Display Name for your site
		RPID:          "aa0c-99-228-180-110.ngrok-free.app",                   // Generally the FQDN for your site
		RPOrigins:     []string{"https://aa0c-99-228-180-110.ngrok-free.app"}, // The origin URLs allowed for WebAuthn requests
	}

	session = &webauthn.SessionData{}
	user    = &LoginUser{}
)

type API struct {
	AuthenticationService Service
}

func ProvideAuthenticationAPI(as Service) API {
	return API{AuthenticationService: as}
}
