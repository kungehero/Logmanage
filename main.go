package main

import (
	"Logmanage/logcollect"
)

type Tweet struct {
	User     string
	Message  string
	Retweets int
}

func main() {

	logcollect.ReadLogWithEs()
	/* 	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200/"))
	   	exists, err := client.IndexExists("twitter").Do(context.Background())
	   	if err != nil {
	   		// Handle error
	   		panic(err)
	   	}
	   	if !exists {
	   		// Create a new index.
	   		mapping := `fedbgfrthty`
	   		createIndex, err := client.CreateIndex("twitter").BodyString(mapping).Do(context.Background())
	   		if err != nil {
	   			// Handle error
	   			panic(err)
	   		}
	   		if !createIndex.Acknowledged {
	   			// Not acknowledged
	   		}
	   	}

	   	// Index a tweet (using JSON serialization)

	   	tweet1 := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	   	put1, err := client.Index().
	   		Index("twitter").
	   		Type("doc").
	   		Id("1").
	   		BodyJson(tweet1).
	   		Do(context.Background())
	   	if err != nil {
	   		// Handle error
	   		panic(err)
	   	}
	   	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type) */

}
