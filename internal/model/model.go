package model

type Model struct {
	Id        int    `json:"id"`
	Urn       string `json:"urn"`
	Struct    Json   `json:"struct"`
	Name      string `json:"name"`
	IsDataset bool   `json:"isDataset"`
}

type Json map[string]interface{}

type ModelsGet struct {
	ModelId        int    `json:"id"`
	ModelName      string `json:"name"`
	ModelUrn       string `json:"urn"`
	ModelIsDataset bool   `json:"isDataset"`
}
