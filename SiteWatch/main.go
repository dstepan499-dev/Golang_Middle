package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Result struct {
	URL string
	Success bool
	Message string
}

func checkSite(url string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		results <- Result{URL: url, Success: false, Message: err.Error()}
		return
	}
	
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		results <- Result{URL: url, Success: true, Message: "OK"}
	} else {
		results <- Result{URL: url, Success: false, Message: fmt.Sprintf("Status: %d", resp.StatusCode)}
	}
}

//можно использовать локальную переменную, чтобы не делать results <- Result{...}
/*
var res Result
res.URL = url
resp, err := http.Get(url)
    if err != nil {
        res.Success = false
        res.Message = "Ошибка: " + err.Error()
        ch <- res // Отправляем и выходим
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        res.Success = true
        res.Message = "OK"
    } else {
        res.Success = false
        res.Message = fmt.Sprintf("Статус: %d", resp.StatusCode)
    }

    ch <- res // Отправляем финальный результат
}
*/

func main() {
	var wg sync.WaitGroup
	results := make(chan Result)
	sites := []string{
		"https://google.com",
		"https://github.com",
		"https://non-existent-site.qwe",
	}

	for _, site := range sites {
		wg.Add(1)
		go checkSite(site, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}
