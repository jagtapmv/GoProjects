package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	dc "github.com/dchest/validator"
)

type Result struct {
	Message    string
	StatusList [][]string
}

var Results = Result{}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/results", statusChecker)
	http.ListenAndServe(":8000", nil)
}

func getStatus(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	var err error
	if !strings.Contains(url, "https://") && !strings.Contains(url, "http://") {
		elems := []string{"https://", url}
		url = strings.Join(elems, "")
	}
	err = dc.ValidateDomainByResolvingIt(strings.Split(strings.Split(url, "//")[1], "/")[0])

	if err != nil {
		log.Print(err)
	} else {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		Results.StatusList = append(Results.StatusList, []string{url, fmt.Sprint(res.StatusCode)})
	}
}

func statusChecker(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	r.ParseMultipartForm(1 << 20)
	file, handler, err := r.FormFile("file")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	if handler.Size > 1<<20 {
		fmt.Println("File size is too big")
		return
	}

	res := csv.NewReader(file)

	urls := []string{}
	for {
		record, err := res.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		urls = append(urls, record...)

	}
	// urls := []string{
	// 	"https://www.google.com/",
	// 	"https://www.github.com",
	// 	"https://support.google.com",
	// 	"https://www.kroll.com/thisisnoturl",
	// 	"https://www.modis.com",
	// 	"https://www.youtube.com",
	// 	"https://www.gmail.com/thisisnoturl",
	// 	"https://open.spotify.com",
	// 	"https://www.w3schools.com",
	// 	"https://www.w3.org",
	// 	"www.geeksforgeeks.org",
	// 	"https://in.linkedin.com/thisisnoturl",
	// }

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go getStatus(&wg, url)
	}
	wg.Wait()
	timePassed := time.Since(t)
	operationTime := fmt.Sprintf("Operation completed in: %s", timePassed)

	Results.Message = operationTime
	temp, _ := template.ParseFiles("result.html")
	temp.Execute(w, Results)

	csvWriter := csv.NewWriter(os.Stdout)

	csvWriter.WriteAll(Results.StatusList)
	if err := csvWriter.Error(); err != nil {
		log.Fatalf("Error writing to CSV: %s", err)
	}
	Results = Result{}
}
