
package a


type A2 struct {}


func NewA2() *A2 {
	return new(A2)
}

func (a *A2) Get() string {
	return "success"
}
