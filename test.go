package main

import (
	"time"
	"fmt"
	"net/http"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:

				fmt.Println("Tick at", t)

				resp, err := http.Get("http://localhost:9090")
				if err != nil {
					fmt.Println(err)
					return
				}
				defer resp.Body.Close()
				for true {

					bs := make([]byte, 1014)
					n, err := resp.Body.Read(bs)
					fmt.Println(string(bs[:n]))

					if n == 0 || err != nil{
						break
					}
				}
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true

	fmt.Println("Ticker stopped")

	//keepDoingSomething()

}


func keepDoingSomething() (bool, error) {
	timeout := time.After(5 * time.Second)
	tick := time.Tick(500 * time.Millisecond)
	// Keep trying until we're timed out or got a result or got an error
	for {
		select {
		// Got a timeout! fail with a timeout error
		case <-timeout:
			fmt.Println("Timeout")
			//return false, errors.New("timed out")
			// Got a tick, we should check on doSomething()
		case <-tick:

			fmt.Println("Tick at next")

			// doSomething() didn't work yet, but it didn't fail, so let's try again
			// this will exit up to the for loop
		}
	}
}