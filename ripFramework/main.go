package ripframework

import (
	"fmt"
	"net"
	"ripframework/ripFramework/constants"
)

type Framework struct {
	handlers map[string]map[string]handleFunction
	status   int
}

/*
	@instance Framework
	@return Framework instance with httpMethod handlers
	@create handlers hash map
*/

func New() *Framework {
	framework := &Framework{
		handlers: make(map[string]map[string]handleFunction),
		status:   constants.HTTP_SUCCESS,
	}

	for _, httpVerb := range constants.VERBS {
		framework.handlers[httpVerb] = make(map[string]handleFunction)
	}

	return framework
}

/*
@param port string -> port were server is going to be listening
*/
func (f *Framework) Listen(port string) {
	listener, err := net.Listen("tcp", "localhost:"+port)

	if err != nil {
		fmt.Println("Error ->", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server listening on port " + port)

	fmt.Println(f.handlers)
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error -> ", err)
			continue
		}

		go f.handleClient(conn)
	}
}
