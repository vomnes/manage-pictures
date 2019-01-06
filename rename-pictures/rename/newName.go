package rename

import (
	"fmt"
	"os"
	"strings"

	"github.com/xor-gate/goexif2/exif"
	"github.com/xor-gate/goexif2/mknote"
)

func GetNewName(fileName, path string) string {
	f, err := os.Open(path + fileName)
	if err != nil {
		fmt.Println("Error - open file:", err)
		return ""
	}

	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		fmt.Println("Error - decode exif file:", err)
		return ""
	}

	var newName string
	// FORMAT DATE TIME
	time, error := x.DateTime()
	if error != nil {
		fmt.Println("Error - DateTime:", err)
		return ""
	}
	formatTime := fmt.Sprintf(
		"%4d_%02d_%02d_%02d%02d%02d",
		time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second(),
	)
	newName += formatTime

	// FORMAT LAT-LONG
	lat, long, _ := x.LatLong()
	if lat != 0 && long != 0 {
		formatGPS := fmt.Sprintf(
			"%f&%f",
			lat,
			long,
		)
		newName += FormatLocation(GetLocationName(lat, long)) + "&" + formatGPS
	}

	newName = strings.Replace(newName, "\"", "", -1)
	newName = strings.Replace(newName, " ", "_", -1)
	newName = strings.Replace(newName, "/", "+", -1)
	// newName = strings.Replace(newName, ".", "-", -1)
	return newName + ".jpg"
}
