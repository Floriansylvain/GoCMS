package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type pingFormat struct {
	Message string `json:"message"`
}

func TestPing(t *testing.T) {
	godotenv.Load("../.env")

	r := gin.Default()
	r.GET("/ping/", Ping)
	go r.Run(":" + os.Getenv("API_PORT"))

	req, err := http.Get(fmt.Sprintf("http://localhost:%v/ping/", os.Getenv("API_PORT")))
	if err != nil {
		t.Fatal(err.Error())
	}
	defer req.Body.Close()

	var exceptedResp = pingFormat{}
	body, _ := io.ReadAll(req.Body)
	json.Unmarshal(body, &exceptedResp)

	if exceptedResp.Message != "pong" {
		t.Fatalf(`/ping/ resulted in "%v" instead of "pong"`, exceptedResp.Message)
	}
}
