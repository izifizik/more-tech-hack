package services

import (
	"context"
	"github.com/machinebox/graphql"
	"log"
	"more-tech-hack/internal/config"
	"strings"
)

var DataHub *graphql.Client

func Init() {
	// create a client (safe to share across requests)
	DataHub = graphql.NewClient(config.DataHubUrl)
}

func GetDataset(urn string) DatasetResp {

	// make a request
	req := graphql.NewRequest(`
  {
  dataset(urn: "` + urn + `") {
    urn
    type
    name
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
	//req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Cookie", config.CookieDataHub)

	// run it and capture the response
	var respData DatasetResp
	if err := DataHub.Run(context.Background(), req, &respData); err != nil {
		log.Println(err)
	}

	return respData
}

func Browse(path string) BrowseResp {

	var res string
	var array []string
	if path != "" {
		array = strings.Split(path, ".")
		res = "["
		for _, v := range array {
			res += "\"" + v + "\","
		}
		res = res[:len(res)-1]
		res += "]"
	} else {
		res = "[]"
	}

	// make a request
	req := graphql.NewRequest(`
 {
  browse(input: {type: DATASET, path: ` + res + `, start: 0, count: 10, filters: null}) {
    entities {
      urn
      type
    }
    groups {
      name
      count
    }
    start
    count
    total
    metadata {
      path
      totalNumEntities
    }
  }
}
`)

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Cookie", config.CookieDataHub)

	// run it and capture the response
	var respData BrowseResp
	if err := DataHub.Run(context.Background(), req, &respData); err != nil {
		log.Println(err)
	}

	for i, v := range respData.Browse.Entities {
		if v.Urn != "" {
			name := strings.ReplaceAll(v.Urn, ",", ":")
			name = strings.ReplaceAll(name, "(", ":")
			name = strings.ReplaceAll(name, ")", ":")
			a := strings.Split(name, ":")
			for i2, v := range a {
				if v == array[len(array)-1] {
					respData.Browse.Entities[i].Name = a[i2+1]
				}
			}
		}
	}

	return respData
}

type BrowseResp struct {
	Browse struct {
		Entities []struct {
			Urn  string `json:"urn"`
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"entities"`
		Groups []struct {
			Name  string `json:"name"`
			Count int    `json:"count"`
		} `json:"groups"`
		Start    int `json:"start"`
		Count    int `json:"count"`
		Total    int `json:"total"`
		Metadata struct {
			Path             []interface{} `json:"path"`
			TotalNumEntities int           `json:"totalNumEntities"`
		} `json:"metadata"`
	} `json:"browse"`
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
