package services

import (
	"context"
	"log"
	"net/http"

	"github.com/shurcooL/graphql"
)

type DataHub struct {
	// urn:li:corpuser:datahub
	// ConnectString string
	Client *graphql.Client
}
type headerTransport struct {
	base    http.RoundTripper
	headers map[string]string
}

func ConnectToDH(ConnectString string) *DataHub {
	client := graphql.NewClient("http://datahub.yc.pbd.ai:9002", nil)

	return &DataHub{
		Client: client,
	}
}

func (DataHub *DataHub) GetBrowses() {
	query := map[string]interface{}{
		"operationName": "getBrowseResults",
		"variables": map[string]interface{}{
			"input": map[string]interface{}{
				"type":  "DATASET",
				"path":  []string{"prod"},
				"start": 0,
				"count": 10,
				// "filters": null
				"filters": "",
			},
		},
	}
	variables := map[string]interface{}{}
	err := DataHub.Client.Query(context.Background(), &query, variables)
	if err != nil {
		log.Println(err)
	}
}
