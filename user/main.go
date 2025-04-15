package main

import (
	"log"
	"os"

	"github.com/syauqeesy/accounting-service/application"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run main.go [http|grpc|migration|seeder]")
	}

	if os.Args[1] == application.ApplicationMigration && len(os.Args) < 3 {
		log.Fatal("usage: go run main.go [generate|up|down]")
	}

	applicationType := os.Args[1]

	err := application.Run(applicationType)
	if err != nil {
		log.Fatal(err.Error())
	}
}
