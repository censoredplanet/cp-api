package entities

import "time"

type Hyperquack struct{

	Protocol string
	Domain string
	DomainIsControl bool
	Date time.Time
	StartTime time.Time
	EndTime time.Time
	ServerIp string
	ServerNetblock string
	ServerAsn string
	ServerAsName string
	ServerAsFullName string
	ServerAsClass string
	Server_country string
	ServerOrganization string
	Source string
	SeceivedError string
	SeceivedStatus string
	ReceivedHeaders string
	ReceivedBody string
	ReceivedTlsVersion string
	ReceviedTlsCipherSuite string
	ReceivedTlsCert string
	ReceivedTlsCertCommonName string
	ReceivedTlsCertAlternativeNames string
	ReceivedTlsCertIssuer string
	MatchesTemplate bool
	NoResponseInMeasurementMatchesTemplate bool
	ControlsFailed bool
	StatefulBlock bool

}
