package main

import (
	"flag"
	"fmt"
	"github.com/SpecterTeam/LegacyDB/database"
	"os"
	"github.com/SpecterTeam/LegacyDB/utils"
	"bufio"
	"github.com/SpecterTeam/LegacyDB"
)

func init() {
	fmt.Println("Initializing database...")
	database.InitDB()
	fmt.Println("Database loaded.")
	fmt.Println("Generating Access key...")
	utils.AccessKey = "example-password-69-lol"
	os.Setenv("LEGACY_DB_ACCESS_KEY", utils.AccessKey)
	fmt.Println("Access key loaded.")
}

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "set port for the http server.")
	flag.Parse()

	srv := LegacyDB.Start(port)

	fmt.Println("LegacyDB has been started. Press ENTER to quit.")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println("Stopping the http server...")
	srv.Close()
	fmt.Println("Successfully shutdown LegacyDB.")
}
