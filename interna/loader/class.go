package loader

type Class struct {
	Name   string
	Inners []*Class
	Types  map[string]string
	Tags   map[string]int
}
