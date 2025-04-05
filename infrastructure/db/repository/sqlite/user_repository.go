package sqlite

import (
	"github/mauljurassicia/nutrifeat/entity"
	"github/mauljurassicia/nutrifeat/infrastructure/db/repository"
)

type sqliteUserRepository struct{
	sqliteRepository[entity.User]
}

func NewUserRepository() repository.UserRepository {
	return &sqliteUserRepository{}
}

