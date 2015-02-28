package vcapi

type Person struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	EmailOne    string `json:"email_1"`
	EmailTwo    string `json:"email_2"`
	PersonPk    int    `json:"person_pk"`
	HouseholdFk int    `json:"household_fk"`
	Username    string `json:"username"`
}
