package xdg

import (
	"deedles.dev/wl/bin"
	wl "deedles.dev/wl/client"
	"deedles.dev/wl/wire"
)

type Toplevel struct {
	Configure       func(w, h int32, states []ToplevelState)
	Close           func()
	ConfigureBounds func(width, height int32)

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

func (tl *Toplevel) SetMaximized(v bool) {
	if v {
		tl.display.Enqueue(tl.obj.SetMaximized())
		return
	}
	tl.display.Enqueue(tl.obj.UnsetMaximized())
}

func (tl *Toplevel) SetMinimized() {
	tl.display.Enqueue(tl.obj.SetMinimized())
}

type toplevelListener struct {
	tl *Toplevel
}

func (lis toplevelListener) Configure(width, height int32, states []byte) {
	if lis.tl.Configure != nil {
		s := make([]ToplevelState, 0, len(states)/4)
		for i := 4; i <= len(states); i += 4 {
			s = append(s, ToplevelState(bin.Value[uint32](*(*[4]byte)(states[i-4 : i]))))
		}
		lis.tl.Configure(width, height, s)
	}
}

func (lis toplevelListener) Close() {
	if lis.tl.Close != nil {
		lis.tl.Close()
	}
}

func (lis toplevelListener) ConfigureBounds(width, height int32) {
	if lis.tl.ConfigureBounds != nil {
		lis.tl.ConfigureBounds(width, height)
	}
}

func (lis toplevelListener) WmCapabilities(capabilities []byte) {
	// TODO
}
