package infrastructure

import (
	"database/sql"
	"fmt"
	"gin-project/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL(config config.Config) (*sql.DB, error) {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		config.Database.User,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
	)

	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	return db, err
}
