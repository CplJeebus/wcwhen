package groups

import (
	"encoding/json"
	"fmt"
	"strconv"
)



type Groups []struct {
	 Group struct {
		ID     int    `json:"id"`
		Letter string `json:"letter"`
		Teams  []struct {
			Team struct {
				ID               int    `json:"id"`
				Country          string `json:"country"`
				FifaCode         string `json:"fifa_code"`
				Points           int    `json:"points"`
				Wins             int    `json:"wins"`
				Draws            int    `json:"draws"`
				Losses           int `json:"losses"`
				GamesPlayed      int    `json:"games_played"`
				GoalsFor         int    `json:"goals_for"`
				GoalsAgainst     int    `json:"goals_against"`
				GoalDifferential int    `json:"goal_differential"`
			} `json:"team"`
		} `json:"teams"`
	} `json:"group"`
}

func GetGroup (body []byte,group string ){

	var groups Groups
	err := json.Unmarshal(body, &groups)
	if err != nil {
		fmt.Printf("%s", err)
	}
	//fmt.Println(string(body))
	//fmt.Println(groups[0])
	for i := range groups {
		if groups[i].Group.Letter == group {
		fmt.Printf("\n\n============================================\n")
		fmt.Println("Group: " + groups[i].Group.Letter)
		fmt.Printf("============================================\n")
		for x := range groups[i].Group.Teams {
		fmt.Println(groups[i].Group.Teams[x].Team.Country +
			"\t\t" + strconv.Itoa(groups[i].Group.Teams[x].Team.GamesPlayed) +
			"\t"	+ strconv.Itoa(groups[i].Group.Teams[x].Team.Wins) +
			"\t" + strconv.Itoa(groups[i].Group.Teams[x].Team.Draws) +
			"\t" + strconv.Itoa(groups[i].Group.Teams[x].Team.Losses) +
			"\t" + strconv.Itoa(groups[i].Group.Teams[x].Team.Points))
		}
	}
	}
}
