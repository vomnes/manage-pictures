package main

import (
	"fmt"
	"os"

	"./rename"
	"github.com/kr/pretty"
)

func main() {
	progArgs := os.Args
	if len(progArgs) <= 1 {
		fmt.Printf("./rename-pictures <directory>\n")
		return
	}
	directoryName := progArgs[1]
	loading := rename.StatFolders(directoryName)
	pretty.Print(loading)
	go rename.Run(directoryName, &loading)
	for loading.Done != loading.Total {
		pretty.Println(loading)
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
