package internalapi

type WazeroOnly interface {
	gowasmOnly()
}

type WazeroOnlyType struct{}

func (WazeroOnlyType) gowasmOnly() {}
