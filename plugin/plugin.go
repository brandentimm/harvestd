package plugin

import (
	"io"
	"time"
)

type Plugin struct {
	Name []byte
}

type Reaper interface {
	io.Reader
}

type Shipper interface {
	io.Writer
	Init() error
}

type CheckPacket struct {
	timestamp time.Time
	payload   []byte
}
