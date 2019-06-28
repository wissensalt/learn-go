package dao

type Scaffolding interface {
	findAll() ([]interface{}, error)
	findById() (interface{}, error)
	insert(interface{}) (interface{}, error)
	update(interface{}) (interface{}, error)
	delete(interface{}) (interface{}, error)
}
