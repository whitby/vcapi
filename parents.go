package vcapi

import "fmt"

const (
	parentsBasePath = "parents"
)

type ParentService struct {
	client *Client
}

// Parent represents a Veracross API parent
type Parent struct {
	Address1             string      `json:"address_1"`
	Address2             interface{} `json:"address_2"`
	Address3             interface{} `json:"address_3"`
	AddressType          int         `json:"address_type"`
	BusinessPhone        interface{} `json:"business_phone"`
	City                 string      `json:"city"`
	Country              string      `json:"country"`
	DisplayAddress       bool        `json:"display_address"`
	DisplayBusinessPhone bool        `json:"display_business_phone"`
	DisplayCity          bool        `json:"display_city"`
	DisplayCountry       bool        `json:"display_country"`
	DisplayEmail1        bool        `json:"display_email_1"`
	DisplayEmail2        bool        `json:"display_email_2"`
	DisplayHomePhone     bool        `json:"display_home_phone"`
	DisplayJobTitle      bool        `json:"display_job_title"`
	DisplayMobilePhone   bool        `json:"display_mobile_phone"`
	DisplayPostalCode    bool        `json:"display_postal_code"`
	DisplayStateProvince bool        `json:"display_state_province"`
	Email1               interface{} `json:"email_1"`
	Email2               interface{} `json:"email_2"`
	Employer             interface{} `json:"employer"`
	FirstName            string      `json:"first_name"`
	FirstNickName        string      `json:"first_nick_name"`
	Gender               string      `json:"gender"`
	HeadOfHousehold      bool        `json:"head_of_household"`
	HomePhone            string      `json:"home_phone"`
	HouseholdFk          int         `json:"household_fk"`
	JobTitle             interface{} `json:"job_title"`
	LastName             string      `json:"last_name"`
	MiddleName           interface{} `json:"middle_name"`
	MobilePhone          string      `json:"mobile_phone"`
	NamePrefix           string      `json:"name_prefix"`
	NameSuffix           string      `json:"name_suffix"`
	Occupation           interface{} `json:"occupation"`
	PersonPk             int         `json:"person_pk"`
	PostalCode           string      `json:"postal_code"`
	PreferredName        string      `json:"preferred_name"`
	Role                 string      `json:"role"`
	Spouse               string      `json:"spouse"`
	StateProvince        string      `json:"state_province"`
	UpdateDate           string      `json:"update_date"`
	UpdateLink           string      `json:"update_link"`
	Username             interface{} `json:"username"`
}

// Relationships returns related persons
func (s ParentService) Relationships(p Parent) (*[]Relationship, error) {
	path := fmt.Sprintf("%s/%v/relationships?format=%v", parentsBasePath, p.PersonPk, format)
	var relationships []Relationship
	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, nil
	}
	resp, err := s.client.Do(req, &relationships)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &relationships, nil
}

// ID returns an individual parent record based on person id.
func (s ParentService) ID(id string) (*Parent, error) {
	type aParent struct {
		Parent `json:"parent"`
	}
	var a aParent
	path := fmt.Sprintf("%s/%s?format=%v", parentsBasePath, id, format)
	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, nil
	}
	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &a.Parent, nil
}

func (s ParentService) list(opt *ListOptions, basePath string) ([]Parent, error) {
	// build url
	path := addOptions(basePath, format, opt)

	var parents = []Parent{}
	req, err := s.client.NewRequest(path)
	fmt.Println(req)
	if err != nil {
		return nil, nil
	}
	resp, err := s.client.Do(req, &parents)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// handle pagination
	paginate(resp, opt)

	return parents, nil
}

// List requests all parents from API
func (s ParentService) List(opt *ListOptions) ([]Parent, error) {
	parents, err := s.list(opt, parentsBasePath)
	if err != nil {
		return nil, err
	}

	return parents, nil
}

// Recent requests recently updated parents from API
func (s ParentService) Recent(opt *ListOptions) ([]Parent, error) {
	parents, err := s.list(opt, parentsBasePath+"/recent")
	if err != nil {
		return nil, err
	}

	return parents, nil
}
