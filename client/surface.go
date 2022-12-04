package xdg

import wl "deedles.dev/wl/client"

type Surface struct {
	Configure func()

	id[surfaceObject]
	display *wl.Display
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
