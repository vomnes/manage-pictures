package rename

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func GetLocationName(lat, long float64) Location {
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

func FormatLocation(data Location) string {
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
