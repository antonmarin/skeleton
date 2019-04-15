package osFamily

import (
	"fmt"
	"testing"
)

func Test_ShouldConvertToString(t *testing.T) {
	var osName string

	osName = fmt.Sprint(new(OsFamily))
	if osName != "Undefined" {
		t.Errorf("Undefined name should be Undefined. Got %s", osName)
	}
	osName = fmt.Sprint(Linux)
	if osName != "Linux" {
		t.Errorf("Linux name should be Linux. Got %s", osName)
	}
	osName = fmt.Sprint(MacOs)
	if osName != "macOS" {
		t.Errorf("MacOs name should be macOS. Got %s", osName)
	}
	osName = fmt.Sprint(Windows)
	if osName != "Windows" {
		t.Errorf("Windows name should be Windows. Got %s", osName)
	}
}

func Test_ShouldCreateFromGOOSString(t *testing.T) {
	var osType OsFamily

	osType = NewFromGOOS("darwin")
	if osType != MacOs {
		t.Errorf("Should be macOS from darwin. Got %s", osType)
	}
	osType = NewFromGOOS("linux")
	if osType != Linux {
		t.Errorf("Should be Linux from linux. Got %s", osType)
	}
	osType = NewFromGOOS("windows")
	if osType != Windows {
		t.Errorf("Should be Windows from windows. Got %s", osType)
	}
	osType = NewFromGOOS("unknown_os")
	if osType != Undefined {
		t.Errorf("Should be Undefined from unknown_os. Got %s", osType)
	}
}
