package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/HouzuoGuo/hello-grpc/hzglexamplesvc"
	"google.golang.org/grpc"
)

const ListenAddr = "127.0.0.1:61212"

type RPCServiceImpl struct {
	hzglexamplesvc.UnimplementedDataTypeExperimentsServer
}

// GetSingleResponse handles a request consisting of a single message input, and responds by a single message as well.
func (svc *RPCServiceImpl) GetSingleResponse(_ context.Context, req *hzglexamplesvc.RequestOfManyTypes) (resp *hzglexamplesvc.ResponseOfManyTypes, err error) {
	log.Printf("GetSingleResponse received request: %s", req)
	respMap := req.GetAMap()
	if respMap == nil {
		return nil, fmt.Errorf("AMap must not be nil")
	}
	respMap["regards-from-server"] = true
	resp = &hzglexamplesvc.ResponseOfManyTypes{
		AString: req.GetAString() + "-regards-from-server",
		AnInt:   req.GetAnInt() + 9999,
		AFloat:  req.GetAFloat() + 9999.99,
		ADouble: req.GetADouble() + 9999.99,
		ABool:   !req.GetABool(),
		ABlob:   bytes.Repeat(req.ABlob, 3),
		AnEnum:  hzglexamplesvc.ResponseOfManyTypes_Third,
		AnArray: append(req.AnArray, 99, 99),
		AMap:    respMap,
	}
	return
}

// GetSingleResponse handles a request consisting of a single message input, and responds by a stream of three messages.
func (svc *RPCServiceImpl) GetStreamResponse(req *hzglexamplesvc.RequestOfManyTypes, stream hzglexamplesvc.DataTypeExperiments_GetStreamResponseServer) error {
	// Respond with AnInt + 100, +200, and +300
	for i := 0; i < 3; i++ {
		if err := stream.Send(&hzglexamplesvc.ResponseOfManyTypes{AnInt: req.GetAnInt() + int32(i+1)*100}); err != nil {
			return err
		}
	}
	return nil
}

// GetSingleResponse handles a request consisting of a stream of messages, and responds by a single message.
func (svc *RPCServiceImpl) GetSingleResponseFromStream(stream hzglexamplesvc.DataTypeExperiments_GetSingleResponseFromStreamServer) error {
	// Respond with a calculated sum of AnInt from all of the messages presented in request
	var sumAnInt int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&hzglexamplesvc.ResponseOfManyTypes{
				AnInt: sumAnInt,
			})
		} else if err != nil {
			return err
		}
		sumAnInt += req.GetAnInt()
	}
}

// GetSingleResponse handles a request consisting of a stream of messages, and responds by a stream of messages.
func (svc *RPCServiceImpl) GetStreamResponseFromStream(stream hzglexamplesvc.DataTypeExperiments_GetStreamResponseFromStreamServer) error {
	// Respond with requested AnInt +100 for each message presented in request
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		if err := stream.Send(&hzglexamplesvc.ResponseOfManyTypes{AnInt: req.AnInt + 100}); err != nil {
			return err
		}
	}
}

func main() {
	serverReady := make(chan struct{}, 1)
	go func() {
		// Start gRPC server
		listener, err := net.Listen("tcp", ListenAddr)
		if err != nil {
			log.Panic(err)
			return
		}
		svcImpl := &RPCServiceImpl{}
		server := grpc.NewServer()
		defer server.Stop()
		hzglexamplesvc.RegisterDataTypeExperimentsServer(server, svcImpl)
		serverReady <- struct{}{}
		if err := server.Serve(listener); err != nil {
			log.Panic(err)
			return
		}
	}()

	// Start gRPC client
	clientConn, err := grpc.Dial(ListenAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Panic(err)
		return
	}
	defer clientConn.Close()
	client := hzglexamplesvc.NewDataTypeExperimentsClient(clientConn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Send a single-message request and receive a single-message response
	singleMsgResp, err := client.GetSingleResponse(ctx, &hzglexamplesvc.RequestOfManyTypes{
		AString: "adam",
		AnInt:   100,
		AFloat:  100.100,
		ADouble: 100.001,
		ABool:   true,
		ABlob:   []byte{0, 1, 2, 3},
		AnEnum:  hzglexamplesvc.RequestOfManyTypes_Second,
		AnArray: []int64{},
		AMap:    map[string]bool{"client": true},
	})
	if err != nil {
		log.Panic(err)
		return
	}
	log.Printf("GetSingleResponse response: %s", singleMsgResp)

	// Send a single-message request and receive a stream of messages in response
	streamResp, err := client.GetStreamResponse(ctx, &hzglexamplesvc.RequestOfManyTypes{
		AnInt: 100,
	})
	for {
		resp, err := streamResp.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Panic(err)
			return
		}
		// There shall be three responses, AnInt: 100+100, 100+200, 100+300
		log.Printf("GetStreamResponse item: %s", resp)
	}

	// Send a stream of messages in request and receive a single message in response
	streamReq, err := client.GetSingleResponseFromStream(ctx)
	if err != nil {
		log.Panic(err)
		return
	}
	for i := 0; i < 3; i++ {
		if err := streamReq.Send(&hzglexamplesvc.RequestOfManyTypes{AnInt: int32(i+1) * 100}); err != nil {
			log.Panic(err)
			return
		}
	}
	sumResp, err := streamReq.CloseAndRecv()
	if err != nil {
		log.Panic(err)
		return
	}
	// There shall be the sum of three integers from the request (100 + 200 + 300)
	log.Printf("GetSingleResponseFromStream response: %s", sumResp)

	// Send a stream of messages in request and receive a stream of messages in response
	streamReqWithStreamResp, err := client.GetStreamResponseFromStream(ctx)
	if err != nil {
		log.Panic(err)
		return
	}
	for i := 0; i < 3; i++ {
		if err := streamReqWithStreamResp.Send(&hzglexamplesvc.RequestOfManyTypes{AnInt: int32(i+1) * 1000}); err != nil {
			if err != nil {
				log.Panic(err)
				return
			}
		}
		resp, err := streamReqWithStreamResp.Recv()
		if err != nil {
			log.Panic(err)
			return
		}
		// There shall be three responses: AnInt: 1000+100, 2000+100, 3000+100
		log.Printf("GetStreamResponseFromStream item: %s", resp)
	}
	if err := streamReqWithStreamResp.CloseSend(); err != nil {
		log.Panic(err)
		return
	}
}
