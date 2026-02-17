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
	"github.com/Khan/genqlient/graphql"
	"net/http"
	"strings"
	"math"
	"math/rand"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) StartScan() {
	saveSearch()
}

var progress float64
var progressPointer *float64 = &progress

func (a *App) CheckProgress() float64 {
	rounded := math.Round((progress * 100))
	
	return rounded
}

// need a variant of listfiles/listfolders that takes in a specific parameter first e.g. retrosync or emulator name, so it can do a quick search

// will keep code for full search 

// when user starts scan; it first does the faster search with the parameters, then there will be an prompt or something that the user can click on which will initiate a full scan for saves (or the user can choose where the save is)

// also add some error handling/message in case no games are found

func searchResolver(console string, consoleFolders map[string]string) ([]string) {
	var results []string
	resultsPointer := &results
	fmt.Println("Parsing", console,"folders!")
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
	fmt.Println("Finished", console,"folders!", elapsed)
	
	return results
}


/*func searchResolver(console string) ([]string) {
	// switch case for each console; within each one; they'll call listfiles/listfolders and use different dirs depending on the console and it will return whatever listfiles/listfolders return
	var results []string
	resultsPointer := &results
	switch console {
	case "retro":
		fmt.Println("Parsing", console,"folders!")
		start := time.Now()
		retroarchFolders := consoleSearch("/", "retroarch")
		for _, dir := range retroarchFolders {
			listFiles, err := listFiles(dir)
			if err != nil {
				log.Fatal(err)
			}
			for _, path := range listFiles {
				*resultsPointer = append(*resultsPointer, path)
			}
		}
		elapsed := time.Since(start)
		fmt.Println("Finished", console,"folders!", elapsed)
		return results
	case "wii":
		fmt.Println("Parsing", console,"folders!")
		start := time.Now()
		wiiFolders := consoleSearch("/", "dolphin")
		for _, dir := range wiiFolders {
			for _, path := range listFolders(dir, console) {
				*resultsPointer = append(*resultsPointer, path)
			}
		}
		elapsed := time.Since(start)
		fmt.Println("Finished", console,"folders!", elapsed)
		return results
	case "psp":
		fmt.Println("Parsing", console,"folders!")
		start := time.Now()
		pspFolders := consoleSearch("/", "ppsspp")
		for _, dir := range pspFolders {
			for _, path := range listFolders(dir, console) {
				*resultsPointer = append(*resultsPointer, path)
			}
		}
		elapsed := time.Since(start)
		fmt.Println("Finished", console,"folders!", elapsed)
		return results
	case "ps3":
		fmt.Println("Parsing", console,"folders!")
		start := time.Now()
		ps3Folders := consoleSearch("/", "rpcs3")
		for _, dir := range ps3Folders {
			for _, path := range listFolders(dir, console) {
				*resultsPointer = append(*resultsPointer, path)
			}
		}
		elapsed := time.Since(start)
		fmt.Println("Finished", console,"folders!", elapsed)
		return results
	case "n3ds":
		fmt.Println("Parsing", console,"folders!")
		start := time.Now()
		n3dsFolders := consoleSearch("/", "azahar")
		for _, dir := range n3dsFolders {
			for _, path := range listFolders(dir, console) {
				*resultsPointer = append(*resultsPointer, path)
			}
		}
		elapsed := time.Since(start)
		fmt.Println("Finished", console,"folders!", elapsed)
		return results
	default:
		return nil
	}

	return nil
}*/


