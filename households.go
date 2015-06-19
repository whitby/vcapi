package vcapi

import "fmt"

const (
	householdBasePath = "households"
)

type HouseholdService struct {
	client *Client
}

// Household represents a Veracross API household
type Household struct {
	Address1 string      `json:"address_1"`
	Address2 interface{} `json:"address_2"`
	Address3 interface{} `json:"address_3"`
	Alumni   []struct {
		HeadOfHousehold bool `json:"head_of_household"`
		PersonPk        int  `json:"person_pk"`
	} `json:"alumni"`
	City             string `json:"city"`
	Country          string `json:"country"`
	CurrentHousehold bool   `json:"current_household"`
	DisplayAddress   bool   `json:"display_address"`
	Email            string `json:"email"`
	FutureHousehold  bool   `json:"future_household"`
	HouseholdName    string `json:"household_name"`
	HouseholdPk      int    `json:"household_pk"`
	Parents          []struct {
		HeadOfHousehold bool `json:"head_of_household"`
		PersonPk        int  `json:"person_pk"`
	} `json:"parents"`
	Phone         string `json:"phone"`
	PostalCode    string `json:"postal_code"`
	StateProvince string `json:"state_province"`
	Students      []struct {
		PersonPk int `json:"person_pk"`
	} `json:"students"`
	UpdateDate string `json:"update_date"`
	UpdateLink string `json:"update_link"`
}

// ID returns an individual household record based on person id.
func (s HouseholdService) ID(id string) (*Household, error) {
	type aHousehold struct {
		Household `json:"household"`
	}
	var a aHousehold
	path := fmt.Sprintf("%s/%s?format=%v", householdBasePath, id, format)
	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, nil
	}
	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &a.Household, nil
}

func (s HouseholdService) list(opt *ListOptions, basePath string) ([]Household, error) {
	// build url
	path := addOptions(basePath, format, opt)

	var households = []Household{}
	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, nil
	}
	resp, err := s.client.Do(req, &households)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// handle pagination
	paginate(resp, opt)

	return households, nil
}

// List requests all Households from API
func (s HouseholdService) List(opt *ListOptions) ([]Household, error) {
	households, err := s.list(opt, householdBasePath)
	if err != nil {
		return nil, err
	}

	return households, nil
}

// Recent requests recently updated households from API
func (s HouseholdService) Recent(opt *ListOptions) ([]Household, error) {
	households, err := s.list(opt, householdBasePath+"/recent")
	if err != nil {
		return nil, err
	}

	return households, nil
}
