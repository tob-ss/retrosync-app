package main

import (
	"context"
	"fmt"
	"maps"
	"net/http"
	"slices"
	"time"

	"github.com/Khan/genqlient/graphql"
	//"math/rand"
)

func (a *App) GetSaves() []map[string]interface{} {
	fmt.Println("Started Get saves function!")

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	client := graphql.NewClient("http://localhost:8080/query", http.DefaultClient)

	userID := 420

	fmt.Println("Querying the API server!")

	resp, err := getLocalSaves(ctx, client, userID)

	if err != nil {
		fmt.Println("failed to get local saves")
		return nil
	}

	fmt.Println("Successfully got local saves! Parsing saves...")

	sliceOfSaves := createMaps(resp.GetLocalSaves.IDs, resp.GetLocalSaves.UserIDs, resp.GetLocalSaves.Names, resp.GetLocalSaves.Consoles, resp.GetLocalSaves.Devices, resp.GetLocalSaves.TimeMods, resp.GetLocalSaves.Paths)

	fmt.Println("Saves parsed! Sending Saves to frontend...")

	fmt.Println("Here are the saves for debugging reasons", sliceOfSaves)

	return sliceOfSaves
}

func createMaps(ids []int, userIDs []int, names []string, consoles []string, devices []string, timeMods []int, paths []string) []map[string]interface{} {
	id_map := maps.Collect(slices.All(ids))
	user_map := maps.Collect(slices.All(userIDs))
	name_map := maps.Collect(slices.All(names))
	console_map := maps.Collect(slices.All(consoles))
	device_map := maps.Collect(slices.All(devices))
	timeStrings := timeConverter(timeMods)
	time_map := maps.Collect(slices.All(timeStrings))
	path_map := maps.Collect(slices.All(paths))

	sliceOfSaves := []map[string]interface{}{}
	slicePoint := &sliceOfSaves

	for index, id := range id_map {
		gameSave := make(map[string]interface{})
		gameSave["ID"] = id
		userID := user_map[index]
		gameSave["User_ID"] = userID
		name := name_map[index]
		gameSave["Game_Name"] = name
		console := console_map[index]
		gameSave["Console"] = console
		device := device_map[index]
		gameSave["Device"] = device
		timeMod := time_map[index]
		gameSave["Time_Modified"] = timeMod
		path := path_map[index]
		gameSave["Save_Path"] = path

		*slicePoint = append(*slicePoint, gameSave)
	}

	return sliceOfSaves
}

func timeConverter(timeMod []int) []string {
	timeStrings := []string{}
	timePoint := &timeStrings
	for _, timemod := range timeMod {
		timeNow := (time.Now()).Unix()
		difference := timeNow - int64(timemod)
		if difference <= 60 {
			*timePoint = append(*timePoint, fmt.Sprintf("%ds ago", difference))
			continue
		} else if difference <= 3600 && difference > 60 {
			difference = difference / 60
			*timePoint = append(*timePoint, fmt.Sprintf("%dm ago", difference))
			continue
		} else if difference <= 86400 && difference > 3600 {
			difference = difference / 3600
			*timePoint = append(*timePoint, fmt.Sprintf("%dh ago", difference))
			continue
		} else if difference <= 604800 && difference > 86400 {
			difference = difference / 86400
			*timePoint = append(*timePoint, fmt.Sprintf("%dd ago", difference))
			continue
		} else if difference <= 31449600 && difference > 604800 {
			difference = difference / 604800
			*timePoint = append(*timePoint, fmt.Sprintf("%dw ago", difference))
			continue
		} else if difference >= 31449600 {
			difference = difference / 31449600
			*timePoint = append(*timePoint, fmt.Sprintf("%dy ago", difference))
			continue
		}

	}

	return timeStrings
}
