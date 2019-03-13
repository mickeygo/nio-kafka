package main

import _ "net/http/pprof"

// Version app version
const Version string = "1.0.0"

func main() {
	poll()

	select {}
}
