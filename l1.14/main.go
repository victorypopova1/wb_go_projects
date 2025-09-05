package main

import (
	"fmt"
	"reflect"
)

func detectType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}:
		return "chan interface{}"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	case chan bool:
		return "chan bool"
	default:
		return reflect.TypeOf(v).String()
	}
}

func main() {
	variables := []interface{}{
		42,
		"hello",
		true,
		make(chan int),
		make(chan string),
		make(chan bool),
		make(chan interface{}),
		3.14,
		[]int{1, 2, 3},
	}

	for _, v := range variables {
		fmt.Printf("Значение: %v, Тип: %s\n", v, detectType(v))
	}
}
