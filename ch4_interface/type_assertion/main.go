package main

type Inter interface {
	Ping()
	Pang()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (St) Ping() {
	println("ping")
}

func (*St) Pang() {
	println("pang")
}

func main() {
	st := &St{"andes"}
	var i interface{} = st

	// 判断i绑定的实例是否实现了接口类型Inter
	o := i.(Inter)
	o.Ping()
	o.Pang()
}
