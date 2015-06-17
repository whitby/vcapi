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
	fmt.Println(req)
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
