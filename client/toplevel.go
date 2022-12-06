package xdg

import (
	wl "deedles.dev/wl/client"
	"deedles.dev/wl/wire"
)

type Toplevel struct {
	obj     toplevelObject
	display *wl.Display
}

func (tl *Toplevel) Object() wire.Object {
	return &tl.obj
}

func (tl *Toplevel) SetTitle(title string) {
	tl.display.Enqueue(tl.obj.SetTitle(title))
}

func (tl *Toplevel) Move(seat *wl.Seat, serial uint32) {
	tl.display.Enqueue(tl.obj.Move(seat.Object().ID(), serial))
}

type toplevelListener struct {
	tl *Toplevel
}

func (lis toplevelListener) Configure(width, height int32, states []byte) {
	// TODO
}

func (lis toplevelListener) Close() {
	// TODO
}

func (lis toplevelListener) ConfigureBounds(width, height int32) {
	// TODO
}

func (lis toplevelListener) WmCapabilities(capabilities []byte) {
	// TODO
}
