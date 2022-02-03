package main

import (
	"simpleWebApp/logger"
	"simpleWebApp/webserver"
	"sync"
)

const WEB_PORT = "8080"
const LISTEN_PORT = "8007"

func main() {
	// create a WaitGroup
	wg := new(sync.WaitGroup)
	wg.Add(2)

	//start request logger
	go func() {
		logger.ListenForRequest(LISTEN_PORT)
		wg.Done()
	}()

	//start webserver
	go func() {
		webserver.Webserver(WEB_PORT)
		wg.Done()
	}()

	wg.Wait()
}
