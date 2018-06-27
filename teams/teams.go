package teams

import (
	"encoding/json"
	"fmt"
)

type team struct {
	Id            int64  `json:"id"`
	Country       string `json:"country"`
	AlternateName string `json:"alternate_name"`
	FifaCode      string `json:"fifa_code"`
	GroupId       int64  `json:"group_id"`
	GroupLetter   string `json:"group_letter"`
}


func GetTeams(body []byte) {

	var teams []team
	err := json.Unmarshal(body, &teams)
	if err != nil {
		fmt.Printf("%s", err)
	}
	for i := range teams {
	fmt.Println(teams[i].Country + " " + teams[i].GroupLetter)
	}

}

