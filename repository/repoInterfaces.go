package repository

import (
	

	"fmt"


	"github.com/nidhish1/go-api/dbCon"	
	"github.com/nidhish1/go-api/models"	
)


func GetJourney(jid string) [] models.Journey  {
	var journeys [] models.Journey
	result, err := dbCon.Db.Query("SELECT JourneyId, OrderRequestID, Status from Lynk_Journey where JourneyId = " + " '"+jid + " ' ")

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

