package ob_dto

type CreateParams struct {
	LowCase LowCase `json:"lowerCaseTableNames"`
}

type ObDBRole struct {
	Database   string             `json:"database"`
	Role       ObDataBaseRoleType `json:"role"`
	Privileges string             `json:"privileges"` // get from ObDataBasePrivilegeType
}
