# PromptPay.Go
สร้างการจ่ายเงินด้วย PromptPay ในภาษา Golang

## สร้างการจ่ายเงินด้วย Tag 29 (Credit Transfer)
```go
creditPayload := &promptpay.Payload{
	Amount: 50.00,
}
creditInfo := &promptpay.CreditTransfer{
	MobileNumber:        "0914185401",
	CustomerPresentedQR: false,
}
qrCode1 := promptpay.CreateCreditTransferQrCode(creditPayload, creditInfo)
fmt.Println("CreditTransfer: " + qrCode1)
```
> **Output**  
CreditTransfer: 0002010102125802TH5303764540550.0029370016A00000067701011101130066914185401630461EF

## สร้างการจ่ายเงินด้วย Tag 30 (Bill Payment)
```go
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
```
> **Output**  
BillPayment: 0002010102125802TH5303764540550.0030510016A000000677012006011901150105523009350080204TestD0B8
