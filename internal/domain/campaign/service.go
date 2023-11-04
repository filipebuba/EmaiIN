package campaign

import (
	"emailn/internal/contact"
)

type Service struct {
	repository Repository
}

func (s *Service) Create(newCampaign contact.NewCompaignDTO) (string, error) {

	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Email)
	s.repository.Save(campaign)

	return campaign.ID, nil
}
