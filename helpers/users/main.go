package users

import (
	"karma_auth/database"
	"log"
	_ "github.com/lib/pq"
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