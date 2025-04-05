package repository

import "github/mauljurassicia/nutrifeat/entity"

type UserRepository interface {
	Repository[entity.User]
}

