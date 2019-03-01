package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var csvPath string
var wg sync.WaitGroup
var reader *csv.Reader
var totalHealthCheckTime time.Duration
var totalCount, successCount, failCount int = 0, 0, 0
var healthcheckReportURL = "https://hiring-challenge.appspot.com/healthcheck/report"
var accessToken = "eyJhbGciOiJIUzI1NiJ9.mZiiSx9dAHaaGhhtugBiuWIrbYZngAJWVokCL4dcL8YzELpY2HsN56WAG965-kkVlIv7M4_GPbU4imAzQ5O1qmh5r0M2uXu2Tv5PB7on8w1lNDZdc4sAjX0JyHSDSQJIPOr2gdeR39teF-8iaRa7KmPnwns4PpUEMxNsZRwXTg8.gExuUY4Ka01eOoSny1WF62ELS16U1XS5dAe3pFuS6Z0"

// DEBUG for turn on/off loging individual HealCheck Request
var DEBUG = false

func main() {
	CheckCSVExist()
	ReadCSV()
	HealthChecking()
	DisplayHealthSummary()
	SendHealthReport()
}

// CheckCSVExist for check if The CSV exist
func CheckCSVExist() {
	csvPath = os.Args[1]
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		log.Fatal(err)
	}
}

// ReadCSV for read csv file that contain url to website
func ReadCSV() {
	csvURL, _ := os.Open(csvPath)
	reader = csv.NewReader(bufio.NewReader(csvURL))
	reader.Comma = ','
}

// HealthChecking perform concurrency health checking to all URL in CSV file
func HealthChecking() (int, int, int, time.Duration) {
	startHealthCheck := time.Now()
	fmt.Println("\nPerform website checking...")
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			PerformHealthCheck(line[0])
		}()
		totalCount++
	}
	wg.Wait()
	fmt.Println("Done!" + "\n")
	totalHealthCheckTime = time.Since(startHealthCheck)

	return totalCount, successCount, failCount, totalHealthCheckTime
}

// PerformHealthCheck is function for checking individual URL
func PerformHealthCheck(currentURL string) {
	currentURL = strings.TrimLeft(currentURL, "'")
	currentURL = strings.TrimRight(currentURL, "'")

	resp, err := http.Get(currentURL)
	if err != nil {
		if DEBUG {
			fmt.Println(err.Error())
		}

		if strings.Contains(err.Error(), "timed out") || strings.Contains(err.Error(), "timeout") {
			// time out here
			failCount++
		}
	} else {
		if DEBUG {
			fmt.Println(currentURL + " => " + resp.Status)
		}

		if resp.Status == "524 A timeout occurred" {
			// time out here
			failCount++
		} else {
			// all http status code is valid.
			successCount++
		}
	}
}

// DisplayHealthSummary result of health checking
func DisplayHealthSummary() {
	fmt.Printf("Checked webistes: %d\n", totalCount)
	fmt.Printf("Successful websites: %d\n", successCount)
	fmt.Printf("Failure websites: %d\n", failCount)
	fmt.Printf("Total times to finished checking website: %s\n", totalHealthCheckTime)
}

// SendHealthReport perform Http Post to Healthcheck Report API with Authorization Header
func SendHealthReport() {
	url := healthcheckReportURL
	if DEBUG {
		url = "http://ptsv2.com/t/mikael_toliet/post"
	}

	message := map[string]interface{}{
		"total_websites": totalCount,
		"success":        successCount,
		"failure":        failCount,
		"total_time":     totalHealthCheckTime,
	}
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	if DEBUG {
		fmt.Println("Healthcheck Report API response Status: ", resp.Status)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}
}
