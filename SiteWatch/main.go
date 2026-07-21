package main

import (
	"fmt"
	"net/http"
	"sync"
)

func checkSite(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[FAIL] %s (ошибка: %v)\n", url, err)
		return
	}
	
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("[OK]  %s\n", url)
	} else {
		fmt.Printf("[FAIL] %s (статус: %d)\n", url, resp.StatusCode)
	}
}

func main() {
	var wg sync.WaitGroup
	sites := []string{
		"https://google.com",
		"https://github.com",
		"https://non-existent-site.qwe",
	}

	for _, site := range sites {
		wg.Add(1)
		go checkSite(site, &wg)
	}
	
	wg.Wait()
}