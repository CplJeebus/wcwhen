package games

import (
	"encoding/json"
	"fmt"
)

type team struct {
	Country string `json:"country"`
	Code    string `json:"code"`
	Goals   int    `json:"goals"`
}

type event struct {
	Id            int64  `json:"id"`
	Player        string `json:"player"`
	Type_of_event string `json:"type_of_event"`
	Time          string `json:"time"`
}
type games struct {
	Venue                string  `json:"venue"`
	Location             string  `json:"location"`
	Status               string  `json:"status"`
	Time                 string  `json:"time"`
	Fifa_id              string  `json:"fifa_id"`
	Datetime             string  `json:"datetime"`
	Last_event_update_at string  `json:"last_event_update_at"`
	Last_score_update_at string  `json:"last_score_update_at"`
	Home_team            team    `json:"home_team"`
	Away_team            team    `json:"away_team"`
	Country              string  `json:"teams"`
	Winner               string  `json:"winner"`
	Winner_code          string  `json:"winner_code"`
	Home_team_events     []event `json:"home_team_events"`
	Away_team_events     []event `json:"away_team_events"`
}

func Getgames(bodyBytes []byte, Country string) {

	var evts []games
	err := json.Unmarshal(bodyBytes, &evts)
	if err != nil {
		fmt.Printf("%s", err)
	}

	for f := range evts {

		if evts[f].Home_team.Country == Country ||
			evts[f].Away_team.Country == Country {
			fmt.Printf("\n\n============================================\n")
			fmt.Println(evts[f].Home_team.Country + " Vs " + evts[f].Away_team.Country)

			if len(evts[f].Last_event_update_at) == 0 {
				fmt.Println("\nScheduled \t" + evts[f].Datetime)
			} else {
				fmt.Println("\nLast updated \t" + evts[f].Last_event_update_at)
			}

			fmt.Printf("============================================\n")

			for x := range evts[f].Home_team_events {
				fmt.Println("\t" + evts[f].Home_team_events[x].Time + " " +
					evts[f].Home_team_events[x].Type_of_event +
					" " + evts[f].Home_team_events[x].Player)
			}

			for x := range evts[f].Away_team_events {
				fmt.Println(evts[f].Away_team_events[x].Time + " " +
					evts[f].Away_team_events[x].Type_of_event +
					" " + evts[f].Away_team_events[x].Player)
			}

		}
	}
}
