package main

import (
	"fmt"
	"strings"
)

type AllPlayers struct {
	Internal struct {
		PubDateTime string `json:"pubDateTime"`
		XSLT        string `json:"xslt"`
		EventName   string `json:"eventName"`
	}
	League struct {
		Standard   []Player `json:"standard"`
		Africa     []Player `json:"africa"`
		Sacramento []Player `json:"sacramento"`
		Vegas      []Player `json:"vegas"`
		Utah       []Player `json:"utah"`
	}
}

type Player struct {
	FirstName       string           `json:"firstName"`
	LastName        string           `json:"lastName"`
	PersonID        string           `json:"personId"`
	TeamID          string           `json:"teamId"`
	Jersey          string           `json:"jersey"`
	IsActive        bool             `json:"isActive"`
	Pos             PosEnum          `json:"pos"`
	HeightFeet      string           `json:"heightFeet"`
	HeightInches    string           `json:"heightInches"`
	HeightMeters    HeightMetersEnum `json:"heightMeters"`
	WeightPounds    string           `json:"weightPounds"`
	WeightKilograms string           `json:"weightKilograms"`
	DateOfBirthUTC  string           `json:"dateOfBirthUTC"`
	Teams           []Team           `json:"teams"`
	Draft           Draft            `json:"draft"`
	NbaDebutYear    string           `json:"nbaDebutYear"`
	YearsPro        string           `json:"yearsPro"`
	CollegeName     string           `json:"collegeName"`
	LastAffiliation string           `json:"lastAffiliation"`
	Country         string           `json:"country"`
}

type Draft struct {
	TeamID     string `json:"teamId"`
	PickNum    string `json:"pickNum"`
	RoundNum   string `json:"roundNum"`
	SeasonYear string `json:"seasonYear"`
}

type Team struct {
	TeamID      string `json:"teamId"`
	SeasonStart string `json:"seasonStart"`
	SeasonEnd   string `json:"seasonEnd"`
}

type HeightMetersEnum string

const (
	HeightMeters HeightMetersEnum = ""
	The175       HeightMetersEnum = "1.75"
	The178       HeightMetersEnum = "1.78"
	The18        HeightMetersEnum = "1.8"
	The183       HeightMetersEnum = "1.83"
	The185       HeightMetersEnum = "1.85"
	The188       HeightMetersEnum = "1.88"
	The19        HeightMetersEnum = "1.9"
	The193       HeightMetersEnum = "1.93"
	The196       HeightMetersEnum = "1.96"
	The198       HeightMetersEnum = "1.98"
	The201       HeightMetersEnum = "2.01"
	The203       HeightMetersEnum = "2.03"
	The206       HeightMetersEnum = "2.06"
	The208       HeightMetersEnum = "2.08"
	The211       HeightMetersEnum = "2.11"
	The213       HeightMetersEnum = "2.13"
	The216       HeightMetersEnum = "2.16"
	The218       HeightMetersEnum = "2.18"
	The221       HeightMetersEnum = "2.21"
)

type PosEnum string

const (
	C   PosEnum = "C"
	CF  PosEnum = "C-F"
	F   PosEnum = "F"
	FC  PosEnum = "F-C"
	FG  PosEnum = "F-G"
	G   PosEnum = "G"
	GF  PosEnum = "G-F"
	Pos PosEnum = ""
)

func getAllPlayers(endpoint string) AllPlayers {
	NBAInfo := new(AllPlayers)
	getNBAJSON(endpoint, NBAInfo)
	return *NBAInfo
}

func getPlayerID(name string, players []Player) string {
	var result Player

	for _, player := range players {
		fmt.Println(player.LastName)
		if strings.ToLower(name) == strings.ToLower(player.FirstName+" "+player.LastName) {
			result = player
			return result.PersonID
		}
	}
	return "No player found"
}
