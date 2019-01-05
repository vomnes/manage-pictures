package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/xor-gate/goexif2/exif"
	"github.com/xor-gate/goexif2/mknote"
)

type Address struct {
	Village     string `json:"village,omitempty"`
	City        string `json:"city,omitempty"`
	County      string `json:"county,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type Location struct {
	Name     string  `json:"name,omitempty"`
	FullName string  `json:"display_name,omitempty"`
	Address  Address `json:"address,omitempty"`
}

func getLocationName(lat, long float64) Location {
	response, err := http.Get(
		fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=%f&lon=%f",
			lat,
			long))

	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)
	var data Location
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error - Failed to get location - ", err)
	}
	return data
}

func formatLocation(data Location) string {
	var result string
	if data.Name != "" {
		result += data.Name
	}
	if result != "" {
		result += "-"
	}
	if data.Address.Village != "" && data.Address.Village != data.Name {
		result += data.Address.Village
	} else {
		result += data.Address.County
	}
	result += "-" + data.Address.CountryCode
	result = strings.Replace(result, " ", "", -1)
	return result
}

func listFolderFile(directoryName string) []string {
	var filesNames []string

	files, err := ioutil.ReadDir(directoryName)
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		filesNames = append(filesNames, f.Name())
	}
	return filesNames
}

func getNewName(fileName, path string) string {
	f, err := os.Open(path + fileName)
	if err != nil {
		fmt.Println("1", err)
		return ""
	}

	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		fmt.Println("2", err)
		return ""
	}

	var newName string
	// FORMAT DATE TIME
	time, error := x.DateTime()
	if error != nil {
		fmt.Println("3", err)
		return ""
	}
	formatTime := fmt.Sprintf(
		"%4d_%02d_%02d_%02d%02d",
		time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(),
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
		newName += formatLocation(getLocationName(lat, long)) + "&" + formatGPS
	}

	newName = strings.Replace(newName, "\"", "", -1)
	newName = strings.Replace(newName, " ", "_", -1)
	newName = strings.Replace(newName, "/", "+", -1)
	// newName = strings.Replace(newName, ".", "-", -1)
	return newName + ".jpg"
}

func main() {
	progArgs := os.Args
	if len(progArgs) <= 1 {
		fmt.Printf("./rename-pictures <directory>\n")
		return
	}
	directoryName := progArgs[1]
	if !strings.HasSuffix(directoryName, "/") {
		directoryName += "/"
	}

	files := listFolderFile(directoryName)
	var newName string
	for _, file := range files {
		file = strings.ToLower(file)
		if strings.HasSuffix(file, ".jpg") || strings.HasSuffix(file, ".jpeg") {
			newName = getNewName(file, directoryName)
			if newName != "" {
				os.Rename(directoryName+file, directoryName+newName)
			}
		}
	}
}

// yyyy-mm-dd-hhmm-lat-xxxx-long-xxxx-position-name-xxxx-model-OneplusA5010-iso-125-speed-1/243-f/1.7-white-balance-auto

// camModel, _ := x.Get(exif.Model)
// fmt.Print("Camera:")
// fmt.Println(camModel.StringVal())
//
// focal, _ := x.Get(exif.FocalLength)
// fmt.Println("focal", focal)
// exposure, _ := x.Get(exif.ExposureTime)
// fmt.Println("exposure time:", exposure)
// aperture, err := x.Get(exif.ApertureValue)
// if err != nil {
//   aperture1, aperture2, _ := aperture.Rat2(0)
//   fValue := math.Round(math.Pow(2.0, float64(aperture1)/float64(aperture2)/2.0)*100) / 100
//   fmt.Println("f/", fValue)
// } else {
//   fmt.Println(err)
// }
// numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
// fmt.Printf("%.4fmm\n", float64(numer)/float64(denom))
// iso, _ := x.Get(exif.ISOSpeedRatings)
// fmt.Println("iso:", iso)
// whitebalance, _ := x.Get(exif.WhiteBalance)
// fmt.Println("white balance:", whitebalance)

// model := fmt.Sprintf("&Model%s", camModel)
// newName += model
// isoValue := fmt.Sprintf("&ISO%s", iso)
// newName += isoValue
// speed := fmt.Sprintf("&Speed%s", exposure)
// newName += speed

// fVal := fmt.Sprintf("&F%.2f", fValue)
// newName += fVal
