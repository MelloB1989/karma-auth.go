package organisation

import (
	"karma_auth/database"
	"log"
)

func GetOrganisationApiToken (oid string) database.ApiToken {
	db, err := database.DBConn()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	token := database.ApiToken{}
	r := db.QueryRow(`SELECT oid, tid, token, "createdAt", "expiresAt", details FROM api_tokens WHERE oid = $1`, oid).Scan(&token.OID, &token.TID, &token.Token, &token.CreatedAt, &token.ExpiresAt, &token.Details)
	log.Println(r)

	return token
}