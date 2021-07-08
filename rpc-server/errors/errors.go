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

type SurlErr struct {
	Msg string
}

func (s *SurlErr) Error() string {
	return s.Msg
}
func (s *SurlErr) New(msg string) {
	s.Msg = msg
}
