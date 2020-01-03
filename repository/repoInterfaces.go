package repository

import (
	
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/nidhish1/go-api/models"
	"github.com/nidhish1/go-api/dbCon"
	
)


func models.Journey() [] getJourney(jid string) {
	var journeys []models.Journey
	result, err := dbCon.Db.Query("SELECT JourneyId, OrderRequestID, Status from Lynk_Journey where JourneyId =  "+ jid)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var journey models.Journey
		err := result.Scan(&journey.Jid, &journey.Ordrqst,&journey.Status )
		if err != nil {
			panic(err.Error())
		}
		journeys = append(journeys, journey)
	}
	fmt.Println("Working")
	
	return journeys

}

