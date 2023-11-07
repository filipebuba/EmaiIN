package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/internalError"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Email)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalError.ErrInternal
	}

	return campaign.ID, nil
}
