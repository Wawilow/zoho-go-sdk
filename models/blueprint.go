package models

type Blueprint struct {
	Blueprint struct {
		ProcessInfo struct {
			FieldId      string      `json:"field_id"`
			Escalation   interface{} `json:"escalation"`
			IsContinuous bool        `json:"is_continuous"`
			ApiName      string      `json:"api_name"`
			Continuous   bool        `json:"continuous"`
			FieldLabel   string      `json:"field_label"`
			Name         string      `json:"name"`
			ColumnName   string      `json:"column_name"`
			FieldValue   string      `json:"field_value"`
			Id           string      `json:"id"`
			FieldName    string      `json:"field_name"`
		} `json:"process_info"`
		Transitions []struct {
			NextTransitions []interface{} `json:"next_transitions"`
			Data            struct {
				MultiUserLookup           []interface{} `json:"Multi_User_Lookup,omitempty"`
				MultiSelectLookupContacts []interface{} `json:"Multi_Select_Lookup_Contacts,omitempty"`
				Widget                    struct {
					Name string `json:"name"`
					Id   string `json:"id"`
				} `json:"widget,omitempty"`
			} `json:"data"`
			NextFieldValue  string `json:"next_field_value"`
			Name            string `json:"name"`
			CriteriaMatched bool   `json:"criteria_matched"`
			Id              string `json:"id"`
			Fields          []struct {
				SystemMandatory bool        `json:"system_mandatory,omitempty"`
				Private         interface{} `json:"private"`
				Webhook         bool        `json:"webhook,omitempty"`
				JsonType        string      `json:"json_type,omitempty"`
				Crypt           interface{} `json:"crypt"`
				FieldLabel      string      `json:"field_label,omitempty"`
				Tooltip         interface{} `json:"tooltip"`
				CreatedSource   string      `json:"created_source,omitempty"`
				Layouts         *struct {
					Name string `json:"name"`
					Id   string `json:"id"`
				} `json:"layouts"`
				FieldReadOnly      bool          `json:"field_read_only,omitempty"`
				Content            []interface{} `json:"content,omitempty"`
				DisplayLabel       *string       `json:"display_label"`
				DisplayType        int           `json:"display_type,omitempty"`
				UiType             int           `json:"ui_type,omitempty"`
				ValidationRule     interface{}   `json:"validation_rule"`
				ReadOnly           bool          `json:"read_only,omitempty"`
				AssociationDetails interface{}   `json:"association_details"`
				MultiModuleLookup  struct {
				} `json:"multi_module_lookup,omitempty"`
				Currency struct {
				} `json:"currency,omitempty"`
				Id              string `json:"id"`
				Multiuserlookup struct {
					DisplayLabel           string      `json:"display_label"`
					LinkingModule          string      `json:"linking_module"`
					LookupApiname          string      `json:"lookup_apiname"`
					ConnectedModule        interface{} `json:"connected_module"`
					ApiName                string      `json:"api_name"`
					ConnectedlookupApiname string      `json:"connectedlookup_apiname"`
					Id                     string      `json:"id"`
				} `json:"multiuserlookup,omitempty"`
				CustomField bool `json:"custom_field,omitempty"`
				Lookup      struct {
				} `json:"lookup,omitempty"`
				ConvertMapping struct {
					Contacts interface{} `json:"Contacts"`
					Deals    interface{} `json:"Deals"`
					Accounts interface{} `json:"Accounts"`
				} `json:"convert_mapping,omitempty"`
				Visible  bool `json:"visible,omitempty"`
				Profiles []struct {
					PermissionType string `json:"permission_type"`
					Name           string `json:"name"`
					Id             string `json:"id"`
				} `json:"profiles,omitempty"`
				Length     int     `json:"length,omitempty"`
				ColumnName *string `json:"column_name"`
				Type       string  `json:"_type"`
				ViewType   struct {
					View        bool `json:"view"`
					Edit        bool `json:"edit"`
					QuickCreate bool `json:"quick_create"`
					Create      bool `json:"create"`
				} `json:"view_type,omitempty"`
				PickListValuesSortedLexically bool        `json:"pick_list_values_sorted_lexically,omitempty"`
				Sortable                      bool        `json:"sortable,omitempty"`
				TransitionSequence            int         `json:"transition_sequence"`
				External                      interface{} `json:"external"`
				ApiName                       string      `json:"api_name,omitempty"`
				Unique                        struct {
				} `json:"unique,omitempty"`
				HistoryTracking interface{} `json:"history_tracking"`
				DataType        string      `json:"data_type"`
				Formula         struct {
				} `json:"formula,omitempty"`
				DecimalPlace   interface{}   `json:"decimal_place"`
				PickListValues []interface{} `json:"pick_list_values,omitempty"`
				AutoNumber     struct {
				} `json:"auto_number,omitempty"`
				Multiselectlookup struct {
					DisplayLabel           string `json:"display_label"`
					LinkingModule          string `json:"linking_module"`
					LookupApiname          string `json:"lookup_apiname"`
					ConnectedModule        string `json:"connected_module"`
					ApiName                string `json:"api_name"`
					ConnectedlookupApiname string `json:"connectedlookup_apiname"`
					Id                     string `json:"id"`
				} `json:"multiselectlookup,omitempty"`
				PersonalityName interface{} `json:"personality_name"`
				Mandatory       bool        `json:"mandatory,omitempty"`
			} `json:"fields"`
			Type               string      `json:"type"`
			CriteriaMessage    interface{} `json:"criteria_message"`
			PercentPartialSave int         `json:"percent_partial_save"`
			ExecutionTime      interface{} `json:"execution_time"`
		} `json:"transitions"`
	} `json:"blueprint"`
}

type UpdateBlueprint struct {
	Blueprint []struct {
		TransitionId string `json:"transition_id"`
		Data         struct {
			Attachments []struct {
				FileId  []string `json:"$file_id,omitempty"`
				LinkUrl string   `json:"$link_url,omitempty"`
			} `json:"Attachments"`
			Phone    int64  `json:"Phone"`
			Notes    string `json:"Notes"`
			Listings []struct {
				InterestedListings struct {
					Id string `json:"id"`
				} `json:"Interested_Listings"`
			} `json:"Listings"`
			MultiUser []struct {
				MultiUser struct {
					Name string `json:"name,omitempty"`
					Id   string `json:"id"`
				} `json:"Multi_user"`
			} `json:"Multi_user"`
			Widget struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"widget"`
		} `json:"data"`
		ProcessInfo struct {
			Escalation struct {
				Days   int    `json:"days"`
				Status string `json:"status"`
			} `json:"escalation"`
		} `json:"process_info"`
		Transitions []struct {
			Type            string `json:"type"`
			NextTransitions []struct {
				Name string `json:"name"`
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"next_transitions"`
		} `json:"transitions"`
	} `json:"blueprint"`
}

type UpdateBlueprintResponse struct {
	Code    string `json:"code"`
	Details struct {
	} `json:"details"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
