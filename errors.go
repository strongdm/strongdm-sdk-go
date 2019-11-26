package sdm

// Error is a generic wrapper that indicates an unknown internal error in the SDK.
type Error struct {
	Wrapped error
}

func (e *Error) Error() string {
	return e.Wrapped.Error()
}

func (e *Error) Unwrap() error {
	return e.Wrapped
}

// RPCError is a generic RPC error indicating something went wrong at the
// transport layer. Use Code() and Unwrap() to inspect the actual failed
// condition.
type RPCError interface {
	// Code returns the gRPC error code
	Code() int
	error
}

type DeadlineExceededError struct {
	Wrapped error
}

func (e *DeadlineExceededError) Error() string {
	return "deadline exceeded"
}

func (e *DeadlineExceededError) Unwrap() error {
	return e.Wrapped
}

func (e *DeadlineExceededError) Code() int {
	return 4
}

type ContextCanceledError struct {
	Wrapped error
}

func (e *ContextCanceledError) Error() string {
	return "context canceled"
}

func (e *ContextCanceledError) Unwrap() error {
	return e.Wrapped
}

func (e *ContextCanceledError) Code() int {
	return 1
}
// AlreadyExistsError is used when an entity already exists in the system
type AlreadyExistsError struct {
	Message string
	Entity string
}

func (e *AlreadyExistsError) Error() string {
	return e.Message
}

// NotFoundError is used when an entity does not exist in the system
type NotFoundError struct {
	Message string
	Entity string
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
