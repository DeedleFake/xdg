package xdg

import wl "deedles.dev/wl/client"

type WmBase struct {
	obj     wmBaseObject
	display *wl.Display
}

func IsWmBase(i wl.Interface) bool {
	return i.Is(wmBaseInterface, wmBaseVersion)
}

func BindWmBase(display *wl.Display, name uint32) *WmBase {
	wm := WmBase{display: display}
	wm.obj.listener = wmBaseListener{wm: &wm}
	display.AddObject(&wm.obj)

	registry := display.GetRegistry()
	registry.Bind(name, wmBaseInterface, wmBaseVersion, wm.obj.id)

	return &wm
}

func (wm *WmBase) GetXdgSurface(surface *wl.Surface) *Surface {
	panic("Not implemented.")
}

type wmBaseListener struct {
	wm *WmBase
}

func (lis wmBaseListener) Ping(serial uint32) {
	lis.wm.display.Enqueue(lis.wm.obj.Pong(serial))
}
