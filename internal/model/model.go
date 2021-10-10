package model

type Model struct {
	Urn       string
	Struct    Json
	Name      string
	IsDataset bool
}

type Json map[string]interface{}

