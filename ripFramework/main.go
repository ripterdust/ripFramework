package ripframework

import (
	"fmt"
	"net"
	"ripframework/ripFramework/verbs"
)

type handleFunction func(string) interface{}

type Framework struct {
	handlers map[string]map[string]handleFunction
}

func New() *Framework {
	framework := &Framework{
		handlers: make(map[string]map[string]handleFunction),
	}

	for _, httpVerb := range verbs.VERBS {
		framework.handlers[httpVerb] = make(map[string]handleFunction)
	}

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
