package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

const concurrent = 5

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	jobChan := make(chan string, 100)

	for w := 1; w <= concurrent; w++ {
		go worker(w, jobChan)
		wg.Add(1)
	}

	for scanner.Scan() {

		jobChan <- scanner.Text()

	}

	close(jobChan)
	wg.Wait()
}

func worker(id int, jobChan <-chan string) {
	defer wg.Done()
	for job := range jobChan {
		fetch(job)
	}
}

func fetch(slug string) {
	response, err := http.Get("https://wordpress.org/plugins/" + slug)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 404 {
		fmt.Printf("%s - page 404 \n", slug)
		return
	}

	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)

	if strings.Contains(pageContent, "Nothing Found") {
		fmt.Printf("%s - Plugin not found \n", slug)
		return
	}

	titleStartIndex := strings.Index(pageContent, "userInteractionCount")
	if titleStartIndex == -1 {
		fmt.Printf("%s - Bad slug or missing data \n", slug)
		return
	}

	titleStartIndex += 23

	// Find the index of the closing tag
	titleEndIndex := strings.Index(pageContent, "\"image\":")
	if titleEndIndex == -1 {
		fmt.Printf("%s - Bad slug or missing data \n", slug)
		return
	}

	titleEndIndex -= 20
	count := []byte(pageContent[titleStartIndex:titleEndIndex])

	fmt.Printf("%s - userInteractionCount: %s\n", slug, count)
}
