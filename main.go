package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	"strings"
)


type team struct {
	Country string `json:"country"`
	Code  string `json:"code"`
	Goals  int `json:"goals"`

}

type event struct {
	Id int64 `json:"id"`
	Player string `json:"player"`
	Type_of_event string `json:"type_of_event"`
	Time string `json:"time"`
}
type games struct {
	Venue string `json:"venue"`
	Location string `json:"location"`
	Status string `json:"status"`
	Time string `json:"time"`
	Fifa_id string `json:"fifa_id"`
	Datetime string `json:"datetime"`
	Last_event_update_at string `json:"last_event_update_at"`
	Last_score_update_at string `json:"last_score_update_at"`
	Home_team team `json:"home_team"`
	Away_team team `json:"away_team"`
	Country string `json:"country"`
	Winner string `json:"winner"`
	Winner_code string `json:"winner_code"`
	Home_team_events []event `json:"home_team_events"`
	Away_team_events []event `json:"away_team_events"`
}


func main() {
	url := "http://worldcup.sfg.io/matches"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var evts []games
	err = json.Unmarshal(bodyBytes, &evts)
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println(strings.Join(os.Args[1:]," "))

	for i := range evts {
		if evts[i].Home_team.Country == strings.Join(os.Args[1:]," ") ||
				evts[i].Away_team.Country == strings.Join(os.Args[1:]," ") {
			fmt.Printf("\n\n============================================\n")
			fmt.Println(evts[i].Home_team.Country + " Vs " + evts[i].Away_team.Country)

			if len(evts[i].Last_event_update_at) == 0 {
				fmt.Println("\nScheduled \t" + evts[i].Datetime)
			} else {
				fmt.Println("\nLast updated \t" + evts[i].Last_event_update_at)
			}

			fmt.Printf("============================================\n")

			for x := range evts[i].Home_team_events {
				fmt.Println("\t" + evts[i].Home_team_events[x].Time + " " +
					evts[i].Home_team_events[x].Type_of_event +
					" " + evts[i].Home_team_events[x].Player)
			}

			for x := range evts[i].Away_team_events {
				fmt.Println(evts[i].Away_team_events[x].Time + " " +
					evts[i].Away_team_events[x].Type_of_event +
					" " + evts[i].Away_team_events[x].Player)
			}

		}
	}

}