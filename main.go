package main

import (
	"fmt"
	"os"
)

// nlp_text_analyzer - NLP toolkit for sentiment analysis
func nlp_text_analyzer(path string) {
	fmt.Println("========================================")
	fmt.Println("  NLP-Text-Analyzer")
	fmt.Println("  NLP toolkit for sentiment analysis")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	nlp_text_analyzer(path)
}
