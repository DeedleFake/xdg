package xdg

import (
	wl "deedles.dev/wl/client"
	"deedles.dev/wl/wire"
)

type Surface struct {
	Configure func()

	obj     surfaceObject
	display *wl.Display
}

func (s *Surface) Object() wire.Object {
	return &s.obj
}

func (s *Surface) GetToplevel() *Toplevel {
	tl := Toplevel{display: s.display}
	tl.obj.listener = toplevelListener{tl: &tl}
	s.display.AddObject(&tl)

	s.display.Enqueue(s.obj.GetToplevel(tl.obj.id))

	return &tl
}

type surfaceListener struct {
	surface *Surface
}

func (lis surfaceListener) Configure(serial uint32) {
	lis.surface.display.Enqueue(lis.surface.obj.AckConfigure(serial))

	if lis.surface.Configure != nil {
		lis.surface.Configure()
	}
}
