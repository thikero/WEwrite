package common

//return type
const C_TYPE_CLIENT = 1
const C_TYPE_CLIENT_RES = 2
const C_TYPE_SERVER = 3
const C_TYPE_SERVER_RES = 4
const C_TYPE_FAIL = 0

//results:
const SUCCESS = 0
const FAIL = -1
const EXCEPT = 1

type Change struct {
	Line int    `json:"Line"`
	Text string `json:"Text"`
}
type Change_Req struct {
	Session_id  string `json:"Session_id"`
	Change_type int    `json:"Change_type"`
	Changes     []Change
}
