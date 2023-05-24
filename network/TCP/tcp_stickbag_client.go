package main

import (
	"fmt"
	"net"
	"ranzhouol/go_study/network/TCP/proto"
	"strconv"
)

/*
	tcp粘包
*/

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	// 直接发送数据
	//for i := 0; i < 20; i++ {
	//	msg := `Hello, Hello. How are you?`
	//	conn.Write([]byte(msg))
	//}

	// 进行封包操作，避免粘包
	for i := 0; i < 20; i++ {
		msg := strconv.Itoa(i) + ` Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
