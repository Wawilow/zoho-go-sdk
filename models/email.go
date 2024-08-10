package models

import "time"

type AddEmail2Record struct {
	Emails []struct {
		From struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"from"`
		To []struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"to"`
		Cc []struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"cc"`
		Bcc []struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"bcc"`
		Subject     string    `json:"subject"`
		Content     string    `json:"content"`
		MailFormat  string    `json:"mail_format"`
		DateTime    time.Time `json:"date_time"`
		Sent        bool      `json:"sent"`
		Attachments []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"attachments"`
		LinkedRecord struct {
			Id     string `json:"id"`
			Name   string `json:"name"`
			Module struct {
				ApiName string `json:"api_name"`
				Id      string `json:"id"`
			} `json:"module"`
		} `json:"linked_record"`
		OriginalMessageId string `json:"original_message_id"`
	} `json:"Emails"`
}

type AddEmail2RecordResponse struct {
	Emails []struct {
		From struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"from"`
		To []struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"to"`
		Cc []struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"cc"`
		Bcc []struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"bcc"`
		Subject     string    `json:"subject"`
		Content     string    `json:"content"`
		MailFormat  string    `json:"mail_format"`
		DateTime    time.Time `json:"date_time"`
		Sent        bool      `json:"sent"`
		Attachments []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"attachments"`
		LinkedRecord struct {
			Id     string `json:"id"`
			Name   string `json:"name"`
			Module struct {
				ApiName string `json:"api_name"`
				Id      string `json:"id"`
			} `json:"module"`
		} `json:"linked_record"`
		OriginalMessageId string `json:"original_message_id"`
	} `json:"Emails"`
}

type EmailSharedDetails struct {
	EmailsSharingDetails struct {
		AvailableTypes []string `json:"available_types"`
		ShareFromUsers []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"_type"`
		} `json:"share_from_users"`
	} `json:"__emails_sharing_details"`
}

type RecordEmailsList struct {
	Emails []struct {
		Cc                  interface{} `json:"cc"`
		HasThreadAttachment bool        `json:"has_thread_attachment"`
		Summary             interface{} `json:"summary"`
		Owner               struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"owner"`
		Read            bool        `json:"read"`
		Subject         string      `json:"subject"`
		MessageId       string      `json:"message_id"`
		HasAttachment   bool        `json:"has_attachment"`
		Source          string      `json:"source"`
		Sent            bool        `json:"sent"`
		Intent          interface{} `json:"intent"`
		SentimentInfo   interface{} `json:"sentiment_info"`
		LinkedRecord    interface{} `json:"linked_record"`
		RelatedToRecord interface{} `json:"related_to_record"`
		Emotion         interface{} `json:"emotion"`
		From            struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"from"`
		To []struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"to"`
		Time   time.Time `json:"time"`
		Status []struct {
			Type          string    `json:"type"`
			BouncedTime   time.Time `json:"bounced_time,omitempty"`
			SubCategory   string    `json:"sub_category,omitempty"`
			BouncedReason string    `json:"bounced_reason,omitempty"`
			Category      string    `json:"category,omitempty"`
		} `json:"status"`
	} `json:"Emails"`
	Info struct {
		PerPage     int    `json:"per_page"`
		NextIndex   string `json:"next_index"`
		Count       int    `json:"count"`
		PrevIndex   string `json:"prev_index"`
		MoreRecords bool   `json:"more_records"`
	} `json:"info"`
}

type AllowedFromEmail struct {
	FromAddresses []struct {
		UserName string  `json:"user_name"`
		Type     string  `json:"type"`
		Email    string  `json:"email"`
		Id       *string `json:"id,omitempty"`
	} `json:"from_addresses"`
}

type SendEmail struct {
	Data []struct {
		From struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"from"`
		To []struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"to"`
		Cc []struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"cc"`
		Bcc []struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"bcc"`
		ReplyTo struct {
			UserName string `json:"user_name"`
			Email    string `json:"email"`
		} `json:"reply_to"`
		OrgEmail  bool `json:"org_email"`
		InReplyTo struct {
			MessageId string `json:"message_id"`
			Owner     struct {
				Id string `json:"id"`
			} `json:"owner"`
		} `json:"in_reply_to"`
		ScheduledTime time.Time `json:"scheduled_time"`
		Subject       string    `json:"subject"`
		Content       string    `json:"content"`
		MailFormat    string    `json:"mail_format"`
		Attachments   []struct {
			Id string `json:"id"`
		} `json:"attachments"`
		Template struct {
			Id string `json:"id"`
		} `json:"template"`
	} `json:"data"`
}

type SendEmailResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			MessageId string `json:"message_id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

type UnblockEmail struct {
	UnblockFields []string `json:"unblock_fields"`
}
