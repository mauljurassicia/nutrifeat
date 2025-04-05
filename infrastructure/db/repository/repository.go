package repository

type Repository[T any] interface {
	Get(id int) (*T, error)
	GetAll() ([]T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id int) error
}
