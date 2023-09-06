package api

import (
	"GohCMS2/adapters/secondary/gateways/models"
	"GohCMS2/useCases"
	"github.com/glebarez/sqlite"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type LocalContainer struct {
	CreatePostUseCase *useCases.CreatePostUseCase
	GetPostUseCase    *useCases.GetPostUseCase
	ListPostsUseCase  *useCases.ListPostsUseCase
	GetUserUseCase    *useCases.GetUserUseCase
	CreateUserUseCase *useCases.CreateUserUseCase
	ListUsersUseCase  *useCases.ListUsersUseCase
	GetPageUseCase    *useCases.GetPageUseCase
}

var Container *LocalContainer

func setContainer(
	createPost *useCases.CreatePostUseCase,
	getPost *useCases.GetPostUseCase,
	listPost *useCases.ListPostsUseCase,
	getUser *useCases.GetUserUseCase,
	createUser *useCases.CreateUserUseCase,
	listUsers *useCases.ListUsersUseCase,
	getPage *useCases.GetPageUseCase,
) *LocalContainer {
	Container = &LocalContainer{
		CreatePostUseCase: createPost,
		GetPostUseCase:    getPost,
		ListPostsUseCase:  listPost,
		GetUserUseCase:    getUser,
		CreateUserUseCase: createUser,
		ListUsersUseCase:  listUsers,
		GetPageUseCase:    getPage,
	}
	return Container
}

func InitContainer() {
	if Container != nil {
		return
	}

	digContainer := dig.New()

	dbName := os.Getenv("DB_FILE")
	if err := os.MkdirAll(filepath.Dir(dbName), os.ModePerm); err != nil {
		panic("Unable to create necessary subdirectories: " + err.Error())
	}

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("Unable to open the database: " + err.Error())
	}

	_ = db.AutoMigrate(&models.Post{}, &models.User{})

	_ = digContainer.Provide(func() *gorm.DB { return db })

	_ = digContainer.Provide(func(db *gorm.DB) *useCases.CreatePostUseCase { return useCases.NewCreatePostUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *useCases.GetPostUseCase { return useCases.NewGetPostUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *useCases.ListPostsUseCase { return useCases.NewListPostsUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *useCases.GetUserUseCase { return useCases.NewGetUserUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *useCases.CreateUserUseCase { return useCases.NewCreateUserUseCase(db) })
	_ = digContainer.Provide(func(db *gorm.DB) *useCases.ListUsersUseCase { return useCases.NewListUsersUseCase(db) })
	_ = digContainer.Provide(useCases.NewGetPageUseCase)

	_ = digContainer.Invoke(setContainer)
}
