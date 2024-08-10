package cerror

import "fmt"

type Option func(*Error)

func WithCode(code Code) Option {
	return func(e *Error) {
		e.code = code
	}
}

func WithInternalCode() Option {
	return func(e *Error) {
		e.code = Internal
	}
}

func WithInvalidArgumentCode() Option {
	return func(e *Error) {
		e.code = InvalidArgument
	}
}

func WithNotFoundCode() Option {
	return func(e *Error) {
		e.code = NotFound
	}
}

func WithAlreadyExistsCode() Option {
	return func(e *Error) {
		e.code = AlreadyExists
	}
}

func WithSQLiteCode() Option {
	return func(e *Error) {
		e.code = SQLite
	}
}

func WithUnauthorizedCode() Option {
	return func(e *Error) {
		e.code = Unauthorized
	}
}

func WithForbiddenCode() Option {
	return func(e *Error) {
		e.code = Forbidden
	}
}

func WithInOpportuneTimeCode() Option {
	return func(e *Error) {
		e.code = InOpportuneTime
	}
}

func WithNoRowsCode() Option {
	return func(e *Error) {
		e.code = NoRows
	}
}

func WithClientMsg(format string, args ...any) Option {
	return func(e *Error) {
		e.clientMsg = fmt.Sprintf(format, args...)
	}
}

func WithEncodingJSONCode() Option {
	return func(e *Error) {
		e.code = EncodingJSON
	}
}

func WithIOCode() Option {
	return func(e *Error) {
		e.code = IO
	}
}

func WithDoExternalHTTPRequestCode() Option {
	return func(e *Error) {
		e.code = DoExternalHTTPRequest
	}
}

func WithCreateExternalHTTPRequestCode() Option {
	return func(e *Error) {
		e.code = CreateExternalHTTPRequest
	}
}

func WithTimeParseCode() Option {
	return func(e *Error) {
		e.code = TimeParse
	}
}

func WithTimeLoadLocationCode() Option {
	return func(e *Error) {
		e.code = TimeLoadLocation
	}
}

func WithStorageAPICode() Option {
	return func(e *Error) {
		e.code = StorageAPI
	}
}
