package rename

import (
	"os"
	"strings"
)

type Loading struct {
	Done  int
	Total int
}

func StatFolders(directoryPath string) Loading {
	files := ListFolderFile(directoryPath, true)
	return Loading{
		Total: len(files),
		Done:  0,
	}
}

func Run(directoryPath string, loading *Loading) {
	if !strings.HasSuffix(directoryPath, "/") {
		directoryPath += "/"
	}

	files := ListFolderFile(directoryPath, true)
	var newName string
	(*loading).Total = int(len(files))
	(*loading).Done = 0
	for _, file := range files {
		newName = GetNewName(file, directoryPath)
		if newName != "" {
			os.Rename(directoryPath+file, directoryPath+newName)
		}
		(*loading).Done++
	}
}
