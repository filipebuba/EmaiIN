package campaign

import (
	"emailn/internal/contact"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	newCampaign = contact.NewCompaignDTO{
		Name:    "Test v",
		Content: "Body",
		Email:   []string{"test@example.com"},
	}
	mockRepasitory = new(MockRepository)
	service        = Service{}
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(campaign *Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign_Success(t *testing.T) {
	// Arrange
	asserts := assert.New(t)
	newCampaign := contact.NewCompaignDTO{
		Name:    "Test v",
		Content: "Body",
		Email:   []string{"test@example.com"},
	}
	// Act
	id, err := service.Create(newCampaign)
	asserts.NotNil(id)
	asserts.Nil(err)
}

func Test_Create_Save_Compaign(t *testing.T) {

	newCampaign = contact.NewCompaignDTO{
		Name:    "Test v",
		Content: "Body",
		Email:   []string{"test@example.com"},
	}
	mockRepasitory := new(MockRepository)
	mockRepasitory.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Email) {
			return false
		}
		return true
	})).Return(nil)
	service.repository = mockRepasitory

	service.Create(newCampaign)
	mockRepasitory.AssertExpectations(t)
}
