package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/LettsDoSomeCoding/ranko/logging"
	"github.com/LettsDoSomeCoding/ranko/persistence"
)

func createPlayerHandler(w http.ResponseWriter, r *http.Request) {
	logging.GetLogger().Print("Create Player endpoint hit")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logHTTPError(w, err, http.StatusBadRequest)
		return
	}

	var player persistence.Player
	json.Unmarshal(reqBody, &player)
	// ID shouldn't be provided on create. If it is, ignore it (should this be an error?)
	player.ID = 0

	result, err := player.Save()

	if err != nil {
		logHTTPError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func returnAllPlayers(w http.ResponseWriter, r *http.Request) {
	logging.GetLogger().Print("Return All Players endpoint hit")
	json.NewEncoder(w).Encode(persistence.GetAllPlayers())
}
