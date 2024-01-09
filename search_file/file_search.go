package file_search

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type fileSearchUtility struct {
	matchedFiles []string
	sync.Mutex
	sync.WaitGroup
}

// Create the FileSearchUtility struct
func NewFileSearchUtility() fileSearchUtility {
	return fileSearchUtility{
		matchedFiles: make([]string, 0),
	}
}

func (fs *fileSearchUtility) MatchedFiles() []string {
	return fs.matchedFiles
}

// Search the filename in the given directory.
// Note Filename is case-sensitive
func (fs *fileSearchUtility) SearchFile(dirName, filename string) {
	directories, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Printf("Failed to get read directory (%s) with error [%s] \n", dirName, err.Error())
	}
	for _, path := range directories {
		currentPath := filepath.Join(dirName, path.Name())
		if path.IsDir() {
			fs.Add(1)
			go fs.SearchFile(currentPath, filename)
		} else if path.Name() == filename {
			fs.addFileInMatchedPath(currentPath)
		}
	}
	defer fs.Done()
}

// Add matched file in the matchedFiles array
func (fs *fileSearchUtility) addFileInMatchedPath(filename string) {
	fs.Lock()
	fs.matchedFiles = append(fs.matchedFiles, filename)
	fs.Unlock()
}

// Print all the matched path for the given fileName
func (fs *fileSearchUtility) PrintMatchedPaths() {
	fmt.Println("Found the below file paths by given filename")
	for _, path := range fs.matchedFiles {
		fmt.Println(path)
	}
}
