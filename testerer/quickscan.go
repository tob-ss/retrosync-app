package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Khan/genqlient/graphql"
	//"github.com/robfig/cron/v3"
)

func quickScan(device string, userID int, scanPath string) error {
	x := 0
	for {

		err := postToDB(device, userID)

		if err != nil {
			//fmt.Print("Unexpected error", err)
			return err
		}
		// 150 loops should take at least 25 mins
		x += 1
		if (x%150 == 0) && (x/150 >= 1) {
			startFullScan(scanPath)
		}
	}
}

func startFullScan(scanPath string) {
	fmt.Println("Starting full scan!")

	fullScan(scanPath)

}

func postToDB(device string, userID int) error {
	time.Sleep(10 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client := graphql.NewClient("http://localhost:8080/query", http.DefaultClient)

	resp, err := getPaths(ctx, client, device, userID)

	if err != nil {
		log.Println("json.Compact:", err)
	}
	defer cancel()

	fmt.Println(resp)

	updated := make(map[int]int)

	id_paths, err := joinLists(resp.GetPaths.Paths, resp.GetPaths.IDs)

	if err != nil {
		fmt.Println("failed to join lists!", err)
		return err
	}

	for id, path := range id_paths {
		dir, timemod, exists := handleFiles(path)
		if !exists {
			// send to API server to delete
			fmt.Println("path no longer exists!", dir)

		} else {
			updated[id] = timemod
			fmt.Println("adding id and timemod to map to be sent to api:", id, timemod)
		}
	}

	// send map with dirs and timemods to insert/update quickscan table
	fmt.Println("sending map of dirs and timemods:", updated)
	for id, timeMod := range updated {
		// send path and timeMod in function
		postTime(id, timeMod, ctx, client)
	}
	return nil
}

func deleteID(id int, ctx context.Context, client graphql.Client) {
	resp, err := deleteLocalID(ctx, client, id)

	if err != nil {
		fmt.Println("Tried to delete metadata but got error...", err)

	} else {
		fmt.Println("Successfully deleted metadata with no errors!")
	}

	_ = resp
}

func joinLists(paths []string, ids []int) (map[int]string, error) {
	dir_timemod := make(map[int]string)

	for idx := range paths {
		dir_timemod[ids[idx]] = paths[idx]
	}
	return dir_timemod, nil
}

func postTime(id int, timeMod int, ctx context.Context, client graphql.Client) {

	resp, err := updateTime(ctx, client, id, timeMod)

	if err != nil {
		fmt.Println("Posted metadata but got error...", err)

	} else {
		fmt.Println("Successfully posted metadata with no errors!")
	}

	_ = resp

}

func handleFiles(path string) (string, int, bool) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("got an error checking the dir", err)
		return path, 0, false
	}

	if fileInfo.IsDir() {
		// call function to loop through files in folder and capture the file with the most recent time mod
		timeMods := searchFolder(path)
		timeMod := getTimemod(timeMods)
		return path, timeMod, true
	} else {
		// call get info
		dir, timemod := getTime(path)
		return dir, timemod, true
	}
}

func getTime(path string) (string, int) {
	fileInfo, err := os.Stat(path)
	modTime := fileInfo.ModTime()
	if err != nil {
		fmt.Println("Unexpected error:", err)
		return path, 0
	}
	return path, int(modTime.Unix())
}

func searchFolder(dir string) []int {
	timeSlice := []int{}
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			dir, timemod := getTime(path)
			_ = dir
			timeSlice = append(timeSlice, timemod)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Unexpected error:", err)
	}
	return timeSlice
}

func getTimemod(timeMods []int) int {
	biggest := timeMods[0]

	for _, v := range timeMods {
		if v > biggest {
			biggest = v
		}
	}
	return biggest
}
