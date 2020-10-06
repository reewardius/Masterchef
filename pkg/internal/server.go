package internal

// ====================
//  IMPORTS
// ====================

import (
	"context"
)

// ====================
//  TYPES
// ====================

type Server interface {
	Listen(context.Context, context.CancelFunc)
}
