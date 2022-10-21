package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/titusdishon/go-docker-mysql/entity"
	"testing"
)

var (
	emptyUser    = "user object is empty"
	emptyName    = "name field is empty"
	emptyEmail   = "email field is empty"
	emptySummary = "summary field is empty"
)

type mockRepository struct {
	mock.Mock
}

func (mock *mockRepository) Save(user *entity.User) (*entity.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.User), args.Error(1)
}
func (mock *mockRepository) Update(user *entity.User, id int64) (*entity.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.User), args.Error(1)
}
func (mock *mockRepository) FindAll() ([]entity.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.User), args.Error(1)
}
func (mock *mockRepository) FindById(id int64) (*entity.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.User), args.Error(1)
}
func (mock *mockRepository) Delete(id int64) (int64, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(int64), args.Error(1)
}

func TestServiceFindAll(t *testing.T) {
	mockRepo := new(mockRepository)
	var identifier int = 1
	var name string = "John Doe"
	var email string = "example@test.com"
	var summary string = "example summary"
	user := entity.User{
		ID:      identifier,
		Name:    name,
		Email:   email,
		Summary: summary,
	}
	mockRepo.On("FindAll").Return([]entity.User{user}, nil)
	testService := NewUserService(mockRepo)
	result, _ := testService.FindAll()
	// behavioral
	mockRepo.AssertExpectations(t)
	//data assertion
	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, name, result[0].Name)
	assert.Equal(t, email, result[0].Email)
	assert.Equal(t, summary, result[0].Summary)
}

func TestServiceFindById(t *testing.T) {
	mockRepo := new(mockRepository)
	var identifier int = 1
	var name string = "John Doe"
	var email string = "example@test.com"
	var summary string = "example summary"
	user := entity.User{
		ID:      identifier,
		Name:    name,
		Email:   email,
		Summary: summary,
	}
	mockRepo.On("FindById").Return(&user, nil)
	testService := NewUserService(mockRepo)
	result, _ := testService.FindById(int64(identifier))
	// behavioral
	mockRepo.AssertExpectations(t)
	//data assertion
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, name, result.Name)
	assert.Equal(t, email, result.Email)
	assert.Equal(t, summary, result.Summary)
}

func TestServiceUpdate(t *testing.T) {
	mockRepo := new(mockRepository)
	var identifier int = 1
	var name string = "John Doe"
	var email string = "example@test.com"
	var summary string = "example summary"
	var name1 string = "Loki kiki"
	var email1 string = "loki@test.com"
	var summary1 string = "loki's summary"
	user1 := entity.User{
		ID:      identifier,
		Name:    name,
		Email:   email,
		Summary: summary,
	}
	user2 := entity.User{
		ID:      identifier,
		Name:    name1,
		Email:   email1,
		Summary: summary1,
	}
	mockRepo.On("Update").Return(&user2, nil)
	testService := NewUserService(mockRepo)
	result, _ := testService.Update(&user1, int64(identifier))
	// behavioral
	mockRepo.AssertExpectations(t)
	//data assertion
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, name1, result.Name)
	assert.Equal(t, email1, result.Email)
	assert.Equal(t, summary1, result.Summary)
}
func TestServiceSave(t *testing.T) {
	mockRepo := new(mockRepository)
	var identifier int = 1
	var name string = "John Doe"
	var email string = "example@test.com"
	var summary string = "example summary"
	user := entity.User{
		ID:      identifier,
		Name:    name,
		Email:   email,
		Summary: summary,
	}
	mockRepo.On("Save").Return(&user, nil)
	testService := NewUserService(mockRepo)
	result, _ := testService.Save(&user)
	// behavioral
	mockRepo.AssertExpectations(t)
	//data assertion
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, name, result.Name)
	assert.Equal(t, email, result.Email)
	assert.Equal(t, summary, result.Summary)
}

func TestValidateEmptyUser(t *testing.T) {
	testService := NewUserService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), emptyUser)
}

func TestValidateEmptyUserName(t *testing.T) {
	testService := NewUserService(nil)
	user := entity.User{
		Name:    "",
		Email:   "example@test.com",
		Summary: "example summary",
	}
	err := testService.Validate(&user)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), emptyName)
}

func TestValidateEmptyUserEmail(t *testing.T) {
	testService := NewUserService(nil)
	user := entity.User{
		Name:    "john doe",
		Email:   "",
		Summary: "example summary",
	}
	err := testService.Validate(&user)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), emptyEmail)
}
func TestValidateEmptyUserSummary(t *testing.T) {
	testService := NewUserService(nil)
	user := entity.User{
		Name:    "john doe",
		Email:   "example@test.com",
		Summary: "",
	}
	err := testService.Validate(&user)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), emptySummary)
}
