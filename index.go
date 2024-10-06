package main

import (
        "fmt"
        "regexp"
        "strings"
)

type SearchEngineIndexer struct {
        index map[string][]int // Word -> list of doc_ids
}

func NewSearchEngineIndexer() *SearchEngineIndexer {
        return &SearchEngineIndexer{index: make(map[string][]int)}
}

// AddDocument adds a document to the index
func (sei *SearchEngineIndexer) AddDocument(docID int, text string) {
        // Normalize text to lowercase and tokenize
        words := tokenize(text)
        wordSet := make(map[string]struct{}) // To avoid duplicate words in the same doc

        for _, word := range words {
                wordSet[word] = struct{}{}
        }

        // Add unique words to the index
        for word := range wordSet {
                sei.index[word] = append(sei.index[word], docID)
        }
}

// BuildIndex processes all documents
func (sei *SearchEngineIndexer) BuildIndex(docs map[int]string) {
        for docID, content := range docs {
                sei.AddDocument(docID, content)
        }
}

// Tokenize splits the text into words and returns them in lowercase
func tokenize(text string) []string {
        re := regexp.MustCompile(`\w+`)
        words := re.FindAllString(strings.ToLower(text), -1)
        return words
}

func main() {
        documents := map[int]string{
                1: "Search engines are programs that search the web.",
                2: "A web crawler is used to gather data for search engines.",
                3: "Indexing is crucial for search engine efficiency.",
        }

        indexer := NewSearchEngineIndexer()
        indexer.BuildIndex(documents)

        // Print the inverted index
        for word, docIDs := range indexer.index {
                fmt.Printf("%s: %v\n", word, docIDs)
        }
}
