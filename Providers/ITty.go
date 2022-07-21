package Providers

import (
	"github.com/bhbosman/goCommsSshListener/common"
	"github.com/gdamore/tcell/v2"
)

type ITty interface {
	tcell.Tty
	common.ISetSize
}
