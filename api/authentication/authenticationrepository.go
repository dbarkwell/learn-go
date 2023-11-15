package authentication

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type Repository struct {
	db *sqlx.DB
}

func ProvideAuthenticationRepository(DB *sqlx.DB) Repository {
	return Repository{db: DB}
}

func (ar *Repository) SaveCredential(userID uuid.UUID, c credential) (bool, error) {
	insert := `INSERT INTO passkey_credential (id, user_id, public_key, attestation_type, transport, flags, authenticator) 
			VALUES (?, UUID_TO_BIN(?), ?, ?, CAST(? AS JSON), CAST(? AS JSON), CAST(? AS JSON))`
	result := ar.db.MustExec(insert, c.ID, userID, c.PublicKey, c.AttestationType, c.Transport, c.Flags, c.Authenticator)

	if rows, err := result.RowsAffected(); err != nil || rows == 0 {
		return false, err
	}

	return true, nil
}

func (ar *Repository) GetUserCredentials(userID uuid.UUID) ([]credential, error) {
	sel := `SELECT id, BIN_TO_UUID(id) user_id, 
       	public_key, attestation_type, transport, flags, authenticator 
		FROM passkey_credential WHERE user_id = UUID_TO_BIN(?)`
	stmt, err := ar.db.Preparex(sel)
	if err != nil {
		return make([]credential, 0), err
	}

	defer stmt.Close()

	var cred []credential
	err = stmt.Select(&cred, userID)

	return cred, nil
}

func (ar *Repository) SaveCredentialSession(userID uuid.UUID, s credentialSession) (bool, error) {
	ac, _ := json.Marshal(s.AllowedCredentialIDS)
	ae, _ := json.Marshal(s.AuthenticationExtensions)

	insert := `INSERT INTO passkey_session (user_id, challenge, allowed_credential_ids, expires, user_verification, authentication_extensions) 
			VALUES (UUID_TO_BIN(?), ?, CAST(? AS JSON), ?, ?, CAST(? AS JSON))`
	result := ar.db.MustExec(insert, userID, s.Challenge, string(ac), toNullTime(s.Expires), s.UserVerification, string(ae))

	//user := User{Username: username, FirstName: first, LastName: last, Email: email}

	if rows, err := result.RowsAffected(); err != nil || rows == 0 {
		return false, err
	}

	return true, nil
}

func (ar *Repository) GetCredentialSession(userID uuid.UUID) (credentialSession, error) {
	sel := `SELECT challenge, allowed_credential_ids , expires, user_verification, authentication_extensions 
		FROM passkey_session WHERE user_id = UUID_TO_BIN(?)`
	stmt, err := ar.db.Preparex(sel)
	var session credentialSession
	if err != nil {
		return session, nil
	}

	defer stmt.Close()

	err = stmt.Get(&session, userID)

	return session, nil
}

func (ar *Repository) RemoveCredentialSession(userID uuid.UUID) (bool, error) {
	del := `DELETE FROM passkey_session WHERE user_id = UUID_TO_BIN(?)`
	stmt, err := ar.db.Preparex(del)

	if err != nil {
		return false, nil
	}

	defer stmt.Close()

	_, err = stmt.Exec(userID)

	return true, nil
}

func toNullTime(expires time.Time) sql.NullTime {
	if expires.IsZero() {
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  expires,
		Valid: true,
	}
}
