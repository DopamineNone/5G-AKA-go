package main

import (
	"context"
)

// ProtocolServiceImpl implements the last service interface defined in the IDL.
type ProtocolServiceImpl struct{}

// HandleConnection implements the ProtocolServiceImpl interface.
func (s *ProtocolServiceImpl) HandleConnection(ctx context.Context, data string) (err error) {
	// TODO: Your code here...

	return
}

// Authenticate implements the ProtocolServiceImpl interface.
func (s *ProtocolServiceImpl) Authenticate(ctx context.Context) (err error) {
	// TODO: Your code here...
	return
}
