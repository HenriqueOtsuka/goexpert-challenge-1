package handler

import "github.com/gin-gonic/gin"

type ContentType string

const (
	CONTENT_TYPE_JSON ContentType = "application/json"
	CONTENT_TYPE_XML  ContentType = "application/xml"
	CONTENT_TYPE_M3U8 ContentType = "application/vnd.apple.mpegurl"
	CONTENT_TYPE_TEXT ContentType = "text/plain"
)

// ResponseController is an interface that defines methods for handling HTTP responses.
type ResponseController interface {
	// IsAbort returns a boolean value indicating whether the response should be aborted.
	IsAbort() bool

	// GetStatusCode returns the HTTP status code of the response.
	GetStatusCode() int

	// GetErrors returns a slice of errors associated with the response.
	GetErrors() []error

	// GetResponse returns the response data.
	GetResponse() interface{}

	// SetStatusCode sets the HTTP status code of the response.
	SetStatusCode(statusCode int)

	// SetErrors sets the errors associated with the response.
	SetErrors(errors []error)

	// AddError adds an error to the response.
	AddError(err error)

	// SetResponse sets the response data.
	SetResponse(response interface{})

	// SetResult sets the HTTP status code and response data.
	SetResult(status int, response interface{})

	// Write writes the response to the given Gin context.
	Write(ctx *gin.Context)

	// SetContentType sets the content type of the response.
	SetContentType(contentType ContentType)
}

type baseResponseController struct {
	contentType ContentType
	statusCode  int
	err         []error
	response    interface{}
}

func (b *baseResponseController) IsAbort() bool {
	return b.statusCode >= 400 || len(b.err) > 0
}

func (b *baseResponseController) GetStatusCode() int {
	return b.statusCode
}

func (b *baseResponseController) GetErrors() []error {
	return b.err
}

func (b *baseResponseController) GetResponse() interface{} {
	return b.response
}

func (b *baseResponseController) SetStatusCode(statusCode int) {
	b.statusCode = statusCode
}

func (b *baseResponseController) SetErrors(errors []error) {
	b.err = append(b.err, errors...)
}

func (b *baseResponseController) AddError(err error) {
	b.err = append(b.err, err)
}

func (b *baseResponseController) SetResponse(response interface{}) {
	b.response = response
}

func (b *baseResponseController) SetResult(status int, response interface{}) {
	b.statusCode = status
	b.response = response
}

func (b *baseResponseController) Write(ctx *gin.Context) {
	ctx.Header("Content-Type", string(b.contentType))
	switch b.contentType {
	case CONTENT_TYPE_M3U8:
		ctx.Writer.Write([]byte(b.response.(string)))
	case CONTENT_TYPE_TEXT:
		ctx.String(b.statusCode, b.response.(string))
	default:
		ctx.JSON(b.statusCode, b.response)
	}
}

func (b *baseResponseController) SetContentType(contentType ContentType) {
	b.contentType = contentType
}

func NewJsonResponseController() ResponseController {
	return &baseResponseController{
		contentType: CONTENT_TYPE_JSON,
	}
}
