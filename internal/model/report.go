package model

import (
	"sync"
	"time"
)

type Report struct {
	TotalSpentTime int64
	TotalRequests  uint64
	RequestMap     map[uint]uint64
}

func NewReport() *Report {
	return &Report{
		TotalSpentTime: 0,
		TotalRequests:  0,
		RequestMap:     make(map[uint]uint64),
	}
}

func (r *Report) SpentTime(initTime time.Time) {
	r.TotalSpentTime = time.Now().Sub(initTime).Milliseconds()
}

var updateMtx = sync.Mutex{}

func (r *Report) UpdateReport(statusCode uint) {
	updateMtx.Lock()
	defer updateMtx.Unlock()

	r.TotalRequests += 1
	_, ok := r.RequestMap[statusCode]

	if !ok {
		r.RequestMap[statusCode] = 1
		return
	}

	r.RequestMap[statusCode] += 1

}

func (r *Report) TotalRequestByStatusCode(statusCode uint) uint64 {
	total, ok := r.RequestMap[statusCode]

	if !ok {
		return 0
	}

	return total
}
