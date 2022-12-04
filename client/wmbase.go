package xdg

import wl "deedles.dev/wl/client"

type WmBase struct {
	id[wmBaseObject]
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
	s := Surface{display: wm.display}
	s.obj.listener = surfaceListener{surface: &s}
	wm.display.AddObject(&s.obj)

	wm.display.Enqueue(wm.obj.GetXdgSurface(s.obj.id, surface.ID()))

	return &s
}

type wmBaseListener struct {
	wm *WmBase
}

func (lis wmBaseListener) Ping(serial uint32) {
	lis.wm.display.Enqueue(lis.wm.obj.Pong(serial))
}
