package exporter

import (
	"fmt"
	"strconv"

	"github.com/bean1310/remo-exporter/types"
)

type AirConditionerInfo struct {
	TargetTemp   int
	Mode         string
	AirVolume    string
	AirDirection string
	PowerState   string
}

func getAirConditioners(apps []*types.Appliance) []*(Appliance) {
	airConditioners := make([]*types.Appliance, 0)
	for _, app := range apps {
		if app.Type == "AC" {
			airConditioners = append(smartMeters, app)
		}
	}
	return airConditioners
}

func airConditionerInfo(ac *types.Appliance) (*AirConditionerInfo, error) {
	if ac.Settings == nil {
		return nil, fmt.Errorf("'%s' does not have settings field", ac.Device.Name)
	}

	var info AirConditionerInfo
	var err error

	info.TargetTemp, err = strconv.Atoi(ac.Settings.Temp)
	if err != nil {
		return nil, err
	}

	switch ac.Settings.Mode {
	case "cool", "warm", "dry", "blow", "auto":
		info.Mode = ac.Settings.Mode
	default:
		info.Mode = "unknown"
	}

	switch ac.Settings.Vol {
	case "auto", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10":
		info.AirVolume = ac.Settings.Vol
	case "":
		info.AirVolume = "auto"
	default:
		info.AirVolume = "unknown"
	}

	if ac.Settings.Dir == "" {
		info.AirVolume = "auto"
	} else {
		info.AirVolume = ac.Settings.Dir
	}

	if ac.Settings.Button == "power-off" {
		info.PowerState = "PowerOff"
	} else {
		info.PowerState = "PowerOn"
	}

	return &info, nil
}
