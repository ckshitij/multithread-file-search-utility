package config

import (
	"flag"
	"log"
	"strings"

	"github.com/vharitonsky/iniflags"
)

var (
	directory = flag.String("directory", "", "Directory name under which file need to find.")
	filename  = flag.String("filename", "", "Name of the file which need to search filename is case-sensitive")
)

func init() {
	// Parse flags from cli arguments/ini file
	iniflags.Parse()

	*directory = strings.TrimSpace(*directory)
	*filename = strings.TrimSpace(*filename)

	if *directory == "" {
		log.Fatalf("error: directory argument is missing\n")
	}
	if *filename == "" {
		log.Fatalf("error: filename argument is missing\n")
	}
}

func GetDirectory() string {
	return *directory
}
func GetFileName() string {
	return *filename
}
