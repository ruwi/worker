package main

import "fmt"
import "time"
import . "./worker_lib"

func main() {
	worker_chan := make(chan Worker, 9)
	generate_workers(worker_chan, 9)
	data_chan := make(chan Data)
	end := make(chan bool)
	go func() { generate_data(data_chan); end <- true }()
	go func() { DoWork(data_chan, worker_chan); end <- true }()
	<-end
	close(data_chan) // must be after DoWork
	<-end
}

func generate_data(data_chan chan Data) {
	for i := 0; i < 20; i++ {
		data_chan <- i
	}
}

func generate_workers(worker_chan chan Worker, n int) {
	for i := 0; i < n; i++ {
		worker_chan <- func(i interface{}) {
			func(i int) {
				time.Sleep(1 * time.Second)
				fmt.Println(i)
			}(i.(int))
		}
	}
}
