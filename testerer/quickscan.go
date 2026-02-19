package main

import (
	"context"
	"time"
	"log"
	"fmt"
	"net/http"
	"github.com/Khan/genqlient/graphql"
	"os"
	"path/filepath"
	"io/fs"
)

func quickScan(device string, userID int) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()
		client := graphql.NewClient("http://localhost:8080/query", http.DefaultClient)

		resp, err := getPaths(ctx, client, device, userID)

		if err != nil {
			log.Println("json.Compact:", err)
		}

		fmt.Println(resp)

		updated := make(map[string]int)

		for _, path := range resp.GetPaths.Paths {
			dir, timemod, exists := handleFiles(path)
			if !exists {
				// send to API server to delete
				fmt.Println("path no longer exists!", dir)
			} else {
				updated[dir] = timemod
				fmt.Println("adding dir and timemod to map to be sent to api:", dir, timemod)
			}
		}

		// send map with dirs and timemods to insert/update quickscan table
		fmt.Println("sending map of dirs and timemods:", updated)
	}
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

func getTimemod(timeMods []int) (int) {
	biggest := timeMods[0]

	for _, v := range timeMods {
		if v > biggest {
			biggest = v
		}
	}
	return biggest
}
