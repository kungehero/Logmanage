package models

import (
	"fmt"

	"github.com/olivere/elastic"
)

var EsClient *elastic.Client

func NewEs(path string) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(path))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}
	EsClient = client
	fmt.Println("conn es succ")

}
