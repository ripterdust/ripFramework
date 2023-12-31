package ripframework

import (
	"log"
	"net"
	"ripframework/ripFramework/requester"
	"ripframework/ripFramework/responser"
	"ripframework/ripFramework/util"
)

func (f *Framework) handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	clientRequest, err := conn.Read(buffer)

	if err != nil {
		log.Printf("Error -> %v", err)
		return
	}

	request := string(buffer[:clientRequest])
	verb := util.GetHttpVerb(request)
	uri := util.GetUri(request)

	handler := f.handlers[verb][uri]

	if handler == nil {
		routeNotFound(verb, uri, conn)

		return
	}

	responser := responser.New(conn)
	requester := requester.New()

	handler(requester, responser)
}
