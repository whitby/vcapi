package vcapi

// Relationship represents a Veracross person relationship
type Relationship struct {
	CustodyStatus      string `json:"custody_status"`
	EmergencyContact   bool   `json:"emergency_contact"`
	FirstName          string `json:"first_name"`
	IsParent           bool   `json:"is_parent"`
	LastName           string `json:"last_name"`
	LegalCustody       bool   `json:"legal_custody"`
	Notes              string `json:"notes"`
	ParentPortalAccess bool   `json:"parent_portal_access"`
	PickUp             bool   `json:"pick_up"`
	RelatedPersonFk    int    `json:"related_person_fk"`
	Relationship       string `json:"relationship"`
	Resident           bool   `json:"resident"`
}
