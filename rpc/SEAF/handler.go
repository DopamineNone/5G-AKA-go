package main

import (
	"context"
)

// ProtocolServiceImpl implements the last service interface defined in the IDL.
type ProtocolServiceImpl struct{}

// Authenticate implements the ProtocolServiceImpl interface.
func (s *ProtocolServiceImpl) Authenticate(ctx context.Context, data string) (resp string, err error) {
	// TODO: Your code here...
	return
}
