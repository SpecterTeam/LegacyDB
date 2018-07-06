package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/jinzhu/gorm"
)

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
	db.Create(Player{Player:name})
}

func SetPlayerRank(db *gorm.DB, name string, rank int) {
	db.Update(&Player{Player:name,Rank:rank})
}

func SetPlayerIp(db *gorm.DB, name string, ip string) {
	db.Update(&Player{Player:name,Ip:ip})
}

func BanPlayer(db *gorm.DB, name string) {
	db.Update(&Player{Player:name,Banned:1})
}

func PardonPlayer(db *gorm.DB, name string) {
	db.Update(&Player{Player:name,Banned:0})
}