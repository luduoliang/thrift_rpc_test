package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"thrift_rpc_test/proto/sessions"
)

func main() {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket("localhost:5555")
	if err != nil {
		fmt.Errorf("NewTSocket failed. err: [%v]\n", err)
		return
	}

	transport, err = thrift.NewTBufferedTransportFactory(8192).GetTransport(transport)
	if err != nil {
		fmt.Errorf("NewTransport failed. err: [%v]\n", err)
		return
	}
	defer transport.Close()

	if err := transport.Open(); err != nil {
		fmt.Errorf("Transport.Open failed. err: [%v]\n", err)
		return
	}

	protocolFactory := thrift.NewTCompactProtocolFactory()
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	client := sessions.NewWaiterClient(thrift.NewTStandardClient(iprot, oprot))

	res, err := client.GetPddSessionsList(context.Background(), &sessions.RequestGetPddSessionsList{
		Page:    1,
		PerPage: 10,
	})
	if err != nil {
		fmt.Errorf("client echo failed. err: [%v]", err)
		return
	}

	fmt.Println(*res.PddSessions[0].Token)
}
