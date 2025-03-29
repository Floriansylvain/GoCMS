package api

import (
	"GoCMS/adapters/secondary/gateways"
	"GoCMS/adapters/secondary/gateways/models"
	domainGateways "GoCMS/domain/gateways"
	"GoCMS/useCases"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/glebarez/sqlite"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type UseCaseDefinition struct {
	Constructor interface{}
	FieldName   string
}

type RepositoryDefinition struct {
	Constructor interface{}
	Interface   interface{}
}

var useCaseDefinitions = []UseCaseDefinition{
	{useCases.NewCreatePostUseCase, "CreatePostUseCase"},
	{useCases.NewGetPostUseCase, "GetPostUseCase"},
	{useCases.NewListPostsUseCase, "ListPostsUseCase"},
	{useCases.NewUpdatePostUseCase, "UpdatePostUseCase"},
	{useCases.NewDeletePostUseCase, "DeletePostUseCase"},
	{useCases.NewGetUserUseCase, "GetUserUseCase"},
	{useCases.NewCreateUserUseCase, "CreateUserUseCase"},
	{useCases.NewUpdateUserUseCase, "UpdateUserUseCase"},
	{useCases.NewDeleteUserUseCase, "DeleteUserUseCase"},
	{useCases.NewListUsersUseCase, "ListUsersUseCase"},
	{useCases.NewGetPageUseCase, "GetPageUseCase"},
	{useCases.NewSendMailUseCase, "SendMailUseCase"},
	{useCases.NewCreateImageUseCase, "CreateImageUseCase"},
	{useCases.NewDeleteImageUseCase, "DeleteImageUseCase"},
}

var repositoryDefinitions = []RepositoryDefinition{
	{gateways.NewPostRepository, new(domainGateways.IPostRepository)},
	{gateways.NewUserRepository, new(domainGateways.IUserRepository)},
	{gateways.NewImageRepository, new(domainGateways.IImageRepository)},
	{gateways.NewMailRepository, new(domainGateways.IMailRepository)},
	{gateways.NewPageRepository, new(domainGateways.IPageRepository)},
}

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
	if err := database.AutoMigrate(&models.Post{}, &models.User{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	if err := digContainer.Provide(func() *gorm.DB { return database }); err != nil {
		panic("Failed to provide database: " + err.Error())
	}

	provideRepositories(digContainer)

	for _, def := range useCaseDefinitions {
		if err := digContainer.Provide(def.Constructor); err != nil {
			funcName := runtime.FuncForPC(reflect.ValueOf(def.Constructor).Pointer()).Name()
			panic(fmt.Sprintf("Failed to provide %s: %v", funcName, err))
		}
	}

	constructorType := buildConstructorType()
	constructor := reflect.MakeFunc(constructorType, buildUseCasesInstance)

	if err := digContainer.Provide(constructor.Interface()); err != nil {
		panic("Failed to provide container constructor: " + err.Error())
	}

	if err := digContainer.Invoke(func(useCases *UseCases) { Container = useCases }); err != nil {
		panic("Unable to invoke container: " + err.Error())
	}
}

func buildConstructorType() reflect.Type {
	var paramTypes []reflect.Type
	for _, def := range useCaseDefinitions {
		constructorType := reflect.TypeOf(def.Constructor)
		returnType := constructorType.Out(0)
		paramTypes = append(paramTypes, returnType)
	}

	return reflect.FuncOf(paramTypes, []reflect.Type{reflect.TypeOf(&UseCases{})}, false)
}

func buildUseCasesInstance(args []reflect.Value) []reflect.Value {
	container := &UseCases{}
	containerValue := reflect.ValueOf(container).Elem()

	for i, arg := range args {
		fieldName := useCaseDefinitions[i].FieldName
		field := containerValue.FieldByName(fieldName)
		if field.IsValid() && field.CanSet() {
			field.Set(arg)
		} else {
			panic(fmt.Sprintf("Failed to set field %s", fieldName))
		}
	}

	return []reflect.Value{reflect.ValueOf(container)}
}

func provideRepositories(container *dig.Container) {
	for _, def := range repositoryDefinitions {
		if err := container.Provide(def.Constructor, dig.As(def.Interface)); err != nil {
			funcName := runtime.FuncForPC(reflect.ValueOf(def.Constructor).Pointer()).Name()
			panic(fmt.Sprintf("Failed to provide repository %s: %v", funcName, err))
		}
	}
}
