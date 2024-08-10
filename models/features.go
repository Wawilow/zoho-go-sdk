package models

type FeaturesList struct {
	Features []struct {
		Components []struct {
			ApiName         string `json:"api_name"`
			ModuleSupported bool   `json:"module_supported"`
			Details         *struct {
				Limits struct {
					EditionLimit int `json:"edition_limit"`
					Total        int `json:"total"`
				} `json:"limits"`
			} `json:"details"`
			FeatureLabel string `json:"feature_label"`
		} `json:"components"`
		ApiName       *string `json:"api_name"`
		ParentFeature *struct {
			ApiName string `json:"api_name"`
		} `json:"parent_feature"`
		ModuleSupported bool `json:"module_supported"`
		Details         *struct {
			Limits struct {
				EditionLimit int `json:"edition_limit"`
				Total        int `json:"total"`
			} `json:"limits"`
			UsedCount struct {
				Total int `json:"total"`
			} `json:"used_count,omitempty"`
		} `json:"details"`
		FeatureLabel *string `json:"feature_label"`
	} `json:"__features"`
	Info struct {
		PerPage     int  `json:"per_page"`
		Count       int  `json:"count"`
		Page        int  `json:"page"`
		MoreRecords bool `json:"more_records"`
	} `json:"info"`
}

type UserLicensesCount struct {
	Features []struct {
		Components      interface{} `json:"components"`
		ApiName         string      `json:"api_name"`
		ParentFeature   interface{} `json:"parent_feature"`
		ModuleSupported bool        `json:"module_supported"`
		Details         struct {
			AvailableCount struct {
				Total int `json:"total"`
			} `json:"available_count"`
			UsedCount struct {
				Total int `json:"total"`
			} `json:"used_count"`
			Limits struct {
				Total int `json:"total"`
			} `json:"limits"`
		} `json:"details"`
		FeatureLabel string `json:"feature_label"`
	} `json:"__features"`
}

type UploadFile struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}
