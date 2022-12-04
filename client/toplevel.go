package xdg

import wl "deedles.dev/wl/client"

type Toplevel struct {
	id[toplevelObject]
	display *wl.Display
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
