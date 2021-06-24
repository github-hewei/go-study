package p1

//定义接口
type It1 interface {
	f1() int    // 不可以在外部访问
	F2() string // 可以在外部访问
}

//定义结构体
type St1 struct {
	v1 int
	v2 string
}

//实现接口
func (s *St1) f1() int {
	return s.v1
}

func (s *St1) F2() string {
	return s.v2
}

//定义公开函数
func Run2() It1 {
	s := St1{111, "qqqqq"}
	return &s
}
