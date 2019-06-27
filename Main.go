package main

import (
	"fmt"
	"time"

	"github.com/shurcooL/graphql"
)

type MyClient struct {
	graphql.Client
}

// Create new connection to localhost graphql server
func (client *MyClient) Connect() {
	client.Client = *graphql.NewClient("http://localhost:5000", nil)
}

func main() {
	client := MyClient{}
	client.Connect()

	// query := statsQuery{}
	rQuery := RootInfoQuery{}

	vars := map[string]interface{}{
		"id": graphql.String("Root"),
	}

	mutationVars := map[string]interface{}{
		"target": graphql.String("Root"),
	}

	mutation := CreateRecord{}

	client.RunMutation(&mutation, mutationVars)

	lQuery := LastRecordQuery{}
	client.RunQuery(&lQuery, map[string]interface{}{"id": graphql.String("Root")})
	fmt.Println(lQuery.LastRecord)

	client.RunQuery(&rQuery, vars)
	fmt.Printf("%v created at %v\n", rQuery.RootInfo.RootID, time.Unix(rQuery.RootInfo.CreatedAt/1000, 0))

	sQuery := StatsQuery{}
	client.RunQuery(&sQuery)
	fmt.Printf("Number of roots :  %v\n", sQuery.Stats.NumberOfRoots)
}
