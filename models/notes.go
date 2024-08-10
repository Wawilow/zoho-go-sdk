package models

import "time"

type NewNote struct {
	Data []struct {
		ParentId struct {
			Module struct {
				ApiName string `json:"api_name"`
				Id      string `json:"id"`
			} `json:"module"`
			Id string `json:"id"`
		} `json:"Parent_Id"`
		NoteContent string `json:"Note_Content"`
	} `json:"data"`
}

type NewNoteResponse struct {
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

type Note struct {
	Owner struct {
		Name  string `json:"name"`
		Id    string `json:"id"`
		Email string `json:"email"`
	} `json:"Owner"`
	CreatedTime time.Time `json:"Created_Time"`
	ParentId    struct {
		Module struct {
			ApiName string `json:"api_name"`
			Id      string `json:"id"`
		} `json:"module"`
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"Parent_Id"`
	Id        string  `json:"id"`
	NoteTitle *string `json:"Note_Title"`
}

type NotesList struct {
	Data []Note `json:"data"`
	Info struct {
		PerPage           int         `json:"per_page"`
		NextPageToken     interface{} `json:"next_page_token"`
		Count             int         `json:"count"`
		SortBy            string      `json:"sort_by"`
		Page              int         `json:"page"`
		PreviousPageToken interface{} `json:"previous_page_token"`
		PageTokenExpiry   interface{} `json:"page_token_expiry"`
		SortOrder         string      `json:"sort_order"`
		MoreRecords       bool        `json:"more_records"`
	} `json:"info"`
}

type UpdateNote struct {
	Data []struct {
		NoteTitle   string `json:"Note_Title"`
		NoteContent string `json:"Note_Content"`
	} `json:"data"`
}

type UpdateNoteResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			CreatedTime  time.Time `json:"created_time"`
			ModifiedTime time.Time `json:"modified_time"`
			ModifiedBy   struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"modified_by"`
			Id        string `json:"id"`
			CreatedBy struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"created_by"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type DelNote struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}
