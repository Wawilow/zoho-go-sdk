package models

import "time"

type AttachmentsList struct {
	Data []struct {
		Owner struct {
			Name  string `json:"name"`
			Id    string `json:"id"`
			Email string `json:"email"`
		} `json:"Owner"`
		FileName    string      `json:"File_Name"`
		CreatedTime time.Time   `json:"Created_Time"`
		ParentId    interface{} `json:"Parent_Id"`
		Id          string      `json:"id"`
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

type NewAttachment struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			ModifiedTime time.Time `json:"Modified_Time"`
			ModifiedBy   struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"Modified_By"`
			CreatedTime time.Time `json:"Created_Time"`
			Id          string    `json:"id"`
			CreatedBy   struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"Created_By"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type DelAttachment struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type NewPhoto struct {
	Code    string `json:"code"`
	Details struct {
	} `json:"details"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type DelPhoto struct {
	Code    string `json:"code"`
	Details struct {
	} `json:"details"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
