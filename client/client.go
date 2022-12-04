package xdg

//go:generate go run deedles.dev/wl/cmd/wlgen -client -out protocol.go -pkg xdg -prefix xdg_ -xml ../protocol/protocol.xml

type id[T interface{ ID() uint32 }] struct {
	obj T
}

func (i id[T]) ID() uint32 {
	return i.obj.ID()
}
