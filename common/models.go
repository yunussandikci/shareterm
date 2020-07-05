package common

type PasteCreateRequest struct {
	Content string `json:"content"`
}

type PasteCreateResponse struct {
	URL string `json:"url"`
}

type PasteReadResponse struct {
	Content string `json:"content"`
}

