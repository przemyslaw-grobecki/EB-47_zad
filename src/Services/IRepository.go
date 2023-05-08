package Services

type IRepository[Entity any] interface {
	GetAll() []Entity
	GetById(uint) Entity
	GetWithFilter() []Entity
	Add(Entity) Entity
	Delete(uint) Entity
	Patch(uint, any) Entity
}
