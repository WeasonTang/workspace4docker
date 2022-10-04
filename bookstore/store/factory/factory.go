// store/factory/factory.go

package factory

import (
	"sync"

	"github.com/WeasonTang/workspace4docker/module/bookstore/store"
)

var (
	providersMu sync.RWMutex
	prividers   = make(map[string]store.Store)
)
