package responser

import (
	"fmt"
	"net"
	"ripframework/ripFramework/constants"
)

type Responser struct {
	status int
	conn   *net.Conn
}

func New(conn *net.Conn) *Responser {
	responser := &Responser{status: constants.HTTP_SUCCESS, conn: conn}

	fmt.Println(responser)
	return responser
}

func (r *Responser) Json(json interface{}) {
	response := "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\n\r\n{\"age\": 12}"

	fmt.Println(response)

}
