package ripframework

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func getHttpVerb(request string) string {

	splitText := strings.Fields(request)

	if len(splitText) > 0 {
		return splitText[0]
	}

	return ""
}

func (f *Framework) handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	clientRequest, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error -> %v", err)
		return
	}

	request := string(buffer[:clientRequest])

	verb := getHttpVerb(request)

	fmt.Println(verb)
	response := "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\n\r\n{\"age\": 12}"

	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Error -> %v", err)
	}
}
