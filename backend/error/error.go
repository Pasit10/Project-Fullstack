package error

import (
	"errors"
)

type TemplateError interface {
	Code() string
	Error() string
	HttpStatusCode() int
	ThaiDescription() string
	EnglishDescription() string
	Response() (int, HttpErrorResponse)
}

type HttpError struct {
	error string
}

type HttpErrorResponse struct {
	Code               string `json:"code"`
	Error              string `json:"error"`
	ThaiDescription    string `json:"thai_description"`
	EnglishDescription string `json:"english_description"`
}

func New(error string) *HttpError {
	return &HttpError{
		error: error,
	}
}

func InitError(err error) error {
	if templateError, ok := err.(TemplateError); ok {
		return templateError
	} else {
		return errors.New(err.Error())
	}
}

func (e *HttpError) Code() string {
	if errorCode, ok := CodesError[e]; ok {
		return errorCode
	} else {
		return InternalServerError.Error()
	}
}

func (e *HttpError) Error() string {
	return e.error
}

func (e *HttpError) ThaiDescription() string {
	if thaiDescription, ok := ThaiDescription[e]; ok {
		return thaiDescription
	} else {
		return InternalServerError.ThaiDescription()
	}
}

func (e *HttpError) EnglishDescription() string {
	if englishDescription, ok := EnglishDescription[e]; ok {
		return englishDescription
	} else {
		return InternalServerError.EnglishDescription()
	}
}

func (e *HttpError) HttpStatusCode() int {
	if httpStatusCode, ok := HttpStatusCodes[e]; ok {
		return httpStatusCode
	} else {
		return InternalServerError.HttpStatusCode()
	}
}

func (e *HttpError) Response() (int, HttpErrorResponse) {
	if httpStatusCode, ok := HttpStatusCodes[e]; ok {
		return httpStatusCode, HttpErrorResponse{
			Code:               e.Code(),
			Error:              e.Error(),
			ThaiDescription:    e.ThaiDescription(),
			EnglishDescription: e.EnglishDescription(),
		}
	} else {
		return InternalServerError.HttpStatusCode(), HttpErrorResponse{
			Code:               InternalServerError.Code(),
			Error:              InternalServerError.Error(),
			ThaiDescription:    InternalServerError.ThaiDescription(),
			EnglishDescription: InternalServerError.EnglishDescription(),
		}
	}
}

func FindError(errorName string) (err error) {
	for err, _ := range EnglishDescription {
		if err.Error() == errorName {
			return err
		}
	}
	return InternalServerError
}

var (
	InternalServerError      TemplateError = New("internal_server_error")
	DatabaseConnectedError   TemplateError = New("database_connected_error")
	BadrequestError          TemplateError = New("Bad_request_Error")
	UnauthorizedError        TemplateError = New("Unauthorized_Error")
	UsernotfoundError        TemplateError = New("User_not_found_Error")
	WrongUserOrPasswordError TemplateError = New("Wrong_user_or_password_Error")
	EmailAlreadyExistError   TemplateError = New("Email_already_exist_Error")
	MissingOrMalformedToken  TemplateError = New("Missing_or_malformed_token")
	InvalidOrExpiredToken    TemplateError = New("Invalid_or_expired_token")
	EmailInvaildFormatError  TemplateError = New("Email_invelid_format")
)

var HttpStatusCodes = map[error]int{
	InternalServerError:      500,
	DatabaseConnectedError:   500,
	BadrequestError:          400,
	UnauthorizedError:        401,
	UsernotfoundError:        404,
	WrongUserOrPasswordError: 401,
	EmailAlreadyExistError:   400,
	MissingOrMalformedToken:  401,
	InvalidOrExpiredToken:    401,
	EmailInvaildFormatError:  400,
}

var CodesError = map[error]string{
	BadrequestError:          "4000",
	EmailAlreadyExistError:   "4001",
	EmailInvaildFormatError:  "4002",
	UnauthorizedError:        "4010",
	WrongUserOrPasswordError: "4011",
	MissingOrMalformedToken:  "4012",
	InvalidOrExpiredToken:    "4013",
	InternalServerError:      "5000",
	DatabaseConnectedError:   "5001",
}

var ThaiDescription = map[error]string{
	InternalServerError:      "เกิดข้อผิดพลาดที่เซิฟเวอร์",
	UnauthorizedError:        "ไม่มีสิทธ์ในการร้องขอ",
	DatabaseConnectedError:   "ไม่สามารถเชื่อมต่อฐานข้อมูลได้",
	BadrequestError:          "คำขอไม่ถูกต้อง",
	UsernotfoundError:        "ไม่พบผู้ใช้",
	WrongUserOrPasswordError: "ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง",
	EmailAlreadyExistError:   "อีเมลนี้มีอยู่แล้ว",
	MissingOrMalformedToken:  "ขาดหรือรูปแบบโทเคนไม่ถูกต้อง",
	InvalidOrExpiredToken:    "โทเคนไม่ถูกต้องหรือหมดอายุ",
	EmailInvaildFormatError:  "รูปแบบ email ไม่ถูกต้อง",
}

var EnglishDescription = map[error]string{
	InternalServerError:      "There was an error on the server,",
	UnauthorizedError:        "Unauthorized request",
	DatabaseConnectedError:   "Unable to connect to database",
	BadrequestError:          "Invalid request",
	UsernotfoundError:        "User not found",
	WrongUserOrPasswordError: "Wrong username or password",
	EmailAlreadyExistError:   "Email already exist",
	MissingOrMalformedToken:  "Missing or malformed token",
	InvalidOrExpiredToken:    "Invalid or expired token",
	EmailInvaildFormatError:  "Email invalid format",
}

func GetErrorResponse(err error) (int, HttpErrorResponse) {
	var mapError bool
	var httpStatusCode int
	var errorCode string
	var thaiDescription string
	var englishDescription string

	if httpStatusCode, mapError = HttpStatusCodes[err]; mapError {
		if errorCode, mapError = CodesError[err]; mapError {
			if thaiDescription, mapError = ThaiDescription[err]; mapError {
				if englishDescription, mapError = EnglishDescription[err]; mapError {
					return httpStatusCode, HttpErrorResponse{
						Code:               errorCode,
						Error:              err.Error(),
						ThaiDescription:    thaiDescription,
						EnglishDescription: englishDescription,
					}
				}
			}
		}
	}
	return InternalServerError.HttpStatusCode(), HttpErrorResponse{
		Code:               InternalServerError.Code(),
		Error:              InternalServerError.Error(),
		ThaiDescription:    InternalServerError.ThaiDescription(),
		EnglishDescription: InternalServerError.EnglishDescription(),
	}
}
