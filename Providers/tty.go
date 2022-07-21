package Providers

import (
	"context"
	"github.com/bhbosman/goCommsSshListener/common"
	"io"
)

type tty struct {
	CancelCtx         context.Context
	CancelFunc        context.CancelFunc
	SshChannel        common.IChannel
	rwc               io.ReadWriteCloser
	cols              int
	rows              int
	callBackForResize func()
}

func newTty(rwc io.ReadWriteCloser) ITty {
	return &tty{
		rwc: rwc,
	}
}

func (self *tty) SetSize(cols int, rows int) error {
	self.cols = cols
	self.rows = rows
	if self.callBackForResize != nil {
		self.callBackForResize()
	}
	return nil
}
func (self *tty) Start() error {
	return nil
}

func (self *tty) Stop() error {
	return nil
}

func (self *tty) Drain() error {
	return self.rwc.Close()
}

func (self *tty) NotifyResize(cb func()) {
	self.callBackForResize = cb
}

func (self *tty) WindowSize() (width int, height int, err error) {
	return self.cols, self.rows, nil
}

func (self *tty) Read(p []byte) (n int, err error) {
	return self.rwc.Read(p)
}

func (self *tty) Write(p []byte) (n int, err error) {
	return self.rwc.Write(p)
}

func (self *tty) Close() error {
	return self.rwc.Close()
}
