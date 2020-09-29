package domain

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type ServerError struct {
	Status   int        `json:"-"`
	Code     string     `json:"code"`
	GRPCCode codes.Code `json:"grpccode"`
	Message  string     `json:"message"`
}

func (e ServerError) Error() string {
	return e.Message
}

func (e ServerError) Map() map[string]interface{} {
	return map[string]interface{}{
		"status":   e.Status,
		"code":     e.Code,
		"grpccode": e.GRPCCode,
		"message":  e.Message,
	}
}

// Custom erros
var (
	// 400
	ErrInvalidInput  = ServerError{Code: "400001", Message: "One of the request inputs is not valid.", Status: http.StatusBadRequest, GRPCCode: codes.InvalidArgument}
	ErrInvalidAmount = ServerError{Code: "400002", Message: "This amount is not allow.", Status: http.StatusBadRequest, GRPCCode: codes.InvalidArgument}

	// 403
	ErrNotAllowed                                  = ServerError{Code: "403001", Message: "The request is understood, but it has been refused or access is not allowed.", Status: http.StatusForbidden, GRPCCode: codes.PermissionDenied}
	ErrInsufficientAccountPermissionsWithOperation = ServerError{Code: "403002", Message: "The account being accessed does not have sufficient permissions to execute this operation.", Status: http.StatusForbidden, GRPCCode: codes.PermissionDenied}
	ErrUnauthorized                                = ServerError{Code: "403003", Message: "Unauthorized operation", Status: http.StatusForbidden, GRPCCode: codes.PermissionDenied}

	// 404
	ErrOrderNotFound     = ServerError{Code: "404001", Message: "Can't find order", Status: http.StatusNotFound, GRPCCode: codes.NotFound}
	ErrCreateOrderFailed = ServerError{Code: "404002", Message: "Create order failed", Status: http.StatusNotFound, GRPCCode: codes.NotFound}

	// 500
	ErrInternalError = ServerError{Code: "500001", Message: "The server encountered an internal error. Please retry the request.", Status: http.StatusInternalServerError, GRPCCode: codes.Internal}
	ErrDbOperation   = ServerError{Code: "500002", Message: "Operation failed", Status: http.StatusInternalServerError, GRPCCode: codes.Internal}
	ErrWallet        = ServerError{Code: "500003", Message: "Internal service unavaliabe", Status: http.StatusInternalServerError, GRPCCode: codes.Internal}
	ErrBB            = ServerError{Code: "500004", Message: "Internal service unavaliabe", Status: http.StatusInternalServerError, GRPCCode: codes.Internal}

	// 503
	ErrGAM   = ServerError{Code: "503001", Message: "channel unavailable", Status: http.StatusServiceUnavailable, GRPCCode: codes.Unavailable}
	ErrATP   = ServerError{Code: "503002", Message: "channel unavailable", Status: http.StatusServiceUnavailable, GRPCCode: codes.Unavailable}
	ErrShan6 = ServerError{Code: "503003", Message: "channel unavailable", Status: http.StatusServiceUnavailable, GRPCCode: codes.Unavailable}
)
