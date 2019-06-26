package model

type Employee struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func (e Employee) GetId() int {
	return e.Id
}

func (e Employee) GetName() string {
	return e.Name
}

func (e Employee) GetRole() string {
	return e.Role
}

func (e Employee) SetId(id int) Employee {
	e.Id = id
	return e
}

func (e Employee) SetName(name string) Employee {
	e.Name = name
	return e
}

func (e Employee) SetRole(role string) Employee {
	e.Role = role
	return e
}
