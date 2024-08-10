package models

import "time"

type Organization struct {
	Country              interface{} `json:"country"`
	HierarchyPreferences struct {
		Type string `json:"type"`
	} `json:"hierarchy_preferences"`
	PhotoId             interface{} `json:"photo_id"`
	City                interface{} `json:"city"`
	Description         interface{} `json:"description"`
	Type                string      `json:"type"`
	McStatus            bool        `json:"mc_status"`
	GappsEnabled        bool        `json:"gapps_enabled"`
	LiteUsersEnabled    bool        `json:"lite_users_enabled"`
	DomainName          string      `json:"domain_name"`
	TranslationEnabled  bool        `json:"translation_enabled"`
	Street              interface{} `json:"street"`
	Alias               interface{} `json:"alias"`
	Currency            string      `json:"currency"`
	DeletableOrgAccount bool        `json:"deletable_org_account"`
	Id                  string      `json:"id"`
	State               interface{} `json:"state"`
	Fax                 interface{} `json:"fax"`
	EmployeeCount       string      `json:"employee_count"`
	Zip                 interface{} `json:"zip"`
	CreatedTime         time.Time   `json:"created_time"`
	Website             interface{} `json:"website"`
	CurrencySymbol      string      `json:"currency_symbol"`
	Mobile              interface{} `json:"mobile"`
	CurrencyLocale      string      `json:"currency_locale"`
	PrimaryZuid         string      `json:"primary_zuid"`
	ZiaPortalId         string      `json:"zia_portal_id"`
	TimeZone            string      `json:"time_zone"`
	Zgid                string      `json:"zgid"`
	CountryCode         string      `json:"country_code"`
	LicenseDetails      struct {
		PaidExpiry                  time.Time   `json:"paid_expiry"`
		UsersLicensePurchased       int         `json:"users_license_purchased"`
		TrialType                   interface{} `json:"trial_type"`
		TrialExpiry                 interface{} `json:"trial_expiry"`
		Paid                        bool        `json:"paid"`
		PaidType                    string      `json:"paid_type"`
		PortalUsersLicensePurchased int         `json:"portal_users_license_purchased"`
	} `json:"license_details"`
	Phone                  string `json:"phone"`
	CompanyName            string `json:"company_name"`
	PrivacySettings        bool   `json:"privacy_settings"`
	PrimaryEmail           string `json:"primary_email"`
	HipaaComplianceEnabled bool   `json:"hipaa_compliance_enabled"`
	IsoCode                string `json:"iso_code"`
}

type NewOrganizationPhoto struct {
	Message string `json:"message"`
	Details struct {
	} `json:"details"`
	Status string `json:"status"`
	Code   string `json:"code"`
}

type DelOrganizationPhoto struct {
	Code    string `json:"code"`
	Details struct {
	} `json:"details"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
