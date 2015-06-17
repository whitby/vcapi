package vcapi

import "fmt"

const (
	studentsBasePath = "students"
	format           = "json"
)

type studentService struct {
	client *Client
}

// Student represents a Veracross API student
type Student struct {
	AdvisorFk                 int         `json:"advisor_fk"`
	AdvisorName               string      `json:"advisor_name"`
	BedNumber                 int         `json:"bed_number"`
	Birthday                  Date        `json:"birthday"`
	Campus                    interface{} `json:"campus"`
	CampusApplyingFor         bool        `json:"campus_applying_for"`
	CurrentGrade              string      `json:"current_grade"`
	DisplayEmail1             bool        `json:"display_email_1"`
	DisplayEmail2             bool        `json:"display_email_2"`
	DisplayHomePhone          bool        `json:"display_home_phone"`
	DisplayMobilePhone        bool        `json:"display_mobile_phone"`
	Dorm                      interface{} `json:"dorm"`
	Email1                    string      `json:"email_1"`
	Email2                    string      `json:"email_2"`
	EnrollmentStatus          string      `json:"enrollment_status"`
	FirstName                 string      `json:"first_name"`
	FirstNickName             string      `json:"first_nick_name"`
	FloorNumber               int         `json:"floor_number"`
	Gender                    string      `json:"gender"`
	GradeApplyingFor          string      `json:"grade_applying_for"`
	GraduationYear            int         `json:"graduation_year"`
	HomePhone                 string      `json:"home_phone"`
	Homeroom                  int         `json:"homeroom"`
	HomeroomTeacherFk         int         `json:"homeroom_teacher_fk"`
	HomeroomTeacherName       interface{} `json:"homeroom_teacher_name"`
	HouseholdFk               int         `json:"household_fk"`
	LastName                  string      `json:"last_name"`
	MailboxNumber             interface{} `json:"mailbox_number"`
	MiddleName                interface{} `json:"middle_name"`
	MobilePhone               interface{} `json:"mobile_phone"`
	NamePrefix                interface{} `json:"name_prefix"`
	NameSuffix                interface{} `json:"name_suffix"`
	Parent1Fk                 int         `json:"parent_1_fk"`
	Parent2Fk                 int         `json:"parent_2_fk"`
	Parent3Fk                 int         `json:"parent_3_fk"`
	Parent4Fk                 int         `json:"parent_4_fk"`
	PersonPk                  int         `json:"person_pk"`
	PreferredName             interface{} `json:"preferred_name"`
	ResidentStatus            interface{} `json:"resident_status"`
	ResidentStatusApplyingFor bool        `json:"resident_status_applying_for"`
	Role                      string      `json:"role"`
	RoomNumber                interface{} `json:"room_number"`
	SchoolLevel               string      `json:"school_level"`
	StudentGroup              interface{} `json:"student_group"`
	StudentGroupApplyingFor   bool        `json:"student_group_applying_for"`
	UpdateDate                string      `json:"update_date"`
	Username                  string      `json:"username"`
	YearApplyingFor           int         `json:"year_applying_for"`
}

// Relationships returns related persons
func (s studentService) Relationships(student *Student) (*[]Relationship, error) {
	path := fmt.Sprintf("%s/%v/relationships?format=%v", studentsBasePath, student.PersonPk, format)
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

// ID returns an individual student record based on person id.
func (s studentService) ID(id string) (*Student, error) {
	type aStudent struct {
		Student `json:"student"`
	}
	var a aStudent
	path := fmt.Sprintf("%s/%s?format=%v", studentsBasePath, id, format)
	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, nil
	}
	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &a.Student, nil
}

func (s studentService) list(opt *ListOptions, basePath string) ([]Student, error) {
	// build url
	path := addOptions(basePath, format, opt)

	var students = []Student{}
	req, err := s.client.NewRequest(path)
	fmt.Println(req)
	if err != nil {
		return nil, nil
	}
	resp, err := s.client.Do(req, &students)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// handle pagination
	paginate(resp, opt)

	return students, nil
}

// List requests all students from API
func (s studentService) List(opt *ListOptions) ([]Student, error) {
	students, err := s.list(opt, studentsBasePath)
	if err != nil {
		return nil, err
	}

	return students, nil
}

// Recent requests recently updated students from API
func (s studentService) Recent(opt *ListOptions) ([]Student, error) {
	students, err := s.list(opt, studentsBasePath+"/recent")
	if err != nil {
		return nil, err
	}

	return students, nil
}
