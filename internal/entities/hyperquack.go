package entities

import "time"

type Hyperquack struct{

	protocol string
	domain string
	domainIsControl bool
	date time.Time
	startTime time.Time
	endTime time.Time
	serverIp string
	serverNetblock string
	serverAsn string
	serverAsName string
	serverAsFullName string
	serverAsClass string
	server_country string
	serverOrganization string
	source string
	receivedError string
	receivedStatus string
	receivedHeaders string
	receivedBody string
	receivedTlsVersion string
	ReceviedTlsCipherSuite string
	ReceivedTlsCert string
	ReceivedTlsCertCommonName string
	ReceivedTlsCertAlternativeNames string
	ReceivedTlsCertIssuer string
	matchesTemplate bool
	noResponseInMeasurementMatchesTemplate bool
	controlsFailed bool
	statefulBlock bool

}
