package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: nlp-analyzer <text-file>")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	text := string(data)
	words := strings.Fields(text)
	sentences := strings.Split(text, ".")
	uniqueWords := map[string]bool{}
	for _, w := range words {
		w = strings.ToLower(strings.Trim(w, ".,!?;:"))
		if len(w) > 0 {
			uniqueWords[w] = true
		}
	}
	fmt.Println("NLP Text Analyzer")
	fmt.Println("=================")
	fmt.Printf("Words:        %d\n", len(words))
	fmt.Printf("Unique words: %d\n", len(uniqueWords))
	fmt.Printf("Sentences:    %d\n", len(sentences))
	fmt.Printf("Avg word len: %.1f chars\n", avgWordLen(words))
	fmt.Printf("Vocab rich:   %.1f%%\n", float64(len(uniqueWords))/float64(len(words)+1)*100)
}

func avgWordLen(words []string) float64 {
	total := 0
	for _, w := range words {
		total += len(w)
	}
	return float64(total) / float64(len(words)+1)
}