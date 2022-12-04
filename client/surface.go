package xdg

import wl "deedles.dev/wl/client"

type Surface struct {
	id[surfaceObject]
	display *wl.Display
}

type surfaceListener struct {
	surface *Surface
}

func (lis surfaceListener) Configure(serial uint32) {
	// TODO?
	lis.surface.display.Enqueue(lis.surface.obj.AckConfigure(serial))
}
