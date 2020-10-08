package main

//参考 https://qiita.com/tutuz/items/ab1fd3c0ee6fa01e08b6

import (
	"bufio"
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

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	fmt.Println("<<< end")


	return nil
}
