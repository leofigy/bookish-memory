package disk

type Space struct {
	Total int64
	Used  int64
}

type Unit interface {
	GetSpace() Space
}
