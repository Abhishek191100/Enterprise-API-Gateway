package logger

import "testing"

func TestLog(t *testing.T) {
	log("Coming through","INFO")
	log("Coming through","DEBUG")
	log("Coming through","WARN")
	log("Coming through","ERROR")
	log("Coming through","INVALID")
}