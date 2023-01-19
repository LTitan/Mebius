package function

type ExecFunction func()
type ExecFunctionError func() error

type FunctionLinkInterface interface {
	Add(fc ExecFunction) FunctionLinkInterface
	AddE(fc ExecFunctionError) FunctionLinkInterface
	Do() // ignore error
	DoErr() error
}

type FunctionLink struct {
	fle *[]ExecFunctionError
	fl  *[]ExecFunction
}

func (f *FunctionLink) Add(fc ExecFunction) FunctionLinkInterface {
	(*f.fl) = append((*f.fl), fc)
	return f
}

func (f *FunctionLink) AddE(fc ExecFunctionError) FunctionLinkInterface {
	(*f.fle) = append((*f.fle), fc)
	return f
}

func NewFunctionLink(fcs ...ExecFunction) *FunctionLink {
	return &FunctionLink{
		fl: &fcs,
	}
}

func NewFunctionLinkErr(fces ...ExecFunctionError) *FunctionLink {
	return &FunctionLink{
		fle: &fces,
	}
}

func (f *FunctionLink) Do() {
	if f.fl != nil {
		for _, fc := range *f.fl {
			fc()
		}
	}
	if f.fle != nil {
		for _, fc := range *f.fle {
			fc()
		}
	}
}

func (f *FunctionLink) DoErr() error {
	if f.fl != nil {
		for _, fc := range *f.fl {
			fc()
		}
	}
	if f.fle != nil {
		for _, fc := range *f.fle {
			if err := fc(); err != nil {
				return err
			}
		}
	}
	return nil
}
