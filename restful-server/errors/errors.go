package errors

const (
	// data status
	Exist   = 2000
	UnExist = 2001
	// stauts
	ERROR   = 1002
	SUCCESS = 1000
	Fail    = 1001
)

var CodeMap = map[int]string{
	Exist:   "exit",
	UnExist: "unexit",
	SUCCESS: "success",
	Fail:    "fail",
	ERROR:   "error",
}
