package user

import (
	"context"
	"fmt"

	"github.com/hussar_company/vstazherstve_back/dbcontext"
	"github.com/hussar_company/vstazherstve_back/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Репозиторий инкапсулирует логику доступа к пользователям
type Repository interface {
	// Получение пользователя по идентификатору
	Get(ctx context.Context, id string) (entity.User, error)

	// Сохранение пользователя в базе данных
	Create(ctx context.Context, album entity.User) error
}

// Репозиторий хранит альбомы в базе данных
type repository struct {
	db *dbcontext.DB
}

// Создает новый экземпляр репозитория
func NewRepository(db *dbcontext.DB) Repository {
	return repository{db}
}

// Получение пользователя по идентификатору
func (r repository) Get(ctx context.Context, id string) (entity.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Invalid id")
	}

	var result entity.User
	err = r.db.DB().Collection("users").FindOne(ctx, bson.M{"_id": objectId}).Decode(&result)
	return result, err
}

// Сохранение пользователя в базе данных
func (r repository) Create(ctx context.Context, user entity.User) error {
	_, err := r.db.DB().Collection("users").InsertOne(ctx, bson.M{"name": user.Name, "password": user.Password})
	return err
}
