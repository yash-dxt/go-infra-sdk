package sqs

// shared btw enrich & pdf2json service.
type Pdf2JsonQueueMsg struct {
	Type       string `json:"type"`   // required.
	RefId      string `json:"ref_id"` // required.
	TpResultId int    `json:"result_id"`

	PdfUrl      string `json:"pdf_url"`
	JsonUrl     string `json:"json_url"`
	UploadToUrl string `json:"upload_to_url"`
}
