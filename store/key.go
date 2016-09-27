package store

type Key interface {
	Name() string
	Deleted() bool
	Value() interface{}
	SetValue(interface{}) error
}

type BaseKey struct {
	deleted bool
	name    string
}

func (bk *BaseKey) Name() string {
	return bk.name
}

func (bk *BaseKey) Deleted() bool {
	return bk.deleted
}
