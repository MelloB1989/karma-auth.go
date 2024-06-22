package helpers

import (
	"karma_auth/database"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Correct the function signature to include a return type of ([]database.User, error)
func GetUsers() ([]database.User, error) {
    db, err := sqlx.Connect("postgres", "user=postgres dbname=karma-auth sslmode=require password=pEOUSii7EkypE4u08bNe host=database-1.ckrzvsguulcs.ap-south-1.rds.amazonaws.com")
    if err != nil {
        log.Fatalln(err)
        return nil, err // Return nil slice and error
    }
  
    defer db.Close()

    // Test the connection to the database
    if err := db.Ping(); err != nil {
        log.Fatal(err)
        return nil, err // Return nil slice and error
    } else {
        log.Println("Successfully Connected")
    }

    users := []database.User{}

    rows, err := db.Queryx("SELECT name, email FROM organisations")
    if err != nil {
        log.Fatalln(err)
        return nil, err // Return nil slice and error
    }

    for rows.Next() {
        var user database.User // Temporary variable to hold each user
        err := rows.StructScan(&user)
        if err != nil {
            log.Fatalln(err)
            return nil, err // Return nil slice and error
        }
        users = append(users, user) // Append the user to the slice
    }

    return users, nil // Return the populated slice and nil error
}