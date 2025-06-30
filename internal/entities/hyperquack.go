package entities

import "time"

type Hyperquack struct {
	Domain                                 string     `json:"domain" ch:"domain"`
	DomainCategory                         *string    `json:"domainCategory,omitempty" ch:"domain_category"`
	DomainIsControl                        *bool      `json:"domainIsControl,omitempty" ch:"domain_is_control"`
	Date                                   *time.Time `json:"date,omitempty" ch:"date"`
	StartTime                              *time.Time `json:"startTime,omitempty" ch:"start_time"`
	EndTime                                *time.Time `json:"endTime,omitempty" ch:"end_time"`
	Retry                                  *uint8     `json:"retry,omitempty" ch:"retry"`
	ServerIp                               *string    `json:"serverIp,omitempty" ch:"server_ip"`
	ServerNetblock                         *string    `json:"serverNetblock,omitempty" ch:"server_netblock"`
	ServerAsn                              *uint32    `json:"serverAsn,omitempty" ch:"server_asn"`
	ServerAsName                           *string    `json:"serverAsName,omitempty" ch:"server_as_name"`
	ServerAsFullName                       *string    `json:"serverAsFullName,omitempty" ch:"server_as_full_name"`
	ServerAsClass                          *string    `json:"serverAsClass,omitempty" ch:"server_as_class"`
	ServerCountry                          *string    `json:"serverCountry,omitempty" ch:"server_country"`
	ServerOrganization                     *string    `json:"serverOrganization,omitempty" ch:"server_organization"`
	ReceivedError                          *string    `json:"receivedError,omitempty" ch:"received_error"`
	ReceivedTlsVersion                     *uint16    `json:"receivedTlsVersion,omitempty" ch:"received_tls_version"`
	ReceivedTlsCipherSuite                 *uint16    `json:"receivedTlsCipherSuite,omitempty" ch:"received_tls_cipher_suite"`
	ReceivedTlsCert                        *string    `json:"receivedTlsCert,omitempty" ch:"received_tls_cert"`
	ReceivedTlsCertMatchesDomain           *bool      `json:"receivedTlsCertMatchesDomain,omitempty" ch:"received_tls_cert_matches_domain"`
	ReceivedTlsCertCommonName              *string    `json:"receivedTlsCertCommonName,omitempty" ch:"received_tls_cert_common_name"`
	ReceivedTlsCertIssuer                  *string    `json:"receivedTlsCertIssuer,omitempty" ch:"received_tls_cert_issuer"`
	ReceivedTlsCertAlternativeNames        []string   `json:"receivedTlsCertAlternativeNames,omitempty" ch:"received_tls_cert_alternative_names"`
	ReceivedStatus                         *string    `json:"receivedStatus,omitempty" ch:"received_status"`
	ReceivedHeaders                        []string   `json:"receivedHeaders,omitempty" ch:"received_headers"`
	ReceivedBody                           *string    `json:"receivedBody,omitempty" ch:"received_body"`
	IsKnownBlockpage                       *bool      `json:"isKnownBlockpage,omitempty" ch:"is_known_blockpage"`
	PageSignature                          *string    `json:"pageSignature,omitempty" ch:"page_signature"`
	Outcome                                *string    `json:"outcome,omitempty" ch:"outcome"`
	MatchesTemplate                        *bool      `json:"matchesTemplate,omitempty" ch:"matches_template"`
	NoResponseInMeasurementMatchesTemplate *bool      `json:"noResponseInMeasurementMatchesTemplate,omitempty" ch:"no_response_in_measurement_matches_template"`
	ControlsFailed                         *bool      `json:"controlsFailed,omitempty" ch:"controls_failed"`
	StatefulBlock                          *bool      `json:"statefulBlock,omitempty" ch:"stateful_block"`
	MeasurementId                          *string    `json:"measurementId,omitempty" ch:"measurement_id"`
	Source                                 *string    `json:"source,omitempty" ch:"source"`
}

type HyperquackFilter struct {
	Protocol  string    `json:"protocol"`
	Domain    string    `json:"domain"`
	Country   string    `json:"country"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type HyperquackFilterCH struct {
	Protocol  string `json:"protocol"`
	Domain    string `json:"domain"`
	Country   string `json:"country"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
