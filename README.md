[![license](https://img.shields.io/github/license/RediSearch/redisearch-go.svg)](https://github.com/RediSearch/redisearch-go)


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
	client := redisai.Connect("localhost:6379")

	// Set a tensor
	// AI.TENSORSET foo FLOAT 2 2 VALUES 1 2 3 4
	client.TensorSet( "foo" , TypeFloat, []int{2,2}, []int{1,2,3,4} )
	
	// Get a tensor content as a slice of values
	err, _, _, fooTensorValues, _ := client.TensorGetValues("foo" )

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fooTensorValues)
	// Output: [ 1 2 3 4 ]
}
```
