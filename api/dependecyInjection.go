package api

import (
	"GoCMS/adapters/secondary/gateways"
	"GoCMS/adapters/secondary/gateways/models"
	domainGateways "GoCMS/domain/gateways"
	"GoCMS/useCases"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
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

type Repositories struct {
	PostRepo  domainGateways.IPostRepository
	UserRepo  domainGateways.IUserRepository
	ImageRepo domainGateways.IImageRepository
	MailRepo  domainGateways.IMailRepository
	PageRepo  domainGateways.IPageRepository
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

	if err := db.AutoMigrate(&models.Post{}, &models.User{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	return db
}

func initRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		PostRepo:  gateways.NewPostRepository(db),
		UserRepo:  gateways.NewUserRepository(db),
		ImageRepo: gateways.NewImageRepository(db),
		MailRepo:  gateways.NewMailRepository(),
		PageRepo:  gateways.NewPageRepository(),
	}
}

func initUseCases(repos *Repositories) *UseCases {
	return &UseCases{
		CreatePostUseCase:  useCases.NewCreatePostUseCase(repos.PostRepo),
		GetPostUseCase:     useCases.NewGetPostUseCase(repos.PostRepo),
		ListPostsUseCase:   useCases.NewListPostsUseCase(repos.PostRepo),
		UpdatePostUseCase:  useCases.NewUpdatePostUseCase(repos.PostRepo),
		DeletePostUseCase:  useCases.NewDeletePostUseCase(repos.PostRepo),
		GetUserUseCase:     useCases.NewGetUserUseCase(repos.UserRepo),
		CreateUserUseCase:  useCases.NewCreateUserUseCase(repos.UserRepo),
		UpdateUserUseCase:  useCases.NewUpdateUserUseCase(repos.UserRepo),
		DeleteUserUseCase:  useCases.NewDeleteUserUseCase(repos.UserRepo),
		ListUsersUseCase:   useCases.NewListUsersUseCase(repos.UserRepo),
		GetPageUseCase:     useCases.NewGetPageUseCase(repos.PageRepo),
		SendMailUseCase:    useCases.NewSendMailUseCase(repos.MailRepo),
		CreateImageUseCase: useCases.NewCreateImageUseCase(repos.ImageRepo),
		DeleteImageUseCase: useCases.NewDeleteImageUseCase(repos.ImageRepo),
	}
}

func InitContainer() {
	db := getDb()
	repos := initRepositories(db)
	Container = initUseCases(repos)
}
