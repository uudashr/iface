package mask

import (
	"context"
	"errors"
	"strings"
)

type Doer interface {
	Do()
	DoItAgain()
}

type DoerImpl struct{}

func (di DoerImpl) Do() {
}

func (di *DoerImpl) DoItAgain() {
}

type DoerV2Impl struct{}

func (d2i DoerV2Impl) Do() {
}

func (d2i *DoerV2Impl) DoItAgain() {
}

func NewDoer() Doer { // want "'NewDoer' function return 'Doer' interface at the 1st result, abstract a single concrete implementation of '\\*DoerImpl'"
	return &DoerImpl{}
}

func MultiDoer(spec string) Doer {
	if spec == "v2" {
		return &DoerV2Impl{}
	}

	if spec == "v1" {
		return &DoerImpl{}
	}

	return &DoerImpl{}
}

func NewDoerWithError() (Doer, error) {
	return nil, nil
}

func NewDoerWithOK() (Doer, bool) {
	Inspect(func(in interface{}) bool {
		return true
	})

	x := func() bool {
		return true
	}

	if x() {
		return nil, true
	}

	return nil, false
}

func NewDoerWithContext(ctx context.Context) (context.Context, Doer) {
	if ctx == nil {
		return newDoerWithContext()
	}

	return context.TODO(), nil
}

func newDoerWithContext() (context.Context, Doer) {
	return nil, nil
	// return newDoerImplWithContext()
}

func newDoerImplWithContext() (context.Context, *DoerImpl) {
	return context.TODO(), nil
}

func NewDoerImpl() *DoerImpl {
	return &DoerImpl{}
}

func NewDoerStruct() DoerImpl {
	return DoerImpl{}
}

func Inspect(fn func(in interface{}) bool) {
	_ = fn(DoerImpl{})
}

func NewContext() context.Context {
	return &myContextImpl{}
}

type myContextImpl struct {
	context.Context
}

func SplitCompleteVer(ver string) (major, minor, patch string, err error) {
	parts := strings.Split(ver, ".")
	if len(parts) != 3 {
		return "", "", "", errors.New("invalid complete version")
	}

	return parts[0], parts[1], parts[2], nil
}
