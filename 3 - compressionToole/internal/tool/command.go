package tool

import (
	"io"
)

type Command interface {
	Execute(file io.Reader) (io.Reader, error)
}
