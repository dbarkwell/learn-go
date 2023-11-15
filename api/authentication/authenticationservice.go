package authentication

import (
	"fmt"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	"learn-go.barkwell.com/user"
)

type Service struct {
	AuthenticationRepository Repository
	UserService              user.Service
}

func ProvideAuthenticationService(ar Repository, us user.Service) Service {
	return Service{
		AuthenticationRepository: ar,
		UserService:              us,
	}
}

func (as Service) SaveCredential(userID uuid.UUID, c CredentialDTO) {
	cred := credential{
		ID:              c.ID,
		PublicKey:       c.PublicKey,
		AttestationType: c.AttestationType,
		Transport:       c.Transport,
		Flags:           c.Flags,
		Authenticator:   c.Authenticator,
	}

	as.AuthenticationRepository.SaveCredential(userID, cred)
}

func (as Service) GetUser(username string) (*LoginUser, error) {
	user, err := as.UserService.Get(username)
	if err != nil {
		return new(LoginUser), err
	}

	loginUser := &LoginUser{
		Name:        user.Username,
		DisplayName: fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		ID:          user.ID,
	}

	return loginUser, nil
}

func (as Service) GetUserAndCredentials(username string) (*LoginUser, error) {
	user, err := as.GetUser(username)

	if err != nil {

	}

	userCreds, err := as.AuthenticationRepository.GetUserCredentials(user.ID)
	if err != nil {

	}

	creds := make([]webauthn.Credential, len(userCreds))
	for i, v := range userCreds {
		creds[i] = webauthn.Credential{
			ID:              v.ID[:],
			PublicKey:       v.PublicKey,
			AttestationType: v.AttestationType,
			Transport:       convertTransport(v.Transport),
			Flags: webauthn.CredentialFlags{
				UserPresent:    v.Flags.UserPresent,
				UserVerified:   v.Flags.UserVerified,
				BackupEligible: v.Flags.BackupEligible,
				BackupState:    v.Flags.BackupState,
			},
			Authenticator: webauthn.Authenticator{
				AAGUID:       v.Authenticator.AAGUID,
				SignCount:    v.Authenticator.SignCount,
				CloneWarning: v.Authenticator.CloneWarning,
				Attachment:   protocol.AuthenticatorAttachment(v.Authenticator.Attachment),
			},
		}
	}

	user.Credential = creds
	return user, nil
}

func (as Service) SaveCredentialSession(userID uuid.UUID, s *webauthn.SessionData) {
	cred := credentialSession{
		Challenge:                s.Challenge,
		AllowedCredentialIDS:     s.AllowedCredentialIDs,
		Expires:                  s.Expires,
		UserVerification:         string(s.UserVerification),
		AuthenticationExtensions: s.Extensions,
	}

	as.AuthenticationRepository.RemoveCredentialSession(userID)
	as.AuthenticationRepository.SaveCredentialSession(userID, cred)
}

func (as Service) GetCredentialSession(userID uuid.UUID) (webauthn.SessionData, error) {
	session, err := as.AuthenticationRepository.GetCredentialSession(userID)
	if err != nil {
		return webauthn.SessionData{}, err
	}

	sessionData := webauthn.SessionData{
		UserID:               userID[:],
		Challenge:            session.Challenge,
		Expires:              session.Expires,
		AllowedCredentialIDs: session.AllowedCredentialIDS,
		Extensions:           session.AuthenticationExtensions,
		UserVerification:     protocol.UserVerificationRequirement(session.UserVerification),
	}

	return sessionData, nil
}

func (as Service) RemoveCredentialSession(userID uuid.UUID) (bool, error) {
	session, err := as.AuthenticationRepository.RemoveCredentialSession(userID)
	if err != nil {
		return false, err
	}

	return session, nil
}

func convertTransport(transport []string) []protocol.AuthenticatorTransport {
	at := make([]protocol.AuthenticatorTransport, len(transport))
	for _, v := range transport {
		at = append(at, protocol.AuthenticatorTransport(v))
	}

	return at
}
