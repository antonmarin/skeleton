package enum

import (
	"fmt"
	"testing"
)

func TestOsType_ShouldConvertToString(t *testing.T) {
	var osName string

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
