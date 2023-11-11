package endpoints

import (
	"emailn/internal/contract"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (r *CampaignServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (r *CampaignServiceMock) GetBy(id string) (*contract.CampaignResponse, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contract.CampaignResponse), args.Error(1)
}

func TestHandler_CampaignGetById_should_return_campain(t *testing.T) {
	asserts := assert.New(t)
	campaign := contract.CampaignResponse{
		ID:      "343",
		Name:    "name",
		Content: "content",
		Status:  "Panding",
	}
	service := new(CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(&campaign, nil)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(rr, req)

	asserts.Equal(200, status)
	asserts.Equal(campaign.ID, response.(*contract.CampaignResponse).ID)
	asserts.Equal(campaign.Name, response.(*contract.CampaignResponse).Name)
}

func Test_CampaignsGetById_should_return_error_when_something_wrong(t *testing.T) {
	asserts := assert.New(t)
	service := new(CampaignServiceMock)
	errExpected := errors.New("something wrong")
	service.On("GetBy", mock.Anything).Return(nil, errExpected)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, _, errReturned := handler.CampaignGetById(rr, req)

	asserts.Equal(errExpected.Error(), errReturned.Error())
}
