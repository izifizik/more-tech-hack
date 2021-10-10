package model

type Model struct {
	Id int
	Urn       string
	Struct    Json
	Name      string
	IsDataset bool
}

type Json map[string]interface{}

type ModelsGet struct {
	ModelId        int
	ModelName      string
	ModelUrn       string
	ModelIsDataset bool
}