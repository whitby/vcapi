package vcapi

import "fmt"

const (
	facstaffBasePath = "facstaff"
)

type FacStaffService struct {
	client *Client
}

// FacStaff represents a Veracross API facstaff
type FacStaff struct {
	Biography            interface{} `json:"biography"`
	Birthday             string      `json:"birthday"`
	BusinessPhone        string      `json:"business_phone"`
	Campus               interface{} `json:"campus"`
	DateHired            string      `json:"date_hired"`
	Department           string      `json:"department"`
	DisplayBusinessPhone bool        `json:"display_business_phone"`
	DisplayEmail1        bool        `json:"display_email_1"`
	DisplayFacultyType   bool        `json:"display_faculty_type"`
	DisplayHomePhone     bool        `json:"display_home_phone"`
	DisplayJobTitle      bool        `json:"display_job_title"`
	DisplayMobilePhone   bool        `json:"display_mobile_phone"`
	Email1               string      `json:"email_1"`
	FacultyType          string      `json:"faculty_type"`
	FirstName            string      `json:"first_name"`
	Gender               string      `json:"gender"`
	HomePhone            string      `json:"home_phone"`
	HouseholdFk          int         `json:"household_fk"`
	JobTitle             string      `json:"job_title"`
	LastName             string      `json:"last_name"`
	MiddleName           interface{} `json:"middle_name"`
	MobilePhone          string      `json:"mobile_phone"`
	NamePrefix           string      `json:"name_prefix"`
	NameSuffix           interface{} `json:"name_suffix"`
	NickFirstName        string      `json:"nick_first_name"`
	NickName             interface{} `json:"nick_name"`
	PersonPk             int         `json:"person_pk"`
	Role                 string      `json:"role"`
	Roles                string      `json:"roles"`
	SchoolLevel          string      `json:"school_level"`
	SecurityHash         string      `json:"security_hash"`
	SecurityRoles        string      `json:"security_roles"`
	UpdateDate           string      `json:"update_date"`
	Username             string      `json:"username"`
}

// ID returns an individual facstaff record based on person id.
func (s FacStaffService) ID(id string) (*FacStaff, error) {
	type aFacStaff struct {
		FacStaff `json:"facstaff"`
	}
	var a aFacStaff
	path := fmt.Sprintf("%s/%s?format=%v", facstaffBasePath, id, format)
	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, nil
	}
	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &a.FacStaff, nil
}

func (s FacStaffService) list(opt *ListOptions, basePath string) ([]FacStaff, error) {
	// build url
	path := addOptions(basePath, format, opt)

	var facstaffs = []FacStaff{}
	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, nil
	}
	resp, err := s.client.Do(req, &facstaffs)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// handle pagination
	paginate(resp, opt)

	return facstaffs, nil
}

// List requests all Faculty and Staff from API
func (s FacStaffService) List(opt *ListOptions) ([]FacStaff, error) {
	facstaffs, err := s.list(opt, facstaffBasePath)
	if err != nil {
		return nil, err
	}

	return facstaffs, nil
}

// Recent requests recently updated Faculty and Staff from API
func (s FacStaffService) Recent(opt *ListOptions) ([]FacStaff, error) {
	facstaffs, err := s.list(opt, facstaffBasePath+"/recent")
	if err != nil {
		return nil, err
	}

	return facstaffs, nil
}