func consoleSearch(dir string) (map[string]string) {
	fmt.Println("Starting console search")
	start := time.Now()
	check := ""
	var m map[string]string
	m = make(map[string]string)

	var dirSize float64

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
       	dirSize += 1
		return nil
	   	})
    if err != nil {
		fmt.Println(err)
    }

	fmt.Println("Directory size is", dirSize)

	var x float64

    err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		x += 1
		*progressPointer = (0.7 * x) / dirSize
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
					
				} else if filepath.Base(path) == "save" && console == "azahar" {
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
						
					} else if filepath.Base(path) == "save" && console == "n3ds" {
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
			if !d.IsDir() || d.IsDir() {
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

func getInfo(console string, files []string) ([]string, []int64) {
	directories := []string{}
	dir_pointer := &directories 
	timeModified := []int64{}
	time_pointer := &timeModified
	for _, save := range files {
		fileInfo, err := os.Stat(save)
		modTime := fileInfo.ModTime()
		if err != nil {
			log.Fatal(err.Error())
		}
		*dir_pointer = append(*dir_pointer, save)
		*time_pointer = append(*time_pointer, modTime.Unix())
	}
	return directories, timeModified
}

func saveSearch() {

	start := time.Now()

	fmt.Println("doing consoleSearch, current elapsed time is,", time.Since(start))

	consoleFolders := consoleSearch("/")

	// support for custom saves
	// support for user to select a folder to do a custom search
	// progress bar should start off at 0

	retro := searchResolver("retroarch", consoleFolders)
	wii := searchResolver("dolphin", consoleFolders)
	psp := searchResolver("ppsspp", consoleFolders)
	ps3 := searchResolver("rpcs3", consoleFolders)
	n3ds := searchResolver("azahar", consoleFolders)

	// custom := searchResolver("custom", consoleFolders)
	// general idea is that we want to prompt the user to pick the folder they want to use, then we call list folders on that path

	fmt.Println("doing getInfo, current elapsed time is,", time.Since(start), retro, wii)

	retro_dirs, retro_time := getInfo("retro", retro)
	wii_dirs, wii_time := getInfo("wii", wii)
	psp_dirs, psp_time := getInfo("psp", psp)
	ps3_dirs, ps3_time := getInfo("ps3", ps3)
	n3ds_dirs, n3ds_time := getInfo("n3ds", n3ds)

	

	fmt.Println("doing postsaves, current elapsed time is,", time.Since(start))

	postSaves("Desktop", "retro", retro_dirs, retro_time)
	postSaves("Desktop", "wii", wii_dirs, wii_time)
	postSaves("Desktop", "psp", psp_dirs, psp_time)
	postSaves("Desktop", "ps3", ps3_dirs, ps3_time)
	postSaves("Desktop", "n3ds", n3ds_dirs, n3ds_time)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func postSaves(device string, console string, dirs []string, timemods []int64) {
	ctx := context.Background()
	client := graphql.NewClient("http://localhost:8080/query", http.DefaultClient)

	resp, err := createSaves(ctx, client, device, console, dirs, timemods)

	if err != nil {
		log.Println("json.Compact:", err)
	}

	fmt.Printf("Posted metadata", resp)

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(30)

	time.Sleep(time.Duration(n)*time.Second)
	
	*progressPointer += 0.06

	fmt.Printf("Progress is now", progress)
	
}

/*var (
    wg *sync.WaitGroup = &sync.WaitGroup{}
    q  chan string     = make(chan string, 1024)
)

func processFiles(ctx context.Context, files []string) {
    defer wg.Done()
    for {
        select {
        case path := <-q:
            if path == "" {
                return
            }
			//if filepath.Ext(path) == ".srm" || filepath.Ext(path) == ".dsv" || filepath.Ext(path) == ".ps2" || filepath.Ext(path) == ".gci" {
          		//files = append(files, path)
			fmt.Println(path)
			//}
            
        case <-ctx.Done():
            return
        }
    }
}

func walk(dir string) error {
    err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {

		
        if err != nil {
            return err
        }

        /*if d.IsDir() {
            return nil
        }
		
		fmt.Println(path)

        q <- path
        return nil
    })

	fmt.Println(err)

	return nil
}

func parseFiles(files map[int]string) []string {
	var dirs []string
	dirs_pointer := &dirs
	for key, value := range files {
		fmt.Println(key)
		fileExtPattern := regexp.MustCompile(`/\.[0-9a-z]+$/i`)
		fileExt := fileExtPattern.FindString(value)
		if fileExt == ".srm" || fileExt == ".dsv" || fileExt == ".ps2" || fileExt == ".gci" {
			*dirs_pointer = append(*dirs_pointer, value)
		}
	}
	return dirs
}

func test() {

	start := time.Now()

    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    fs := "/"

	var retro []string
	files_pointer := &retro

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go processFiles(ctx, *files_pointer)
    }
    retro, err := listFiles(fs)

	if err != nil {
			fmt.Println("idk")
		}

    close(q)

    wg.Wait()

	fmt.Println("doing getInfo, current elapsed time is,", time.Since(start), retro)

	//parseFiles(retro)

	//retro_dirs, retro_time := getInfo("retro", retro)

	//postSaves("Desktop", "retro", retro_dirs, retro_time)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}



func scaning_recursive(dir_path string) ([]string, []string) {

	folders := []string{}
	files := []string{}

	filepath.WalkDir(dir_path, func(path string, f fs.FileInfo, err error) error {

		f, err = fs.Stat(path)

		if err != nill {
			log.Fatal(err)
		}

		f_mode := f.FileMode()

		if f_mode.IsDir() {

			folders = append(folders, path)

		} else if f_mode.IsRegular() {

			files = append(files, path)

		}

	})
	
	return folders, files

}


func scan_recursive(dir_path string, ignore []string) ([]string, []string) {

	fmt.Println("Starting scan recursive")

	folders := []string{}
	files   := []string{}

	root := dir_path
	fileSystem := os.DirFS(root)

	// Scan
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {

		fmt.Println("Checking following file info", f)

		_continue := false

		// Loop : Ignore Files & Folders
		for _, i := range ignore {

			// If ignored path
			if strings.Index(path, i) != -1 {

				// Continue
				_continue = true
			}
		}

		if _continue == false {

			fmt.Println("This should display last")

			//f, err = os.Stat(path)

			fmt.Println("this does show up")

			// If no error
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("does this show up last?")


			// File & Folder Mode
			f_mode := f.Mode()

			// Is folder
			if f_mode.IsDir() {

				// Append to Folders Array
				folders = append(folders, path)

			// Is file
			} else if f_mode.IsRegular(){

				// Append to Files Array
				files = append(files, path)
			}
		}

		

		return nil
	})

	

	return folders, files
}



func searchFileSaves() {

	fmt.Println("Starting function")

	folders, files := scan_recursive("/home", []string{".srm", ".dsv", ".ps2", ".gci"})

	fmt.Println("did the scan recursive funciton")

	// Files
	for _, i := range files {
		fmt.Println(i)
	}

	// Folders
	for _, i := range folders {
		fmt.Println(i)
	}
}

func searchFolderWii3DS() {

	folders, files := scan_recursive("/home/tobs", []string{"/title"})

	// Files
	for _, i := range files {
		fmt.Println(i)
	}

	// Folders
	for _, i := range folders {
		fmt.Println(i)
	}
}

func moretesting() {
	root := "/"
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})
} */