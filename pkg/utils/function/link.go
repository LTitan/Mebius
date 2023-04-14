package function

import (
	"container/list"
	"sync"
)

type ExecFunction func()
type ExecFunctionError func() error

type FunctionLinkInterface interface {
	Add(fc ExecFunction) FunctionLinkInterface
	AddE(fc ExecFunctionError) FunctionLinkInterface
	Do() // ignore error
	DoErr() error
}

type FunctionLink struct {
	fle *list.List
	mux *sync.Mutex
}

func (f *FunctionLink) Add(fc ExecFunction) FunctionLinkInterface {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.fle.PushBack(ExecFunctionError(func() error {
		fc()
		return nil
	}))
	return f
}

func (f *FunctionLink) AddE(fc ExecFunctionError) FunctionLinkInterface {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.fle.PushBack(ExecFunctionError(fc))
	return f
}

func NewFunctionLink(fcs ...ExecFunction) *FunctionLink {
	f := &FunctionLink{
		fle: list.New(),
		mux: &sync.Mutex{},
	}
	for _, fc := range fcs {
		f.Add(fc)
	}
	return f
}

func NewFunctionLinkErr(fces ...ExecFunctionError) *FunctionLink {
	f := &FunctionLink{
		fle: list.New(),
		mux: &sync.Mutex{},
	}
	for _, fc := range fces {
		f.AddE(fc)
	}
	return f
}

func (f *FunctionLink) Do() {
	f.mux.Lock()
	defer f.mux.Unlock()
	iter := f.fle.Front()
	for iter != nil {
		_ = iter.Value.(ExecFunctionError)()
		iter = iter.Next()
	}
}

func (f *FunctionLink) DoErr() error {
	f.mux.Lock()
	defer f.mux.Unlock()
	iter := f.fle.Front()
	for iter != nil {
		if err := iter.Value.(ExecFunctionError)(); err != nil {
			return err
		}
		iter = iter.Next()
	}
	return nil
}
