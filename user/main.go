package main

import (
	"log"
	"os"

	"github.com/syauqeesy/accounting-service/user/application"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run main.go [http|grpc|migration|seeder]")
	}

	if os.Args[1] == application.ApplicationMigration && len(os.Args) < 3 {
		log.Fatal("usage: go run main.go [generate|up|down]")
	}

	applicationType := os.Args[1]

	arguments := make([]string, 0, 2)

	if os.Args[1] == application.ApplicationMigration && os.Args[2] == application.MigrationGenerate && len(os.Args) > 3 {
		arguments = append(arguments, os.Args[2], os.Args[3])
	} else if os.Args[1] == application.ApplicationMigration && os.Args[2] != application.MigrationGenerate && len(os.Args) < 3 {
		log.Fatal("usage: go run main.go [generate {migration_name}]")
	} else if os.Args[1] == application.ApplicationMigration && os.Args[2] != application.MigrationGenerate {
		arguments = append(arguments, os.Args[2])
	}

	err := application.Run(applicationType, arguments)
	if err != nil {
		log.Fatal(err.Error())
	}
}
