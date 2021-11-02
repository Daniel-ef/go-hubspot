/*
 * Domains
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package domains

type Domain struct {
	PortalId                                     int32            `json:"portalId"`
	Id                                           int64            `json:"id"`
	Created                                      int64            `json:"created"`
	Updated                                      int64            `json:"updated"`
	Domain                                       string           `json:"domain"`
	PrimaryLandingPage                           bool             `json:"primaryLandingPage"`
	PrimaryEmail                                 bool             `json:"primaryEmail"`
	PrimaryBlog                                  bool             `json:"primaryBlog"`
	PrimaryBlogPost                              bool             `json:"primaryBlogPost"`
	PrimarySitePage                              bool             `json:"primarySitePage"`
	PrimaryKnowledge                             bool             `json:"primaryKnowledge"`
	PrimaryLegacyPage                            bool             `json:"primaryLegacyPage"`
	PrimaryClickTracking                         bool             `json:"primaryClickTracking"`
	FullCategoryKey                              string           `json:"fullCategoryKey"`
	SecondaryToDomain                            string           `json:"secondaryToDomain"`
	IsResolving                                  bool             `json:"isResolving"`
	IsDnsCorrect                                 bool             `json:"isDnsCorrect"`
	ManuallyMarkedAsResolving                    bool             `json:"manuallyMarkedAsResolving"`
	ConsecutiveNonResolvingCount                 int32            `json:"consecutiveNonResolvingCount"`
	SslCname                                     string           `json:"sslCname"`
	IsSslEnabled                                 bool             `json:"isSslEnabled"`
	IsSslOnly                                    bool             `json:"isSslOnly"`
	CertificateId                                int64            `json:"certificateId"`
	SslRequestId                                 int64            `json:"sslRequestId"`
	IsUsedForBlogPost                            bool             `json:"isUsedForBlogPost"`
	IsUsedForSitePage                            bool             `json:"isUsedForSitePage"`
	IsUsedForLandingPage                         bool             `json:"isUsedForLandingPage"`
	IsUsedForEmail                               bool             `json:"isUsedForEmail"`
	IsUsedForKnowledge                           bool             `json:"isUsedForKnowledge"`
	SetupTaskId                                  int64            `json:"setupTaskId"`
	IsSetupComplete                              bool             `json:"isSetupComplete"`
	SetUpLanguage                                string           `json:"setUpLanguage"`
	TeamIds                                      []int64          `json:"teamIds"`
	ActualCname                                  string           `json:"actualCname"`
	CorrectCname                                 string           `json:"correctCname"`
	ActualIp                                     string           `json:"actualIp"`
	ApexResolutionStatus                         string           `json:"apexResolutionStatus"`
	ApexDomain                                   string           `json:"apexDomain"`
	PublicSuffix                                 string           `json:"publicSuffix"`
	ApexIpAddresses                              []string         `json:"apexIpAddresses"`
	SiteId                                       int64            `json:"siteId"`
	BrandId                                      int64            `json:"brandId"`
	Deletable                                    bool             `json:"deletable"`
	DomainCdnConfig                              *DomainCdnConfig `json:"domainCdnConfig"`
	SetupInfo                                    *DomainSetupInfo `json:"setupInfo"`
	DerivedBrandName                             string           `json:"derivedBrandName"`
	CreatedById                                  int32            `json:"createdById"`
	UpdatedById                                  int32            `json:"updatedById"`
	Label                                        string           `json:"label"`
	IsAnyPrimary                                 bool             `json:"isAnyPrimary"`
	IsLegacyDomain                               bool             `json:"isLegacyDomain"`
	IsInternalDomain                             bool             `json:"isInternalDomain"`
	IsResolvingInternalProperty                  bool             `json:"isResolvingInternalProperty"`
	IsResolvingIgnoringManuallyMarkedAsResolving bool             `json:"isResolvingIgnoringManuallyMarkedAsResolving"`
	IsUsedForAnyContentType                      bool             `json:"isUsedForAnyContentType"`
	IsLegacy                                     bool             `json:"isLegacy"`
	AuthorAt                                     int64            `json:"authorAt"`
	CosObjectType                                string           `json:"cosObjectType"`
	CdnPurgeEmbargoTime                          int64            `json:"cdnPurgeEmbargoTime"`
	IsStagingDomain                              bool             `json:"isStagingDomain"`
}
