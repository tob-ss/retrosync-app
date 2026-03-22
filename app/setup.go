package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	//"strings"
	"os"
	//"io"
	"time"
	//"sync"
	//"regexp"
	"math"
	"net/http"
	"strings"

	"github.com/Khan/genqlient/graphql"
	//"math/rand"
)

func startScan() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	saveSearch(homeDir)
}

func StartQuickScan() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	err = quickScan("Desktop", 420, homeDir)

	if err != nil {
		fmt.Println("Unexpected error:", err)
	}
}

var progress float64
var progressPointer *float64 = &progress

func CheckProgress() float64 {
	rounded := math.Round((progress * 100))

	return rounded
}

// need a variant of listfiles/listfolders that takes in a specific parameter first e.g. retrosync or emulator name, so it can do a quick search

// will keep code for full search

// when user starts scan; it first does the faster search with the parameters, then there will be an prompt or something that the user can click on which will initiate a full scan for saves (or the user can choose where the save is)

// also add some error handling/message in case no games are found

func searchResolver(console string, consoleFolders map[string]string) []string {
	var results []string
	resultsPointer := &results
	fmt.Println("Parsing", console, "folders!")
	start := time.Now()
	for key, value := range consoleFolders {
		if value == "retroarch" && console == "retroarch" {
			listFiles, err := listFiles(key)
			if err != nil {
				log.Fatal(err)
			}

			for _, path := range listFiles {
				*resultsPointer = append(*resultsPointer, path)
			}
		}

		if value == console && console != "retroarch" {
			for _, path := range listFolders(key, console, true) {
				*resultsPointer = append(*resultsPointer, path)
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Finished", console, "folders!", elapsed)

	return results
}

func consoleSearch(dir string) map[string]string {
	fmt.Println("Starting console search")
	start := time.Now()
	check := ""
	var m map[string]string
	m = make(map[string]string)

	//var dirSize float64

	/*err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
	       	dirSize += 1
			return nil
		   	})
	    if err != nil {
			fmt.Println(err)
	    }

		fmt.Println("Directory size is", dirSize)

	*/

	//var x float64

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		//x += 1
		//*progressPointer = (0.7 * x) / dirSize
		if d.IsDir() {
			base := strings.ToLower(filepath.Base(path))
			check = "retroarch"
			if strings.Contains(base, check) {
				m[path] = check
			}
			check = "dolphin"
			if strings.Contains(base, check) {
				m[path] = check
			}
			check = "ppsspp"
			if strings.Contains(base, check) {
				m[path] = check
			}
			check = "rpcs3"
			if strings.Contains(base, check) {
				m[path] = check
			}
			check = "azahar"
			if strings.Contains(base, check) {
				m[path] = check
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	elapsed := time.Since(start)
	fmt.Println("Finished console search, time elapsed", elapsed)

	return m
}

func listFiles(dir string) ([]string, error) {
	//fmt.Println("Starting listFiles search for", dir)
	//start := time.Now()
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {

		if filepath.Ext(path) == ".srm" || filepath.Ext(path) == ".dsv" || filepath.Ext(path) == ".ps2" || filepath.Ext(path) == ".gci" {
			files = append(files, path)
		}

		/*if d.IsDir() {
		    return nil
		}*/

		//files = append(files, path)

		//q <- path

		return nil
	})

	_ = err

	//elapsed := time.Since(start)
	//fmt.Println("Finished", dir, "time elapsed", elapsed)

	return files, nil
}

func listFolders(dir string, console string, quick bool) []string {
	//fmt.Println("Starting listFolders search for", dir)
	//start := time.Now()
	var folders []string

	if quick {
		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				if filepath.Base(path) == "title" && console == "dolphin" {
					folders = append(folders, path)

				} else if filepath.Base(path) == "SAVEDATA" && console == "ppsspp" {
					folders = append(folders, path)

				} else if filepath.Base(path) == "savedata" && console == "rpcs3" {
					folders = append(folders, path)

				} else if filepath.Base(path) == "title" && console == "azahar" {
					folders = append(folders, path)

				}
			}

			return nil
		})
		_ = err
	} else {
		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				if filepath.Base(path) == "title" && console == "wii" {
					folders = append(folders, path)

				} else if filepath.Base(path) == "SAVEDATA" && console == "psp" {
					folders = append(folders, path)

				} else if filepath.Base(path) == "savedata" && console == "ps3" {
					folders = append(folders, path)

				} else if filepath.Base(path) == "title" && console == "n3ds" {
					folders = append(folders, path)

				}
			}
			_ = err
			return nil
		})
		_ = err
	}

	parsedFolders := searchFolders(folders)

	//elapsed := time.Since(start)
	fmt.Println("Finished", dir)

	return parsedFolders
}

