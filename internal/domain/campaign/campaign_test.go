package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	name     = "Minha campanha"
	content  = "Body: Olá, %s! Você foi convidado para participar da minha campanha"
	contacts = []string{"emai1@gmail.com", "emai2@gmail.com"}
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

func Test_NewCampaign_CreatedOnIsMustValidateName(t *testing.T) {
	// AAA -> Arrange, Act, Assert é uma convenção para escrever testes unitários e de integração de software
	// Arrange é a preparação do teste
	asserts := assert.New(t)

	// Act é o que estamos testando
	_, err := NewCampaign("", content, contacts)

	// Assert é a validação do teste
	asserts.Equal("name is required", err.Error())
}

func Test_NewCampaign_ContactsMustValidateContent(t *testing.T) {
	// AAA -> Arrange, Act, Assert é uma convenção para escrever testes unitários e de integração de software
	// Arrange é a preparação do teste
	asserts := assert.New(t)

	// Act é o que estamos testando
	_, err := NewCampaign(name, "", contacts)

	// Assert é a validação do teste
	asserts.Equal("content is required", err.Error())
}

func Test_NewCampaign_ContactsMustValidateContacts(t *testing.T) {
	// AAA -> Arrange, Act, Assert é uma convenção para escrever testes unitários e de integração de software
	// Arrange é a preparação do teste
	asserts := assert.New(t)

	// Act é o que estamos testando
	_, err := NewCampaign(name, content, []string{})

	// Assert é a validação do teste
	asserts.Equal("contacts is required", err.Error())
}
