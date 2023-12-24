package common

import "encoding/json"

func JsonToChangeReq(request string) (data Change_Req, err error) {
	//change json data to change_req struct.
	err = json.Unmarshal([]byte(request), &data)
	return
}
