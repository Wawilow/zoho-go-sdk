package models

import "time"

type Record struct {
	ConvertedDateTime time.Time   `json:"Converted_Date_Time"`
	Email             interface{} `json:"Email"`
	LastName          string      `json:"Last_Name"`
	Id                string      `json:"id"`
	RecordStatusS     string      `json:"Record_Status__s"`
	ConvertedS        bool        `json:"Converted__s"`
}

type RecordsList struct {
	Data []Record `json:"data"`
	Info struct {
		Call              bool        `json:"call"`
		PerPage           int         `json:"per_page"`
		NextPageToken     string      `json:"next_page_token"`
		Count             int         `json:"count"`
		SortBy            string      `json:"sort_by"`
		Page              int         `json:"page"`
		PreviousPageToken interface{} `json:"previous_page_token"`
		PageTokenExpiry   time.Time   `json:"page_token_expiry"`
		SortOrder         string      `json:"sort_order"`
		Email             bool        `json:"email"`
		MoreRecords       bool        `json:"more_records"`
	} `json:"info"`
}

type NewRecords struct {
	Data []struct {
		Layout struct {
			Id string `json:"id"`
		} `json:"Layout"`
		LeadSource           string   `json:"Lead_Source"`
		Company              string   `json:"Company"`
		LastName             string   `json:"Last_Name"`
		FirstName            string   `json:"First_Name"`
		Email                string   `json:"Email"`
		State                string   `json:"State"`
		WizardConnectionPath []string `json:"$wizard_connection_path,omitempty"`
		Wizard               struct {
			Id string `json:"id"`
		} `json:"Wizard,omitempty"`
	} `json:"data"`
	ApplyFeatureExecution []struct {
		Name string `json:"name"`
	} `json:"apply_feature_execution"`
	Trigger []string `json:"trigger"`
}

type NewRecordsResponse struct {
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
			ApprovalState string `json:"$approval_state"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type UpdateRecord struct {
	Data []struct {
		Id               string   `json:"id"`
		DealName         string   `json:"Deal_Name"`
		Stage            string   `json:"Stage"`
		Pipeline         string   `json:"Pipeline"`
		ForeignLanguages []string `json:"Foreign_Languages"`
		AppendValues     struct {
			ForeignLanguages bool `json:"Foreign_Languages"`
		} `json:"$append_values"`
		ImageUpload []struct {
			FileIdS string      `json:"File_Id__s,omitempty"`
			Id      string      `json:"id,omitempty"`
			Delete  interface{} `json:"_delete"`
		} `json:"Image_Upload"`
		Listings []struct {
			InterestedListings struct {
				Id string `json:"id"`
			} `json:"Interested_Listings,omitempty"`
			Id     string      `json:"id,omitempty"`
			Delete interface{} `json:"_delete"`
		} `json:"Listings"`
		AssociatedUser []struct {
			User struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"User,omitempty"`
			Id     string      `json:"id,omitempty"`
			Delete interface{} `json:"_delete"`
		} `json:"AssociatedUser"`
	} `json:"data"`
}

