package main

import "xsis/assignment-test/configuration"

func main() {
	configuration := configuration.Init()

	configuration.Start()

	defer configuration.Stop()
}
