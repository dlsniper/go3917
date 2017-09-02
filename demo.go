package go3917

import (
	"github.com/dlsniper/go3917/appCore"
	"github.com/dlsniper/go3917/database"
)

type appServices struct {
	database database.Service
}

func (s *appServices) Cleanup() *appCore.Error {
	return s.database.Cleanup()
}
