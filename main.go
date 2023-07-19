package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshuaetim/gallery_go/route"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	log.Fatal(route.RunAPI(":" + os.Getenv("PORT")))
}
