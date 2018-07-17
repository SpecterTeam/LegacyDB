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
package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/jinzhu/gorm"
)

type PlayerExistJson struct {
	Exist bool `json:"exist"`
}

var Instance *gorm.DB

func InitDB() {
	Instance = NewSqlite3("legacy.db")
}


func NewSqlite3(file string) *gorm.DB {
	db, err := gorm.Open("sqlite3", file)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Player{}, &HCFStats{}, &PracticeStats{})

	return db
}

func GetPlayers(db *gorm.DB) []Player {
	var players []Player
	db.First(&players)
	return players
}

func GetPlayer(db *gorm.DB, name string) Player {
	var player Player
	db.Find(&player, Player{Player:name})
	return player
}

func CreatePlayer(db *gorm.DB, name string) {
	db.Create(&Player{Player:name})
	db.Create(&HCFStats{Player:name})
	db.Create(&PracticeStats{Player:name})
}

func SetPlayerRank(db *gorm.DB, name string, rank int) {
	db.Model(&Player{Player:name}).Update("Rank", rank)
}

func SetPlayerIp(db *gorm.DB, name string, ip string) {
	db.Model(&Player{Player:name}).Update("Ip", ip)
}

func BanPlayer(db *gorm.DB, name string) {
	db.Model(&Player{Player:name}).Update("Banned", 1)
}

func PardonPlayer(db *gorm.DB, name string) {
	db.Model(&Player{Player:name}).Update("Banned", 0)
}

func ExistPlayer(db *gorm.DB, name string) PlayerExistJson {
	var player []Player
	db.Find(&player, Player{Player:name})
	return PlayerExistJson{Exist: len(player) != 0}
}

func UpdatePracticeStats(db *gorm.DB, name string, stats map[string]int) {
	for key,value := range stats {
		db.Model(&PracticeStats{Player:name}).Update(key, value)
	}
}

func GetPracticeStats(db *gorm.DB, name string) PracticeStats {
	var stats PracticeStats
	db.Find(&stats, PracticeStats{Player:name})
	return stats
}

func GetAllPracticeStats(db *gorm.DB) []PracticeStats {
	var stats []PracticeStats
	db.Find(&stats)
	return stats
}