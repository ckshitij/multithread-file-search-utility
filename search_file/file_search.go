package file_search

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type FileSearchUtility struct {
	//created set for getting the unique file path
	matchedFiles map[string]struct{}
	sync.Mutex
	sync.WaitGroup
}

// Create the FileSearchUtility struct
func NewFileSearchUtility() FileSearchUtility {
	return FileSearchUtility{
		matchedFiles: make(map[string]struct{}, 0),
	}
}

func (fs *FileSearchUtility) MatchedFiles() map[string]struct{} {
	return fs.matchedFiles
}

// Search the filename in the given directory.
// Note Filename is case-sensitive
func (fs *FileSearchUtility) SearchFile(dirName, filename string) {
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

// Search the filename in the given directory.
// Note Filename is case-sensitive
func (fs *FileSearchUtility) SyncSearchFile(dirName, filename string) {
	directories, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Printf("Failed to get read directory (%s) with error [%s] \n", dirName, err.Error())
	}
	for _, path := range directories {
		currentPath := filepath.Join(dirName, path.Name())
		if path.IsDir() {
			fs.SyncSearchFile(currentPath, filename)
		} else if path.Name() == filename {
			fs.matchedFiles[currentPath] = struct{}{}
		}
	}
}

// Add matched file in the matchedFiles array
func (fs *FileSearchUtility) addFileInMatchedPath(filename string) {
	fs.Lock()
	fs.matchedFiles[filename] = struct{}{}
	fs.Unlock()
}

// Print all the matched path for the given fileName
func (fs *FileSearchUtility) PrintMatchedPaths() {
	fmt.Println("Found the below file paths by given filename")
	for path, _ := range fs.matchedFiles {
		fmt.Println(path)
	}
}
