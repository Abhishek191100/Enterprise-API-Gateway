package logger

import "testing"

func TestLog(t *testing.T) {
	Log("Coming through","INFO")
	Log("Coming through","DEBUG")
	Log("Coming through","WARN")
	Log("Coming through","ERROR")
	Log("Coming through","INVALID")
}