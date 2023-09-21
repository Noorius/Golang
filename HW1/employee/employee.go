package employee

type Employee interface {
	GetPosition() string
	GetSalary() float64
	GetAddress() string

	SetPosition(position string)
	SetSalary(salary float64)
	SetAddress(address string)
}