func searchFolders(dirs []string) []string {
	var files []string

	for _, dir := range dirs {

		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	return files
}

func getInfo(console string, files []string) ([]string, []int) {
	directories := []string{}
	dir_pointer := &directories
	timeModified := []int{}
	time_pointer := &timeModified
	for _, save := range files {
		fileInfo, err := os.Stat(save)
		modTime := fileInfo.ModTime()
		if err != nil {
			log.Fatal(err.Error())
		}
		*dir_pointer = append(*dir_pointer, save)
		*time_pointer = append(*time_pointer, int(modTime.Unix()))
	}
	return directories, timeModified
}

func saveSearch(dir string) {

	start := time.Now()

	*progressPointer = 0

	fmt.Println("doing consoleSearch, current elapsed time is,", time.Since(start))

	consoleFolders := consoleSearch(dir)

	*progressPointer = 0.25

	// support for custom saves
	// support for user to select a folder to do a custom search
	// progress bar should start off at 0

	retro := searchResolver("retroarch", consoleFolders)
	wii := searchResolver("dolphin", consoleFolders)
	psp := searchResolver("ppsspp", consoleFolders)
	ps3 := searchResolver("rpcs3", consoleFolders)
	n3ds := searchResolver("azahar", consoleFolders)

	*progressPointer = 0.5

	// custom := searchResolver("custom", consoleFolders)
	// general idea is that we want to prompt the user to pick the folder they want to use, then we call list folders on that path

	//fmt.Println("doing getInfo, current elapsed time is,", time.Since(start), retro, wii)

	retro_dirs, retro_time := getInfo("retro", retro)
	wii_dirs, wii_time := getInfo("wii", wii)
	psp_dirs, psp_time := getInfo("psp", psp)
	ps3_dirs, ps3_time := getInfo("ps3", ps3)
	fmt.Println("the n3ds variable is:", n3ds)
	n3ds_dirs, n3ds_time := getInfo("n3ds", n3ds)

	*progressPointer = 0.75

	//fmt.Println("doing postsaves, current elapsed time is,", time.Since(start))

	flushSaves("Desktop", 420)

	*progressPointer = 0.79

	postSaves("Desktop", "retro", retro_dirs, retro_time)
	*progressPointer = 0.83
	postSaves("Desktop", "wii", wii_dirs, wii_time)
	*progressPointer = 0.87
	postSaves("Desktop", "psp", psp_dirs, psp_time)
	*progressPointer = 0.91
	postSaves("Desktop", "ps3", ps3_dirs, ps3_time)
	*progressPointer = 0.95
	fmt.Println("the n3ds directory list is:", n3ds_dirs)
	postSaves("Desktop", "n3ds", n3ds_dirs, n3ds_time)

	elapsed := time.Since(start)
	fmt.Println(elapsed)

	*progressPointer = 1

	err := quickScan("Desktop", 420, dir)

	if err != nil {
		fmt.Println("Unexpected error:", err)
	}
}

func flushSaves(device string, userID int) {
	ctx := context.Background()
	client := graphql.NewClient("http://localhost:8080/query", http.DefaultClient)

	resp, err := deleteLocal(ctx, client, device, userID)

	if err != nil {
		log.Println("json.Compact:", err)
	}

	_ = resp
}

func postSaves(device string, console string, dirs []string, timemods []int) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	client := graphql.NewClient("http://localhost:8080/query", http.DefaultClient)

	resp, err := createSaves(ctx, client, device, console, dirs, timemods)

	if err != nil {
		log.Println("json.Compact:", err)
	}

	_ = resp

	//fmt.Printf("Posted metadata", resp)

	//rand.Seed(time.Now().UnixNano())

	//n := rand.Intn(30)

	//time.Sleep(time.Duration(n)*time.Second)
}
