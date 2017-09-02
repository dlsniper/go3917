package appCore

type Error struct{}

type Cleanup interface {
	Cleanup() *Error
}
