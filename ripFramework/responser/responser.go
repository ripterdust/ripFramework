package responser

import (
	"encoding/json"
	"fmt"
	"net"
	"ripframework/ripFramework/constants"
)

type Responser struct {
	status int
	conn   net.Conn
}

func New(conn net.Conn) *Responser {
	responser := &Responser{status: constants.HTTP_SUCCESS, conn: conn}

	fmt.Println(responser)
	return responser
}

func (r *Responser) Json(jsonElement interface{}) {
	jsonBytes, err := json.Marshal(jsonElement)

	if err != nil {
		fmt.Println("Error -> ", err)
		return
	}

	response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/json\r\n\r\n%s", jsonBytes)

	_, err = r.conn.Write([]byte(response))

	if err != nil {
		fmt.Println("Error -> ", err)
	}

}
