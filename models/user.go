package models

import "time"

type User struct {
	Users []struct {
		Country            *string     `json:"country"`
		NameFormatS        string      `json:"name_format__s,omitempty"`
		Language           string      `json:"language"`
		Microsoft          bool        `json:"microsoft"`
		ShiftEffectiveFrom interface{} `json:"$shift_effective_from"`
		Currency           string      `json:"Currency"`
		Id                 string      `json:"id"`
		State              *string     `json:"state"`
		Fax                interface{} `json:"fax"`
		CountryLocale      string      `json:"country_locale"`
		SandboxDeveloper   bool        `json:"sandboxDeveloper"`
		Zip                interface{} `json:"zip"`
		DecimalSeparator   string      `json:"decimal_separator"`
		CreatedTime        time.Time   `json:"created_time"`
		TimeFormat         string      `json:"time_format"`
		Offset             int         `json:"offset"`
		Profile            struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"profile"`
		CreatedBy struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"created_by"`
		Zuid                 *string     `json:"zuid"`
		FullName             string      `json:"full_name"`
		Phone                interface{} `json:"phone"`
		Dob                  interface{} `json:"dob"`
		SortOrderPreferenceS string      `json:"sort_order_preference__s"`
		Status               string      `json:"status"`
		Role                 struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"role"`
		CustomizeInfo struct {
			NotesDesc       bool        `json:"notes_desc"`
			ShowRightPanel  interface{} `json:"show_right_panel"`
			BcView          interface{} `json:"bc_view"`
			ShowHome        bool        `json:"show_home"`
			ShowDetailView  bool        `json:"show_detail_view"`
			UnpinRecentItem interface{} `json:"unpin_recent_item"`
		} `json:"customize_info,omitempty"`
		City            interface{} `json:"city"`
		Signature       interface{} `json:"signature"`
		Locale          string      `json:"locale"`
		PersonalAccount bool        `json:"personal_account,omitempty"`
		Isonline        bool        `json:"Isonline"`
		DefaultTabGroup string      `json:"default_tab_group,omitempty"`
		ModifiedBy      struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Modified_By"`
		Street       interface{} `json:"street"`
		CurrentShift interface{} `json:"$current_shift"`
		Alias        interface{} `json:"alias"`
		Theme        struct {
			NormalTab struct {
				FontColor  string `json:"font_color"`
				Background string `json:"background"`
			} `json:"normal_tab"`
			SelectedTab struct {
				FontColor  string `json:"font_color"`
				Background string `json:"background"`
			} `json:"selected_tab"`
			NewBackground interface{} `json:"new_background"`
			Background    string      `json:"background"`
			Screen        string      `json:"screen"`
			Type          string      `json:"type"`
		} `json:"theme,omitempty"`
		FirstName       *string     `json:"first_name"`
		Email           string      `json:"email"`
		StatusReasonS   interface{} `json:"status_reason__s"`
		Website         interface{} `json:"website"`
		ModifiedTime    time.Time   `json:"Modified_Time"`
		NextShift       interface{} `json:"$next_shift"`
		Mobile          interface{} `json:"mobile"`
		LastName        string      `json:"last_name"`
		TimeZone        string      `json:"time_zone"`
		NumberSeparator string      `json:"number_separator"`
		Confirm         bool        `json:"confirm"`
		DateFormat      string      `json:"date_format"`
		Category        string      `json:"category"`
	} `json:"users"`
	Info struct {
		PerPage     int  `json:"per_page"`
		Count       int  `json:"count"`
		Page        int  `json:"page"`
		MoreRecords bool `json:"more_records"`
	} `json:"info"`
}

type NewUser struct {
	Role                 string `json:"role"`
	FirstName            string `json:"first_name"`
	Email                string `json:"email"`
	Profile              string `json:"profile"`
	LastName             string `json:"last_name"`
	NameFormatS          string `json:"name_format__s"`
	SortOrderPreferenceS string `json:"sort_order_preference__s"`
}

type NewUserResponse struct {
	Code    string `json:"code"`
	Details struct {
		Id string `json:"id"`
	} `json:"details"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type UpdateUser struct {
	Id                   string `json:"id"`
	Phone                string `json:"phone"`
	Dob                  string `json:"dob"`
	Role                 string `json:"role"`
	Profile              string `json:"profile"`
	CountryLocale        string `json:"country_locale"`
	TimeFormat           string `json:"time_format"`
	TimeZone             string `json:"time_zone"`
	NameFormatS          string `json:"name_format__s"`
	SortOrderPreferenceS string `json:"sort_order_preference__s"`
}

type UpdateUserResponse struct {
	Code    string `json:"code"`
	Details struct {
		Id string `json:"id"`
	} `json:"details"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type DelUsers struct {
	Code    string `json:"code"`
	Details struct {
	} `json:"details"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type UserCount struct {
	Count int `json:"count"`
}

type UserTerritories struct {
	Territories []struct {
		Id      string `json:"id"`
		Manager struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Manager"`
		Name        string `json:"Name"`
		ReportingTo *struct {
			Id   string `json:"id"`
			Name string `json:"Name"`
		} `json:"Reporting_To"`
	} `json:"territories"`
	Info struct {
		PerPage     int  `json:"per_page"`
		Count       int  `json:"count"`
		Page        int  `json:"page"`
		MoreRecords bool `json:"more_records"`
	} `json:"info"`
}

type AddUserTerritories struct {
	Territories []struct {
		Id string `json:"id"`
	} `json:"territories"`
}

type AddUserTerritoriesResponse struct {
	Territories []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"territories"`
}

type TransferAndDelete struct {
	Id       string `json:"id"`
	Transfer struct {
		Id         string `json:"id"`
		Records    bool   `json:"records"`
		Assignment bool   `json:"assignment"`
		Criteria   bool   `json:"criteria"`
	} `json:"transfer"`
	MoveSubordinate struct {
		Id string `json:"id"`
	} `json:"move_subordinate"`
}

type TransferAndDeleteResponse struct {
	Code    string `json:"code"`
	Details struct {
		JobId string `json:"jobId"`
		Id    string `json:"id"`
	} `json:"details"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
