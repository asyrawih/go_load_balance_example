package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestWithCancel(t *testing.T) {
	fmt.Println("Total Routine", runtime.NumGoroutine())
	parent := context.Background()
	ctx , cancel := context.WithTimeout(parent , 100 * time.Second)
	defer cancel()
	dest := CreateContext(ctx)

	person := DataContext(ctx)

	fmt.Println("Total Routine Running" , runtime.NumGoroutine())

	for n := range dest {
		fmt.Println("Counter " , n)
		// Out Kan Data Dari Channel Routine
		fmt.Println(<-person)
	}


	time.Sleep(2 * time.Second)
	fmt.Println("Total Routine " , runtime.NumGoroutine())
}

type Person struct {
	Id int
	Name string
}

func CreateContext(ctx context.Context) chan int {
	dest := make(chan int)
	go func() {
		defer close(dest)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				dest <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return dest
}

func DataContext(ctx context.Context) chan interface{}{
	dest := make(chan interface{})
	go func() {
		defer close(dest)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				dest <- Person{
					Id:   counter,
					Name: "Test",
				}
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return dest
}
