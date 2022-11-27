package promptpay_test

import (
	"testing"

	"saladpuk.com/promptpay/promptpay"
)

func TestCreateCreditTransfer(t *testing.T) {
	pl := &promptpay.Payload{
		Amount: 50.00,
	}
	tx := &promptpay.CreditTransfer{
		MobileNumber:        "0914185401",
		CustomerPresentedQR: false,
	}
	actual := promptpay.CreateCreditTransferQrCode(pl, tx)
	var expected = "0002010102125802TH5303764540550.0029370016A00000067701011101130066914185401630461EF"
	validate(actual, expected, t)
}

func TestCreateCreditTransferSegmentWithMobileNumberAndCustomerPresented(t *testing.T) {
	tf := &promptpay.CreditTransfer{
		MobileNumber:        "0914185401",
		CustomerPresentedQR: true,
	}
	actual := promptpay.CreateCreditTransferSegment(tf)
	var expected = "29370016A00000067701011401130066914185401"
	validate(actual, expected, t)
}
func TestCreateCreditTransferSegmentWithMobileNumberAndMerchantPresented(t *testing.T) {
	tf := &promptpay.CreditTransfer{
		MobileNumber:        "0914185401",
		CustomerPresentedQR: false,
	}
	actual := promptpay.CreateCreditTransferSegment(tf)
	var expected = "29370016A00000067701011101130066914185401"
	validate(actual, expected, t)
}
func TestCreateCreditTransferSegmentWithMobileAndCountryCode(t *testing.T) {
	tf := &promptpay.CreditTransfer{
		MobileNumber:        "66914185401",
		CustomerPresentedQR: true,
	}
	actual := promptpay.CreateCreditTransferSegment(tf)
	var expected = "29370016A00000067701011401130066914185401"
	validate(actual, expected, t)
}
func TestCreateCreditTransferSegmentWithNationalId(t *testing.T) {
	tf := &promptpay.CreditTransfer{
		NationalIdOrTaxId:   "1234567890123",
		CustomerPresentedQR: true,
	}
	actual := promptpay.CreateCreditTransferSegment(tf)
	var expected = "29370016A00000067701011402131234567890123"
	validate(actual, expected, t)
}
func TestCreateCreditTransferSegmentWithEWallet(t *testing.T) {
	tf := &promptpay.CreditTransfer{
		EWalletId:           "123456789012345",
		CustomerPresentedQR: true,
	}
	actual := promptpay.CreateCreditTransferSegment(tf)
	var expected = "29390016A0000006770101140315123456789012345"
	validate(actual, expected, t)
}
func TestCreateCreditTransferSegmentWithBankAccount(t *testing.T) {
	tf := &promptpay.CreditTransfer{
		BankAccount:         "12345678901234567890",
		CustomerPresentedQR: true,
	}
	actual := promptpay.CreateCreditTransferSegment(tf)
	var expected = "29440016A000000677010114042012345678901234567890"
	validate(actual, expected, t)
}
func TestCreateCreditTransferSegmentWithOTA(t *testing.T) {
	tf := &promptpay.CreditTransfer{
		OTA:                 "1234567890",
		CustomerPresentedQR: true,
	}
	actual := promptpay.CreateCreditTransferSegment(tf)
	var expected = "29340016A00000067701011405101234567890"
	validate(actual, expected, t)
}

func validate(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Logf("Actual: " + actual)
		t.Logf("Expected: " + expected)
		t.Error()
	}
}
