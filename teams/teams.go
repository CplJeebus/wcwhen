package teams

import (
	"encoding/json"
	"fmt"
)

type team struct {
	Id int64 `json:"id"`
	Country string `json:"country"`
	Alternate_name string `json:"alternate_name"`
	Fifa_code string `json:"fifa_code"`
	Group_id int64 `json:"group_id"`
	Group_letter string `json:"group_letter"`
}


func GetTeams(body []byte) {

	var teams []team
	err := json.Unmarshal(body, &teams)
	if err != nil {
		fmt.Printf("%s", err)
	}
	for i := range teams {
	fmt.Println(teams[i].Country)
	}

}

