package worker_lib

type Worker func(interface{})
type Data interface{}

func DoWork(data_chan chan Data, worker_chan chan Worker) {
	e := make(chan bool)    // worker end work signal
	quit := make(chan bool) // data chan close; all workes end
	n := 0
	// Hear create gorutine to decrement working process
	go func() {
		loop := true
		for loop { // when data are generated
			select {
			case <-e:
				n--
			case <-quit: // data chain is close
				loop = false
			}
		}
		for n != 0 { // when data chain is close
			<-e
			n--
		}
		quit <- true // all workers finish job

	}()
	for data := range data_chan {
		worker := <-worker_chan
		n++
		go func(worker_chan chan Worker, worker Worker, data Data) {
			worker(data)
			worker_chan <- worker
			e <- true
		}(worker_chan, worker, data)
	}
	quit <- true // data chan is closed
	<-quit       // all workers finish job
}
