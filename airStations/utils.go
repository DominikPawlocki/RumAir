package airStations

import (
	"fmt"
	"time"
)

func DoHttpCallWithConsoleDots(fn func() []byte) (bytesRead []byte) {
	ticker := time.NewTicker(300 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				fmt.Printf(".")
			}
		}
	}()

	bytesRead = fn()

	ticker.Stop()
	done <- true
	return
}

func DeserializeWithConsoleDots(fn func([]byte, interface{}) error, bytes []byte, result interface{}) error {
	ticker := time.NewTicker(50 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				fmt.Printf(".")
			}
		}
	}()

	err := fn(bytes, result)

	ticker.Stop()
	done <- true
	return err
}
