package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	anypb "google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	myproto "ranzhouol/go_study/microService/grpc/hellowrld/proto"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "")
)

func main() {
	flag.Parse()
	// 1. 连接到服务端, grpc.WithTransportCredentials(insecure.NewCredentials()) 表示使用不安全的连接，即不使用SSL
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	// 2. 建立连接，创建 server的客户端对象
	client := myproto.NewGreeterClient(conn)

	// 3. 执行rpc调用（这个方法在服务端来实现并返回结果）
	// 一元调用
	//sayHello(client)

	// 客户端流
	//sayHelloClientStream(client)

	// 双向流
	sayHelloTwoWayStream(client)
}

// 设置请求
func getHelloRequest() *myproto.HelloRequest {
	birthday := timestamppb.New(time.Now())
	any1, _ := anypb.New(birthday)
	in := &myproto.HelloRequest{
		Name:     "nick",
		Gender:   0,
		Age:      18,
		Birthday: birthday,
		Hobbys:   []string{"篮球", "羽毛球"},
		Addr: &myproto.Address{
			Province: "湖南",
			City:     "长沙",
		},
		Data: map[string]*anypb.Any{
			"a": any1,
		},
	}

	return in
}

// 一元调用
func sayHello(client myproto.GreeterClient) {
	ctx := context.Background()
	in := getHelloRequest()
	resp, err := client.SayHello(ctx, in)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(resp.Msg)
}

// 客户端流
func sayHelloClientStream(client myproto.GreeterClient) {
	ctx := context.Background()
	list := []*myproto.HelloRequest{
		getHelloRequest(), getHelloRequest(), getHelloRequest(),
	}
	// 创建客户端流对象
	stream, err := client.SayHelloClientStream(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	//发送流数据
	for _, in := range list {
		err := stream.Send(in)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	// 发送完毕后，关闭客户端，接受响应
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("client recv: %v\n", resp)
}

// 双向流
func sayHelloTwoWayStream(client myproto.GreeterClient) {
	ctx := context.Background()
	list := []*myproto.HelloRequest{
		getHelloRequest(), getHelloRequest(), getHelloRequest(),
	}
	// 创建客户端对象
	stream, err := client.SayHelloTwoWayStream(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	var done = make(chan struct{}, 0)
	// 启动协程来接受服务端的响应
	go func() {
		for {
			// 接受服务端的响应
			reply, err := stream.Recv()
			// 接受完毕，关闭通道
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Println(err)
				return
			}

			fmt.Printf("client recv: %v\n", reply.Msg)
		}
	}()

	//发送流数据
	for _, in := range list {
		err := stream.Send(in)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	// 发送完毕后，关闭客户端
	stream.CloseSend()

	// 未接收完服务端的响应时，会一直阻塞
	// 接收完后，关闭通道会发送过来一个零值，此时不再阻塞
	<-done
}
