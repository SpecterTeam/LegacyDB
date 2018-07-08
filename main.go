/**
 *     LegacyDB  Copyright (C) 2018  SpecterTeam
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"flag"
	"fmt"
	"github.com/SpecterTeam/LegacyDB/database"
	"os"
	"github.com/SpecterTeam/LegacyDB/utils"
	"bufio"
	"github.com/SpecterTeam/LegacyDB/LegacyDB"
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
