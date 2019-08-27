[![license](https://img.shields.io/github/license/RediSearch/redisearch-go.svg)](https://github.com/itamarhaber/redisai-go)
[![CircleCI](https://circleci.com/gh/itamarhaber/redisai-go/tree/master.svg?style=svg)](https://circleci.com/gh/itamarhaber/redisai-go/tree/master)
[![GitHub issues](https://img.shields.io/github/release/itamarhaber/redisai-go.svg)](https://github.com/itamarhaber/redisai-go/releases/latest)
[![Codecov](https://codecov.io/gh/itamarhaber/redisai-go/branch/master/graph/badge.svg)](https://codecov.io/gh/itamarhaber/redisai-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/itamarhaber/redisai-go)](https://goreportcard.com/report/github.com/itamarhaber/redisai-go)
[![GoDoc](https://godoc.org/github.com/itamarhaber/redisai-go?status.svg)](https://godoc.org/github.com/itamarhaber/redisai-go)

# RedisAI Go Client

Go client for [RedisAI](http://redisai.io), based on redigo.

# Installing 

```sh
go get github.com/itamarhaber/redisai-go/redisai
```

# Usage Example

```go

import (
	"fmt"
	"log"
	"github.com/itamarhaber/redisai-go/redisai"
)

func ExampleClient() {

	// Create a client. 
	client := redisai.Connect("localhost:6379", nil )

	// Set a tensor
	// AI.TENSORSET foo FLOAT 2 2 VALUES 1 2 3 4
	client.TensorSet( "foo" , TypeFloat, []int{2,2}, []float32{1,2,3,4} )
	
	// Get a tensor content as a slice of values
	// returned in format
	// dt DataType, shape []int, data interface{}, err error
	_, _, fooTensorValues, err := client.TensorGetValues( "foo" )

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fooTensorValues)
	// Output: [ 1 2 3 4 ]
}
```
