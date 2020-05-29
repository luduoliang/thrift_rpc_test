package main

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"thrift_rpc_test/config"
	"thrift_rpc_test/controllers"
	"thrift_rpc_test/proto/sessions"
)

func main() {
	appPort := config.Default("APP_PORT", "5555")
	transport, err := thrift.NewTServerSocket(":" + appPort)
	if err != nil {
		panic(err)
	}

	processor := sessions.NewWaiterProcessor(&controllers.Server{})
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		thrift.NewTBufferedTransportFactory(8192),
		thrift.NewTCompactProtocolFactory(),
	)

	fmt.Println("Now listening on: http://127.0.0.1:" + appPort)
	fmt.Println("Application started. Press CTRL+C to shut down.")

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
