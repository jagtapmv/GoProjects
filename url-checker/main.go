package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	dc "github.com/dchest/validator"
)

const (
	numberOfThreads = 4
)

var (
	wg sync.WaitGroup
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error occured: ", err)
	}
}

func getData() []string {
	var data []string
	file, err := os.Open("urls.csv")
	checkError(err)
	defer file.Close()

	csvData := csv.NewReader(file)

	for {
		record, err := csvData.Read()
		if err == io.EOF {
			break
		}
		checkError(err)
		data = append(data, record...)
	}
	return data
}

func getstatus(inputChannel chan string) {

	for url := range inputChannel {
		var err error
		var sc int = 500
		if !strings.Contains(url, "https://") && !strings.Contains(url, "http://") {
			elems := []string{"https://", url}
			url = strings.Join(elems, "")
		}
		err = dc.ValidateDomainByResolvingIt(strings.Split(strings.Split(url, "//")[1], "/")[0])

		if err != nil {
			log.Print(err)
		} else {
			res, err := http.Head(url)
			if err != nil {
				log.Print(err)
			}
			sc = res.StatusCode
		}
		fmt.Printf("%v : %v\n", url, sc)
	}
	wg.Done()
}

func main() {

	csvData := getData()

	inputChannel := make(chan string, 1000)

	for i := 0; i < numberOfThreads; i++ {
		go getstatus(inputChannel)
	}
	wg.Add(numberOfThreads)
	t := time.Now()
	for _, data := range csvData {
		inputChannel <- data
	}
	close(inputChannel)
	wg.Wait()
	fmt.Println("Total elapsed time: ", time.Since(t))
}
