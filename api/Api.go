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
package api

import (
	"net/http"
	"github.com/SpecterTeam/LegacyDB/utils"
	"github.com/gorilla/mux"
	"github.com/SpecterTeam/LegacyDB/database"
	"strconv"
)

type ApiResponse struct {
	Name    string  `json:"Name"`
	Version float64 `json:"Version"`
	Author  string  `json:"Author"`
	Status  bool    `json:"Status"`
}

func HandleAPI(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	utils.WriteJson(w, ApiResponse{Name: "LegacyDB", Version: 1.0, Author: "Fris", Status: true})
}

func HandlePlayer(w http.ResponseWriter, r *http.Request) {
	if utils.IsAccessTokenValid(r) {
		name := mux.Vars(r)["name"]
		w.WriteHeader(http.StatusOK)
		utils.WriteJson(w, database.GetPlayer(database.Instance, name))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func HandlePlayers(w http.ResponseWriter, r *http.Request) {
	if utils.IsAccessTokenValid(r) {
		w.WriteHeader(http.StatusOK)
		utils.WriteJson(w, database.GetPlayers(database.Instance))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func HandlePlayerCreate(w http.ResponseWriter, r *http.Request) {
	if utils.IsAccessTokenValid(r) {
		name := mux.Vars(r)["player"]
		w.WriteHeader(http.StatusOK)
		database.CreatePlayer(database.Instance, name)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func HandlePlayerRank(w http.ResponseWriter, r *http.Request) {
	if utils.IsAccessTokenValid(r) {
		name := mux.Vars(r)["player"]
		rank,_ := strconv.Atoi(mux.Vars(r)["rank"])
		w.WriteHeader(http.StatusOK)
		database.SetPlayerRank(database.Instance, name, rank)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func HandlePlayerIp(w http.ResponseWriter, r *http.Request) {
	if utils.IsAccessTokenValid(r) {
		name := mux.Vars(r)["player"]
		ip := mux.Vars(r)["ip"]
		w.WriteHeader(http.StatusOK)
		database.SetPlayerIp(database.Instance, name, ip)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func HandleBanPlayer(w http.ResponseWriter, r *http.Request) {
	if utils.IsAccessTokenValid(r) {
		name := mux.Vars(r)["player"]
		w.WriteHeader(http.StatusOK)
		database.BanPlayer(database.Instance, name)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func HandlePardonPlayer(w http.ResponseWriter, r *http.Request) {
	if utils.IsAccessTokenValid(r) {
		name := mux.Vars(r)["player"]
		w.WriteHeader(http.StatusOK)
		database.PardonPlayer(database.Instance, name)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func HandlePlayerExist(w http.ResponseWriter, r *http.Request) {
	if utils.IsAccessTokenValid(r) {
		name := mux.Vars(r)["player"]
		w.WriteHeader(http.StatusOK)
		utils.WriteJson(w, database.ExistPlayer(database.Instance, name))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func HandleUpdatePracticeStats(w http.ResponseWriter, r *http.Request) {
	if utils.IsAccessTokenValid(r) {
		name := mux.Vars(r)["player"]
		w.WriteHeader(http.StatusOK)
		stats := map[string]int{}
		for k,v := range r.PostForm {
			value, _ := strconv.Atoi(v[0])
			switch k {
			//Elo
			case "buhcelo":
				stats["BUHCElo"] = value
				break
			case "nodebuffelo":
				stats["NoDebuffElo"] = value
				break
			case "classicelo":
				stats["ClassicElo"] = value
				break
			case "sumoelo":
				stats["SumoElo"] = value
				break
				//Wins
			case "buhcwins":
				stats["BUHCWins"] = value
				break
			case "nodebuffwins":
				stats["NoDebuffWins"] = value
				break
			case "classicwins":
				stats["ClassicWins"] = value
				break
			case "sumowins":
				stats["SumoWins"] = value
				break
				//Loses
			case "buhcloses":
				stats["BUHCLoses"] = value
				break
			case "nodebuffloses":
				stats["NoDebuffLoses"] = value
				break
			case "classicloses":
				stats["ClassicLoses"] = value
				break
			case "sumoloses":
				stats["SumoLoses"] = value
				break
			}
		}
		database.UpdatePracticeStats(database.Instance, name, stats)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func HandlePracticeStats(w http.ResponseWriter, r *http.Request) {
	if utils.IsAccessTokenValid(r) {
		w.WriteHeader(http.StatusOK)
		utils.WriteJson(w, database.GetAllPracticeStats(database.Instance))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func HandlePracticeStatsPlayer(w http.ResponseWriter, r *http.Request) {
	if utils.IsAccessTokenValid(r) {
		name := mux.Vars(r)["player"]
		w.WriteHeader(http.StatusOK)
		utils.WriteJson(w, database.GetPracticeStats(database.Instance, name))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}