package main

import (
	"context"
	"machinerydemo/worker"
)

func main() {
	worker.SendHelloWorldTask(context.Background())
}
