package zed

import "fmt"

// ArticleLabel struct
type ArticleLabel struct {
	ID        *float64 `json:"id,omitempty"`
	URL       *string  `json:"url,omitempty"`
	Name      *string  `json:"name,omitempty"`
	CreatedAt *string  `json:"created_at,omitempty"`
	UpdatedAt *string  `json:"updated_at,omitempty"`
}

// LabelResponse struct
type LabelResponse struct {
	Label *ArticleLabel `json:"label,omitempty"`
}

// LabelCollectionResponse struct
type LabelCollectionResponse struct {
	Results   []ArticleLabel `json:"labels,omitempty"`
	Count     *int64         `json:"count,omitempty"`
	Next      *string        `json:"next_page,omitempty"`
	Page      *int64         `json:"page,omitempty"`
	PageCount *int64         `json:"page_count,omitempty"`
	PerPage   *int64         `json:"per_page,omitempty"`
	Previous  *string        `json:"previous_page,omitempty"`
}

// LabelService struct
type LabelService struct {
	client *Client
}

// Create func creates a single new article
func (s *LabelService) Create(id *int64, l *ArticleLabel) (*ArticleLabel, error) {
	label := &ArticleLabel{}

	if l.Name == nil {
		return label, fmt.Errorf("missing required field: label name")
	}

	lw := &LabelResponse{Label: l}

	url := fmt.Sprintf("help_center/articles/%v/labels.json", *id)

	req, err := s.client.NewRequest("POST", url, lw)
	if err != nil {
		return label, err
	}

	result := LabelResponse{}
	_, err = s.client.Do(req, result)
	if err != nil {
		return label, err
	}
	label = result.Label

	return label, err
}

// List function lists labels used in all articles
func (s *LabelService) List() ([]ArticleLabel, error) {
	resource := []ArticleLabel{}

	rp, next, _, err := s.getPage("")
	if err != nil {
		return resource, err
	}
	resource = append(resource, rp...)

	for next != nil {
		rp, nx, _, err := s.getPage(*next)
		if err != nil {
			return resource, err
		}
		next = nx
		resource = append(resource, rp...)
	}

	return resource, err
}

// Get function lists lablels used in an article with a given id
func (s *LabelService) Get(id *int64) ([]ArticleLabel, error) {
	resource := []ArticleLabel{}

	url := fmt.Sprintf("help_center/articles/%v/labels.json", *id)
	rp, next, _, err := s.getPage(url)
	if err != nil {
		return resource, err
	}
	resource = append(resource, rp...)

	for next != nil {
		rp, nx, _, err := s.getPage(*next)
		if err != nil {
			return resource, err
		}
		next = nx
		resource = append(resource, rp...)
	}

	return resource, err
}

func (s *LabelService) getPage(url string) ([]ArticleLabel, *string, *Response, error) {

	if url == "" {
		url = "help_center/articles/labels.json"
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	result := LabelCollectionResponse{}
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, resp, err
	}

	next := result.Next
	resource := result.Results
	return resource, next, resp, err
}

// Delete func deletes a single article
func (s *LabelService) Delete(articleID, id *int64) error {
	if articleID == nil {
		return fmt.Errorf("missing required field: article id")
	}

	if id == nil {
		return fmt.Errorf("missing required field: id")
	}

	url := fmt.Sprintf("help_center/articles/%v/labels/%v.json", *articleID, *id)

	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("creating new request failed: %v\n", err)
	}

	_, err = s.client.Do(req, nil)
	if err != nil {
		return fmt.Errorf("request failed with: %v\n", err)
	}

	return err
}
