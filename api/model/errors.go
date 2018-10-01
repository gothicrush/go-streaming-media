package model

type ErrorModel struct {
	Info string `json:"error_info"`
	Code string `json:"error_code"`
}

type ErrorResponseModel struct {
	HttpStatusCode int
	Err ErrorModel
}

var (
    ErrorRequestBodyParseFailed = ErrorResponseModel {
    	HttpStatusCode: 400,
    	Err: ErrorModel {
            Info: "Request body is not correct",
            Code: "001",
    	},
    }

    ErrorNotAuthUser = ErrorResponseModel {
    	HttpStatusCode: 401,
    	Err: ErrorModel {
    		Info: "User authentication failed",
    		Code: "002",
    	},
    }
)