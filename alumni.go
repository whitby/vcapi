package vcapi

import "fmt"

const (
	alumniBasePath = "alumni"
)

type AlumniService struct {
	client *Client
}

// Alumni represents a Veracross API alumnis/former student
type Alumni struct {
	Address1              string      `json:"address_1"`
	Address2              interface{} `json:"address_2"`
	Address3              interface{} `json:"address_3"`
	Birthday              string      `json:"birthday"`
	BusinessPhone         interface{} `json:"business_phone"`
	City                  string      `json:"city"`
	CollegeGraduationYear int         `json:"college_graduation_year"`
	Country               string      `json:"country"`
	DateOfDeath           interface{} `json:"date_of_death"`
	DisplayAddress        bool        `json:"display_address"`
	DisplayBusinessPhone  bool        `json:"display_business_phone"`
	DisplayEmail1         bool        `json:"display_email_1"`
	DisplayEmail2         bool        `json:"display_email_2"`
	DisplayHomePhone      bool        `json:"display_home_phone"`
	DisplayJobTitle       bool        `json:"display_job_title"`
	DisplayMobilePhone    bool        `json:"display_mobile_phone"`
	DisplaySpouseName     bool        `json:"display_spouse_name"`
	Email1                string      `json:"email_1"`
	Email2                string      `json:"email_2"`
	Employer              interface{} `json:"employer"`
	ExitDate              string      `json:"exit_date"`
	FirstName             string      `json:"first_name"`
	Gender                string      `json:"gender"`
	GraduationYear        int         `json:"graduation_year"`
	HeadOfHousehold       bool        `json:"head_of_household"`
	HomePhone             string      `json:"home_phone"`
	HouseholdFk           int         `json:"household_fk"`
	JobTitle              interface{} `json:"job_title"`
	LastName              string      `json:"last_name"`
	MaidenName            interface{} `json:"maiden_name"`
	MiddleName            interface{} `json:"middle_name"`
	MobilePhone           interface{} `json:"mobile_phone"`
	NamePrefix            string      `json:"name_prefix"`
	NameSuffix            interface{} `json:"name_suffix"`
	NickFirstName         string      `json:"nick_first_name"`
	Occupation            interface{} `json:"occupation"`
	OrganizationFk        int         `json:"organization_fk"`
	PersonPk              int         `json:"person_pk"`
	PostalCode            string      `json:"postal_code"`
	PreferredName         interface{} `json:"preferred_name"`
	Role                  string      `json:"role"`
	SendMailTo            string      `json:"send_mail_to"`
	Spouse                interface{} `json:"spouse"`
	StateProvince         string      `json:"state_province"`
	UpdateDate            string      `json:"update_date"`
	Username              interface{} `json:"username"`
}

// ID returns an individual alumni record based on person id.
func (s AlumniService) ID(id string) (*Alumni, error) {
	var a Alumni
	path := fmt.Sprintf("%s/%s?format=%v", alumniBasePath, id, format)
	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, nil
	}
	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &a, nil
}

func (s AlumniService) list(opt *ListOptions, basePath string) ([]Alumni, error) {
	// build url
	path := addOptions(basePath, format, opt)

	var alumni = []Alumni{}
	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, nil
	}
	resp, err := s.client.Do(req, &alumni)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// handle pagination
	paginate(resp, opt)

	return alumni, nil
}

// List requests all alumni from API
func (s AlumniService) List(opt *ListOptions) ([]Alumni, error) {
	alumni, err := s.list(opt, alumniBasePath)
	if err != nil {
		return nil, err
	}

	return alumni, nil
}

// Recent requests recently updated alumni from API
func (s AlumniService) Recent(opt *ListOptions) ([]Alumni, error) {
	alumni, err := s.list(opt, alumniBasePath+"/recent")
	if err != nil {
		return nil, err
	}

	return alumni, nil
}
