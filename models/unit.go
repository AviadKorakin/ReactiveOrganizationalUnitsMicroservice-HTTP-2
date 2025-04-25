package models

// Unit represents an organizational unit in MongoDB
// @Description Unit entity with auto-generated ID and creation date
type Unit struct {
	// @ID
	UnitID string `bson:"unitId" json:"unitId"`
	// @Name of the unit
	Name string `bson:"name" json:"name"`
	// @CreationDate in dd-MM-yyyy format
	CreationDate string `bson:"creationDate" json:"creationDate"`
	// @Description of this unit
	Description string `bson:"description" json:"description"`
	// @Users assigned to this unit (email identifiers)
	Users []string `bson:"users,omitempty" json:"users,omitempty"`
}

// UnitBoundary for JSON input/output
type UnitBoundary struct {
	UnitID       string `json:"unitId,omitempty"`
	Name         string `json:"name"`
	CreationDate string `json:"creationDate,omitempty"`
	Description  string `json:"description"`
}
