package ripframework

import (
	"fmt"
	"net"
	"ripframework/ripFramework/constants"
	"ripframework/ripFramework/util"
	"strconv"
)

func routeNotFound(verb string, uri string, conn net.Conn) {
	json := make(map[string]interface{})

	json["error"] = "Route not found"
	json["status"] = constants.HTTP_NOT_FOUND

	jsonString, err := util.MapToJSON(json)

	if err != nil {
		fmt.Println("Error -> ", err)
		return
	}

	response := fmt.Sprintf("HTTP/1.1 %s\r\nContent-Type: application/json\r\n\r\n%s", strconv.Itoa(constants.HTTP_NOT_FOUND), jsonString)

	_, err = conn.Write([]byte(response))

	if err != nil {
		fmt.Println("Error -> ", err)
		return
	}
}
