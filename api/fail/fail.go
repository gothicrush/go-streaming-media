package fail

type Fail struct {
	FailMessage string
	FailCode    string
}

type FailResponse struct {
	HttpCode int
	Fail  Fail
}

var (
	RequestBodyParseFail = FailResponse{
		HttpCode:400,
		Fail:Fail{
			FailMessage: "Request Body Parse Failed",
			FailCode: "001",
		},
	}

	UserAuthFail = FailResponse{
		HttpCode:403,
		Fail:Fail{
			FailMessage: "User Auth Failed",
			FailCode: "002",
		},
	}
)
