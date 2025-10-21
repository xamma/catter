package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/xamma/catter/internal/fetcher"
)

func main() {
	savePath := flag.String("path", ".", "Path where to save the file")
	numImages := flag.Int("num", 0, "Number of images to fetch")

	fmt.Println("Starting catter...")
	fmt.Println("Your friendly feline image getter 🐾")
	fmt.Println("")
	fmt.Println("Usage: catter -path <path_to_save> -num <number_of_images>")
	fmt.Println("")
	flag.Parse()

	if *numImages <= 0 {
		log.Fatal("Error: -num flag is required and must be greater than 0")
	}

	cats, err := fetcher.FetchCatImages(*numImages)
	if err != nil {
		log.Fatal(err)
	}

	if len(cats) > *numImages {
		cats = cats[:*numImages]
	}

	for i, cat := range cats {
		fmt.Printf("Cat %d: %s\n", i+1, cat.Url)
		err := fetcher.SaveCatImage(cat.Url, *savePath)
		if err != nil {
			log.Fatal(err)
		}
	}
}
