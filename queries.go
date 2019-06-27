package main

import (
	"context"
	"log"
)

type StatsQuery struct {
	Stats struct {
		NumberOfRoots int
	} `graphql:"statsQuery"`
}

type RootInfoQuery struct {
	RootInfo struct {
		RootID    string
		CreatedAt int64
		UpdatedAt int64
	} `graphql:"rootInfoQuery(id : $id )"`
}

type LastRecordQuery struct {
	LastRecord struct {
		Id        string
		TimeStamp int64
		Value     float32
	} `graphql:"lastRecordQuery(id : $id)"`
}

// Run graphql query on server
func (client *MyClient) RunQuery(query interface{}, varsParams ...map[string]interface{}) {
	// run it and capture the response
	ctx := context.Background()
	vars := map[string]interface{}{}
	if len(varsParams) > 0 {
		vars = varsParams[0]
	}

	if err := client.Query(ctx, query, vars); err != nil {
		log.Fatal(err)
	}
}
