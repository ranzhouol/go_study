package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	myproto "ranzhouol/go_study/microService/grpc/hellowrld/proto"
)

var (
	port = flag.Int("port", 50051, "")
)

// 1. 创建 gRPC Server 对象, 内嵌默认server
type server struct {
	myproto.UnimplementedGreeterServer
}

// 重写方法
// 一元调用
func (server) SayHello(ctx context.Context, in *myproto.HelloRequest) (*myproto.HelloReply, error) {
	log.Printf("server recv: %v\n", in)
	return &myproto.HelloReply{
		Msg: "hello client",
	}, nil
}

// 客户端流
func (server) SayHelloClientStream(stream myproto.Greeter_SayHelloClientStreamServer) error {
	i := 0
	// 持续接受客户端请求
	for {
		req, err := stream.Recv()
		// 输入结束
		if err == io.EOF {
			return stream.SendAndClose(&myproto.HelloReply{Msg: fmt.Sprintf("total resv count: %d", i)})
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Printf("server recv : %v\n", req)
		i++
	}
}

// 服务端流，未实现
func (server) SayHelloServerStream(in *myproto.HelloRequest, stream myproto.Greeter_SayHelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStream not implemented")
}

// 双向流
func (server) SayHelloTwoWayStream(stream myproto.Greeter_SayHelloTwoWayStreamServer) error {
	var count = 0
	for {
		// 接收客户端请求
		req, err := stream.Recv()
		// 输入结束
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Printf("server recv : %v\n", req)

		// 返回响应
		stream.Send(&myproto.HelloReply{
			Msg: fmt.Sprintf("%dth request ok\n", count),
		})

		count++
	}

	return nil
}

func main() {
	flag.Parse()

	// 监听端口
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
		return
	}

	//2. 创建grpc服务
	s := grpc.NewServer()

	//3. 在grpc服务端中注册我们的服务
	myproto.RegisterGreeterServer(s, &server{})

	//4. 启动服务
	log.Printf("server listening at %s\n", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
		return
	}
}
