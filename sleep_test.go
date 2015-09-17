package timesleep

import "testing"

func TestSleepWithValidDuration(t *testing.T) {
	actualResult := Sleep(2000)
	if actualResult != true {
		t.Error("Test Passed!")
	}
}

func TestSleepWithInValidDuration(t *testing.T) {
	actualResult := Sleep(-2000)
	if actualResult != false {
		t.Error("Test Passed!")
	}
}