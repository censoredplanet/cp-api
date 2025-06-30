package entities

import "time"

type Satellite struct {
	Domain                                  string       `json:"domain" ch:"domain"`
	DomainCategory                          *string      `json:"domainCategory,omitempty" ch:"domain_category"`
	DomainIsControl                         *bool        `json:"domainIsControl,omitempty" ch:"domain_is_control"`
	Date                                    *time.Time   `json:"date,omitempty" ch:"date"`
	StartTime                               *time.Time   `json:"startTime,omitempty" ch:"start_time"`
	EndTime                                 *time.Time   `json:"endTime,omitempty" ch:"end_time"`
	Retry                                   *uint8       `json:"retry,omitempty" ch:"retry"`
	ResolverIp                              *string      `json:"resolverIp,omitempty" ch:"resolver_ip"`
	ResolverName                            *string      `json:"resolverName,omitempty" ch:"resolver_name"`
	ResolverIsTrusted                       *bool        `json:"resolverIsTrusted,omitempty" ch:"resolver_is_trusted"`
	ResolverNetblock                        *string      `json:"resolverNetblock,omitempty" ch:"resolver_netblock"`
	ResolverAsn                             *uint32      `json:"resolverAsn,omitempty" ch:"resolver_asn"`
	ResolverAsName                          *string      `json:"resolverAsName,omitempty" ch:"resolver_as_name"`
	ResolverAsFullName                      *string      `json:"resolverAsFullName,omitempty" ch:"resolver_as_full_name"`
	ResolverAsClass                         *string      `json:"resolverAsClass,omitempty" ch:"resolver_as_class"`
	ResolverCountry                         *string      `json:"resolverCountry,omitempty" ch:"resolver_country"`
	ResolverOrganization                    *string      `json:"resolverOrganization,omitempty" ch:"resolver_organization"`
	ResolverNonZeroRcodeRate                *float32     `json:"resolverNonZeroRcodeRate,omitempty" ch:"resolver_non_zero_rcode_rate"`
	ResolverPrivateIpRate                   *float32     `json:"resolverPrivateIpRate,omitempty" ch:"resolver_private_ip_rate"`
	ResolverZeroIpRate                      *float32     `json:"resolverZeroIpRate,omitempty" ch:"resolver_zero_ip_rate"`
	ResolverConnectErrorRate                *float32     `json:"resolverConnectErrorRate,omitempty" ch:"resolver_connect_error_rate"`
	ResolverInvalidCertRate                 *float32     `json:"resolverInvalidCertRate,omitempty" ch:"resolver_invalid_cert_rate"`
	ReceivedError                           *string      `json:"receivedError,omitempty" ch:"received_error"`
	ReceivedRcode                           *int8        `json:"receivedRcode,omitempty" ch:"received_rcode"`
	AnswersIp                               []*string    `json:"answersIp,omitempty" ch:"answers_ip"`
	AnswersAsn                              []*uint32    `json:"answersAsn,omitempty" ch:"answers_asn"`
	AnswersAsName                           []*string    `json:"answersAsName,omitempty" ch:"answers_as_name"`
	AnswersIpOrganization                   []*string    `json:"answersIpOrganization,omitempty" ch:"answers_ip_organization"`
	AnswersCensysHttpBodyHash               []*string    `json:"answersCensysHttpBodyHash,omitempty" ch:"answers_censys_http_body_hash"`
	AnswersCensysIpCert                     []*string    `json:"answersCensysIpCert,omitempty" ch:"answers_censys_ip_cert"`
	AnswersMatchesControlIp                 []*bool      `json:"answersMatchesControlIp,omitempty" ch:"answers_matches_control_ip"`
	AnswersMatchesControlCensysHttpBodyHash []*bool      `json:"answersMatchesControlCensysHttpBodyHash,omitempty" ch:"answers_matches_control_censys_http_body_hash"`
	AnswersMatchesControlCensysIpCert       []*bool      `json:"answersMatchesControlCensysIpCert,omitempty" ch:"answers_matches_control_censys_ip_cert"`
	AnswersMatchesControlAsn                []*bool      `json:"answersMatchesControlAsn,omitempty" ch:"answers_matches_control_asn"`
	AnswersMatchesControlAsName             []*bool      `json:"answersMatchesControlAsName,omitempty" ch:"answers_matches_control_as_name"`
	AnswersMatchConfidence                  []*float32   `json:"answersMatchConfidence,omitempty" ch:"answers_match_confidence"`
	AnswersHttpError                        []*string    `json:"answersHttpError,omitempty" ch:"answers_http_error"`
	AnswersHttpResponseStatus               []*string    `json:"answersHttpResponseStatus,omitempty" ch:"answers_http_response_status"`
	AnswersHttpAnalysisIsKnownBlockpage     []*bool      `json:"answersHttpAnalysisIsKnownBlockpage,omitempty" ch:"answers_http_analysis_is_known_blockpage"`
	AnswersHttpAnalysisPageSignature        []*string    `json:"answersHttpAnalysisPageSignature,omitempty" ch:"answers_http_analysis_page_signature"`
	AnswersHttpsError                       []*string    `json:"answersHttpsError,omitempty" ch:"answers_https_error"`
	AnswersHttpsTlsVersion                  []*uint16    `json:"answersHttpsTlsVersion,omitempty" ch:"answers_https_tls_version"`
	AnswersHttpsTlsCipherSuite              []*uint16    `json:"answersHttpsTlsCipherSuite,omitempty" ch:"answers_https_tls_cipher_suite"`
	AnswersHttpsTlsCert                     []*string    `json:"answersHttpsTlsCert,omitempty" ch:"answers_https_tls_cert"`
	AnswersHttpsTlsCertCommonName           []*string    `json:"answersHttpsTlsCertCommonName,omitempty" ch:"answers_https_tls_cert_common_name"`
	AnswersHttpsTlsCertIssuer               []*string    `json:"answersHttpsTlsCertIssuer,omitempty" ch:"answers_https_tls_cert_issuer"`
	AnswersHttpsTlsCertStartDate            []*time.Time `json:"answersHttpsTlsCertStartDate,omitempty" ch:"answers_https_tls_cert_start_date"`
	AnswersHttpsTlsCertEndDate              []*time.Time `json:"answersHttpsTlsCertEndDate,omitempty" ch:"answers_https_tls_cert_end_date"`
	AnswersHttpsTlsCertAlternativeNames     [][]string   `json:"answersHttpsTlsCertAlternativeNames,omitempty" ch:"answers_https_tls_cert_alternative_names"`
	AnswersHttpsTlsCertHasTrustedCa         []*bool      `json:"answersHttpsTlsCertHasTrustedCa,omitempty" ch:"answers_https_tls_cert_has_trusted_ca"`
	AnswersHttpsTlsCertMatchesDomain        []*bool      `json:"answersHttpsTlsCertMatchesDomain,omitempty" ch:"answers_https_tls_cert_matches_domain"`
	AnswersHttpsResponseStatus              []*string    `json:"answersHttpsResponseStatus,omitempty" ch:"answers_https_response_status"`
	AnswersHttpsAnalysisIsKnownBlockpage    []*bool      `json:"answersHttpsAnalysisIsKnownBlockpage,omitempty" ch:"answers_https_analysis_is_known_blockpage"`
	AnswersHttpsAnalysisPageSignature       []*string    `json:"answersHttpsAnalysisPageSignature,omitempty" ch:"answers_https_analysis_page_signature"`
	Success                                 *bool        `json:"success,omitempty" ch:"success"`
	Anomaly                                 *bool        `json:"anomaly,omitempty" ch:"anomaly"`
	DomainControlsFailed                    *bool        `json:"domainControlsFailed,omitempty" ch:"domain_controls_failed"`
	AverageConfidence                       *float32     `json:"averageConfidence,omitempty" ch:"average_confidence"`
	UntaggedControls                        *bool        `json:"untaggedControls,omitempty" ch:"untagged_controls"`
	UntaggedResponse                        *bool        `json:"untaggedResponse,omitempty" ch:"untagged_response"`
	Excluded                                *bool        `json:"excluded,omitempty" ch:"excluded"`
	ExcludeReason                           *string      `json:"excludeReason,omitempty" ch:"exclude_reason"`
	HasTypeA                                *bool        `json:"hasTypeA,omitempty" ch:"has_type_a"`
	MeasurementId                           *string      `json:"measurementId,omitempty" ch:"measurement_id"`
	Source                                  *string      `json:"source,omitempty" ch:"source"`
}

type SatelliteFilter struct {
	Domain    string    `json:"domain"`
	Country   string    `json:"country"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type SatelliteFilterCH struct {
	Domain    string `json:"domain"`
	Country   string `json:"country"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
