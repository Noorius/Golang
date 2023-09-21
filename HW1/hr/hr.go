package hr

type Hr struct {
	position string
	salary   float64
	address  string
}

func (a *Hr) SetSalary(salary float64) {
	a.salary = salary
}

func (a *Hr) SetPosition(position string) {
	a.position = position
}

func (a *Hr) SetAddress(address string) {
	a.address = address
}

func (a *Hr) GetSalary() float64 {
	return a.salary
}

func (a *Hr) GetPosition() string {
	return a.position
}

func (a *Hr) GetAddress() string {
	return a.address
}
