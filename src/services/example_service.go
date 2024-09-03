package services

import (
"go-api/src/repositories"
)

type ExampleService interface{
//Insert your function interface
}

type exampleService struct {
repositoryExample repositories.ExampleRepository
}

func NewExampleService(
repositoryExample repositories.ExampleRepository,
//repo
) ExampleService {
	return &exampleService{
repositoryExample :repositoryExample,
//repo
}
}