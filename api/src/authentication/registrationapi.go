package authentication

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"net/http"
)

func (api *API) BeginRegistration(c *gin.Context) {
	webAuthn, err := webauthn.New(wconfig)
	if err != nil {
		fmt.Println(err)
	}
	
	user.Name = "dbarkwell"
	user.DisplayName = "Dan Barkwell"
	user.ID = []byte("83d973fa-6294-4933-a1fc-9ef42d0b73bd")

	var options = &protocol.CredentialCreation{}
	options, session, err = webAuthn.BeginRegistration(user)
	if session == nil {
		return
	}

	// handle errors if present
	// store the sessionData values
	c.JSON(http.StatusOK, options) // return the options generated
	// options.publicKey contain our registration options
}

func (api *API) FinishRegistration(c *gin.Context) {
	webAuthn, err := webauthn.New(wconfig)
	if err != nil {
		fmt.Println(err)
	}

	//user := datastore.GetUser() // Get the user

	// Get the session data stored from the function above
	//session := datastore.GetSession()

	credential, err := webAuthn.FinishRegistration(user, *session, c.Request)
	if err != nil {
		// Handle Error and return.

		return
	}

	// If creation was successful, store the credential object
	// Pseudocode to add the user credential.
	user.Credential = append(user.Credential, *credential)
	//datastore.SaveUser(user)

	c.JSON(http.StatusOK, "Registration Success") // Handle next steps
}
