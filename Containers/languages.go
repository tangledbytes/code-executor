package languages

import (
	"log"
	"os"
	"path"
)

func getPath(filename string) string {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal("Failed to get the path of working directory")
	}

	return path.Join(wd, filename)
}

// Languages map returns the path for the container for corresponding code executor
var Languages = map[string]string{
	"cpp":    getPath("CPP"),
	"c":      getPath("C"),
	"python": getPath("python"),
}
