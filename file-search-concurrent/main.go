package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("File search started at: ", t)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go fileSearch("/home/mahesh/Documents/GO/GoProjects", ".go", &wg)
	wg.Wait()
	timeElapsed := time.Since(t)
	fmt.Println("Total time elapsed: ", timeElapsed)
}

func fileSearch(dir string, filename string, wg *sync.WaitGroup) {
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), filename) {
			fmt.Println("found at: ", dir+"/"+file.Name())
		} else if file.IsDir() {
			newPath := filepath.Join(dir, file.Name())
			wg.Add(1)
			go fileSearch(newPath, "main.go", wg)
		}
	}
	wg.Done()
}
