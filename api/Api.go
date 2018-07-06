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

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	if utils.IsAccessTokenValid(r) {
		w.WriteHeader(http.StatusOK)
		utils.WriteJson(w, ApiResponse{Name: "LegacyDB", Version: 1.0, Author: "Fris", Status: true})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
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