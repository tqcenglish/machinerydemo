package main

import "machinerydemo/worker"

func main() {
	taskWorker := worker.NewAsyncTaskWorker(0)
	taskWorker.Launch()
}
