package manager

type Manager struct {
	position string
	salary   float64
	address  string
}

func (a *Manager) SetSalary(salary float64) {
	a.salary = salary
}

func (a *Manager) SetPosition(position string) {
	a.position = position
}

func (a *Manager) SetAddress(address string) {
	a.address = address
}

func (a *Manager) GetSalary() float64 {
	return a.salary
}

func (a *Manager) GetPosition() string {
	return a.position
}

func (a *Manager) GetAddress() string {
	return a.address
}
