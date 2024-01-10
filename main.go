package main

import (
	"fmt"
	"time"

	"github.com/ckshitij/multithread-file-search-utility/config"
	file_search "github.com/ckshitij/multithread-file-search-utility/search_file"
)

func GetTimeElapse(callbackFunc func()) {
	startTime := time.Now()
	callbackFunc()
	elapsed := time.Since(startTime)
	fmt.Printf("Total time taken to run the function %d ms\n", elapsed.Milliseconds())
}

func executeMultithreadedSearch() {
	fs := file_search.NewFileSearchUtility()
	fs.Add(1)
	go fs.SearchFile(config.GetDirectory(), config.GetFileName())
	fs.Wait()
	fmt.Printf("executed the multithreaded, found total %d matched files", len(fs.MatchedFiles()))
}

func executeSyncSearch() {
	fs := file_search.NewFileSearchUtility()
	fs.SyncSearchFile(config.GetDirectory(), config.GetFileName())
	fmt.Printf("executed the sync search, found total %d matched files", len(fs.MatchedFiles()))
}

func main() {
	// Compare the time difference between sync and multithread search calls.
	GetTimeElapse(executeMultithreadedSearch)
	GetTimeElapse(executeSyncSearch)
}
