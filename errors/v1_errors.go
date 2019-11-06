package errors
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