package main

import (
	"fmt"

	"saladpuk.com/promptpay/promptpay"
)

func main() {
	// สร้างการจ่ายเงินด้วย Tag 29 (Credit Transfer)
	creditPayload := &promptpay.Payload{
		Amount: 50.00,
	}
	creditInfo := &promptpay.CreditTransfer{
		MobileNumber:        "0914185401",
		CustomerPresentedQR: false,
	}
	qrCode1 := promptpay.CreateCreditTransferQrCode(creditPayload, creditInfo)
	fmt.Println("CreditTransfer: " + qrCode1)

	// สร้างการจ่ายเงินด้วย Tag 30 (Bill Payment)
	billPayload := &promptpay.Payload{
		Amount: 50.00,
	}
	billInfo := &promptpay.BillPayment{
		BillerId:              "01150105523009350",
		Suffix:                "08",
		Reference1:            "Test",
		CrossBorderMerchantQR: true,
	}
	qrCode2 := promptpay.CreateBillPaymentQrCode(billPayload, billInfo)
	fmt.Println("BillPayment: " + qrCode2)
}
