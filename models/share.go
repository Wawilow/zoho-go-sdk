package models

import "time"

type ShareRecordList struct {
	Share []struct {
		SharedWith struct {
			Name string `json:"name"`
			Id   string `json:"id"`
			Type string `json:"type"`
			Zuid string `json:"zuid"`
		} `json:"shared_with"`
		ShareRelatedRecords bool `json:"share_related_records"`
		SharedThrough       struct {
			Module struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"module"`
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"shared_through"`
		SharedTime time.Time `json:"shared_time"`
		Permission string    `json:"permission"`
		SharedBy   struct {
			Name string `json:"name"`
			Id   string `json:"id"`
			Zuid string `json:"zuid"`
		} `json:"shared_by"`
		Type string `json:"type"`
	} `json:"share"`
}

type NewShareRecord struct {
	Share []struct {
		SharedWith struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"shared_with"`
		ShareRelatedRecords bool   `json:"share_related_records"`
		Permission          string `json:"permission"`
		Type                string `json:"type"`
	} `json:"share"`
	NotifySharedMembers bool `json:"notify_shared_members"`
	NotifyOnCompletion  bool `json:"notify_on_completion"`
}

type NewShareRecordResponse struct {
	Share []struct {
		Code    string `json:"code"`
		Details struct {
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"share"`
}

type UpdateSharePermissions struct {
	Share []struct {
		ShareRelatedRecords bool   `json:"share_related_records"`
		Permission          string `json:"permission"`
		Type                string `json:"type"`
		SharedWith          struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"shared_with"`
	} `json:"share"`
	Notify bool `json:"notify"`
}

type UpdateSharePermissionsResponse struct {
	Share []struct {
		Code    string `json:"code"`
		Details struct {
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"share"`
}

type DelSharePermissions struct {
	Share struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"share"`
}
