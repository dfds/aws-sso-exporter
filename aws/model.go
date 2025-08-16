package aws

import "time"

type ConsoleLoginEvent struct {
	EventVersion string `json:"eventVersion"`
	UserIdentity struct {
		Type           string `json:"type"`
		PrincipalID    string `json:"principalId"`
		Arn            string `json:"arn"`
		AccountID      string `json:"accountId"`
		SessionContext struct {
			SessionIssuer struct {
				Type        string `json:"type"`
				PrincipalID string `json:"principalId"`
				Arn         string `json:"arn"`
				AccountID   string `json:"accountId"`
				UserName    string `json:"userName"`
			} `json:"sessionIssuer"`
			WebIDFederationData struct {
			} `json:"webIdFederationData"`
			Attributes struct {
				CreationDate     time.Time `json:"creationDate"`
				MfaAuthenticated string    `json:"mfaAuthenticated"`
			} `json:"attributes"`
		} `json:"sessionContext"`
	} `json:"userIdentity"`
	EventTime         time.Time   `json:"eventTime"`
	EventSource       string      `json:"eventSource"`
	EventName         string      `json:"eventName"`
	AwsRegion         string      `json:"awsRegion"`
	SourceIPAddress   string      `json:"sourceIPAddress"`
	UserAgent         string      `json:"userAgent"`
	RequestParameters interface{} `json:"requestParameters"`
	ResponseElements  struct {
		ConsoleLogin string `json:"ConsoleLogin"`
	} `json:"responseElements"`
	AdditionalEventData struct {
		MobileVersion string `json:"MobileVersion"`
		MFAUsed       string `json:"MFAUsed"`
	} `json:"additionalEventData"`
	EventID            string `json:"eventID"`
	ReadOnly           bool   `json:"readOnly"`
	EventType          string `json:"eventType"`
	ManagementEvent    bool   `json:"managementEvent"`
	RecipientAccountID string `json:"recipientAccountId"`
	EventCategory      string `json:"eventCategory"`
	TLSDetails         struct {
		TLSVersion               string `json:"tlsVersion"`
		CipherSuite              string `json:"cipherSuite"`
		ClientProvidedHostHeader string `json:"clientProvidedHostHeader"`
	} `json:"tlsDetails"`
}

type FederateEvent struct {
	EventVersion string `json:"eventVersion"`
	UserIdentity struct {
		Type        string `json:"type"`
		PrincipalID string `json:"principalId"`
		AccountID   string `json:"accountId"`
		UserName    string `json:"userName"`
		OnBehalfOf  struct {
			UserId string `json:"userId"`
		} `json:"onBehalfOf"`
	} `json:"userIdentity"`
	EventTime           time.Time   `json:"eventTime"`
	EventSource         string      `json:"eventSource"`
	EventName           string      `json:"eventName"`
	AwsRegion           string      `json:"awsRegion"`
	SourceIPAddress     string      `json:"sourceIPAddress"`
	UserAgent           string      `json:"userAgent"`
	RequestParameters   interface{} `json:"requestParameters"`
	ResponseElements    interface{} `json:"responseElements"`
	RequestID           string      `json:"requestID"`
	EventID             string      `json:"eventID"`
	ReadOnly            bool        `json:"readOnly"`
	EventType           string      `json:"eventType"`
	ManagementEvent     bool        `json:"managementEvent"`
	RecipientAccountID  string      `json:"recipientAccountId"`
	ServiceEventDetails struct {
		RoleName  string `json:"role_name"`
		AccountID string `json:"account_id"`
	} `json:"serviceEventDetails"`
	EventCategory string `json:"eventCategory"`
}

type AuthenticateEvent struct {
	EventVersion string `json:"eventVersion"`
	UserIdentity struct {
		Type        string `json:"type"`
		PrincipalID string `json:"principalId"`
		AccountID   string `json:"accountId"`
		UserName    string `json:"userName"`
		OnBehalfOf  struct {
			UserId string `json:"userId"`
		} `json:"onBehalfOf"`
	} `json:"userIdentity"`
	EventTime          time.Time   `json:"eventTime"`
	EventSource        string      `json:"eventSource"`
	EventName          string      `json:"eventName"`
	AwsRegion          string      `json:"awsRegion"`
	SourceIPAddress    string      `json:"sourceIPAddress"`
	UserAgent          string      `json:"userAgent"`
	RequestParameters  interface{} `json:"requestParameters"`
	ResponseElements   interface{} `json:"responseElements"`
	RequestID          string      `json:"requestID"`
	EventID            string      `json:"eventID"`
	ReadOnly           bool        `json:"readOnly"`
	EventType          string      `json:"eventType"`
	ManagementEvent    bool        `json:"managementEvent"`
	RecipientAccountID string      `json:"recipientAccountId"`
	EventCategory      string      `json:"eventCategory"`
}
