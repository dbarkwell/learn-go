package authentication

import (
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
)

type credential struct {
	ID              []byte          `db:"id"`
	UserID          uuid.UUID       `db:"user_id"`
	PublicKey       []byte          `db:"public_key"`
	AttestationType string          `db:"attestation_type"`
	Transport       Transport       `db:"transport"`
	Flags           CredentialFlags `db:"flags"`
	Authenticator   Authenticator   `db:"authenticator"`
}

type Transport []string

type credentialSession struct {
	UserID                   []byte                 `db:"user_id"`
	Challenge                string                 `db:"challenge"`
	AllowedCredentialIDS     [][]byte               `db:"allowed_credential_ids"`
	Expires                  time.Time              `db:"expires"`
	UserVerification         string                 `db:"user_verification"`
	AuthenticationExtensions map[string]interface{} `db:"authentication_extensions"`
}

type CredentialDTO struct {
	ID              []byte          `json:"id"`
	PublicKey       []byte          `json:"publicKey"`
	AttestationType string          `json:"attestationType"`
	Transport       []string        `json:"transport"`
	Flags           CredentialFlags `json:"flags"`
	Authenticator   Authenticator   `json:"authenticator"`
}

func (c *CredentialDTO) RawID() string {
	return base64.RawURLEncoding.EncodeToString(c.ID)
}

type CredentialFlags struct {
	UserPresent    bool
	UserVerified   bool
	BackupEligible bool
	BackupState    bool
}

type Authenticator struct {
	AAGUID       []byte
	SignCount    uint32
	CloneWarning bool
	Attachment   string
}

type UserCredentialDTO struct {
	UserID          uuid.UUID       `json:"userID""`
	Username        string          `json:"username"`
	DisplayName     string          `json:"displayName"`
	PublicKey       []byte          `json:"publicKey"`
	AttestationType string          `json:"attestationType"`
	Transport       []string        `json:"transport"`
	Flags           CredentialFlags `json:"flags"`
	Authenticator   Authenticator   `json:"authenticator"`
}

type CredentialSessionDTO struct {
	UserID                   uuid.UUID              `json:"userID"`
	Challenge                string                 `json:"challenge"`
	AllowedCredentialIDs     [][]byte               `json:"allowedCredentialIDs"`
	Expires                  time.Time              `json:"expires"`
	UserVerification         string                 `json:"userVerification"`
	AuthenticationExtensions map[string]interface{} `json:"authenticationExtensions"`
}

func (t *Transport) Scan(value interface{}) error {
	return scan[*Transport](&t, value)
}

func (t Transport) Value() (driver.Value, error) {
	return value[Transport](t)
}

func (f *CredentialFlags) Scan(value interface{}) error {
	return scan[*CredentialFlags](&f, value)
}

func (f CredentialFlags) Value() (driver.Value, error) {
	return value[CredentialFlags](f)
}

func (a *Authenticator) Scan(value interface{}) error {
	return scan[*Authenticator](&a, value)
}

func (a Authenticator) Value() (driver.Value, error) {
	return value[Authenticator](a)
}

func value[T Transport | CredentialFlags | Authenticator](t T) (driver.Value, error) {
	return json.Marshal(t)
}

func scan[T *Transport | *CredentialFlags | *Authenticator](t *T, value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &t)
}
