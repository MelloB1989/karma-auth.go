package users

import (
	"karma_auth/database"
	"log"
	"time"

	_ "github.com/lib/pq"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GetUsers() ([]database.User, error) {
    db, err := database.DBConn()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

    users := []database.User{}

    rows, err := db.Queryx("SELECT name, email FROM organisations")
    if err != nil {
        log.Fatalln(err)
        return nil, err
    }

    for rows.Next() {
        var user database.User // Temporary variable to hold each user
        err := rows.StructScan(&user)
        if err != nil {
            log.Fatalln(err)
            return nil, err
        }
        users = append(users, user) // Append the user to the slice
    }

    return users, nil
}

func CreateUser(oid string, name string, email string, password string, phone string, details string){
	db, err := database.DBConn()
	if err != nil {
		log.Fatalln(err)
	}
	
	//Generate a uid
	id, _ := gonanoid.Generate("qwertyuiopasdfghjklzxcvbnm1234567890_-", 10);
	uid := oid + "---" + id

	r, err := db.Exec(`INSERT INTO users (uid, oid, name, email, password, phone, details, "createdAt", "email_verified", "phone_verified") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, uid, oid, name, email, password, phone, details, time.Now(), false, false)

	if err != nil || r == nil {
		log.Fatalln(err)
	}
}