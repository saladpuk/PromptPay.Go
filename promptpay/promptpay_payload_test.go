package promptpay_test

import (
	"testing"

	"saladpuk.com/promptpay/promptpay"
)

func TestPayloadWithStaticMode(t *testing.T) {
	model := &promptpay.Payload{
		FormatIndicator: "01",
		Reusable:        false,
		CountryCode:     "TH",
		CurrencyCode:    "764",
		Amount:          50.01,
	}
	actual := promptpay.CreatePayloadSegment(model)
	expected := "0002010102125802TH5303764540550.01"
	validate(actual, expected, t)
}
func TestPayloadWithDynamicMode(t *testing.T) {
	model := &promptpay.Payload{
		FormatIndicator: "01",
		Reusable:        true,
		CountryCode:     "TH",
		CurrencyCode:    "764",
		Amount:          50.01,
	}
	actual := promptpay.CreatePayloadSegment(model)
	expected := "0002010102115802TH5303764540550.01"
	validate(actual, expected, t)
}
func TestPayloadWithDefaultValue(t *testing.T) {
	model := &promptpay.Payload{
		Amount: 50.01,
	}
	actual := promptpay.CreatePayloadSegment(model)
	expected := "0002010102125802TH5303764540550.01"
	validate(actual, expected, t)
}
func TestPayloadWithoutAmount(t *testing.T) {
	model := &promptpay.Payload{}
	actual := promptpay.CreatePayloadSegment(model)
	expected := "0002010102125802TH5303764"
	validate(actual, expected, t)
}
