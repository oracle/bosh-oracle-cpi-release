package test

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}, msg string) {
	if expected == actual {
		return
	}
	if len(msg) <= 0 {
		msg = fmt.Sprintf("Expected %v. Actual %v", expected, actual)
	}
	t.Error(msg)
	t.Fail()
}

func assertNotNil(t *testing.T, value interface{}, msg string) {
	if value == nil {
		if len(msg) <= 0 {
			msg = "Null"
		}
		t.Error(msg)
		t.Fail()
	}
}

func assertIsNil(t *testing.T, value interface{}, msg string) {
	if value != nil {
		if len(msg) <= 0 {
			msg = "Unexpected non null reference"
		}
		t.Error(msg)
		t.Fail()
	}
}
