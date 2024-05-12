package userrepo

import (
	"errors"
	"fmt"
	"graphyy/domain"
	"graphyy/internal/envvar"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepo should i rename it?
type UserRepo struct {
	db *gorm.DB
}

// NewUserRepo ..
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// GetExistingUser fetches a user by the username from the db and returns it.
func (repo *UserRepo) GetExistingUser(username string) (*domain.User, error) {
	var user UserEntity
	result := repo.db.Where("username = ? AND deleted_at IS NULL", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no user found with username: %s", username)
		}
		return nil, result.Error
	}
	return &domain.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// CreateUser creates a new user in the db..
func (repo *UserRepo) CreateUser(user domain.User) (*domain.User, error) {
	hashedPass, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	dbUser := UserEntity{
		Username: user.Username,
		Password: hashedPass,
	}

	result := repo.db.Create(&dbUser)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println("Inserted a user with ID:", dbUser.ID)
	return &domain.User{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Password:  dbUser.Password,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}, nil
}

func HashPassword(password string) (string, error) {
	cost := envvar.HashCost()
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

// use the entity suffix for the entities used for database operations
// UserEntity struct represents the user entity in the db. entities defined in this package should only be used in the repository package.
type UserEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"unique;index;not null"`
	Password  string         `gorm:"type:varchar(1000);not null"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *UserEntity) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (user *UserEntity) TableName() string {
	return "users"
}
