package main

import (
	"fmt"
	"net"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("%+v", err)
	}
 }

func run() error {
	fmt.Println("start tcp listen..")

	listen, err := net.Listen("tcp", "localhost:12345")

	if err != nil {
		return err
	}

	defer listen.Close()


	conn, err := listen.Accept()
	if err != nil {
		return err
	}

	defer conn.Close()

	fmt.Println(">>> start")

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(string(buf[:n]))
	}

	fmt.Println("<<< end")

	return nil
}
