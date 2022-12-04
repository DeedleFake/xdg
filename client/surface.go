package xdg

import wl "deedles.dev/wl/client"

type Surface struct {
	Configure func()

	id[surfaceObject]
	display *wl.Display
}

func (s *Surface) GetToplevel() *Toplevel {
	tl := Toplevel{display: s.display}
	tl.obj.listener = toplevelListener{tl: &tl}
	s.display.AddObject(&tl.obj)

	s.display.Enqueue(s.obj.GetToplevel(tl.obj.id))

	return &tl
}

type surfaceListener struct {
	surface *Surface
}

func (lis surfaceListener) Configure(serial uint32) {
	if lis.surface.Configure != nil {
		lis.surface.Configure()
	}

	lis.surface.display.Enqueue(lis.surface.obj.AckConfigure(serial))
}
