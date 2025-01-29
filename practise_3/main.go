package main

import (
	"fmt"
)

type HealthCheck struct {
	serviceId int
	status    string
}

const PASSED = "pass"
const FAILED = "fail"

func generateCheck() []HealthCheck {
	result := []HealthCheck{}

	for i := 0; i < 5; i++ {
		var checkStatus string
		if i%2 == 0 {
			checkStatus = PASSED
		} else {
			checkStatus = FAILED
		}
		result = append(result, HealthCheck{
			serviceId: i,
			status:    checkStatus,
		})
	}

	return result
}

func main() {
	fmt.Println("Тут будет выведен идентификатор")

	healthChecks := generateCheck()
	for _, hc := range healthChecks {
		if hc.status == PASSED {
			fmt.Println("Service %s has passed", hc.serviceId)
		} else {
			fmt.Println("Service %s has failed", hc.serviceId)
		}
	}
}
