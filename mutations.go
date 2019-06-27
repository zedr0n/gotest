package main

import (
	"context"
	"log"
)

type CreateRoot struct {
	CreateRoot bool `graphql:"createRoot(name : $name)"`
}

type CreateRecord struct {
	CreateRecord bool `graphql:"createRecord(target : $target)"`
}

// Run mutation query on server
func (client *MyClient) RunMutation(mutation interface{}, varsParams ...map[string]interface{}) {
	// run it and capture the response
	ctx := context.Background()
	vars := map[string]interface{}{}
	if len(varsParams) > 0 {
		vars = varsParams[0]
	}

	if err := client.Mutate(ctx, mutation, vars); err != nil {
		log.Fatal(err)
	}
}
