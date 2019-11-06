package errors

import (
	"google.golang.org/grpc/status"

	v1 "github.com/strongdm/strongdm-sdk-go/internal/v1"
)
// AlreadyExistsError is used when an entity already exists in the system
type AlreadyExistsError struct {
	Message string
	Entities []string
}

func (e *AlreadyExistsError) Error() string {
	return e.Message
}

// NotFoundError is used when an entity does not exist in the system
type NotFoundError struct {
	Message string
	Entities []string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

// BadRequestError identifies a bad request sent by the client
type BadRequestError struct {
	Message string
}

func (e *BadRequestError) Error() string {
	return e.Message
}

// AuthenticationError is used to specify an authentication failure condition
type AuthenticationError struct {
	Message string
}

func (e *AuthenticationError) Error() string {
	return e.Message
}

// PermissionError is used to specify a permissions violation
type PermissionError struct {
	Message string
	Permission string
	Entities []string
}

func (e *PermissionError) Error() string {
	return e.Message
}

// InternalError is used to specify an internal system error
type InternalError struct {
	Message string
}

func (e *InternalError) Error() string {
	return e.Message
}

// RateLimitError is used for rate limit excess condition
type RateLimitError struct {
	Message string
}

func (e *RateLimitError) Error() string {
	return e.Message
}

// RPCError is a generic RPC error, use Unwrap to inspect the actual failed condition
type RPCError struct {
	Wrapped error
}

func (e *RPCError) Error() string {
	return e.Wrapped.Error()
}

func (e *RPCError) Unwrap() error {
	return e.Wrapped
}

func New(err error) error {
	if s, ok := status.FromError(err); ok {
		for _, details := range s.Details() {
			switch d := details.(type) {
			// AlreadyExistsError is used when an entity already exists in the system
			case *v1.AlreadyExistsError:
				e := &AlreadyExistsError{}
				e.Message = s.Message()
				e.Entities = d.Entities
				return e

			// NotFoundError is used when an entity does not exist in the system
			case *v1.NotFoundError:
				e := &NotFoundError{}
				e.Message = s.Message()
				e.Entities = d.Entities
				return e

			// BadRequestError identifies a bad request sent by the client
			case *v1.BadRequestError:
				e := &BadRequestError{}
				e.Message = s.Message()
				return e

			// AuthenticationError is used to specify an authentication failure condition
			case *v1.AuthenticationError:
				e := &AuthenticationError{}
				e.Message = s.Message()
				return e

			// PermissionError is used to specify a permissions violation
			case *v1.PermissionError:
				e := &PermissionError{}
				e.Message = s.Message()
				e.Permission = d.Permission
				e.Entities = d.Entities
				return e

			// InternalError is used to specify an internal system error
			case *v1.InternalError:
				e := &InternalError{}
				e.Message = s.Message()
				return e

			// RateLimitError is used for rate limit excess condition
			case *v1.RateLimitError:
				e := &RateLimitError{}
				e.Message = s.Message()
				return e

			}
		}
	}
	return &RPCError{Wrapped: err}
}