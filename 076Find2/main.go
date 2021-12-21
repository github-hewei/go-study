package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var query = "README.md"
var matches int

var workerCount int
var maxWorkerCount = 32
var searchRequest = make(chan string)
var workerDone = make(chan bool)
var foundMatch = make(chan bool)

func main() {
	start := time.Now()
	workerCount = 1
	go search("d:/temp", true)
	waitForWorker()
	fmt.Println("数量：", matches)
	fmt.Println("耗时：", time.Since(start))
}

func waitForWorker() {
	for {
		select {
		case p := <- searchRequest:
			workerCount++
			go search(p, true)
		case <- workerDone:
			workerCount--
			if workerCount == 0 {
				return
			}
		case <- foundMatch:
			matches++
		}
	}
}

func search(p string, master bool) {
	files, err := ioutil.ReadDir(p)
	if err == nil {
		for _, f := range files {
			if f.Name() == query {
				foundMatch <- true
			}
			if f.IsDir() {
				if workerCount < maxWorkerCount {
					searchRequest <- p + "/" + f.Name() + "/"
				} else {
					search(p + "/" + f.Name() + "/", false)
				}
			}
		}
		if master {
			workerDone <- true
		}
	}
}
