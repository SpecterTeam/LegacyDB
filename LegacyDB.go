package LegacyDB

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"fmt"
	"github.com/SpecterTeam/LegacyDB/api"
)

func Start(port int) *http.Server {

	router := mux.NewRouter()

	router.HandleFunc("/api", api.HandleAPI)
	router.HandleFunc("/api/players", api.HandlePlayers)
	router.HandleFunc("/api/player/{player}", api.HandlePlayer)
	router.HandleFunc("/api/player/create/{player}", api.HandlePlayerCreate)
	router.HandleFunc("/api/player/rank/{player}/{rank}", api.HandlePlayerRank)
	router.HandleFunc("/api/player/ip/{player}/{ip}", api.HandlePlayerIp)
	router.HandleFunc("/api/player/ban/{player}", api.HandleBanPlayer)
	router.HandleFunc("/api/player/pardon/{player}", api.HandlePardonPlayer)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:" + fmt.Sprint(port),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	go srv.ListenAndServe()

	return srv
}
