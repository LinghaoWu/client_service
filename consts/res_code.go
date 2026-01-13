package consts

const (
	Success             = 200
	BadRequest          = 400
	Unauthorized        = 401
	Forbidden           = 403
	NotFound            = 404
	InternalServerError = 500
	ServiceUnavailable  = 503
)

var MsgFlags = map[int]string{
	Success:             "success",
	BadRequest:          "bad request",
	Unauthorized:        "unauthorized",
	Forbidden:           "forbidden",
	NotFound:            "not found",
	InternalServerError: "internal server error",
	ServiceUnavailable:  "service unavailable",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[InternalServerError]
}
