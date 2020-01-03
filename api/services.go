package api

import (

	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/nidhish1/go-api/models"
	"github.com/nidhish1/go-api/repository"

)

func JourneyDetails (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	jid := params["jid"]

	
	var DriverJourney[] models.Journey = repository.GetJourney(jid )
	json.NewEncoder(w).Encode(DriverJourney)
}


