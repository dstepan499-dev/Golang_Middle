package main

import (
	"fmt"
	"net/http"
)

func checkSite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	
	defer resp.Body.Close()

	fmt.Println(resp)
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