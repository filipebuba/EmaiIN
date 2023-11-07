package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/internalError"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type repositoryMock struct {
	mock.Mock
}

func (m *repositoryMock) Save(campaign *Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body Hi!",
		Email:   []string{"test@example.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	asserts := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	asserts.NotNil(id)
	asserts.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	asserts := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	asserts.False(errors.Is(internalError.ErrInternal, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Email) {
			return false
		}

		return true
	})).Return(nil)
	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	asserts := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	asserts.True(errors.Is(internalError.ErrInternal, err))
}
