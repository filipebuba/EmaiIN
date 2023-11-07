package campaign

import (
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	name     = "Minha campanha"
	content  = "Body Hi!"
	contacts = []string{"emai1@gmail.com", "emai2@gmail.com"}
	fake     = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	// AAA -> Arrange, Act, Assert é uma convenção para escrever testes unitários e de integração de software
	// Arrange é a preparação do teste
	asserts := assert.New(t)

	// Act é o que estamos testando
	campaign, _ := NewCampaign(name, content, contacts)

	// Assert é a validação do teste
	asserts.Equal(campaign.Name, name)
	asserts.Equal(campaign.Content, content)
	asserts.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	// AAA -> Arrange, Act, Assert é uma convenção para escrever testes unitários e de integração de software
	// Arrange é a preparação do teste
	asserts := assert.New(t)

	// Act é o que estamos testando
	campaign, _ := NewCampaign(name, content, contacts)

	// Assert é a validação do teste
	asserts.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnIsMustBeNow(t *testing.T) {
	// AAA -> Arrange, Act, Assert é uma convenção para escrever testes unitários e de integração de software
	// Arrange é a preparação do teste
	asserts := assert.New(t)
	now := time.Now().Add(-time.Minute)

	// Act é o que estamos testando
	campaign, _ := NewCampaign(name, content, contacts)

	// Assert é a validação do teste
	asserts.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	// AAA -> Arrange, Act, Assert é uma convenção para escrever testes unitários e de integração de software
	// Arrange é a preparação do teste
	asserts := assert.New(t)

	// Act é o que estamos testando
	_, err := NewCampaign("", content, contacts)

	// Assert é a validação do teste
	asserts.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	asserts := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	asserts.Equal("name is required with max 24", err.Error())
}

func Test_NewCampaign_ContactsMustValidateMin(t *testing.T) {
	// AAA -> Arrange, Act, Assert é uma convenção para escrever testes unitários e de integração de software
	// Arrange é a preparação do teste
	asserts := assert.New(t)

	// Act é o que estamos testando
	_, err := NewCampaign(name, "", contacts)

	// Assert é a validação do teste
	asserts.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaign_ContactsMustValidateContentMax(t *testing.T) {
	// AAA -> Arrange, Act, Assert é uma convenção para escrever testes unitários e de integração de software
	// Arrange é a preparação do teste
	asserts := assert.New(t)

	// Act é o que estamos testando
	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	// Assert é a validação do teste
	asserts.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_ContactsMustValidateContactsMin(t *testing.T) {
	// AAA -> Arrange, Act, Assert é uma convenção para escrever testes unitários e de integração de software
	// Arrange é a preparação do teste
	asserts := assert.New(t)

	// Act é o que estamos testando
	_, err := NewCampaign(name, content, nil)

	// Assert é a validação do teste
	asserts.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaign_ContactsMustValidateContacts(t *testing.T) {
	asserts := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalido"})

	asserts.Equal("email is invalid", err.Error())
}
