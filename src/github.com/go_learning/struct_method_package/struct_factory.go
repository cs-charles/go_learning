package struct_method_package

//工厂方法
func NewPeople(name string, age uint8) *People {
	if name == "" || age <= 0 {
		return nil
	}
	return &People{name: name, age: age}
}
