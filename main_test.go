package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"
)

func TestHealthChecking(t *testing.T) {
	csvPath = "test.csv"
	csvURL, _ := os.Open(csvPath)
	reader = csv.NewReader(bufio.NewReader(csvURL))

	total, _, _, totalHealthCheckTime := HealthChecking()

	if csvPath == "test.csv" {
		if total != 5 {
			t.Error("HealthCheck counting is incorrect!")
		}
		if totalHealthCheckTime.Seconds() > 6.0 {
			t.Error("HelthCheck is too slow! on test set")
		}
	} else if csvPath == "tiny-test.csv" {
		if total != 25 {
			t.Error("HealthCheck counting is incorrect!")
		}
		if totalHealthCheckTime.Seconds() > 10.0 {
			t.Error("HelthCheck is too slow! on tiny test set")
		}
	} else if csvPath == "small-test.csv" {
		if total != 40 {
			t.Error("HealthCheck counting is incorrect!")
		}
		if totalHealthCheckTime.Seconds() > 16.0 {
			t.Error("HelthCheck is too slow! on small test set")
		}

	} else if csvPath == "medium-test.csv" {
		if total != 200 {
			t.Error("HealthCheck counting is incorrect!")
		}
		if totalHealthCheckTime.Seconds() > 750 {
			t.Error("HelthCheck is too slow! on medium test set")
		}
	} else if csvPath == "large-test.csv" {
		if total != 1721 {
			t.Error("HealthCheck counting is incorrect!")
		}
		if totalHealthCheckTime.Seconds() > 1500 {
			t.Error("HelthCheck is too slow! on medium test set")
		}
	}
}
