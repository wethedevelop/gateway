package main

import "github.com/wethedevelop/gateway/protocol"

func main() {
	Init()

	r := Route()
	protocol.Http3Run(r)
}
