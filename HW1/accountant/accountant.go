package accountant

type Accountant struct {
	position string
	salary   float64
	address  string
}

func (a *Accountant) SetSalary(salary float64) {
	a.salary = salary
}

func (a *Accountant) SetPosition(position string) {
	a.position = position
}

func (a *Accountant) SetAddress(address string) {
	a.address = address
}

func (a *Accountant) GetSalary() float64 {
	return a.salary
}

func (a *Accountant) GetPosition() string {
	return a.position
}

func (a *Accountant) GetAddress() string {
	return a.address
}
