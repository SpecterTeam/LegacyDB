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

package LegacyDB

import (
	"testing"
	"github.com/SpecterTeam/LegacyDB/utils"
	"github.com/SpecterTeam/LegacyDB/database"
	"net/http"
	"net/url"
)

var srv *http.Server

func TestServer(t *testing.T) {
	utils.AccessKey = "testing"
	database.InitDB()
	srv = Start(8080)
}

func TestRouter(t *testing.T) {	
	r,err := http.PostForm("http://127.0.0.1:8080/api/player/create/test", url.Values{"access_token":{"testing"},})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if r.StatusCode != http.StatusOK {
		t.FailNow()
	}
	
	r,err := http.PostForm("http://127.0.0.1:8080/api/players", url.Values{"access_token":{"testing"},})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if r.StatusCode != http.StatusOK {
		t.FailNow()
	}

	r,err = http.PostForm("http://127.0.0.1:8080/api", url.Values{"access_token":{"testing"},})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if r.StatusCode != http.StatusOK {
		t.FailNow()
	}
	srv.Close()
}
