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
	line int
	text string
}
type Change_Req struct {
	session_id  string
	change_type int
	changes     []Change
}
