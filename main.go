package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/wcwhen/games"
	"flag"
	"github.com/wcwhen/teams"
	"github.com/wcwhen/groups"
)


func main() {
	game := flag.String("g","","Please provide a valid teams use -t to get valid countries")
	team := flag.Bool("t",false,"This will provide you a list of countries")
	groupList := flag.Bool("gl",false,"This will provide you a list of groups")
	group := flag.String("gp","","List of groups")

	flag.Parse()

	if len(*game) != 0 {
		url := "http://worldcup.sfg.io/matches"
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%s", err)
		}
	//	fmt.Println("The Fuck")
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		games.Getgames(bodyBytes, *game)
	}

	if *groupList == true {
		url := "https://worldcup.sfg.io/teams/group_results"
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%s", err)
		}

		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(groups.GetGroupList(bodyBytes))
	}

	if *team == true {
		url := "http://worldcup.sfg.io/teams"
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%s", err)
		}

		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		teams.GetTeams(bodyBytes)
	}

	if len(*group) != 0 {

		url := "https://worldcup.sfg.io/teams/group_results"
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%s", err)
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		if *group == "all" {
			list := groups.GetGroupList(bodyBytes)
			for x := range list {
				groups.GetGroup(bodyBytes, list[x])
			}
		} else {
			groups.GetGroup(bodyBytes, *group)
		}

	}

}