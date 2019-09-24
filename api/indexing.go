package main

import (
	"log"

	"github.com/7onetella/password/api/model"
	"github.com/blevesearch/bleve"
)

// IndexPasswords indexes passwords
func IndexPasswords(passwords []model.Password) {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("passwords.bleve", mapping)
	if err != nil {
		panic(err)
	}
	for _, password := range passwords {
		index.Index(password.ID, password)
	}
	index.Close()
}

// IndexSearch search index for presence of given text
func IndexSearch(text string) (*bleve.SearchResult, error) {
	index, err := bleve.Open("passwords.bleve")
	defer index.Close()

	if err != nil {
		log.Println("error while opening the index:", err)
		return nil, err
	}

	query := bleve.NewQueryStringQuery(text)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		log.Println("error while searching the index:", err)
		return nil, err
	}
	return searchResult, nil
}
