# retry
[![GoDoc](https://godoc.org/github.com/karngyan/retry?status.svg)](https://godoc.org/github.com/karngyan/retry)
[![Go Report Card](https://goreportcard.com/badge/github.com/karngyan/retry)](https://goreportcard.com/report/github.com/karngyan/retry)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)]()

Retry is a simple Golang library for automatically retrying failed operations. It provides configurable mechanisms for retrying operations upon failure, such as maximum retry attempts, exponential backoff, and maximum elapsed time.

Installation
To install the Retry library, simply run:

arduino
Copy code
go get github.com/karngyan/retry
Usage
Below is a basic example demonstrating how to use Retry to automatically retry a failing operation:

go
Copy code
package main

import (
	"errors"
	"fmt"
	"time"
	
	"github.com/karngyan/retry"
)

func main() {
	c := retry.Config{
		MaxRetries: 3,
		Interval:   1 * time.Second,
		Exponential: true,
		MaxElapsedTime: 5 * time.Second,
	}

	// A function that always fails
	f := func() error {
		return errors.New("always fails")
	}

	err := retry.Retry(f, c)
	if err != nil {
		fmt.Println("Operation failed after retries:", err)
	}
}
