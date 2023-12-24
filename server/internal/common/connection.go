package common

import (
	"bufio"
	"fmt"
	"net"
)

func ReadFromConn(conn net.Conn) (data_res string, status int) {
	reader := bufio.NewReader(conn)
	data_res, err := reader.ReadString('\n')
	if err != nil {
		data_res = fmt.Sprintf("Error reading from connection : %v", err)
		status = FAIL
		return
	}
	status = SUCCESS
	return
}
