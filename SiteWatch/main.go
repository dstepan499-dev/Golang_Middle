package main

import (
	"fmt"
	"net/http"
)

func checkSite(url string) {
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
	sites := []string{
		"https://google.com",
		"https://github.com",
		"https://non-existent-site.qwe",
	}

	for _, site := range sites {
		checkSite(site)
	}
}