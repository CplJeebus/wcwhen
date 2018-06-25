package groups

import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/ryanuber/columnize"

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

func GetGroup (body []byte,group string ) {

	var groups Groups
	err := json.Unmarshal(body, &groups)
	if err != nil {
		fmt.Printf("%s", err)
	}

	for i := range groups {
		if groups[i].Group.Letter == group {
			fmt.Printf("\n\n============================================\n")
			fmt.Println("Group: " + groups[i].Group.Letter)
			fmt.Printf("============================================\n")
			var out []string
			out = append(out, "| GP | W | D | L | P | GF | GA | GD")
			for x := range groups[i].Group.Teams {
				out = append(out, groups[i].Group.Teams[x].Team.Country+
					"|"+ strconv.Itoa(groups[i].Group.Teams[x].Team.GamesPlayed)+
					"|"+ strconv.Itoa(groups[i].Group.Teams[x].Team.Wins)+
					"|"+ strconv.Itoa(groups[i].Group.Teams[x].Team.Draws)+
					"|"+ strconv.Itoa(groups[i].Group.Teams[x].Team.Losses)+
					"|"+ strconv.Itoa(groups[i].Group.Teams[x].Team.Points)+
					"|"+ strconv.Itoa(groups[i].Group.Teams[x].Team.GoalsFor)+
					"|"+ strconv.Itoa(groups[i].Group.Teams[x].Team.GoalsAgainst)+
					"|"+ strconv.Itoa(groups[i].Group.Teams[x].Team.GoalDifferential))
			}
			fmt.Println(columnize.SimpleFormat(out))
		}
	}
}

func GetGroupList (body []byte) []string {
	var groups Groups
	err := json.Unmarshal(body, &groups)
	if err != nil {
		fmt.Printf("%s", err)
	}
	var groupList []string

	for i := range groups {
		groupList = append(groupList,groups[i].Group.Letter)
	}

	return groupList
	}



