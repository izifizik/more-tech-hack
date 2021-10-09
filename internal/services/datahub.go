package services

import (
	"context"
	"github.com/machinebox/graphql"
	"log"
	"more-tech-hack/internal/config"
)

var DataHub *graphql.Client

func Init() {
	// create a client (safe to share across requests)
	DataHub = graphql.NewClient("http://datahub.yc.pbd.ai:9002/api/graphql")
}

func GetDataset() DatasetResp {

	// make a request
	req := graphql.NewRequest(`
  {
  dataset(urn: "urn:li:dataset:(urn:li:dataPlatform:hive,fct_users_created,PROD)") {
    urn
    type
    name
    status {
      removed
    }
    properties {
      description
      origin
      customProperties {
        key
        value
      }
    }
    editableProperties {
      description
    }
    datasetProfiles {
      rowCount
      columnCount
    }
    schemaMetadata(version: 0) {
      aspectVersion
      datasetUrn
      hash
      platformSchema {
        __typename
      }
      primaryKeys
      foreignKeys {
        name
        foreignDataset {
          urn
          type
          platform {
            urn
            name
          }
          name
          properties {
            origin
            description
            externalUrl
          }
        }
      }
      fields {
        fieldPath
        nullable
        jsonPath
        description
        type
        nativeDataType
        tags {
          tags {
            tag {
              urn
              type
              name
              editableProperties {
                description
              }
              ownership {
                owners {
                  owner {
                    __typename
                  }
                  type
                  source {
                    type
                    url
                  }
                }
                lastModified {
                  time
                  actor
                }
              }
            }
          }
        }
      }
    }
  }
}
`)

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Cookie", config.CookieDataHub)

	// run it and capture the response
	var respData DatasetResp
	if err := DataHub.Run(context.Background(), req, &respData); err != nil {
		log.Println(err)
	}

	return respData
}

type DatasetResp struct {
	Dataset struct {
		Urn        string      `json:"urn"`
		Type       string      `json:"type"`
		Name       string      `json:"name"`
		Status     interface{} `json:"status"`
		Properties struct {
			Description      string `json:"description"`
			Origin           string `json:"origin"`
			CustomProperties []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"customProperties"`
		} `json:"properties"`
		EditableProperties struct {
			Description string `json:"description"`
		} `json:"editableProperties"`
		DatasetProfiles []interface{} `json:"datasetProfiles"`
		SchemaMetadata  struct {
			AspectVersion  int         `json:"aspectVersion"`
			DatasetUrn     interface{} `json:"datasetUrn"`
			Hash           string      `json:"hash"`
			PlatformSchema struct {
				Typename string `json:"__typename"`
			} `json:"platformSchema"`
			PrimaryKeys interface{} `json:"primaryKeys"`
			ForeignKeys []struct {
				Name           string `json:"name"`
				ForeignDataset struct {
					Urn      string `json:"urn"`
					Type     string `json:"type"`
					Platform struct {
						Urn  string `json:"urn"`
						Name string `json:"name"`
					} `json:"platform"`
					Name       string `json:"name"`
					Properties struct {
						Origin      string      `json:"origin"`
						Description string      `json:"description"`
						ExternalURL interface{} `json:"externalUrl"`
					} `json:"properties"`
				} `json:"foreignDataset"`
			} `json:"foreignKeys"`
			Fields []struct {
				FieldPath      string      `json:"fieldPath"`
				Nullable       bool        `json:"nullable"`
				JSONPath       interface{} `json:"jsonPath"`
				Description    string      `json:"description"`
				Type           string      `json:"type"`
				NativeDataType string      `json:"nativeDataType"`
				Tags           interface{} `json:"tags"`
			} `json:"fields"`
		} `json:"schemaMetadata"`
	} `json:"dataset"`
}
