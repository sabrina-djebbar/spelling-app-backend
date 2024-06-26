package shttp

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/serr"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/validator"
	"log"
	"net/http"
	"reflect"
	"syscall"
)

type responseType string

const (
	JSON responseType = "json"
	XML  responseType = "xml"
)

var (
	errType = reflect.TypeOf((*error)(nil)).Elem()
	ctxType = reflect.TypeOf((*context.Context)(nil)).Elem()
)
var l log.Logger

func (srv *server) RegisterHandler(path string, handler interface{}) {
	// Gets the handlers function type to then get the arguments from and to later use in the .Call
	handlerType := reflect.TypeOf(handler)
	handlerValue := reflect.ValueOf(handler)
	handlerInArgNum := handlerType.NumIn()
	handlerOutArgNum := handlerType.NumOut()

	if handlerInArgNum != 2 || (handlerOutArgNum != 1 && handlerOutArgNum != 2) {

		l.Panicf("Invalid number of input, output params \n Inputs: %d\n Outputs %d\n ", handlerInArgNum, handlerOutArgNum)
	}

	// Ensure that the first parameter is assignable to context.Context
	if !handlerType.In(0).AssignableTo(ctxType) {
		panic("Invalid handler arg. Parameter 1 must be assignable to context.Context")
	}

	// Gets the request type from the context
	requestType := handlerType.In(1)

	// Ensure that the second parameter is assignable to a struct
	if requestType.Kind() != reflect.Struct {
		panic("Invalid handler arg. Parameter 2 must be assignable to struct")
	}

	var outType reflect.Kind

	if handlerOutArgNum == 2 {
		// Ensure that our out type is matching of a pointer or slice
		// We don't allow anything else as our body should be a json response
		outType = handlerType.Out(0).Kind()
		if outType != reflect.Ptr && outType != reflect.Slice && outType != reflect.Map {
			panic("Invalid handler arg. Out 1 must be of type pointer")
		}

		// Ensure that our secondary out is of type error
		// This is to allow graceful serr without having to use recover
		if !handlerType.Out(1).AssignableTo(errType) {
			panic("Invalid handler arg. Out 2 must be assignable to error")
		}
	} else {
		// If we have a single output it's required to be of type error
		if !handlerType.Out(0).AssignableTo(errType) {
			panic("Invalid handler arg. Out 1 must be assignable to error")
		}
	}

	srv.router.POST(path,
		// Perform auth validation before passing to next handler
		// srv.ValidateAuth(requireAuth),
		func(ctx *gin.Context) {
			// Create a new request type and then get the value of it
			requestPointer := reflect.New(requestType)
			requestValue := requestPointer.Interface()

			// Set default request options
			var err error
			reqType := JSON

			// Assign the request body to the struct in the handler
			if ctx.ContentType() == gin.MIMEJSON {
				err = ctx.ShouldBindJSON(requestValue)
			} else if ctx.ContentType() == gin.MIMEPOSTForm {
				reqType = XML
				err = ctx.ShouldBind(requestValue)
			}

			// Reject the request if binding fails
			if err != nil {
				// Check if validation error
				vErr := validator.ValidationError{}
				if ok := errors.As(err, &vErr); ok {
					returnResponse(ctx, http.StatusBadRequest, serr.Error{
						Code:    serr.ErrCodeBadRequest,
						Message: vErr.Message,
						Details: vErr.Details,
					}, reqType)

					return
				}
				returnResponse(ctx, http.StatusBadRequest, serr.Error{
					Code:    serr.ErrCodeBadRequest,
					Message: "invalid_body",
					Details: err.Error(),
				}, reqType)

			}

			// Create a new args collection for passing into the Call
			args := make([]reflect.Value, handlerInArgNum)
			args[0] = reflect.ValueOf(createContext(ctx))
			reqValueElem := reflect.ValueOf(requestValue).Elem()
			args[1] = reqValueElem

			out := handlerValue.Call(args)

			// Get the response out depending on the number of args
			var resp interface{}
			if handlerOutArgNum == 2 {
				// If error response then return 500
				if !(out[1].IsNil()) {
					err = (out[1].Interface()).(error)
				}

				resp = out[0].Interface()

			} else {
				if !out[0].IsNil() {
					err = (out[0].Interface()).(error)
				}
			}

			if err != nil {
				status := http.StatusInternalServerError

				// If the error type is assignable to our error type then just return it
				responseErr := serr.Error{}
				if ok := errors.As(err, &responseErr); ok {
					status = serr.HttpStatus(responseErr)
					returnResponse(ctx, status, responseErr, reqType)
					return
				}

				returnResponse(ctx, status, serr.Error{
					Code:    serr.ErrCodeInternalService,
					Message: err.Error(),
				}, reqType)

				return
			}

			// If no response body then simply return 204
			if resp == nil {
				ctx.Status(http.StatusNoContent)

				return
			}

			returnResponse(ctx, http.StatusOK, resp, reqType)
		})
}

// RegisterGinHandler will register a handler using the gin.HandlerFunc implementation
func (srv *server) RegisterGinHandler(method string, path string, handler gin.HandlerFunc) {
	srv.router.Handle(method, path, handler)
}

func returnResponse(ctx *gin.Context, status int, out interface{}, respType responseType) {
	defer func() {
		if e := recover(); e != nil {
			if err, ok := e.(error); ok {
				// handle "connection reset by peer" and "broken pipe" serr
				// https://gosamples.dev/connection-reset-by-peer/ https://gosamples.dev/broken-pipe/
				if errors.Is(err, syscall.ECONNRESET) || errors.Is(err, syscall.EPIPE) {
					log.Panic("connection error while writing the response", err)
					return
				}
			}

			log.Panic("panic occurred while serializing the response", e)
		}
	}()

	if respType == JSON {
		if e, ok := out.(error); ok {
			_ = ctx.Error(e)
		}
		ctx.JSON(status, out)
	} else if respType == XML {
		if e, ok := out.(error); ok {
			_ = ctx.Error(e)
		}
		ctx.XML(status, out)
	} else {
		panic(serr.New("unknown_response_type", serr.WithCode(serr.ErrCodeInternalService)))
	}
}

// RegisterNativeHandler will register a native shttp.HandlerFunc under gin
func (srv *server) RegisterNativeHandler(method string, path string, handler http.HandlerFunc) {
	srv.router.Handle(method, path, func(ctx *gin.Context) {
		r := ctx.Request
		w := ctx.Writer

		handler(w, r)
	})
}
