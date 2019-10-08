package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	output     = os.Stdout
	outputLock = sync.Mutex{}
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s URL...\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	urls := flag.Args()
	if len(urls) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			catHttp(url)
			wg.Done()
		}(url)
	}
	go waitAndFlush(time.Second)
	wg.Wait()
}

func catHttp(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("%v: %v", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		log.Fatalf("%v: returns %v", url, res.StatusCode)
	}

	lines := bufio.NewScanner(res.Body)
	lines.Buffer(nil, 1e6)
	for lines.Scan() {
		outputLock.Lock()
		output.WriteString(lines.Text())
		output.WriteString("\n")
		outputLock.Unlock()
	}
	if err := lines.Err(); err != nil {
		log.Fatalf("%v: %v", url, err)
	}
}

func waitAndFlush(delay time.Duration) {
	for {
		time.Sleep(delay)

		outputLock.Lock()
		output.Sync()
		outputLock.Unlock()
	}
}
