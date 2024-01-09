package main

import (
	"fmt"
	"time"

	"github.com/ckshitij/multithread-file-search-utility/config"
	file_search "github.com/ckshitij/multithread-file-search-utility/search_file"
)

func main() {
	fs := file_search.NewFileSearchUtility()

	startTime := time.Now()
	fs.Add(1)
	fs.SearchFile(config.GetDirectory(), config.GetFileName())
	fs.Wait()
	totalTime := time.Since(startTime)
	fs.PrintMatchedPaths()
	fmt.Printf("Total matched files found %d and it took %d ms to search\n", len(fs.MatchedFiles()), totalTime.Milliseconds())
}
