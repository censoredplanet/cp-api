package entities

import "time"

type Satellite struct{

  domain string
  domainIsControl bool
  testUrl string
  date time.Time
  startTime time.Time
  endTime time.Time
  resolverIp string
  resolverName string
  resolverIsTrusted bool
  resolverNetblock string
  resolverAsn string
  resolverAsName string
  resolverAsFullName string 
  resolverAsClass string
  resolverCountry string
  resolverOrganization string
  receivedError string
  receivedRcode int
  source string

}
