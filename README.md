[![license](https://img.shields.io/github/license/RediSearch/redisearch-go.svg)](https://github.com/filipecosta90/redisai-go)
[![CircleCI](https://circleci.com/gh/filipecosta90/redisai-go/tree/master.svg?style=svg)](https://circleci.com/gh/filipecosta90/redisai-go/tree/master)
[![GitHub issues](https://img.shields.io/github/release/filipecosta90/redisai-go.svg)](https://github.com/filipecosta90/redisai-go/releases/latest)
[![Codecov](https://codecov.io/gh/filipecosta90/redisai-go/branch/master/graph/badge.svg)](https://codecov.io/gh/filipecosta90/redisai-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/filipecosta90/redisai-go)](https://goreportcard.com/report/github.com/filipecosta90/redisai-go)
[![GoDoc](https://godoc.org/github.com/filipecosta90/redisai-go?status.svg)](https://godoc.org/github.com/filipecosta90/redisai-go)

# RedisAI Go Client

Go client for [RedisAI](http://redisai.io), based on redigo.

# Installing 

```sh
go get github.com/filipecosta90/redisai-go/redisai
```

# Usage Examples
See the [examples](./examples) folder for further feature samples:

## Simple Client 
[(sample code here)](./examples/redisai_simple_client)

```go
package main

import (
	"fmt"
	"github.com/filipecosta90/redisai-go/redisai"
	"log"
)

func main() {

	// Create a client.
	client := redisai.Connect("redis://localhost:6379", nil)

	// Set a tensor
	// AI.TENSORSET foo FLOAT 2 2 VALUES 1.1 2.2 3.3 4.4
	_ = client.TensorSet("foo", redisai.TypeFloat, []int{2, 2}, []float32{1.1, 2.2, 3.3, 4.4})

	// Get a tensor content as a slice of values
	// dt DataType, shape []int, data interface{}, err error
	// AI.TENSORGET foo VALUES
	_, _, fooTensorValues, err := client.TensorGetValues("foo")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fooTensorValues)
	// Output: [1.1 2.2 3.3 4.4]
}
```

## Pipelined Client 
[(sample code here)](./examples/redisai_pipelined_client)
```go
package main

import (
	"fmt"
	"github.com/filipecosta90/redisai-go/redisai"
	"log"
)

func main() {

	// Create a client.
	client := redisai.Connect("redis://localhost:6379", nil)

	// Enable pipeline of commands on the client.
	client.Pipeline(3)

	// Set a tensor
	// AI.TENSORSET foo FLOAT 2 2 VALUES 1.1 2.2 3.3 4.4
	err := client.TensorSet("foo1", redisai.TypeFloat, []int{2, 2}, []float32{1.1, 2.2, 3.3, 4.4})
	if err != nil {
		log.Fatal(err)
	}
	// AI.TENSORSET foo2 FLOAT 1" 1 VALUES 1.1
	err = client.TensorSet("foo2", redisai.TypeFloat, []int{1, 1}, []float32{1.1})
	if err != nil {
		log.Fatal(err)
	}
	// AI.TENSORGET foo2 META
	_, err = client.TensorGet("foo2", redisai.TensorContentTypeMeta)
	if err != nil {
		log.Fatal(err)
	}
	// Ignore the AI.TENSORSET Reply
	_, err = client.Receive()
	if err != nil {
		log.Fatal(err)
	}
	// Ignore the AI.TENSORSET Reply
	_, err = client.Receive()
	if err != nil {
		log.Fatal(err)
	}
	foo2TensorMeta, err := client.Receive()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(foo2TensorMeta)
	// Output: [FLOAT [1 1]]
}
```
