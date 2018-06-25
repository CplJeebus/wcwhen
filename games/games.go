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
	Id          int64  `json:"id"`
	Player      string `json:"player"`
	TypeOfEvent string `json:"type_of_event"`
	Time        string `json:"time"`
}

type games struct {
	Venue             string  `json:"venue"`
	Location          string  `json:"location"`
	Status            string  `json:"status"`
	Time              string  `json:"time"`
	FifaId            string  `json:"fifa_id"`
	Datetime          string  `json:"datetime"`
	LastEventUpdateAt string  `json:"last_event_update_at"`
	LastScoreUpdateAt string  `json:"last_score_update_at"`
	HomeTeam          team    `json:"home_team"`
	AwayTeam          team    `json:"away_team"`
	Country           string  `json:"teams"`
	Winner            string  `json:"winner"`
	WinnerCode        string  `json:"winner_code"`
	HomeTeamEvents    []event `json:"home_team_events"`
	AwayTeamEvents    []event `json:"away_team_events"`
}

func Getgames(bodyBytes []byte, Country string) {

	var evts []games
	err := json.Unmarshal(bodyBytes, &evts)
	if err != nil {
		fmt.Printf("%s", err)
	}

	for f := range evts {

		if evts[f].HomeTeam.Country == Country ||
			evts[f].AwayTeam.Country == Country {
			fmt.Printf("\n\n============================================\n")
			fmt.Println(evts[f].HomeTeam.Country + " Vs " + evts[f].AwayTeam.Country)

			if len(evts[f].LastEventUpdateAt) == 0 {
				fmt.Println("\nScheduled \t" + evts[f].Datetime)
			} else {
				fmt.Println("\nLast updated \t" + evts[f].LastEventUpdateAt)
			}

			fmt.Printf("============================================\n")

			for x := range evts[f].HomeTeamEvents {
				fmt.Println("\t" + evts[f].HomeTeamEvents[x].Time + " " +
					evts[f].HomeTeamEvents[x].TypeOfEvent +
					" " + evts[f].HomeTeamEvents[x].Player)
			}

			for x := range evts[f].AwayTeamEvents {
				fmt.Println(evts[f].AwayTeamEvents[x].Time + " " +
					evts[f].AwayTeamEvents[x].TypeOfEvent +
					" " + evts[f].AwayTeamEvents[x].Player)
			}

		}
	}
}
