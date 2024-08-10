package models

type ContactRole struct {
	SequenceNumber int    `json:"sequence_number"`
	Name           string `json:"name"`
	Id             string `json:"id"`
}

type ContactRolesList struct {
	ContactRoles []ContactRole `json:"contact_roles"`
}

type NewContactRole struct {
	ContactRoles []struct {
		Name           string `json:"name"`
		SequenceNumber int    `json:"sequence_number"`
	} `json:"contact_roles"`
}

type NewContactRoleResponse struct {
	ContactRoles []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"contact_roles"`
}

type UpdateContactRole struct {
	ContactRoles []struct {
		Name           string `json:"name"`
		SequenceNumber int    `json:"sequence_number"`
		Id             string `json:"id"`
	} `json:"contact_roles"`
}

type UpdateContactRoleResponse struct {
	ContactRoles []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"contact_roles"`
}

type DelContactRole struct {
	ContactRoles []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"contact_roles"`
}

type DealContactRole struct {
	Data []struct {
		Department  *string `json:"Department"`
		ContactRole struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Contact_Role"`
		Email string `json:"Email"`
		Id    string `json:"id"`
	} `json:"data"`
	Info struct {
		PerPage           int         `json:"per_page"`
		NextPageToken     interface{} `json:"next_page_token"`
		Count             int         `json:"count"`
		Page              int         `json:"page"`
		PreviousPageToken interface{} `json:"previous_page_token"`
		PageTokenExpiry   interface{} `json:"page_token_expiry"`
		MoreRecords       bool        `json:"more_records"`
	} `json:"info"`
}

type NewDealContactRole struct {
	Data []struct {
		ContactRole struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Contact_Role"`
	} `json:"data"`
}

type NewDealContactRoleResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type DelDealContactRole struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}
