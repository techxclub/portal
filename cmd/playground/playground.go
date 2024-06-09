package main

import (
	"context"
	gohttp "net/http"
	"time"

	"github.com/techx/portal/client/http"
	"github.com/techx/portal/config"
)

type request struct {
	Name string `json:"name"`
	Data data   `json:"data"`
}

type response struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Data      data      `json:"data"`
	CreatedAt time.Time `json:"createdAt"`
}

type data struct {
	Year         int     `json:"year"`
	Price        float64 `json:"price"`
	CPUModel     string  `json:"CPU model"`
	HardDiskSize string  `json:"Hard disk size"`
}

func main() {
	var testResponse response

	testCmd := "test"
	testHTTPConfig := config.HTTPConfig{
		Host:                   "api.restful-api.dev",
		HTTPTimeout:            1000,
		RetryCount:             2,
		HystrixTimeout:         1000,
		MaxConcurrentRequests:  10,
		RequestVolumeThreshold: 100,
	}

	testReq := request{
		Name: "Apple MacBook Pro 16",
		Data: data{
			Year:         2019,
			Price:        1849.99,
			CPUModel:     "Intel Core i9",
			HardDiskSize: "1 TB",
		},
	}
	testDoer := http.DefaultDoer("test", testHTTPConfig)

	_ = http.NewRequest(context.Background(), testCmd).
		SetMethod(gohttp.MethodPost).
		SetHost(testHTTPConfig.Host).
		SetPath("/objects").
		SetBody(testReq).
		Send(testDoer, &testResponse, &testResponse)
}
