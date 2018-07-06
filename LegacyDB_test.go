package LegacyDB

import (
	"testing"
	"github.com/SpecterTeam/LegacyDB/utils"
	"github.com/SpecterTeam/LegacyDB/database"
	"net/http"
	"net/url"
)

func TestServer(t *testing.T) {
	utils.AccessKey = "testing"
	database.InitDB()
	Start(8080)
}

func TestRouter(t *testing.T) {
	r,err := http.PostForm("127.0.0.1:8080/api", url.Values{"access_token":{"testing"},})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if r.StatusCode != http.StatusOK {
		t.FailNow()
	}
}