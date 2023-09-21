package engineer

type Engineer struct {
	position string
	salary   float64
	address  string
}

func (a *Engineer) SetSalary(salary float64) {
	a.salary = salary
}

func (a *Engineer) SetPosition(position string) {
	a.position = position
}

func (a *Engineer) SetAddress(address string) {
	a.address = address
}

func (a *Engineer) GetSalary() float64 {
	return a.salary
}

func (a *Engineer) GetPosition() string {
	return a.position
}

func (a *Engineer) GetAddress() string {
	return a.address
}
