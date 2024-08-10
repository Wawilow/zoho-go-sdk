package models

import "time"

type Modules struct {
	HasMoreProfiles     bool        `json:"has_more_profiles"`
	AccessType          string      `json:"access_type"`
	PrivateProfile      interface{} `json:"private_profile"`
	Deletable           bool        `json:"deletable"`
	Description         *string     `json:"description"`
	Creatable           bool        `json:"creatable"`
	RecycleBinOnDelete  bool        `json:"recycle_bin_on_delete"`
	ModifiedTime        *time.Time  `json:"modified_time"`
	PluralLabel         string      `json:"plural_label"`
	PresenceSubMenu     bool        `json:"presence_sub_menu"`
	ActualPluralLabel   string      `json:"actual_plural_label"`
	Id                  string      `json:"id"`
	Visibility          int         `json:"visibility"`
	Convertable         bool        `json:"convertable"`
	SubMenuAvailable    bool        `json:"sub_menu_available"`
	Editable            bool        `json:"editable"`
	ActualSingularLabel string      `json:"actual_singular_label"`
	Profiles            []struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"profiles"`
	ShowAsTab     bool        `json:"show_as_tab"`
	WebLink       interface{} `json:"web_link"`
	SingularLabel string      `json:"singular_label"`
	Viewable      bool        `json:"viewable"`
	ApiSupported  bool        `json:"api_supported"`
	ApiName       string      `json:"api_name"`
	QuickCreate   bool        `json:"quick_create"`
	ModifiedBy    *struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"modified_by"`
	GeneratedType          string        `json:"generated_type"`
	FeedsRequired          bool          `json:"feeds_required"`
	Arguments              []interface{} `json:"arguments"`
	ModuleName             string        `json:"module_name"`
	ProfileCount           int           `json:"profile_count"`
	BusinessCardFieldLimit int           `json:"business_card_field_limit"`
	ParentModule           struct {
		ApiName string `json:"api_name,omitempty"`
		Id      string `json:"id,omitempty"`
	} `json:"parent_module"`
	Status           string `json:"status"`
	SequenceNumber   int    `json:"sequence_number,omitempty"`
	TrackCurrentData bool   `json:"track_current_data,omitempty"`
}

type ApiList struct {
	Apis []struct {
		Path           string `json:"path"`
		OperationTypes []struct {
			Method     string `json:"method"`
			OauthScope string `json:"oauth_scope"`
			MaxCredits int    `json:"max_credits"`
			MinCredits int    `json:"min_credits"`
		} `json:"operation_types"`
	} `json:"__apis"`
}
