package xdg

import wl "deedles.dev/wl/client"

type Surface struct {
	obj     surfaceObject
	display *wl.Display
}
