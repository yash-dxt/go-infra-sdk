package sqs

// shared btw enrich & pdf2json service.
type Pdf2JsonQueueMsg struct {
	Type       string `json:"type"`   // required.
	RefId      string `json:"ref_id"` // required.
	TpResultId int    `json:"result_id"`

	PdfUrl      string      `json:"pdf_url"`
	JsonUrl     string      `json:"json_url"`
	UploadToUrl string      `json:"upload_to_url"`
	FraudChecks FraudChecks `json:"fraud_checks"`

	Retry int `json:"retry"`
}

type FraudChecks struct {
	Pass            bool             `json:"pass" validate:"required"`
	FraudIndicators []FraudIndicator `json:"fraudIndicators" validate:"required,dive"`
}

type FraudIndicator struct {
	IndicatorName        string `json:"indicatorName" validate:"required"`
	IndicatorDescription string `json:"indicatorDescription" validate:"required"`
	Comment              string `json:"comment"`
	TransactionId        int    `json:"transactionId"`
}
