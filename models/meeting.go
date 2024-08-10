package models

type MeetingCancel struct {
	Data []struct {
		SendCancellingMail bool `json:"send_cancelling_mail"`
	} `json:"data"`
}

type MeetingCancelResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			Id string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}
