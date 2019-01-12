package main

import (
	"context"
	"flag"
)

func main() {
	address := flag.String("address", ":8888",
		"The UDP Server listen address with port, e.g. `:8888` or `0.0.0.0:8888`.")
	flag.Parse()

	ctx, _ := context.WithCancel(context.Background())

	go serve(ctx, *address)

	<-ctx.Done()
}
