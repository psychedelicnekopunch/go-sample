
package importFile

type Func struct {
	Vars FuncVars
}


type FuncVars struct {
	Test string
	Test2 string
}


func NewFunc() *Func {
	return &Func{
		Vars: FuncVars{
			Test: "--test",
			Test2: "--test2",
		},
	}
}

func Test() string {
	return "--success"
}

func (fn *Func) GetFuncVars() FuncVars {
	return fn.Vars
}
