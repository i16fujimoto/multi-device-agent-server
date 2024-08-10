package cerror

import "net/http"

type Code int

func (c Code) String() string {
	return codeMap[c].message
}

const (
	OK Code = iota
	NotFound
	InvalidArgument
	Forbidden
	Unauthorized
	Internal
	AlreadyExists
	TooManyRequests
	SQLite
	Unknown
	Pagination
	NoRows
	InOpportuneTime
	EncodingJSON
	IO
	DoExternalHTTPRequest
	CreateExternalHTTPRequest
	TimeLoadLocation
	TimeParse
	StorageAPI

	ErrorCodeMax // to validate codeMap size
)

type codeDetail struct {
	message    string
	httpStatus int
}

var codeMap = map[Code]codeDetail{ //nolint: gochecknoglobals
	OK:                        {"OK", http.StatusOK},
	NotFound:                  {"Not found", http.StatusNotFound},
	InvalidArgument:           {"Invalid argument", http.StatusBadRequest},
	Forbidden:                 {"Forbidden", http.StatusForbidden},
	Unauthorized:              {"Unauthorized", http.StatusUnauthorized},
	Internal:                  {"Internal", http.StatusInternalServerError},
	AlreadyExists:             {"Already exists", http.StatusConflict},
	TooManyRequests:           {"Too many requests", http.StatusTooManyRequests},
	Unknown:                   {"Unknown", http.StatusInternalServerError},
	SQLite:                    {"SQLite", http.StatusInternalServerError},                       // データベース系のエラー
	Pagination:                {"Pagination", http.StatusBadRequest},                            // ページネーション系のエラー
	NoRows:                    {"No rows", http.StatusNotFound},                                 // データベース系のエラー
	InOpportuneTime:           {"In opportune_time", http.StatusBadRequest},                     // 適した時間でない時
	EncodingJSON:              {"Encoding JSON", http.StatusInternalServerError},                // Marshal/Unmarshal
	IO:                        {"IO", http.StatusInternalServerError},                           // IO系のエラー
	DoExternalHTTPRequest:     {"Do external HTTP Request", http.StatusInternalServerError},     // HTTPリクエスト系のエラー
	CreateExternalHTTPRequest: {"Create external HTTP Request", http.StatusInternalServerError}, // HTTPリクエスト系のエラー
	TimeLoadLocation:          {"Time", http.StatusInternalServerError},                         // タイムゾーンのロード時のエラー
	TimeParse:                 {"Time parse", http.StatusInternalServerError},                   // 文字列を時間に変換しようとした時のエラー
	StorageAPI:                {"Storage API", http.StatusInternalServerError},                  // ストレージAPIのエラー
}

func MapHTTPErrorToCode(httpStatusCode int) Code {
	switch httpStatusCode {
	case http.StatusOK:
		return OK
	case http.StatusNotFound:
		return NotFound
	case http.StatusBadRequest:
		return InvalidArgument
	case http.StatusForbidden:
		return Forbidden
	case http.StatusTooManyRequests:
		return TooManyRequests
	case http.StatusUnauthorized:
		return Unauthorized
	case http.StatusServiceUnavailable:
		return InOpportuneTime
	case http.StatusInternalServerError:
		return Internal
	default:
		return Unknown
	}
}
