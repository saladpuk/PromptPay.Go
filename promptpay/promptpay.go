package promptpay

import (
	"fmt"
	"math"

	"saladpuk.com/promptpay/crc"
)

type Payload struct {
	Reusable        bool
	FormatIndicator string
	CountryCode     string
	CurrencyCode    string
	Amount          float64
}

type CreditTransfer struct {
	Payload
	CustomerPresentedQR bool
	MobileNumber        string
	NationalIdOrTaxId   string
	EWalletId           string
	BankAccount         string
	OTA                 string
}

type BillPayment struct {
	Payload
	CrossBorderMerchantQR bool
	BillerId              string
	Suffix                string
	Reference1            string
	Reference2            string
}

func CreateCreditTransferQrCode(payload *Payload, info *CreditTransfer) string {
	payloadSeg := CreatePayloadSegment(payload)
	txSeg := CreateCreditTransferSegment(info)
	data := payloadSeg + txSeg + "6304"
	crc := crc.GenCRC(data)
	return data + crc
}

func CreateBillPaymentQrCode(payload *Payload, info *BillPayment) string {
	payloadSeg := CreatePayloadSegment(payload)
	txSeg := CreateBillPaymentSegment(info)
	crc := crc.GenCRC(payloadSeg + txSeg)
	return payloadSeg + txSeg + crc
}

func CreatePayloadSegment(info *Payload) string {
	if info.FormatIndicator == "" {
		info.FormatIndicator = "01"
	}
	if info.CountryCode == "" {
		info.CountryCode = "TH"
	}
	if info.CurrencyCode == "" {
		info.CurrencyCode = "764"
	}

	indicatorSeg := createSegment("00", info.FormatIndicator)
	pointOfInitSeg := createPointOfInitiationMethodSegment(info.Reusable)
	countryCodeSeg := createSegment("58", info.CountryCode)
	currencyCodeSeg := createSegment("53", info.CurrencyCode)
	txAmountSeg := createNumberSegment("54", info.Amount)
	return indicatorSeg + pointOfInitSeg + countryCodeSeg + currencyCodeSeg + txAmountSeg
}

func CreateCreditTransferSegment(info *CreditTransfer) string {
	aidSeg := createCreditTransferAidSegment(info.CustomerPresentedQR)
	mobileSeg := createMobileSegment(info.MobileNumber)
	nationalSeg := createSegment("02", info.NationalIdOrTaxId)
	ewalletSeg := createSegment("03", info.EWalletId)
	bankAccSeg := createSegment("04", info.BankAccount)
	otaSeg := createSegment("05", info.OTA)
	return createSegment("29", aidSeg+mobileSeg+nationalSeg+ewalletSeg+bankAccSeg+otaSeg)
}

func CreateBillPaymentSegment(info *BillPayment) string {
	aidSeg := createBillPaymentAidSegment(info.CrossBorderMerchantQR)
	billerIdSeg := createSegment("01", info.BillerId+info.Suffix)
	ref1Seg := createSegment("02", info.Reference1)
	ref2Seg := createSegment("03", info.Reference2)
	return createSegment("30", aidSeg+billerIdSeg+ref1Seg+ref2Seg)
}

func createPointOfInitiationMethodSegment(isReusable bool) string {
	var data string
	if isReusable {
		data = "11"
	} else {
		data = "12"
	}
	return createSegment("01", data)
}

func createBillPaymentAidSegment(isCrossBorderMerchantQR bool) string {
	var data string
	if isCrossBorderMerchantQR {
		data = "A000000677012006"
	} else {
		data = "A000000677010112"
	}
	return createSegment("00", data)
}
func createCreditTransferAidSegment(isCustomerPresented bool) string {
	var data string
	if isCustomerPresented {
		data = "A000000677010114"
	} else {
		data = "A000000677010111"
	}
	return createSegment("00", data)
}
func createMobileSegment(value string) string {
	if value == "" {
		return ""
	}
	if value[0:1] == "0" {
		value = value[1:]
	} else if value[0:1] == "6" {
		value = value[2:]
	}
	data := "0066" + value
	return createSegment("01", data)
}
func createNumberSegment(id string, value float64) string {
	if value == 0 {
		return ""
	}
	return createSegment(id, fmt.Sprintf("%.2f", math.Abs(value)))
}
func createSegment(id string, value string) string {
	if value == "" {
		return ""
	}
	return id + fmt.Sprintf("%02d", len(value)) + value
}
