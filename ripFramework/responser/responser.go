package responser

import (
	"encoding/json"
	"fmt"
	"net"
	"ripframework/ripFramework/constants"
	"strconv"
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

	status := strconv.Itoa(r.status)
	response := fmt.Sprintf("HTTP/1.1 %s \r\nContent-Type: application/json\r\n\r\n%s", status, string(jsonBytes))

	r.Status(constants.HTTP_SUCCESS)

	_, err = r.conn.Write([]byte(response))

	if err != nil {
		fmt.Println("Error -> ", err)
	}
}

func (r *Responser) Status(status int) {
	r.status = status
}
