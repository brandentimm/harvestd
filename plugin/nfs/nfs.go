package nfs

import (
	"time"
)

type NFSReaper struct {
	Name   string
	Buffer []byte
}

func (reaper *NFSReaper) Read(p []byte) (int, error) {
	copy(reaper.Buffer, p)
	time.Sleep(7 * time.Second)
	return len(reaper.Buffer), nil
}

func Init() (*NFSReaper, error) {
	return &NFSReaper{Name: `nfs_reaper`}, nil
}
