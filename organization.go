package zed

import (
	"fmt"
)

// OrganizationResponse struct
type OrganizationResponse struct {
	Organization *Organization `json:"organization"`
}

// OrganizationCollectionResponse struct
type OrganizationCollectionResponse struct {
	Organizations []Organization `json:"organizations,omitempty"`
	NextPage      *string        `json:"next_page,omitempty"`
	PreviousPage  *string        `json:"previous_page,omitempty"`
	Count         *int           `json:"count,omitempty"`
}

// Organization struct
type Organization struct {
	URL                *string           `json:"url,omitempty"`
	ID                 *int              `json:"id,omitempty"`
	Name               *string           `json:"name,omitempty"`
	SharedTickets      *bool             `json:"shared_tickets,omitempty"`
	SharedComments     *bool             `json:"shared_comments,omitempty"`
	ExternalID         *string           `json:"external_id,omitempty"`
	CreatedAt          *string           `json:"created_at,omitempty"`
	UpdatedAt          *string           `json:"updated_at,omitempty"`
	DomainNames        []string          `json:"domain_names,omitempty"`
	Details            *string           `json:"details,omitempty"`
	Notes              *string           `json:"notes,omitempty"`
	GroupID            *int              `json:"group_id,omitempty"`
	Tags               []string          `json:"tags,omitempty"`
	OrganizationFields map[string]string `json:"organization_fields,omitempty"`
}

// OrganizationService struct
type OrganizationService struct {
	client *Client
}

// Get finds an organization in Zendesk by ID
func (s *OrganizationService) Get(organizationID string) (*Organization, *Response, error) {
	org := OrganizationResponse{}

	url := fmt.Sprintf("organizations/%s.json", organizationID)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, &org)
	if err != nil {
		return nil, nil, err
	}

	return org.Organization, resp, err
}

// Update updates and organization by id
func (s *OrganizationService) Update(org *Organization) (*Organization, error) {
	organization := &Organization{}

	url := fmt.Sprintf("organizations/%d.json", *org.ID)
	or := &OrganizationResponse{Organization: org}

	req, err := s.client.NewRequest("PUT", url, or)
	if err != nil {
		return organization, err
	}

	result := OrganizationResponse{}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return organization, err
	}

	organization = result.Organization
	return organization, err
}

//Create creates a new organization
func (s *OrganizationService) Create(org *Organization) (*Organization, error) {
	organization := &Organization{}

	or := &OrganizationResponse{Organization: org}
	url := fmt.Sprintf("organizations.json")

	req, err := s.client.NewRequest("POST", url, or)
	if err != nil {
		return organization, err
	}

	result := OrganizationResponse{}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return organization, err
	}

	organization = result.Organization
	return organization, nil
}
