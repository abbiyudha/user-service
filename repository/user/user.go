package user

import (
	"CodingTestUser/entities"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	database *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

type UserRepositoryInferface interface {
	CreateUser(UserParam entities.User) error
	LoginUser(email string) (entities.User, error)
	GetUserById(id string) (entities.User, error)
}

func (ur *UserRepository) CreateUser(UserParam entities.User) error {
	var ctx context.Context

	UserParam.Id = primitive.NewObjectID()
	encrypt, _ := bcrypt.GenerateFromPassword([]byte(UserParam.Password), bcrypt.DefaultCost)
	UserParam.Password = string(encrypt)

	_, err := ur.database.Collection("User").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	_, err = ur.database.Collection("User").InsertOne(ctx, UserParam)
	if err != nil {
		return errors.New("Email already used")
	}

	return nil
}

func (ur *UserRepository) LoginUser(email string) (entities.User, error) {
	var ctx context.Context
	var user entities.User

	result := ur.database.Collection("User").FindOne(ctx, bson.M{"email": email})
	if result.Err() != nil {
		return user, result.Err()
	}

	err := result.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserById(id string) (entities.User, error) {
	var ctx context.Context
	var UserParam entities.User

	UserParam.Id, _ = primitive.ObjectIDFromHex(id)

	result := ur.database.Collection("User").FindOne(ctx, bson.M{"_id": UserParam.Id})
	if result.Err() != nil {
		return UserParam, errors.New("User not found")
	}

	err := result.Decode(&UserParam)
	if err != nil {
		return entities.User{}, err
	}

	return UserParam, nil
}
