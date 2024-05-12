package factory

import (
	"context"
	"fmt"
	"graphyy/domain"
	"graphyy/repository/userrepo"
	"time"

	"github.com/bluele/factory-go/factory"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbKey struct{}
type passwordKey struct{}

var UserFactory = factory.NewFactory(
	&domain.User{},
).Attr("ID", func(args factory.Args) (interface{}, error) {
	return uuid.New(), nil
}).Attr("Username", func(args factory.Args) (interface{}, error) {
	user := args.Instance().(*domain.User)
	return fmt.Sprintf("user-%s", user.ID.String()), nil
}).Attr("Password", func(args factory.Args) (interface{}, error) {
	password := args.Context().Value(passwordKey{}).(string)
	hashedPassword, _ := userrepo.HashPassword(password)
	return hashedPassword, nil
}).Attr("CreatedAt", func(args factory.Args) (interface{}, error) {
	return time.Now(), nil
}).Attr("UpdatedAt", func(args factory.Args) (interface{}, error) {
	return time.Now(), nil
}).OnCreate(func(args factory.Args) error {
	db := args.Context().Value(dbKey{}).(*gorm.DB)
	return db.Create(args.Instance()).Error
})

func CreateUser(db *gorm.DB) domain.User {
	tx := db.Begin()
	ctx := context.WithValue(context.Background(), dbKey{}, tx)
	v, err := UserFactory.CreateWithContext(ctx)
	if err != nil {
		panic(err)
	}
	user := *v.(*domain.User)
	tx.Commit()
	return user
}

func CreateUsers(db *gorm.DB, n int) []domain.User {
	var users []domain.User
	for i := 0; i < n; i++ {
		tx := db.Begin()
		ctx := context.WithValue(context.Background(), dbKey{}, tx)
		password := fmt.Sprintf("user-%d", i)
		ctx = context.WithValue(ctx, passwordKey{}, password)
		v, err := UserFactory.CreateWithContext(ctx)
		if err != nil {
			panic(err)
		}
		user := *v.(*domain.User)
		user.Password = password
		tx.Commit()
		users = append(users, user)
	}
	return users
}
