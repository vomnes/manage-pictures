package rename

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ListFolderFile(directoryName string, onlyJPG bool) []string {
	var filesNames []string

	files, err := ioutil.ReadDir(directoryName)
	if err != nil {
		fmt.Println(err)
	}
	var file string
	for _, f := range files {
		file = strings.ToLower(f.Name())
		if onlyJPG && (strings.HasSuffix(file, ".jpg") || strings.HasSuffix(file, ".jpeg")) {
			filesNames = append(filesNames, file)
		} else if !onlyJPG {
			filesNames = append(filesNames, file)
		}
	}
	return filesNames
}
