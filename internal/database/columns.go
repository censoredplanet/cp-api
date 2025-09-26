package database

import (
	"reflect"
	"strings"

	"github.com/censoredplanet/cp-api/internal/entities"
)

var GQLToCHHyperquack map[string]string
var GQLToCHSatellite map[string]string
var GQLToCHDashboard map[string]string
var GQLToCHCenAlert map[string]string
var GQLToCHCenAlertEvents map[string]string

func init() {
	GQLToCHHyperquack = make(map[string]string, 32)
	t := reflect.TypeOf(entities.Hyperquack{})
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		jsonTag := strings.Split(sf.Tag.Get("json"), ",")[0]
		chTag := sf.Tag.Get("ch")
		if jsonTag != "" && chTag != "" {
			GQLToCHHyperquack[jsonTag] = chTag
		}
	}
	GQLToCHSatellite = make(map[string]string, 32)
	t = reflect.TypeOf(entities.Satellite{})
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		jsonTag := strings.Split(sf.Tag.Get("json"), ",")[0]
		chTag := sf.Tag.Get("ch")
		if jsonTag != "" && chTag != "" {
			GQLToCHSatellite[jsonTag] = chTag
		}
	}
	GQLToCHDashboard = make(map[string]string, 32)
	t = reflect.TypeOf(entities.Dashboard{})
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		jsonTag := strings.Split(sf.Tag.Get("json"), ",")[0]
		chTag := sf.Tag.Get("ch")
		if jsonTag != "" && chTag != "" {
			GQLToCHDashboard[jsonTag] = chTag
		}
	}
	GQLToCHCenAlert = make(map[string]string, 32)
	t = reflect.TypeOf(entities.CenAlertTimeSeries{})
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		jsonTag := strings.Split(sf.Tag.Get("json"), ",")[0]
		chTag := sf.Tag.Get("ch")
		if jsonTag != "" && chTag != "" {
			GQLToCHCenAlert[jsonTag] = chTag
		}
	}
	GQLToCHCenAlertEvents = make(map[string]string, 32)
	t = reflect.TypeOf(entities.CenAlertEvents{})
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		jsonTag := strings.Split(sf.Tag.Get("json"), ",")[0]
		chTag := sf.Tag.Get("ch")
		if jsonTag != "" && chTag != "" {
			GQLToCHCenAlertEvents[jsonTag] = chTag
		}
	}
}