type UpdateRecordResponse struct {
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

type DeleteRecord struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type UpsertRecord struct {
	Data []struct {
		LastName         string   `json:"Last_Name"`
		Email            string   `json:"Email"`
		Company          string   `json:"Company,omitempty"`
		LeadStatus       string   `json:"Lead_Status"`
		ForeignLanguages []string `json:"Foreign_Languages"`
		AppendValues     struct {
			ForeignLanguages bool `json:"Foreign_Languages"`
		} `json:"$append_values"`
		FirstName string `json:"First_Name,omitempty"`
		Mobile    string `json:"Mobile,omitempty"`
	} `json:"data"`
	DuplicateCheckFields []string `json:"duplicate_check_fields"`
	Trigger              []string `json:"trigger"`
}

type UpsertRecordResponse struct {
	Data []struct {
		Code           string  `json:"code"`
		DuplicateField *string `json:"duplicate_field"`
		Action         string  `json:"action"`
		Details        struct {
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

type CloneReport struct {
	Data []struct {
		LastName       string `json:"Last_Name"`
		Company        string `json:"Company"`
		Email          string `json:"Email"`
		ProjectDetails []struct {
			ProjectName    string `json:"Project_Name"`
			ProjectType    string `json:"Project_Type"`
			ExpectedBudget int    `json:"Expected_Budget"`
			Status         string `json:"Status"`
		} `json:"Project_Details"`
	} `json:"data"`
}

type CloneReportResponse struct {
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
			ApprovalState string `json:"$approval_state"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type RichTextField struct {
	MultiLineRTF1 string `json:"Multi_Line_RTF1"`
	MultiLineRTF2 string `json:"Multi_Line_RTF2"`
	Id            string `json:"id"`
}

type RichTextFieldList struct {
	Data []RichTextField `json:"data"`
}

type SearchRecordCriteria struct {
	Data []struct {
		Account interface{} `json:"Account"`
		Company string      `json:"Company"`
		Owner   struct {
			Name  string `json:"name"`
			Id    string `json:"id"`
			Email string `json:"email"`
		} `json:"Owner"`
		Email             interface{} `json:"Email"`
		CurrencySymbol    string      `json:"$currency_symbol"`
		FieldStates       interface{} `json:"$field_states"`
		SharingPermission string      `json:"$sharing_permission"`
		LastActivityTime  interface{} `json:"Last_Activity_Time"`
		Industry          interface{} `json:"Industry"`
		UnsubscribedMode  interface{} `json:"Unsubscribed_Mode"`
		ProcessFlow       bool        `json:"$process_flow"`
		Street            interface{} `json:"Street"`
		LockedForMe       bool        `json:"$locked_for_me"`
		ZipCode           interface{} `json:"Zip_Code"`
		Id                string      `json:"id"`
		Approval          struct {
			Delegate bool `json:"delegate"`
			Approve  bool `json:"approve"`
			Reject   bool `json:"reject"`
			Resubmit bool `json:"resubmit"`
		} `json:"$approval"`
		CreatedTime          time.Time   `json:"Created_Time"`
		WizardConnectionPath interface{} `json:"$wizard_connection_path"`
		Editable             bool        `json:"$editable"`
		City                 interface{} `json:"City"`
		NoOfEmployees        interface{} `json:"No_of_Employees"`
		Longitude            interface{} `json:"Longitude"`
		ConvertedAccount     interface{} `json:"Converted_Account"`
		State                interface{} `json:"State"`
		Country              interface{} `json:"Country"`
		CreatedBy            struct {
			Name  string `json:"name"`
			Id    string `json:"id"`
			Email string `json:"email"`
		} `json:"Created_By"`
		ZiaOwnerAssignment string      `json:"$zia_owner_assignment"`
		SecondaryEmail     interface{} `json:"Secondary_Email"`
		AnnualRevenue      interface{} `json:"Annual_Revenue"`
		Currency1          interface{} `json:"Currency_1"`
		Description        interface{} `json:"Description"`
		Rating             interface{} `json:"Rating"`
		ReviewProcess      struct {
			Approve  bool `json:"approve"`
			Reject   bool `json:"reject"`
			Resubmit bool `json:"resubmit"`
		} `json:"$review_process"`
		Website            interface{} `json:"Website"`
		Twitter            interface{} `json:"Twitter"`
		FileUpload         interface{} `json:"File_Upload"`
		AssociatedContacts interface{} `json:"Associated_Contacts"`
		Salutation         interface{} `json:"Salutation"`
		CoOwner            interface{} `json:"Co_Owner"`
		FullName           string      `json:"Full_Name"`
		FirstName          string      `json:"First_Name"`
		LeadStatus         interface{} `json:"Lead_Status"`
		RecordImage        interface{} `json:"Record_Image"`
		ConvertedDeal      interface{} `json:"Converted_Deal"`
		ModifiedBy         struct {
			Name  string `json:"name"`
			Id    string `json:"id"`
			Email string `json:"email"`
		} `json:"Modified_By"`
		Review             interface{} `json:"$review"`
		LeadConversionTime interface{} `json:"Lead_Conversion_Time"`
		SkypeID            interface{} `json:"Skype_ID"`
		Phone              interface{} `json:"Phone"`
		ImageUpload1       interface{} `json:"Image_Upload_1"`
		EmailOptOut        bool        `json:"Email_Opt_Out"`
		ZiaVisions         interface{} `json:"$zia_visions"`
		Designation        interface{} `json:"Designation"`
		ModifiedTime       time.Time   `json:"Modified_Time"`
		ConvertedDetail    struct {
		} `json:"$converted_detail"`
		UnsubscribedTime interface{} `json:"Unsubscribed_Time"`
		ConvertedContact interface{} `json:"Converted_Contact"`
		Mobile           interface{} `json:"Mobile"`
		RecordStatusS    string      `json:"Record_Status__s"`
		Orchestration    bool        `json:"$orchestration"`
		LastName         string      `json:"Last_Name"`
		Layout           struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Layout"`
		InMerge       bool          `json:"$in_merge"`
		LockedS       bool          `json:"Locked__s"`
		LeadSource    interface{}   `json:"Lead_Source"`
		Tag           []interface{} `json:"Tag"`
		Fax           interface{}   `json:"Fax"`
		ApprovalState string        `json:"$approval_state"`
		Pathfinder    bool          `json:"$pathfinder"`
	} `json:"data"`
	Info struct {
		PerPage     int    `json:"per_page"`
		Count       int    `json:"count"`
		SortBy      string `json:"sort_by"`
		Page        int    `json:"page"`
		SortOrder   string `json:"sort_order"`
		MoreRecords bool   `json:"more_records"`
	} `json:"info"`
}

type SearchRecordEmail struct {
	Data []struct {
		Account interface{} `json:"Account"`
		Owner   struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Owner"`
		Company          interface{} `json:"Company"`
		Email            string      `json:"Email"`
		CurrencySymbol   string      `json:"$currency_symbol"`
		VisitorScore     interface{} `json:"Visitor_Score"`
		LastActivityTime time.Time   `json:"Last_Activity_Time"`
		Industry         string      `json:"Industry"`
		Converted        bool        `json:"$converted"`
		ProcessFlow      bool        `json:"$process_flow"`
		Street           string      `json:"Street"`
		ZipCode          string      `json:"Zip_Code"`
		Id               string      `json:"id"`
		Approved         bool        `json:"$approved"`
		Approval         struct {
			Delegate bool `json:"delegate"`
			Approve  bool `json:"approve"`
			Reject   bool `json:"reject"`
			Resubmit bool `json:"resubmit"`
		} `json:"$approval"`
		FirstVisitedURL interface{} `json:"First_Visited_URL"`
		DaysVisited     interface{} `json:"Days_Visited"`
		CreatedTime     time.Time   `json:"Created_Time"`
		Editable        bool        `json:"$editable"`
		City            string      `json:"City"`
		NoOfEmployees   int         `json:"No_of_Employees"`
		State           string      `json:"State"`
		Country         string      `json:"Country"`
		LastVisitedTime interface{} `json:"Last_Visited_Time"`
		CreatedBy       struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Created_By"`
		AnnualRevenue           int         `json:"Annual_Revenue"`
		SecondaryEmail          string      `json:"Secondary_Email"`
		Description             string      `json:"Description"`
		NumberOfChats           interface{} `json:"Number_Of_Chats"`
		Rating                  string      `json:"Rating"`
		Website                 string      `json:"Website"`
		Twitter                 string      `json:"Twitter"`
		AverageTimeSpentMinutes interface{} `json:"Average_Time_Spent_Minutes"`
		AssociatedContacts      interface{} `json:"Associated_Contacts"`
		Salutation              string      `json:"Salutation"`
		FirstName               string      `json:"First_Name"`
		LeadStatus              string      `json:"Lead_Status"`
		FullName                string      `json:"Full_Name"`
		RecordImage             interface{} `json:"Record_Image"`
		ModifiedBy              struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Modified_By"`
		SkypeID         string    `json:"Skype_ID"`
		Phone           string    `json:"Phone"`
		EmailOptOut     bool      `json:"Email_Opt_Out"`
		Designation     string    `json:"Designation"`
		ModifiedTime    time.Time `json:"Modified_Time"`
		ConvertedDetail struct {
		} `json:"$converted_detail"`
		Mobile           string        `json:"Mobile"`
		PredictionScore  interface{}   `json:"Prediction_Score"`
		FirstVisitedTime interface{}   `json:"First_Visited_Time"`
		LastName         string        `json:"Last_Name"`
		Referrer         interface{}   `json:"Referrer"`
		LeadSource       string        `json:"Lead_Source"`
		Tag              []interface{} `json:"Tag"`
		Fax              string        `json:"Fax"`
	} `json:"data"`
	Info struct {
		PerPage     int  `json:"per_page"`
		Count       int  `json:"count"`
		Page        int  `json:"page"`
		MoreRecords bool `json:"more_records"`
	} `json:"info"`
}

type SearchRecordPhone struct {
	Data []struct {
		Account interface{} `json:"Account"`
		Owner   struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Owner"`
		Company          interface{} `json:"Company"`
		Email            string      `json:"Email"`
		CurrencySymbol   string      `json:"$currency_symbol"`
		VisitorScore     interface{} `json:"Visitor_Score"`
		LastActivityTime time.Time   `json:"Last_Activity_Time"`
		Industry         string      `json:"Industry"`
		Converted        bool        `json:"$converted"`
		ProcessFlow      bool        `json:"$process_flow"`
		Street           string      `json:"Street"`
		ZipCode          string      `json:"Zip_Code"`
		Id               string      `json:"id"`
		Approved         bool        `json:"$approved"`
		Approval         struct {
			Delegate bool `json:"delegate"`
			Approve  bool `json:"approve"`
			Reject   bool `json:"reject"`
			Resubmit bool `json:"resubmit"`
		} `json:"$approval"`
		FirstVisitedURL interface{} `json:"First_Visited_URL"`
		DaysVisited     interface{} `json:"Days_Visited"`
		CreatedTime     time.Time   `json:"Created_Time"`
		Editable        bool        `json:"$editable"`
		City            string      `json:"City"`
		NoOfEmployees   int         `json:"No_of_Employees"`
		State           string      `json:"State"`
		Country         string      `json:"Country"`
		LastVisitedTime interface{} `json:"Last_Visited_Time"`
		CreatedBy       struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Created_By"`
		AnnualRevenue           int         `json:"Annual_Revenue"`
		SecondaryEmail          string      `json:"Secondary_Email"`
		Description             string      `json:"Description"`
		NumberOfChats           interface{} `json:"Number_Of_Chats"`
		Rating                  string      `json:"Rating"`
		Website                 string      `json:"Website"`
		Twitter                 string      `json:"Twitter"`
		AverageTimeSpentMinutes interface{} `json:"Average_Time_Spent_Minutes"`
		AssociatedContacts      interface{} `json:"Associated_Contacts"`
		Salutation              string      `json:"Salutation"`
		FirstName               string      `json:"First_Name"`
		LeadStatus              string      `json:"Lead_Status"`
		FullName                string      `json:"Full_Name"`
		RecordImage             interface{} `json:"Record_Image"`
		ModifiedBy              struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Modified_By"`
		SkypeID         string    `json:"Skype_ID"`
		Phone           string    `json:"Phone"`
		EmailOptOut     bool      `json:"Email_Opt_Out"`
		Designation     string    `json:"Designation"`
		ModifiedTime    time.Time `json:"Modified_Time"`
		ConvertedDetail struct {
		} `json:"$converted_detail"`
		Mobile           string        `json:"Mobile"`
		PredictionScore  interface{}   `json:"Prediction_Score"`
		FirstVisitedTime interface{}   `json:"First_Visited_Time"`
		LastName         string        `json:"Last_Name"`
		Referrer         interface{}   `json:"Referrer"`
		LeadSource       string        `json:"Lead_Source"`
		Tag              []interface{} `json:"Tag"`
		Fax              string        `json:"Fax"`
	} `json:"data"`
	Info struct {
		PerPage     int  `json:"per_page"`
		Count       int  `json:"count"`
		Page        int  `json:"page"`
		MoreRecords bool `json:"more_records"`
	} `json:"info"`
}

type SearchRecordSingle struct {
	Data []struct {
		Type         string `json:"Type"`
		CampaignName string `json:"Campaign_Name"`
		Id           string `json:"id"`
	} `json:"data"`
	Info struct {
		PerPage     int  `json:"per_page"`
		Count       int  `json:"count"`
		Page        int  `json:"page"`
		MoreRecords bool `json:"more_records"`
	} `json:"info"`
}

type RecordsDeletedList struct {
	Data []struct {
		DeletedBy *struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"deleted_by"`
		Id          string  `json:"id"`
		DisplayName *string `json:"display_name"`
		Type        string  `json:"type"`
		CreatedBy   *struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"created_by"`
		DeletedTime time.Time `json:"deleted_time"`
	} `json:"data"`
	Info struct {
		PerPage     int  `json:"per_page"`
		Count       int  `json:"count"`
		Page        int  `json:"page"`
		MoreRecords bool `json:"more_records"`
	} `json:"info"`
}

type RecordsCount struct {
	Count int `json:"count"`
}

type RecordsTimeline struct {
	Timeline []struct {
		DoneBy struct {
			Profile struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"profile"`
			Name  string `json:"name"`
			Id    string `json:"id"`
			TypeS string `json:"type__s"`
		} `json:"done_by"`
		RelatedRecord *struct {
			Module struct {
				ApiName string `json:"api_name"`
				Id      string `json:"id"`
			} `json:"module"`
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"related_record"`
		AutomationDetails interface{} `json:"automation_details"`
		Record            struct {
			Module struct {
				ApiName string `json:"api_name"`
				Id      string `json:"id"`
			} `json:"module"`
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"record"`
		AuditedTime  time.Time `json:"audited_time"`
		Action       string    `json:"action"`
		Id           string    `json:"id"`
		Source       string    `json:"source"`
		FieldHistory []struct {
			ApiName          string `json:"api_name"`
			FieldLabel       string `json:"field_label"`
			EnableColourCode bool   `json:"enable_colour_code"`
			DataType         string `json:"data_type"`
			Value            struct {
				New string  `json:"new"`
				Old *string `json:"old"`
			} `json:"_value"`
			Id             string `json:"id"`
			PickListValues []struct {
				DisplayValue   string      `json:"display_value"`
				SequenceNumber int         `json:"sequence_number"`
				ColourCode     interface{} `json:"colour_code"`
				ActualValue    string      `json:"actual_value"`
				Id             string      `json:"id"`
				Type           string      `json:"type"`
			} `json:"pick_list_values"`
		} `json:"field_history"`
	} `json:"__timeline"`
	Info struct {
		PerPage           int         `json:"per_page"`
		NextPageToken     interface{} `json:"next_page_token"`
		Count             int         `json:"count"`
		Page              int         `json:"page"`
		PreviousPageToken interface{} `json:"previous_page_token"`
		MoreRecords       bool        `json:"more_records"`
	} `json:"info"`
}

type RelatedRecordsList struct {
	Data []struct {
		Owner struct {
			Name  string `json:"name"`
			Id    string `json:"id"`
			Email string `json:"email"`
		} `json:"Owner"`
		ParentId struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Parent_Id"`
		Id string `json:"id"`
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

type UpdateRelatedRecord struct {
	Data []struct {
		Id           string `json:"id"`
		MemberStatus string `json:"Member_Status"`
	} `json:"data"`
}

type UpdateRelatedRecordResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type DelRelatedRecords struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type ExternalRecord struct {
	Data []struct {
		FirstName         interface{} `json:"First_Name"`
		FullName          string      `json:"Full_Name"`
		LastName          string      `json:"Last_Name"`
		ExternalContactID string      `json:"External_Contact_ID"`
		AccountName       struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Account_Name"`
		Id          string `json:"id"`
		OrgExternal string `json:"OrgExternal"`
	} `json:"data"`
}

type ExternalRecordsList struct {
	Data []struct {
		FirstName         interface{} `json:"First_Name"`
		FullName          string      `json:"Full_Name"`
		LastName          string      `json:"Last_Name"`
		ExternalContactID string      `json:"External_Contact_ID"`
		AccountName       struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Account_Name"`
		Id          string `json:"id"`
		OrgExternal string `json:"OrgExternal"`
	} `json:"data"`
	Info struct {
		PerPage     int  `json:"per_page"`
		Count       int  `json:"count"`
		Page        int  `json:"page"`
		MoreRecords bool `json:"more_records"`
	} `json:"info"`
}

type NewExternalRecordContacts struct {
	Data []struct {
		LastName    string `json:"Last_Name"`
		AccountName struct {
			ExternalAccountID string `json:"External_Account_ID"`
		} `json:"Account_Name"`
	} `json:"data"`
}

type NewExternalRecordContactsResponse struct {
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

type NewExternalRecordQuote struct {
	Data []struct {
		ExternalQuoteID string `json:"External_Quote_ID"`
		Subject         string `json:"Subject"`
		QuotedItems     []struct {
			Product struct {
				ExternalProductID string `json:"External_Product_ID"`
			} `json:"product"`
			Book struct {
				ExternalPriceBookID string `json:"External_Price_book_ID"`
			} `json:"book"`
			Quantity int `json:"quantity"`
		} `json:"Quoted_Items"`
	} `json:"data"`
}

type NewExternalRecordQuoteResponse struct {
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

type NewExternalRecordEvents struct {
	Data []struct {
		SeModule string `json:"se_module"`
		WhatId   struct {
			ExternalAccountID string `json:"External_Account_ID"`
		} `json:"What_Id"`
		EventTitle    string    `json:"Event_Title"`
		StartDateTime time.Time `json:"Start_DateTime"`
		EndDateTime   time.Time `json:"End_DateTime"`
	} `json:"data"`
}

type NewExternalRecordEventsResponse struct {
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

type UpdateExternalRecordContact struct {
	Data []struct {
		LastName string `json:"Last_Name"`
	} `json:"data"`
}

type UpdateExternalRecordContactResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			ModifiedTime time.Time `json:"Modified_Time"`
			ModifiedBy   struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"Modified_By"`
			CreatedTime       time.Time `json:"Created_Time"`
			ExternalContactID string    `json:"External_Contact_ID"`
			Id                string    `json:"id"`
			CreatedBy         struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"Created_By"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type UpdateExternalRecordQuote struct {
	Data []struct {
		ProductDetails []struct {
			Id      string `json:"id,omitempty"`
			Product struct {
				ExternalProductID string `json:"External_Product_ID"`
			} `json:"product,omitempty"`
			Book struct {
				ExternalPricebookID string `json:"External_Pricebook_ID"`
			} `json:"book,omitempty"`
			Quantity int `json:"quantity,omitempty"`
		} `json:"Product_Details"`
	} `json:"data"`
}

type UpdateExternalRecordQuoteResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			ModifiedTime time.Time `json:"Modified_Time"`
			ModifiedBy   struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"Modified_By"`
			CreatedTime     time.Time `json:"Created_Time"`
			ExternalQuoteID string    `json:"External_Quote_ID"`
			Id              string    `json:"id"`
			CreatedBy       struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"Created_By"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type DelExternalRecord struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			ExternalContactID string `json:"External_Contact_ID"`
			Id                string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type UpsertExternalRecord struct {
	Data []struct {
		LastName          string `json:"Last_Name"`
		Email             string `json:"Email"`
		ExternalContactID string `json:"External_Contact_ID"`
	} `json:"data"`
}

type UpsertExternalRecordResponse struct {
	Data []struct {
		Code           string `json:"code"`
		DuplicateField string `json:"duplicate_field"`
		Action         string `json:"action"`
		Details        struct {
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

type SearchExternalRecord struct {
	Data []struct {
		Email             string `json:"Email"`
		ExternalContactID string `json:"External_Contact_ID"`
		Id                string `json:"id"`
	} `json:"data"`
	Info struct {
		PerPage     int  `json:"per_page"`
		Count       int  `json:"count"`
		Page        int  `json:"page"`
		MoreRecords bool `json:"more_records"`
	} `json:"info"`
}

type ExternalRelatedRecordsList struct {
	Data []struct {
		ContactRole int64  `json:"Contact_Role"`
		Stage       string `json:"Stage"`
		AccountName struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"Account_Name"`
		Id          string `json:"id"`
		ClosingDate string `json:"Closing_Date"`
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

type UpdateExternalRelatedRecord struct {
	Data []struct {
		ExternalDealID string `json:"External_Deal_ID"`
		ContactRole    string `json:"Contact_Role"`
	} `json:"data"`
}

type UpdateExternalRelatedRecordResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			ExternalDealID string `json:"External_Deal_ID"`
			Id             string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type DelUpdateExternalRelated struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type MergeRecords struct {
	Merge []struct {
		MasterRecordFields []struct {
			ApiName string `json:"api_name"`
			Data    []struct {
				Id string `json:"id"`
			} `json:"_data,omitempty"`
		} `json:"master_record_fields"`
		Data []struct {
			Id     string `json:"id"`
			Fields []struct {
				ApiName string `json:"api_name"`
			} `json:"_fields"`
		} `json:"data"`
	} `json:"merge"`
}

type MergeRecordsResponse struct {
	Merge []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"merge"`
}

type ScheduledMergeRecordsResponse struct {
	Merge []struct {
		Code    string `json:"code"`
		Details struct {
			JobId string `json:"job_id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"merge"`
}

type MergeRecordsStatus struct {
	Merge []struct {
		JobId  string `json:"job_id"`
		Status string `json:"status"`
	} `json:"merge"`
}
