package promptpay_test

import (
	"testing"

	"saladpuk.com/promptpay/promptpay"
)

func TestCreateBillPayment(t *testing.T) {
	pl := &promptpay.Payload{
		Amount: 50.00,
	}
	tx := &promptpay.BillPayment{
		BillerId:              "01150105523009350",
		Suffix:                "08",
		Reference1:            "Test",
		CrossBorderMerchantQR: true,
	}
	actual := promptpay.CreateBillPaymentQrCode(pl, tx)
	var expected = "0002010102125802TH5303764540550.0030510016A000000677012006011901150105523009350080204Test63041B9A"
	validate(actual, expected, t)
}

func TestCreateBillPaymentSegmentWithCrossBorderMerchant(t *testing.T) {
	tf := &promptpay.BillPayment{
		BillerId:              "01150105523009350",
		Suffix:                "08",
		Reference1:            "Test",
		CrossBorderMerchantQR: true,
	}
	actual := promptpay.CreateBillPaymentSegment(tf)
	var expected = "30510016A000000677012006011901150105523009350080204Test"
	validate(actual, expected, t)
}
func TestCreateBillPaymentSegmentWithDomesticMerchant(t *testing.T) {
	tf := &promptpay.BillPayment{
		BillerId:              "01150105523009350",
		Suffix:                "08",
		Reference1:            "Test",
		CrossBorderMerchantQR: false,
	}
	actual := promptpay.CreateBillPaymentSegment(tf)
	var expected = "30510016A000000677010112011901150105523009350080204Test"
	validate(actual, expected, t)
}
func TestCreateBillPaymentSegmentWithRef2(t *testing.T) {
	tf := &promptpay.BillPayment{
		BillerId:              "01150105523009350",
		Suffix:                "08",
		Reference1:            "Test",
		Reference2:            "Other",
		CrossBorderMerchantQR: false,
	}
	actual := promptpay.CreateBillPaymentSegment(tf)
	var expected = "30600016A000000677010112011901150105523009350080204Test0305Other"
	validate(actual, expected, t)
}
