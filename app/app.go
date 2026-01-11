package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	//"strings"
	//"os"
	//"io"
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
func (a *App) Greet(name string) string {
	return fmt.Sprintf("ello %s, It's show time!", name)
}



func listFiles(dir string) []string {
    var files []string

    err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
       if !d.IsDir() {
			if filepath.Ext(path) == ".srm" || filepath.Ext(path) == ".dsv" || filepath.Ext(path) == ".ps2" || filepath.Ext(path) == ".gci" {
          		files = append(files, path)
       }
	   }  
       return nil
    })
    if err != nil {
       log.Fatal(err)
    }

    return files
}

func listFolders(dir string, console string) []string {
	var folders []string

    err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
       if d.IsDir() {
			if filepath.Base(path) == "title" && console == "wii" {
				folders = append(folders, path)
			} else if filepath.Base(path) == "SAVEDATA" && console == "ps" {
				folders = append(folders, path)
			} else if filepath.Base(path) == "savedata" && console == "ps" {
				folders = append(folders, path)
			} else if filepath.Base(path) == "save" && console == "wiiu" {
				folders = append(folders, path)
			}
	   }  
       return nil
    })
    if err != nil {
       log.Fatal(err)
    }

	parsedFolders := searchFolders(folders)

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

func saveSearch() {

	dir := "/"

	retro := listFiles(dir)
	wii := listFolders(dir, "wii")
	ps := listFolders(dir, "ps")
	wiiu  := listFolders(dir, "wiiu")

    for _, save := range retro {
       fmt.Println("Found a retro folder! :", save)
    }

	for _, save := range wii {
       fmt.Println("Found a Wii folder! :", save)
    }

	for _, save := range ps {
       fmt.Println("Found a PS3 or PSP folder! :", save)
    }

	for _, save := range wiiu {
       fmt.Println("Found a Wii U or 3DS folder! :", save)
    }
}



/*func scaning_recursive(dir_path string) ([]string, []string) {

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

/*func moretesting() {
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