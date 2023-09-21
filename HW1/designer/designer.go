package designer

type Designer struct {
	position string
	salary   float64
	address  string
}

func (a *Designer) SetSalary(salary float64) {
	a.salary = salary
}

func (a *Designer) SetPosition(position string) {
	a.position = position
}

func (a *Designer) SetAddress(address string) {
	a.address = address
}

func (a *Designer) GetSalary() float64 {
	return a.salary
}

func (a *Designer) GetPosition() string {
	return a.position
}

func (a *Designer) GetAddress() string {
	return a.address
}
