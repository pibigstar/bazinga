package db

type Tracer interface {
	i()
	Before()
	After()
}

type DBTrace struct {
}

func (*DBTrace) i() {}

func (*DBTrace) Before() {

}

func (*DBTrace) After() {

}
