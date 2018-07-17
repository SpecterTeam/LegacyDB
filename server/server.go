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
package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"fmt"
	"github.com/SpecterTeam/LegacyDB/api"
)

func Start(port int) *http.Server {

	router := mux.NewRouter()

	router.HandleFunc("/", api.HandleAPI)
	router.HandleFunc("/api", api.HandleAPI)
	router.HandleFunc("/api/players", api.HandlePlayers)
	router.HandleFunc("/api/player/{player}", api.HandlePlayer)
	router.HandleFunc("/api/player/create/{player}", api.HandlePlayerCreate)
	router.HandleFunc("/api/player/rank/{player}/{rank}", api.HandlePlayerRank)
	router.HandleFunc("/api/player/ip/{player}/{ip}", api.HandlePlayerIp)
	router.HandleFunc("/api/player/ban/{player}", api.HandleBanPlayer)
	router.HandleFunc("/api/player/pardon/{player}", api.HandlePardonPlayer)
	router.HandleFunc("/api/player/exist/{player}", api.HandlePlayerExist)
	router.HandleFunc("/api/practice/stats", api.HandlePracticeStats)
	router.HandleFunc("/api/practice/stats/{player}", api.HandlePracticeStatsPlayer)
	router.HandleFunc("/api/practice/stats/update/{player}", api.HandleUpdatePracticeStats)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:" + fmt.Sprint(port),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	go srv.ListenAndServe()

	return srv
}
