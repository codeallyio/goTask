package main

import (
	"errors"
	"fmt"
	"time"
)

var mockConnections = []connection{
	{ip: "111.111.111.111"},
	{ip: "222.222.222.222"},
	{ip: "333.333.333.333"},
}

var activeConnections = []connection{}
var activeRequests = make(chan request, 9)

type request struct {
	rid            int8 // request id
	ip             string
	processingTime int32 // in miliseconds
}

type connection struct {
	ip        string
	holdUntil time.Time
}

func getConnection(r request) (connection, error) {
	for _, c := range activeConnections {
		if r.ip == c.ip {
			return c, nil
		}
	}

	err := errors.New("there is no such connection")
	return connection{}, err
}

func (c *connection) refreshConnection() {
	c.holdUntil = time.Now().Add(3 * time.Second)
}

func runMockServer() {
	for _, c := range mockConnections {
		c.refreshConnection()
		activeConnections = append(activeConnections, c)

		fmt.Println("New connection from:", c.ip)

		time.Sleep(600 * time.Millisecond)
	}
	fmt.Println()

	generateSomeRequests()
}

func generateSomeRequests() {
	activeRequests <- request{rid: 1, ip: mockConnections[0].ip, processingTime: 1111}
	activeRequests <- request{rid: 2, ip: mockConnections[0].ip, processingTime: 2111}
	activeRequests <- request{rid: 3, ip: mockConnections[0].ip, processingTime: 3111}

	activeRequests <- request{rid: 4, ip: mockConnections[1].ip, processingTime: 1444}
	activeRequests <- request{rid: 5, ip: mockConnections[1].ip, processingTime: 2444}
	activeRequests <- request{rid: 6, ip: mockConnections[1].ip, processingTime: 3444}

	activeRequests <- request{rid: 7, ip: mockConnections[2].ip, processingTime: 1777}
	activeRequests <- request{rid: 8, ip: mockConnections[2].ip, processingTime: 2777}
	activeRequests <- request{rid: 9, ip: mockConnections[2].ip, processingTime: 3777}

	close(activeRequests)
}

var requestsCompleted = 0

func checkSuccess() {
	fmt.Println()

	if requestsCompleted == 9 {
		fmt.Println("=== Success ===")
	} else {
		fmt.Println("=== Failure ===")
	}

	fmt.Printf("Completed %d out of 9 requests\n\n", requestsCompleted)
}

func handleRequest(r request) {
	c, err := getConnection(r)

	fmt.Println("RID:", r.rid, ">> From:", r.ip, ">> New request")

	if err != nil {
		fmt.Println("RID:", r.rid, ">> From:", r.ip, "Error:", err)
		return
	}

	if !c.isActive() {
		fmt.Println("RID:", r.rid, ">> From:", r.ip, ">> Request failure, connection no longer active!")
		return
	}

	time.Sleep(time.Duration(r.processingTime) * time.Millisecond)

	requestsCompleted++
	fmt.Println("RID:", r.rid, ">> From:", r.ip, ">> Request completed")
	c.refreshConnection()
}
