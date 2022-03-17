package venstarapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// ThermostatInfo contains the full info values from the thermostat api
type ThermostatInfo struct {
	Name            string `json:"name"`
	Mode            int    `json:"mode"`
	State           int    `json:"state"`
	Fan             int    `json:"fan"`
	FanState        int    `json:"fanstate"`
	TempUnits       int    `json:"tempunits"`
	Schedule        int    `json:"schedule"`
	SchedulePart    int    `json:"schedulepart"`
	Away            int    `json:"away"`
	Holiday         int    `json:"holiday"`
	Override        int    `json:"override"`
	OverrideTime    int    `json:"overridetime"`
	ForceUnoccupied int    `json:"forceunocc"`
	SpaceTemp       int    `json:"spacetemp"`
	HeatTemp        int    `json:"heattemp"`
	CoolTemp        int    `json:"cooltemp"`
	CoolTempMin     int    `json:"cooltempmin"`
	CoolTempMax     int    `json:"cooltempmax"`
	HeatTempMin     int    `json:"heattempmin"`
	HeatTempMax     int    `json:"heattempmax"`
	SetPointDelta   int    `json:"setpointdelta"`
	Humidity        int    `json:"hum"`
	AvaliableModes  int    `json:"avaliablemodes"`
}

// GetThermostatInfo queries the current thermostat info, and then returns it as a ThermostatInfo struct
func GetThermostatInfo(thermostatIP string) (ThermostatInfo, error) {
	currentInfo := ThermostatInfo{}
	req, err := http.Get(fmt.Sprintf("http://%s/query/info", thermostatIP))
	if err != nil {
		log.Println(err)
		return currentInfo, err
	}
	defer req.Body.Close()
	by, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return currentInfo, err
	}
	json.Unmarshal(by, &currentInfo)
	return currentInfo, nil

}

// SetThermostatMode sets the thermostat mode
func SetThermostatMode(ipaddress string, mode int, currentInfo ThermostatInfo) bool {
	data := url.Values{}
	data.Set("heattemp", fmt.Sprintf("%d", currentInfo.HeatTemp))
	data.Set("cooltemp", fmt.Sprintf("%d", currentInfo.CoolTemp))
	data.Set("mode", fmt.Sprintf("%d", mode))
	_, reqErr := http.PostForm(fmt.Sprintf("http://%s/control", ipaddress), url.Values(data))
	return reqErr == nil
}

// SetCoolTemp sets the "cool to" temp on the thermostat
func SetCoolTemp(ipaddress string, coolTemp int, currentInfo ThermostatInfo) bool {
	data := url.Values{}
	data.Set("heattemp", fmt.Sprintf("%d", currentInfo.HeatTemp))
	data.Set("cooltemp", fmt.Sprintf("%d", coolTemp))
	_, reqErr := http.PostForm(fmt.Sprintf("http://%s/control", ipaddress), url.Values(data))
	return reqErr == nil
}

// SetFanMode sets whether the fan is "auto" or "on"
func SetFanMode(ipaddress string, fanMode int, currentInfo ThermostatInfo) bool {
	data := url.Values{}
	data.Set("heattemp", fmt.Sprintf("%d", currentInfo.HeatTemp))
	data.Set("cooltemp", fmt.Sprintf("%d", currentInfo.CoolTemp))
	data.Set("fan", fmt.Sprintf("%d", fanMode))
	_, reqErr := http.PostForm(fmt.Sprintf("http://%s/control", ipaddress), url.Values(data))
	return reqErr == nil
}
