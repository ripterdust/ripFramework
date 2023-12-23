package ripframework

import (
	"fmt"
	"net"
)

type Framework struct {
}

func New() *Framework {
	framework := &Framework{}

	return framework
}

func (f *Framework) Listen(port string) {

	listener, err := net.Listen("tcp", "localhost:"+port)

	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	defer listener.Close()

	fmt.Println("Server listening on port " + port)

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		go f.handleClient(conn)
	}
}
