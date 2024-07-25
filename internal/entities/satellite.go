package entities

import "time"

type Satellite struct{

  Domain string
  DomainIsControl bool
  TestUrl string
  Date time.Time
  StartTime time.Time
  EndTime time.Time
  ResolverIp string
  ResolverName string
  ResolverIsTrusted bool
  ResolverNetblock string
  ResolverAsn string
  ResolverAsName string
  ResolverAsFullName string 
  ResolverAsClass string
  ResolverCountry string
  ResolverOrganization string
  ReceivedError string
  ReceivedRcode int
  Source string

}
