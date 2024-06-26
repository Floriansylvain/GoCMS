package api

import (
	"GoCMS/adapters/secondary/gateways/models"
	"GoCMS/useCases"
	"github.com/glebarez/sqlite"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type UseCases struct {
	CreatePostUseCase  *useCases.CreatePostUseCase
	GetPostUseCase     *useCases.GetPostUseCase
	ListPostsUseCase   *useCases.ListPostsUseCase
	UpdatePostUseCase  *useCases.UpdatePostUseCase
	DeletePostUseCase  *useCases.DeletePostUseCase
	GetUserUseCase     *useCases.GetUserUseCase
	CreateUserUseCase  *useCases.CreateUserUseCase
	UpdateUserUseCase  *useCases.UpdateUserUseCase
	DeleteUserUseCase  *useCases.DeleteUserUseCase
	ListUsersUseCase   *useCases.ListUsersUseCase
	GetPageUseCase     *useCases.GetPageUseCase
	SendMailUseCase    *useCases.SendMailUseCase
	CreateImageUseCase *useCases.CreateImageUseCase
	DeleteImageUseCase *useCases.DeleteImageUseCase
}

var Container *UseCases

func getDb() *gorm.DB {
	dbName := os.Getenv("DB_FILE")
	if err := os.MkdirAll(filepath.Dir(dbName), os.ModePerm); err != nil {
		panic("Unable to create necessary subdirectories: " + err.Error())
	}
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("Unable to open the database: " + err.Error())
	}
	return db
}

func InitContainer() {
	digContainer := dig.New()

	database := getDb()
	_ = database.AutoMigrate(&models.Post{}, &models.User{})
	_ = digContainer.Provide(func() *gorm.DB { return database })

	_ = digContainer.Provide(func(db *gorm.DB) *UseCases {
		return &UseCases{
			CreatePostUseCase:  useCases.NewCreatePostUseCase(db),
			GetPostUseCase:     useCases.NewGetPostUseCase(db),
			ListPostsUseCase:   useCases.NewListPostsUseCase(db),
			UpdatePostUseCase:  useCases.NewUpdatePostUseCase(db),
			DeletePostUseCase:  useCases.NewDeletePostUseCase(db),
			GetUserUseCase:     useCases.NewGetUserUseCase(db),
			CreateUserUseCase:  useCases.NewCreateUserUseCase(db),
			UpdateUserUseCase:  useCases.NewUpdateUserUseCase(db),
			DeleteUserUseCase:  useCases.NewDeleteUserUseCase(db),
			ListUsersUseCase:   useCases.NewListUsersUseCase(db),
			GetPageUseCase:     useCases.NewGetPageUseCase(),
			SendMailUseCase:    useCases.NewSendMailUseCase(),
			CreateImageUseCase: useCases.NewCreateImageUseCase(db),
			DeleteImageUseCase: useCases.NewDeleteImageUseCase(db),
		}
	})
	err := digContainer.Invoke(func(useCases *UseCases) { Container = useCases })
	if err != nil {
		panic("Unable to invoke container: " + err.Error())
	}
}
