package main

import (
	"fmt"
	"time"
)

func fullScan(dir string) {

	start := time.Now()

	fmt.Println("doing consoleSearch, current elapsed time is,", time.Since(start))

	consoleFolders := consoleSearch(dir)

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

	//fmt.Println("doing getInfo, current elapsed time is,", time.Since(start), retro, wii)

	retro_dirs, retro_time := getInfo("retro", retro)
	wii_dirs, wii_time := getInfo("wii", wii)
	psp_dirs, psp_time := getInfo("psp", psp)
	ps3_dirs, ps3_time := getInfo("ps3", ps3)
	fmt.Println("the n3ds variable is:", n3ds)
	n3ds_dirs, n3ds_time := getInfo("n3ds", n3ds)

	//fmt.Println("doing postsaves, current elapsed time is,", time.Since(start))

	flushSaves("Desktop")

	postSaves("Desktop", "retro", retro_dirs, retro_time)

	postSaves("Desktop", "wii", wii_dirs, wii_time)

	postSaves("Desktop", "psp", psp_dirs, psp_time)

	postSaves("Desktop", "ps3", ps3_dirs, ps3_time)

	fmt.Println("the n3ds directory list is:", n3ds_dirs)
	postSaves("Desktop", "n3ds", n3ds_dirs, n3ds_time)

	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
