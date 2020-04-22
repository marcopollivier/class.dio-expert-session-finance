package util

import (
	"testing"
)

func TestStringToDate(testing *testing.T) {
	var convertedTime = StringToTime("2019-02-12T10:00:00")

	if convertedTime.Year() != 2019 {
		testing.Errorf("Converter StringToDate is failed. Expected Year %v, got %v", 2019, convertedTime.Year())
	}

	if convertedTime.Month() != 2 {
		testing.Errorf("Converter StringToDate is failed. Expected Month %v, got %v", 2, convertedTime.Month())
	}

	if convertedTime.Hour() != 10 {
		testing.Errorf("Converter StringToDate is failed. Expected Hour %v, got %v", 10, convertedTime.Hour())
	}

}
