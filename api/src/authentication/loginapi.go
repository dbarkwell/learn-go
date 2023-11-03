package authentication

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"net/http"
)

func (api *API) BeginLogin(c *gin.Context) {
	webAuthn, err := webauthn.New(wconfig)
	if err != nil {
		fmt.Println(err)
	}

	//user := datastore.GetUser() // Find the user
	var options = &protocol.CredentialAssertion{}
	options, session, err = webAuthn.BeginLogin(user)
	if err != nil {
		// Handle Error and return.

		return
	}

	// store the session values
	//datastore.SaveSession(session)

	c.JSON(http.StatusOK, options)
}

func (api *API) FinishLogin(c *gin.Context) {
	webAuthn, err := webauthn.New(wconfig)
	if err != nil {
		fmt.Println(err)
	}

	//user := datastore.GetUser() // Get the user

	// Get the session data stored from the function above
	//session := datastore.GetSession()

	// _ credential
	credential, err := webAuthn.FinishLogin(user, *session, c.Request)
	if err != nil {
		// Handle Error and return.
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// If login was successful, handle next steps
	c.JSON(http.StatusOK, fmt.Sprintf("Login Success %s", base64.StdEncoding.EncodeToString(credential.ID)))
}
