package services

import "testing"

func DataHubTestBrowse(t *testing.T) {
	conn := "urn:li:corpuser:datahub"
	dh := ConnectToDH(conn)
	dh.GetBrowses()
}
