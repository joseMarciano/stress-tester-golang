package application

import (
	"fmt"
	"net/http"
	"stress-test/internal/model"
	"sync"
	"time"
)

type StressTesterService struct {
	url            string
	requestsNumber uint64
	concurrency    uint64
}

func NewStressTesterService(url string, requestsNumber, concurrency uint64) (*StressTesterService, error) {
	if concurrency == 0 {
		concurrency = 1
	}

	if requestsNumber == 0 {
		requestsNumber = 1
	}

	if url == "" {
		return nil, fmt.Errorf("url must not be nil")
	}

	return &StressTesterService{
		url:            url,
		requestsNumber: requestsNumber,
		concurrency:    concurrency,
	}, nil
}

func (s *StressTesterService) Test() *model.Report {
	concurrencyChannel := make(chan uint64, s.concurrency)

	result := model.NewReport()
	initTime := time.Now()

	var counter uint64
	var wg sync.WaitGroup

	for counter < s.requestsNumber {
		concurrencyChannel <- 1
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				<-concurrencyChannel
			}()

			result.UpdateReport(s.executeHttpRequest())

		}()

		counter++
	}

	wg.Wait()
	result.SpentTime(initTime)
	return result
}

func (s *StressTesterService) executeHttpRequest() uint {
	client := http.DefaultClient

	response, _ := client.Get(s.url)

	return uint(response.StatusCode)
}
