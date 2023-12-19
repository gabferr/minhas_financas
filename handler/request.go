package handler

import "fmt"

//CreateOpening

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"`
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Company == "" && r.Link == "" && r.Location == "" && r.Remote == nil && r.Role == "" && r.Salary <= 0 {
		return fmt.Errorf("Request body is empty")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}

	return nil
}