package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

const coachesUri = "/coaches"

type Coach struct {
	CoachID      int     `json:"coach_id"`
	TeamID       int     `json:"team_id"`
	CountryID    int     `json:"country_id"`
	CommonName   string  `json:"common_name"`
	FullName     string  `json:"fullname"`
	FirstName    string  `json:"firstname"`
	LastName     string  `json:"lastname"`
	Nationality  string  `json:"nationality"`
	BirthDate    string  `json:"birthdate"`
	BirthCountry string  `json:"birthcountry"`
	Birthplace   *string `json:"birthplace"`
	ImagePath    string  `json:"image_path"`
}

// CoachById sends a request and returns a single Coach resource.
func (c *HTTPClient) CoachById(ctx context.Context, id int) (*Coach, *Meta, error) {
	path := fmt.Sprintf(coachesUri+"/%d", id)

	response := struct {
		Data *Coach `json:"data"`
		Meta *Meta  `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
