package sqlite

type sqliteRepository[T any] struct{}



func (r *sqliteRepository[T]) Get(id int) (*T, error) {
	return nil, nil
}

func (r *sqliteRepository[T]) GetAll() ([]T, error) {
	return nil, nil
}

func (r *sqliteRepository[T]) Create(entity *T) error {
	return nil
}

func (r *sqliteRepository[T]) Update(entity *T) error {
	return nil
}

func (r *sqliteRepository[T]) Delete(id int) error {	
	return nil
}




