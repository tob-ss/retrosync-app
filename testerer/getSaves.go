package main

import (
	"context"
	"fmt"
	"maps"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/Khan/genqlient/graphql"
	//"math/rand"
)

func (a App) GetHeaders(savesSlice []map[string]interface{}) []string {
	headers := createHeaders(savesSlice)

	fmt.Println("headers are ", headers)

	return headers
}

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

	//fmt.Println("Here are the saves for debugging reasons", sliceOfSaves)

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
	epoch_map := maps.Collect(slices.All(timeMods))
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
		epoch := epoch_map[index]
		gameSave["Epoch_Time"] = epoch

		*slicePoint = append(*slicePoint, gameSave)
	}

	sliceOfSaves = addDate(sliceOfSaves)

	return sliceOfSaves
}

func addDate(savesSlice []map[string]interface{}) []map[string]interface{} {
	timeNow := (time.Now().Unix())
	timeYesterday := timeNow - 86400

	yesterdaysDate := time.Unix(timeYesterday, 0).Format(time.RFC822Z)

	todaysDate := time.Unix(timeNow, 0).Format(time.RFC822Z)

	yesterdayFormatted := string(yesterdaysDate[0:9])
	todayFormatted := string(todaysDate[0:9])

	for _, saveMap := range savesSlice {
		if integer, ok := (saveMap["Epoch_Time"]).(int); ok {
			dateEpoch := integer

			date := time.Unix(int64(dateEpoch), 0).Format(time.RFC822Z)

			dateFormatted := string(date[0:9])

			if strings.Compare(todayFormatted, dateFormatted) != 0 && strings.Compare(yesterdayFormatted, dateFormatted) != 0 {
				saveMap["Date_String"] = dateFormatted
			} else if strings.Compare(todayFormatted, dateFormatted) == 0 {
				saveMap["Date_String"] = "Today"
			} else if strings.Compare(yesterdayFormatted, dateFormatted) == 0 {
				saveMap["Date_String"] = "Yesterday"
			}
		}

	}

	return savesSlice
}

func sliceSearcher(targetSlice []string, targetString string) bool {
	var exists bool

	for _, x := range targetSlice {
		if strings.Compare(x, targetString) != 0 {
			exists = false
		} else {
			exists = true
		}
	}

	return exists
}

func createHeaders(savesSlice []map[string]interface{}) []string {
	timeMod := []int{}
	timePoint := &timeMod

	//fmt.Println("savesslice length is", len(savesSlice))

	for _, gameMap := range savesSlice {
		//fmt.Println("gameMap name, epoch is and length is", gameMap["Game_Name"], gameMap["Epoch_Time"], len(gameMap))

		epoch := gameMap["Epoch_Time"]

		if float, ok := (epoch).(float64); ok {
			epoch = int(float)
		}

		//fmt.Println("the type of epoch is", reflect.TypeOf(epoch))

		if integer, ok := (epoch).(int); ok {
			dateEpoch := integer
			*timePoint = append(*timePoint, dateEpoch)
		} else {
			//fmt.Println("ok and integer is", ok, integer)

		}
		//fmt.Println("timemod is now:", timeMod)
	}

	fmt.Println("timeMod is", timeMod)

	slices.Sort(timeMod)
	slices.Reverse(timeMod)
	dayHeaders := []string{}
	dayPoint := &dayHeaders
	timeNow := (time.Now().Unix())
	timeYesterday := timeNow - 86400

	yesterdaysDate := time.Unix(timeYesterday, 0).Format(time.RFC822Z)

	todaysDate := time.Unix(timeNow, 0).Format(time.RFC822Z)

	yesterdayFormatted := string(yesterdaysDate[0:9])
	todayFormatted := string(todaysDate[0:9])

	for _, dateEpoch := range timeMod {

		date := time.Unix(int64(dateEpoch), 0).Format(time.RFC822Z)
		dateFormatted := string(date[0:9])

		fmt.Println(dateFormatted)

		found := sliceSearcher(*dayPoint, dateFormatted)
		todayFound := sliceSearcher(*dayPoint, "Today")
		yesterdayFound := sliceSearcher(*dayPoint, "Yesterday")
		if found {
			fmt.Println("I found a duplicate!", found, dateFormatted)
			continue
		} else {
			fmt.Println("There's no duplicates...", found, dateFormatted)
			if strings.Compare(todayFormatted, dateFormatted) != 0 && strings.Compare(yesterdayFormatted, dateFormatted) != 0 {
				*dayPoint = append(*dayPoint, dateFormatted)
			} else if strings.Compare(todayFormatted, dateFormatted) == 0 && !todayFound {
				*dayPoint = append(*dayPoint, "Today")
			} else if strings.Compare(yesterdayFormatted, dateFormatted) == 0 && !yesterdayFound {
				*dayPoint = append(*dayPoint, "Yesterday")
			}
		}
	}

	return dayHeaders
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
