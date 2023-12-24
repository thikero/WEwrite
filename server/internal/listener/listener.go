package listener

import (
	"fmt"
	"log"
	"net"

	"github.com/thikero/WEwrite/internal/common"
	"github.com/thikero/WEwrite/internal/server_update"
)

func Listen() error {
	l, err := net.Listen("tcp", ":2020")
	if err != nil {
		return err
	}
	defer l.Close()
	for {
		con, err := l.Accept()
		if err != nil {
			log.Printf("Error Accepting connection : %v", err)
			continue
		}
		go func(c net.Conn) {

			data, status := common.ReadFromConn(con)
			if status == common.SUCCESS {
				c_r, err := common.JsonToChangeReq(data)
				if err != nil {
					status = common.FAIL
					data = fmt.Sprintf("Error changing json format. :%v", err)
					goto err_response
				}

				//validate connection
				err = server_update.CheckSessionAuth(c_r)
				if err != nil {
					status = common.FAIL
					data = fmt.Sprintf("Request is not authenticated. :%v", err)
					goto err_response
				}
				//get Changes from connection
				//save changes to local file.
			}
		err_response:
			if status == common.FAIL {
				log.Printf("Error reading data from connection: %v", data)
				//response failure to client
				c.Close()
				return
			}
			//response back to sender client
			//close connection
			//flood changes to other clients
		}(con)

	}
}
