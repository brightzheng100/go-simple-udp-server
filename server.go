package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

const (
	maxBufferSize = 1024
	timeout       = 120 * time.Second
)

/*
 * serve as the UPD Server with context support
 */
func serve(ctx context.Context, address string) (err error) {
	fmt.Printf("UDP Server listening on: \"%s\"\n", address)

	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		return
	}
	defer pc.Close()

	doneChan := make(chan error, 1)
	buffer := make([]byte, maxBufferSize)

	go func() {
		for {
			n, addr, err := pc.ReadFrom(buffer)
			if err != nil {
				doneChan <- err
				return
			}

			fmt.Printf("Packet Received: [%d] bytes from [%s]\n", n, addr.String())

			deadline := time.Now().Add(timeout)
			err = pc.SetWriteDeadline(deadline)
			if err != nil {
				doneChan <- err
				return
			}

			// Write the packets back to the client
			n, err = pc.WriteTo(buffer[:n], addr)
			if err != nil {
				doneChan <- err
				return
			}

			fmt.Printf("Packet Written: [%d] bytes to [%s]\n", n, addr.String())
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("cancelled")
		err = ctx.Err()
	case err = <-doneChan:
	}

	return
}
