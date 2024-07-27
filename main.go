package main

import (
	database "Egy2015/cms_go/database/config"
	"Egy2015/cms_go/router"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = db.SqlDb.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err = router.Run(db); err != nil {
		log.Fatal(err)
	}
}
