package crc_test

import (
	"testing"

	"saladpuk.com/promptpay/crc"
)

func TestCRC(t *testing.T) {
	actual := crc.GenCRC("00020101021229370016A000000677010111011300669141854015303764540550.005802TH6304")
	expected := "01F8"
	if actual != expected {
		t.Logf("Actual: " + actual)
		t.Logf("Expected: " + expected)
		t.Error()
	}
}
