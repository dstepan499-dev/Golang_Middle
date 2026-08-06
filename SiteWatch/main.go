package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL string
	Success bool
	Message string
}

func checkSite(ctx context.Context, url string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- Result{URL: url, Success: false, Message: err.Error()}
		return
	}
	
	http.DefaultClient.Do(req)
	defer req.Body.Close()

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
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		go checkSite(ctx, site, results, &wg)
		defer cancel()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}
