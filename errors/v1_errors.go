package errors

import (
	"google.golang.org/grpc/status"
)

// TODO: generate porcelain errors from protobuf error definitions in spec.proto
// rough process: range through all messages, check if the message has the
// .ModelOptions.Error flag, and if so, generate a struct for it.

func New(err error) error {
	if s, ok := status.FromError(err); ok {
		for _, details := range s.Details() {
			switch details.(type) {
				// TODO: convert protobuf error to porcelain error
			}
		}
	}
	return err
}