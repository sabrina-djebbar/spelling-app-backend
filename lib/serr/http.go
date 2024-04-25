package serr

import (
	"errors"
	"net/http"
	"slices"
)

type ErrCode string

// error code constants.
const (
	ErrCodeBadRequest          ErrCode = "bad_request"
	ErrCodeBadResponse         ErrCode = "bad_response"
	ErrCodeForbidden           ErrCode = "forbidden"
	ErrCodeInternalService     ErrCode = "internal_service"
	ErrCodeBadGateway          ErrCode = "bad_gateway"
	ErrCodeNotFound            ErrCode = "not_found"
	ErrCodePreconditionFailed  ErrCode = "precondition_failed"
	ErrCodeTimeout             ErrCode = "timeout"
	ErrCodeUnauthorized        ErrCode = "unauthorized"
	ErrCodeConflict            ErrCode = "conflict"
	ErrCodeStatusNotAcceptable ErrCode = "status_not_acceptable"
	ErrCodeUnknown             ErrCode = "unknown"

	// todo move those error codes in a better place

	ErrCodeServiceUnavailable      ErrCode = "service_unavailable"
	ErrCodeMotoNotEnabled          ErrCode = "error_moto_not_enabled"
	ErrCodeCOFNotEnabled           ErrCode = "error_card_on_file_not_enabled"
	ErrCodeAccountInactive         ErrCode = "account_inactive"
	ErrCodeAccountInvalid          ErrCode = "account_invalid"
	ErrCodeBillingSurchargeInvalid ErrCode = "billing_surcharge_invalid"
	ErrMissingRequiredField        ErrCode = "missing_required_field"
)

var httpStatusMap = map[ErrCode]int{
	ErrCodeBadRequest:              http.StatusBadRequest,          // 400
	ErrCodeBadResponse:             http.StatusNotAcceptable,       // 406
	ErrCodeForbidden:               http.StatusForbidden,           // 403
	ErrCodeInternalService:         http.StatusInternalServerError, // 500
	ErrCodeBadGateway:              http.StatusBadGateway,          // 502
	ErrCodeServiceUnavailable:      http.StatusServiceUnavailable,  // 503
	ErrCodeNotFound:                http.StatusNotFound,            // 404
	ErrCodePreconditionFailed:      http.StatusPreconditionFailed,  // 412
	ErrCodeTimeout:                 http.StatusGatewayTimeout,      // 504
	ErrCodeUnauthorized:            http.StatusUnauthorized,        // 401
	ErrCodeConflict:                http.StatusConflict,            // 409
	ErrCodeStatusNotAcceptable:     http.StatusNotAcceptable,       // 406
	ErrCodeMotoNotEnabled:          http.StatusBadRequest,          // 400
	ErrCodeCOFNotEnabled:           http.StatusBadRequest,          // 400
	ErrCodeAccountInactive:         http.StatusBadRequest,          // 400
	ErrCodeAccountInvalid:          http.StatusBadRequest,          // 400
	ErrCodeBillingSurchargeInvalid: http.StatusBadRequest,          // 400
	ErrMissingRequiredField:        http.StatusBadRequest,          // 400
}

func HttpStatus(err error) int {
	if err == nil {
		return 0
	}

	e := Error{}
	if errors.As(err, &e) {
		if code, ok := httpStatusMap[e.Code]; ok {
			return code
		}
		return 500
	}

	return 500
}

// HasCode returns true iff the topmost werr.Error in the wrapped stack has an
// error code in the specified set of codes. If there is no wrapped werr.Error
// then the error is treated as having code ErrCodeInternalService.
func HasCode(err error, codes ...ErrCode) bool {
	if err == nil || len(codes) == 0 {
		return false
	}

	if e := (Error{}); errors.As(err, &e) && e.Code != "" {
		return slices.Contains(codes, e.Code)
	}

	// no serr.Error so defaults to internal service error
	return slices.Contains(codes, ErrCodeInternalService)
}

var (
	ErrBadRequest          = New(string(ErrCodeBadRequest), WithCode(ErrCodeBadRequest))
	ErrBadResponse         = New(string(ErrCodeBadResponse), WithCode(ErrCodeBadResponse))
	ErrForbidden           = New(string(ErrCodeForbidden), WithCode(ErrCodeForbidden))
	ErrInternalService     = New(string(ErrCodeInternalService), WithCode(ErrCodeInternalService))
	ErrNotFound            = New(string(ErrCodeNotFound), WithCode(ErrCodeNotFound))
	ErrPreconditionFailed  = New(string(ErrCodePreconditionFailed), WithCode(ErrCodePreconditionFailed))
	ErrTimeout             = New(string(ErrCodeTimeout), WithCode(ErrCodeTimeout))
	ErrUnauthorized        = New(string(ErrCodeUnauthorized), WithCode(ErrCodeUnauthorized))
	ErrConflict            = New(string(ErrCodeConflict), WithCode(ErrCodeConflict))
	ErrStatusNotAcceptable = New(string(ErrCodeStatusNotAcceptable), WithCode(ErrCodeStatusNotAcceptable))
	ErrServiceUnavailable  = New(string(ErrCodeServiceUnavailable), WithCode(ErrCodeServiceUnavailable))
	ErrUnknown             = New(string(ErrCodeUnknown), WithCode(ErrCodeUnknown))
)
