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

//Ip is crypted from the client.
type Player struct {
	Player string `gorm:"Type:TEXT;Column:player;NOT NULL;primary_key;unique" json:"player"`
	Rank   int    `gorm:"Type:INT;Column:rank;NOT NULL;DEFAULT:0" json:"rank"`
	Ip     string `gorm:"Type:TEXT;Column:ip" json:"ip"`
	Banned int    `gorm:"Type:tinyint(1);Column:banned;NOT NULL;DEFAULT:0" json:"banned"`
}

type HCFStats struct {
	Player      string `gorm:"Type:TEXT;Column:player;NOT NULL;primary_key;unique" json:"player"`
	Kills       int    `gorm:"Type:INT;Column:kills;NOT NULL;DEFAULT:0" json:"kills"`
	Deaths      int    `gorm:"Type:INT;Column:deaths;NOT NULL;DEFAULT:0" json:"deaths"`
	DeathBanned int    `gorm:"Type:tinyint(1);Column:deathbanned;NOT NULL;DEFAULT:0" json:"deathbanned"`
}

type PracticeStats struct {
	Player        string `gorm:"Type:TEXT;Column:player;NOT NULL;primary_key;unique" json:"player"`
	//Elo
	BUHCElo       int    `gorm:"Type:INT;Column:buhcelo;NOT NULL;DEFAULT:1200" json:"buhcelo"`
	NoDebuffElo   int    `gorm:"Type:INT;Column:nodebuffelo;NOT NULL;DEFAULT:1200" json:"nodebuffelo"`
	ClassicElo    int    `gorm:"Type:INT;Column:classicelo;NOT NULL;DEFAULT:1200" json:"classicelo"`
	SumoElo       int    `gorm:"Type:INT;Column:sumoelo;NOT NULL;DEFAULT:1200" json:"sumoelo"`
	//Wins
	BUHCWins      int    `gorm:"Type:INT;Column:buhcwins;NOT NULL;DEFAULT:0" json:"buhcwins"`
	NoDebuffWins  int    `gorm:"Type:INT;Column:nodebuffwins;NOT NULL;DEFAULT:0" json:"nodebuffwins"`
	SumoWins      int    `gorm:"Type:INT;Column:sumowins;NOT NULL;DEFAULT:0" json:"sumowins"`
	ClassicWins   int    `gorm:"Type:INT;Column:classicwins;NOT NULL;DEFAULT:0" json:"classicwins"`
	//Loses
	BUHCLoses     int    `gorm:"Type:INT;Column:buhcloses;NOT NULL;DEFAULT:0" json:"buhcloses"`
	NoDebuffLoses int    `gorm:"Type:INT;Column:nodebuffloses;NOT NULL;DEFAULT:0" json:"nodebuffloses"`
	ClassicLoses  int    `gorm:"Type:INT;Column:classicloses;NOT NULL;DEFAULT:0" json:"classicloses"`
	SumoLoses     int    `gorm:"Type:INT;Column:sumoloses;NOT NULL;DEFAULT:0" json:"sumoloses"`
}
