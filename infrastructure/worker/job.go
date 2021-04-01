package worker

type Job interface {
	Apply() bool
	GetData() int
}

type NewsJobFilter struct {
	data int
}

func NewNewsJobFilter(data int) Job {
	return &NewsJobFilter{data: data}
}

func (n NewsJobFilter) Apply() bool {
	if n.data%2 == 0 {
		return true
	}
	return false
}

func (n NewsJobFilter) GetData() int {
	return n.data
}
