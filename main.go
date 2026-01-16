package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/user"
)

type IpLocationInfo struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionName  string  `json:"region_name"`
	CityName    string  `json:"city_name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	ZipCode     string  `json:"zip_code"`
	TimeZone    string  `json:"time_zone"`
	Asn         string  `json:"asn"`
	As          string  `json:"as"`
	IsProxy     bool    `json:"is_proxy"`
	Message     string  `json:"message"`
}

//https://api.ip2location.io/

func main() {
    cuser,err:=user.Current()
	url := "https://api.ip2location.io/"
	resp, err := http.Get(url)
    if err!=nil{
        fmt.Println("Kann Nutzer nicht auslesen.")
    }

	if err != nil {
		fmt.Println("Sorry ein fehler ist aufgetreten: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fehler beim Lesen!")
		return
	}
	var info IpLocationInfo
	if err := json.Unmarshal(body, &info); err != nil {
		fmt.Println("Prasing fehler")
		return
	}
    fmt.Println("\033[37m")
	fmt.Println("Hi:",cuser.Name,"\nDeine aktuelle Ip Adresse ist: \033[31m", info.IP)
	fmt.Println("\033[37mDu lebst in ", info.ZipCode, info.CityName,"(", info.CountryCode,")")
    fmt.Println("Der Längengrad ist: ", info.Longitude," und der Breitengrad ist:", info.Latitude)
	fmt.Println("In der Region: ", info.RegionName)

}
