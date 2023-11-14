// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package fastly

import (
	"Fastly/internal/utils"
	"Fastly/models/components"
	"context"
	"fmt"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://api.fastly.com",
	"https://rt.fastly.com",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// SDK - Fastly API: Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports.
// The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts.
// For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/)
type SDK struct {
	// Supports redirecting traffic for apex domains, subdomains, or wildcard domains to a WWW subdomain.
	//
	// https://developer.fastly.com/reference/api/vcl-services/apex-redirect
	ApexRedirect *ApexRedirect
	// Get information on current and past bills.
	//
	// https://developer.fastly.com/reference/api/account/billing
	Billing *Billing
	// Fastly makes it possible to see which version of a particular URL is cached on each edge server.
	//
	// https://developer.fastly.com/reference/api/utils/content
	Content *Content
	// A Customer is the base object that owns your Users and Services. Some information may be limited depending on access level.
	//
	// https://developer.fastly.com/reference/api/account/customer
	Customer *Customer
	// A user of the Fastly API and web interface. A user is always associated with a customer. Some information may be limited depending on access level.
	//
	// https://developer.fastly.com/reference/api/account/user
	User *User
	// A billing address is used to calculate your bill correctly.
	//
	// https://developer.fastly.com/reference/api/account/billing-address
	BillingAddress *BillingAddress
	// A Customer Contact is the base object that holds the different types of contact information Fastly uses to contact a customer.
	//
	// https://developer.fastly.com/reference/api/account/contact
	Contact *Contact
	// An API Token is used to identify who is making the API call. Users can create multiple tokens to suit their needs.
	//
	// https://developer.fastly.com/reference/api/auth-tokens/user
	Tokens *Tokens
	// List Fastly POPs and their locations.
	//
	// https://developer.fastly.com/reference/api/utils/pops
	Pop *Pop
	// A domain ownership is the record that Fastly keeps after a customer has validated control over a domain.
	//
	// https://developer.fastly.com/reference/api/services/domain-ownerships
	DomainOwnerships *DomainOwnerships
	// These endpoints allow you to enable, disable, and check the enablement status of products on your services.
	//
	// https://developer.fastly.com/reference/api/products/enablement
	EnabledProducts *EnabledProducts
	// [Event logs](https://docs.fastly.com/en/guides/reviewing-service-activity-with-the-event-log) are used to audit actions performed by customers.
	//
	//
	// https://developer.fastly.com/reference/api/account/events
	Events *Events
	// Invitations allow superusers and engineers to invite users to set up accounts as collaborators under a main customer account. Superusers can invite collaborators and assign them any role or permission level on a per-service basis. Engineers with no per-service limitations on their role can only invite new collaborators but cannot modify their permissions.
	//
	// https://developer.fastly.com/reference/api/account/invitations
	Invitations *Invitations
	// A list of available permissions that can be assigned to a custom role.
	//
	// https://developer.fastly.com/reference/api/account/permissions
	IamPermissions *IamPermissions
	// To help you manage firewall rules for connections from Fastly to your origin, we provide access to the [list of Fastly's assigned IP ranges](https://api.fastly.com/public-ip-list). Changes to this list will be announced in advance as an "IP address announcement" along with other service announcements to our [status page](https://fastlystatus.com/), which you can [subscribe](https://docs.fastly.com/en/guides/fastlys-network-status#subscribing-to-notifications) to. This list is exhaustive and includes all Fastly-owned IP ranges, so any client connections, log streaming reports, and origin connections should use these addresses.
	//
	// https://developer.fastly.com/reference/api/utils/public-ip-list
	PublicIPList *PublicIPList
	// Instant Purging removes content from Fastly immediately so it can be refreshed from your origin servers. While the default approach for issuing an individual URL Instant Purge uses the Fastly API, `https://api.fastly.com/`, it is not required.
	//
	// https://developer.fastly.com/reference/api/purging
	Purge *Purge
	// Rate limiters add configurable origin request rate limiting to a service. This information is part of a limited availability release. For more information, see our [product and feature lifecycle](https://docs.fastly.com/products/fastly-product-lifecycle#limited-availability) descriptions. To use this feature you must purchase a Professional or Premier Platform subscription for either [Signal Sciences Cloud WAF](https://docs.fastly.com/products/signal-sciences-cloud-waf) or [Signal Sciences Next-Gen WAF](https://docs.fastly.com/products/signal-sciences-next-gen-waf) and have a [paid account with a contract](https://docs.fastly.com/en/guides/accounts-and-pricing-plans#paid-accounts-with-contractual-commitments) for [full-site delivery](https://docs.fastly.com/products/fastlys-legacy-full-site-delivery-services).
	//
	// https://developer.fastly.com/reference/api/vcl-services/rate-limiter
	RateLimiter *RateLimiter
	// A container that lets you store data in key-value pairs.
	//
	// https://developer.fastly.com/reference/api/services/resources/config-store
	ConfigStore *ConfigStore
	// A key-value pair within a config store.
	//
	// https://developer.fastly.com/reference/api/services/resources/config-store-item
	ConfigStoreItem *ConfigStoreItem
	// An kv store is a persistent, globally consistent key-value store accessible to Compute@Edge services during request processing.
	//
	// https://developer.fastly.com/reference/api/services/resources/kv-store
	KvStore *KvStore
	// An item in an kv store.
	//
	// https://developer.fastly.com/reference/api/services/resources/kv-store-item
	KvStoreItem *KvStoreItem
	// A role is a collection of permissions that can be assigned to a user group.
	//
	// https://developer.fastly.com/reference/api/account/roles
	IamRoles *IamRoles
	// A Service represents the configuration for a website, app, API, or anything else to be served through Fastly. A Service can have many Versions, through which Backends, Domains, and more can be configured.
	//
	// https://developer.fastly.com/reference/api/services/service
	Service *Service
	// A service authorization allows limited users to access only specified services.
	//
	// https://developer.fastly.com/reference/api/account/service-authorization
	ServiceAuthorizations *ServiceAuthorizations
	// A service group is a collection of services that can be assigned to a user group.
	//
	// https://developer.fastly.com/reference/api/account/service-groups
	IamServiceGroups *IamServiceGroups
	// An ACL entry holds an individual IP address or subnet range and is a member of an ACL. ACL entries are versionless, which means they can be created, modified, or deleted without activating a new version of your service.
	//
	//
	// https://developer.fastly.com/reference/api/acls/acl-entry
	ACLEntry *ACLEntry
	// A Dictionary Item is a single key-value pair that makes up an entry in a Dictionary. Dictionary Items can be added, removed and modified without activating a new version of the associated service.
	//
	// https://developer.fastly.com/reference/api/dictionaries/dictionary-item
	DictionaryItem *DictionaryItem
	// See the line-by-line changes in configuration between two different versions of a service.
	//
	// https://developer.fastly.com/reference/api/utils/diff
	Diff *Diff
	// A server is an address (IP address or hostname) to which the Fastly Load Balancer service can forward requests. This service can define multiple servers and assign it to a pool. Fastly can then select any one of these servers based on a selection policy defined for the pool.
	//
	// https://developer.fastly.com/reference/api/load-balancing/pools/server
	Server *Server
	// Publishing sends messages to [Fanout](https://developer.fastly.com/learning/concepts/real-time-messaging/fanout) subscribers. Fanout is designed to be [GRIP-compatible](https://pushpin.org/docs/protocols/grip/), such that `https://api.fastly.com/service/{service_id}` can be used as a GRIP URL in application configurations.
	//
	// https://developer.fastly.com/reference/api/services/service
	Publish *Publish
	// VCL Snippets are blocks of VCL logic inserted into your service's configuration that don't require custom VCL.
	//
	// https://developer.fastly.com/reference/api/vcl-services/snippet
	Snippet *Snippet
	// Stats give you information on the usage and performance of your Service. They can be requested by Service and over a particular time span. Stats are broken down per POP, giving you information on how your Services are being used across the world. There is now a more flexible, and fully featured [Stats API](/reference/api/metrics-stats/historical-stats/) available.
	//
	// https://developer.fastly.com/reference/api/metrics-stats/stats
	Stats *Stats
	// Compare the changes in generated VCL between two versions of a service. This is sometimes called a "diff" because the comparison may highlight "differences" between the versions. To compare the configuration changes between two versions of a service represented in YAML format instead, use the related [diff](/reference/api/utils/diff/#diff-service-versions) endpoint.
	//
	// https://developer.fastly.com/reference/api/vcl-services/diff
	VclDiff *VclDiff
	// A Version represents a specific instance of the configuration for a service. A Version can be cloned, locked, activated, or deactivated.
	//
	// https://developer.fastly.com/reference/api/services/version
	Version *Version
	// An access control list or "ACL" specifies individual IP addresses or subnet ranges and can be accessed and used from Fastly VCL.
	//
	// https://developer.fastly.com/reference/api/acls/acl
	ACL *ACL
	// A backend (also sometimes called an origin server) is a server identified by IP address or hostname, from which Fastly will fetch your content. There can be multiple backends attached to a service, but each backend is specific to one service. By default, the first backend added to a service configuration will be used for all requests (provided it meets any [conditions](/reference/api/vcl-services/condition) attached to it). If multiple backends are defined for a service, the first one that has no attached conditions, or whose condition is satisfied for the current request, will be used, unless that behavior is modified using the `auto_loadbalance` field described below.
	//
	// https://developer.fastly.com/reference/api/services/backend
	Backend *Backend
	// A VCL is a Varnish configuration file used to customize the configuration for a Service.
	//
	// https://developer.fastly.com/reference/api/vcl-services/vcl
	Vcl *Vcl
	// Configures cache lifetime for objects stored in the Fastly cache, overriding cache freshness information that would otherwise be determined from cache-related headers on the HTTP response. When used in conjunction with conditions, cache settings objects provide detailed control over how long content persists in the cache.
	//
	// https://developer.fastly.com/reference/api/vcl-services/cache-settings
	CacheSettings *CacheSettings
	// Conditions are used to control whether logic defined in configured VCL objects is applied for a particular client request. A condition contains a VCL conditional expression that evaluates to either true or false and is used to determine whether the condition is met. The type of the condition determines where it is executed and the VCL variables that can be evaluated as part of the conditional logic.
	//
	// https://developer.fastly.com/reference/api/vcl-services/condition
	Condition *Condition
	// A Dictionary is a VCL data table that stores key-value pairs accessible to VCL during request processing. New, empty dictionaries can be attached to a draft version of a service, which must be activated for the dictionary to be included in VCL. Once installed, a dictionary's items may be updated via API calls without having to activate a new version of the associated service configuration. To remove a dictionary, delete it on a draft version of a service (one that is not locked and not active). Once removed, activate the draft service version without the dictionary.
	//
	// https://developer.fastly.com/reference/api/dictionaries/dictionary
	Dictionary *Dictionary
	// Dictionary Info is a set of metadata describing properties of a dictionary which change as items are added and removed.
	//
	// https://developer.fastly.com/reference/api/dictionaries/dictionary-info
	DictionaryInfo *DictionaryInfo
	// A Director is responsible for balancing requests among a group of Backends. In addition to simply balancing, Directors can be configured to attempt retrying failed requests. Additionally, Directors have a quorum setting which can be used to determine when the Director as a whole is considered "up", in order to prevent "server whack-a-mole" following an outage as servers come back up. Only directors created via the API can be modified via the API. Directors known as "autodirectors" that are created automatically when load balancing groups of servers together cannot be modified or retrieved via the API.
	//
	// https://developer.fastly.com/reference/api/load-balancing/directors/director
	Director *Director
	// Maps and relates backends as belonging to directors. Backends can belong to any number of directors but directors can only hold one reference to a specific backend.
	//
	// https://developer.fastly.com/reference/api/load-balancing/directors/backend
	DirectorBackend *DirectorBackend
	// A domain represents the domain name through which visitors will retrieve content. There can be multiple domains for a service.
	//
	// https://developer.fastly.com/reference/api/services/domain
	Domain *Domain
	// Gzip configuration allows you to choose resources to automatically compress.  For more information about compressing and decompressing data with Fastly, check out our [concept guide to compression](https://developer.fastly.com/learning/concepts/compression/).
	//
	// https://developer.fastly.com/reference/api/vcl-services/gzip
	Gzip *Gzip
	// Header objects are used to add, modify, or delete headers from requests and responses. The header content can be simple strings or be derived from variables inside Varnish. Regular expressions can be used to customize the headers even further.
	//
	// https://developer.fastly.com/reference/api/vcl-services/header
	Header *Header
	// Health checks are used to customize the way Fastly checks on your Backends. If an origin server is marked unhealthy due to health checks, Fastly will stop attempting to send requests to it. If all origin servers are marked unhealthy, Fastly will attempt to serve stale. If no stale object is available, a 503 will be returned to the client.
	//
	// https://developer.fastly.com/reference/api/services/healthcheck
	Healthcheck *Healthcheck
	// Supports the use of the HTTP/3 (QUIC) protocol.
	//
	// https://developer.fastly.com/reference/api/vcl-services/http3
	Http3 *Http3
	// Fastly will upload log messages to the Azure Blob Storage container in the format specified in the Azure Blob object.
	//
	// https://developer.fastly.com/reference/api/logging/azureblob
	LoggingAzureblob *LoggingAzureblob
	// Fastly will upload log messages to the Google BigQuery dataset and table in the format specified in the BigQuery logging object.
	//
	// https://developer.fastly.com/reference/api/logging/bigquery
	LoggingBigquery *LoggingBigquery
	// Fastly will upload log messages to your Rackspace Cloud Files account.
	//
	// https://developer.fastly.com/reference/api/logging/cloudfiles
	LoggingCloudfiles *LoggingCloudfiles
	// Fastly will upload log messages to Datadog in the format specified in the Datadog configuration object.
	//
	// https://developer.fastly.com/reference/api/logging/datadog
	LoggingDatadog *LoggingDatadog
	// Fastly will upload log messages to the DigitalOcean Space in the format specified in the DigitalOcean Spaces object.
	//
	// https://developer.fastly.com/reference/api/logging/digitalocean
	LoggingDigitalocean *LoggingDigitalocean
	// Fastly will upload log messages periodically to the server in the format specified in the Elasticsearch object.
	//
	// https://developer.fastly.com/reference/api/logging/elasticsearch
	LoggingElasticsearch *LoggingElasticsearch
	// Fastly will upload log messages periodically to the server in the format specified in the FTP object.
	//
	// https://developer.fastly.com/reference/api/logging/ftp
	LoggingFtp *LoggingFtp
	// Fastly will upload log messages to the GCS bucket in the format specified in the GCS object.
	//
	// https://developer.fastly.com/reference/api/logging/gcs
	LoggingGcs *LoggingGcs
	// Fastly will stream log messages to the Heroku account in the format specified in the Heroku object.
	//
	// https://developer.fastly.com/reference/api/logging/heroku
	LoggingHeroku *LoggingHeroku
	// Fastly will upload log messages to Honeycomb.io in the format specified in the Honeycomb object.
	//
	// https://developer.fastly.com/reference/api/logging/honeycomb
	LoggingHoneycomb *LoggingHoneycomb
	// Fastly will upload log messages to an HTTPS endpoint in the format specified in the HTTPS object. The HTTPS endpoint requires proof of domain ownership before logs can be received. Learn how to validate your domain in our [HTTPS endpoint documentation](https://docs.fastly.com/en/guides/log-streaming-https).
	//
	// https://developer.fastly.com/reference/api/logging/https
	LoggingHTTPS *LoggingHTTPS
	// Fastly will upload log messages periodically to the server in the format specified in the Kafka object.
	//
	// https://developer.fastly.com/reference/api/logging/kafka
	LoggingKafka *LoggingKafka
	// Fastly will publish log messages to an Amazon Kinesis stream in the format specified in the Amazon Kinesis Data Streams logging object.
	//
	// https://developer.fastly.com/reference/api/logging/kinesis
	LoggingKinesis *LoggingKinesis
	// The Logentries integration has been discontinued.  No new Logentries endpoints can be created.
	//
	// https://developer.fastly.com/reference/api/logging/logentries
	LoggingLogentries *LoggingLogentries
	// Fastly will stream log messages to the Loggly account in the format specified in the Loggly logging object.
	//
	// https://developer.fastly.com/reference/api/logging/loggly
	LoggingLoggly *LoggingLoggly
	// Fastly will upload log messages to the Log Shuttle bucket in the format specified in the logshuttle object.
	//
	// https://developer.fastly.com/reference/api/logging/logshuttle
	LoggingLogshuttle *LoggingLogshuttle
	// Fastly will upload log messages to New Relic Logs in the format specified in the New Relic configuration object.
	//
	// https://developer.fastly.com/reference/api/logging/newrelic
	LoggingNewrelic *LoggingNewrelic
	// Fastly will upload log messages to the OpenStack bucket in the format specified in the openstack object.
	//
	// https://developer.fastly.com/reference/api/logging/openstack
	LoggingOpenstack *LoggingOpenstack
	// Fastly will stream log messages to the Papertrail account in the format specified in the Papertrail object.
	//
	// https://developer.fastly.com/reference/api/logging/papertrail
	LoggingPapertrail *LoggingPapertrail
	// Fastly will publish log messages to a Google Cloud Pub/Sub topic in the format specified in the Pub/Sub logging object.
	//
	// https://developer.fastly.com/reference/api/logging/gcp-pubsub
	LoggingPubsub *LoggingPubsub
	// Fastly will upload log messages to the S3 bucket in the format specified in the S3 object.
	//
	// https://developer.fastly.com/reference/api/logging/s3
	LoggingS3 *LoggingS3
	// Fastly will stream log messages to the Scalyr account in the format specified in the Scalyr object.
	//
	// https://developer.fastly.com/reference/api/logging/scalyr
	LoggingScalyr *LoggingScalyr
	// Fastly will upload log messages periodically to the server in the format specified in the SFTP object.
	//
	// https://developer.fastly.com/reference/api/logging/sftp
	LoggingSftp *LoggingSftp
	// Fastly will POST messages to your Splunk account in the format specified in the Splunk object.
	//
	// https://developer.fastly.com/reference/api/logging/splunk
	LoggingSplunk *LoggingSplunk
	// Fastly will POST messages to the Sumo Logic account in the format specified in the Sumologic object.
	//
	// https://developer.fastly.com/reference/api/logging/sumologic
	LoggingSumologic *LoggingSumologic
	// Fastly will stream log messages to the location in the format specified in the Syslog object.
	//
	// https://developer.fastly.com/reference/api/logging/syslog
	LoggingSyslog *LoggingSyslog
	// Compute@Edge is a computation platform capable of running custom binary packages that you compile on your own systems and upload to Fastly. These packages are associated with a service version and are deployed to Fastly's edge network.
	//
	//
	// https://developer.fastly.com/reference/api/services/package
	Package *Package
	// A pool is responsible for balancing requests among a group of servers. In addition to balancing, pools can be configured to attempt retrying failed requests. Pools have a quorum setting that can be used to determine when the pool as a whole is considered up, in order to prevent problems following an outage as servers come back up.
	//
	// https://developer.fastly.com/reference/api/load-balancing/pools/pool
	Pool *Pool
	// Settings used to customize Fastly's request handling. When used with [Conditions](#condition) the Request Settings object allows you to fine tune how specific types of requests are handled.
	//
	// https://developer.fastly.com/reference/api/vcl-services/request-settings
	RequestSettings *RequestSettings
	// A resource link represents a link between a shared resource (such as an kv store or config store) and a service version.
	//
	// https://developer.fastly.com/reference/api/services/resource
	Resource *Resource
	// Allows you to create synthetic responses that exist entirely on the varnish machine. Useful for creating error or maintenance pages that exists outside the scope of your backend architecture. Best when used with [Condition](#condition) objects.
	//
	// https://developer.fastly.com/reference/api/vcl-services/response-object
	ResponseObject *ResponseObject
	// Handles default settings for a particular version of a service.
	//
	// https://developer.fastly.com/reference/api/vcl-services/settings
	Settings *Settings
	// A star allows users to mark services of interest.
	//
	// https://developer.fastly.com/reference/api/account/star
	Star *Star
	// The Historical Stats API allows you to programmatically retrieve historical caching statistics derived from your Fastly services. You can use these metrics to help you optimize your site’s data caching and analyze your site’s traffic.
	//
	// https://developer.fastly.com/reference/api/metrics-stats/historical-stats
	Historical *Historical
	// TLS activations.
	//
	// https://developer.fastly.com/reference/api/tls/custom-certs/activations
	TLSActivations *TLSActivations
	// Available to Platform TLS customers, these endpoints streamline the upload, deployment and management of large numbers of TLS certificates. A certificate is used to terminate TLS traffic for one or more of your fully qualified domain names (domains). Uploading a new certificate automatically enables TLS for all domains listed as Subject Alternative Names (SAN entries) on the certificate.
	//
	// https://developer.fastly.com/reference/api/tls/platform
	TLSBulkCertificates *TLSBulkCertificates
	// A TLS certificate is used to terminate TLS traffic for one or more of your [TLS domains](#tls_domains).
	//
	// https://developer.fastly.com/reference/api/tls/custom-certs/certificates
	TLSCertificates *TLSCertificates
	// Customers with access to multiple sets of IP pools are able to apply different configuration options to their TLS enabled domains.
	//
	// https://developer.fastly.com/reference/api/tls/configuration
	TLSConfigurations *TLSConfigurations
	// TLS domains are all the domains (including wildcard domains) included in any [TLS certificate](#tls_certificates)'s Subject Alternative Names (SAN) list. Included in the response is information about which certificates reference this domain as well as the [TLS activation](#tls_activations) indicating which certificate is enabled to serve TLS traffic for the domain.
	//
	// https://developer.fastly.com/reference/api/tls/custom-certs/domains
	TLSDomains *TLSDomains
	// The Mutual TLS API allows for client-to-server authentication using client-side X.509 authentication. The main Mutual Authentication object represents the certificate bundle and other configurations which support Mutual TLS for your domains.
	//
	// https://developer.fastly.com/reference/api/tls/mutual-tls/authentication
	MutualAuthentication *MutualAuthentication
	// A private key is used to sign a Certificate. A key can be used to sign multiple certificates.
	//
	// https://developer.fastly.com/reference/api/tls/custom-certs/private-keys
	TLSPrivateKeys *TLSPrivateKeys
	// The TLS subscriptions API allows you to programmatically generate TLS certificates that are procured and renewed by Fastly. Once a subscription is created for a given hostname or wildcard domain, DNS records are checked to ensure that the domain on the subscription is owned by the subscription creator. Provided DNS records are maintained, TLS certificates will automatically renew. If Fastly is unable to issue a certificate, we will retry to issue the certificate for 7 days past subscription creation or the latest certificate's not_after date, whichever is later. If after 7 days Fastly is unable to issue a certificate, the subscription state will change to `failed` and Fastly will stop retrying.
	//
	// https://developer.fastly.com/reference/api/tls/subs
	TLSSubscriptions *TLSSubscriptions
	// A user group is a collection of users with service groups and roles.
	//
	// https://developer.fastly.com/reference/api/account/user-groups
	IamUserGroups *IamUserGroups
	// The real-time analytics API offers a standardized set of data about traffic received by a specified service in one-second time periods up to the last complete second.
	//
	// https://developer.fastly.com/reference/api/metrics-stats/realtime
	Realtime *Realtime
	// A firewall represents a configuration of a Web Application Firewall (WAF) on a service. A firewall has many firewall versions through which you can manage rules.
	//
	// https://developer.fastly.com/reference/api/waf/firewall
	WafFirewalls *WafFirewalls
	// Firewall version objects contain all of the rules and settings for your WAF and remain empty until properly configured. To understand the behavior of thresholds and scores, see [Managing rules](https://docs.fastly.com/en/guides/managing-rules-on-the-fastly-waf#understanding-scoring-active-rule-behavior). Newly created firewall versions are initiated without any associated rules. See [Active Rules](/reference/api/waf/rules/active/) for details. Changes to your WAF's rules and settings can be made by [cloning](/reference/api/waf/firewall-version/#clone-waf-firewall-version) an existing firewall version, making the changes, and then [activating](/reference/api/waf/firewall-version/#deploy-activate-waf-firewall-version) the new firewall version.
	//
	// https://developer.fastly.com/reference/api/waf/firewall/version
	WafFirewallVersions *WafFirewallVersions
	// WAF rule exclusions provide a flexible way to handle false positives, allowing specific variables, rules, and the entire WAF to be excluded on a per request basis. You can configure up to 300 WAF exclusions and each exclusion of type `rule` can have up to 30 rules associated with it.
	//
	//
	// https://developer.fastly.com/reference/api/waf/rules/exclusions
	WafExclusions *WafExclusions
	// An active rule represents a [rule revision](/reference/api/waf/rules/revisions/) added to a particular [firewall version](/reference/api/waf/firewall-version/).
	//
	// https://developer.fastly.com/reference/api/waf/rules/active
	WafActiveRules *WafActiveRules
	// Rules are universally available for every firewall. Rules can have one or multiple [rule revisions](/reference/api/waf/rules/revisions/). You can add rules to your firewall by creating [active rules](/reference/api/waf/rules/active/).
	//
	// https://developer.fastly.com/reference/api/waf/rules
	WafRules *WafRules
	// Rule revisions track new variations of [rules](/reference/api/waf/rules/) as they change over time with enhancements, fixes, and improvements. This object allows you to find a specific variation of a rule for use in your application. An [active rule](/reference/api/waf/rules/active/) on a firewall uses a specific rule revision.
	//
	// https://developer.fastly.com/reference/api/waf/rules/revisions
	WafRuleRevisions *WafRuleRevisions
	// Tags for categorizing WAF [rules](/reference/api/waf/rules/).
	//
	// https://developer.fastly.com/reference/api/waf/tags
	WafTags *WafTags

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*SDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *SDK) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details

func WithSecurity(token string) SDKOption {
	return func(sdk *SDK) {
		security := components.Security{Token: token}
		sdk.sdkConfiguration.Security = withSecurity(&security)
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "1.1.0",
			GenVersion:        "2.187.7",
			UserAgent:         "speakeasy-sdk/go 1.1.0 2.187.7 1.0.0 Fastly",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.ApexRedirect = newApexRedirect(sdk.sdkConfiguration)

	sdk.Billing = newBilling(sdk.sdkConfiguration)

	sdk.Content = newContent(sdk.sdkConfiguration)

	sdk.Customer = newCustomer(sdk.sdkConfiguration)

	sdk.User = newUser(sdk.sdkConfiguration)

	sdk.BillingAddress = newBillingAddress(sdk.sdkConfiguration)

	sdk.Contact = newContact(sdk.sdkConfiguration)

	sdk.Tokens = newTokens(sdk.sdkConfiguration)

	sdk.Pop = newPop(sdk.sdkConfiguration)

	sdk.DomainOwnerships = newDomainOwnerships(sdk.sdkConfiguration)

	sdk.EnabledProducts = newEnabledProducts(sdk.sdkConfiguration)

	sdk.Events = newEvents(sdk.sdkConfiguration)

	sdk.Invitations = newInvitations(sdk.sdkConfiguration)

	sdk.IamPermissions = newIamPermissions(sdk.sdkConfiguration)

	sdk.PublicIPList = newPublicIPList(sdk.sdkConfiguration)

	sdk.Purge = newPurge(sdk.sdkConfiguration)

	sdk.RateLimiter = newRateLimiter(sdk.sdkConfiguration)

	sdk.ConfigStore = newConfigStore(sdk.sdkConfiguration)

	sdk.ConfigStoreItem = newConfigStoreItem(sdk.sdkConfiguration)

	sdk.KvStore = newKvStore(sdk.sdkConfiguration)

	sdk.KvStoreItem = newKvStoreItem(sdk.sdkConfiguration)

	sdk.IamRoles = newIamRoles(sdk.sdkConfiguration)

	sdk.Service = newService(sdk.sdkConfiguration)

	sdk.ServiceAuthorizations = newServiceAuthorizations(sdk.sdkConfiguration)

	sdk.IamServiceGroups = newIamServiceGroups(sdk.sdkConfiguration)

	sdk.ACLEntry = newACLEntry(sdk.sdkConfiguration)

	sdk.DictionaryItem = newDictionaryItem(sdk.sdkConfiguration)

	sdk.Diff = newDiff(sdk.sdkConfiguration)

	sdk.Server = newServer(sdk.sdkConfiguration)

	sdk.Publish = newPublish(sdk.sdkConfiguration)

	sdk.Snippet = newSnippet(sdk.sdkConfiguration)

	sdk.Stats = newStats(sdk.sdkConfiguration)

	sdk.VclDiff = newVclDiff(sdk.sdkConfiguration)

	sdk.Version = newVersion(sdk.sdkConfiguration)

	sdk.ACL = newACL(sdk.sdkConfiguration)

	sdk.Backend = newBackend(sdk.sdkConfiguration)

	sdk.Vcl = newVcl(sdk.sdkConfiguration)

	sdk.CacheSettings = newCacheSettings(sdk.sdkConfiguration)

	sdk.Condition = newCondition(sdk.sdkConfiguration)

	sdk.Dictionary = newDictionary(sdk.sdkConfiguration)

	sdk.DictionaryInfo = newDictionaryInfo(sdk.sdkConfiguration)

	sdk.Director = newDirector(sdk.sdkConfiguration)

	sdk.DirectorBackend = newDirectorBackend(sdk.sdkConfiguration)

	sdk.Domain = newDomain(sdk.sdkConfiguration)

	sdk.Gzip = newGzip(sdk.sdkConfiguration)

	sdk.Header = newHeader(sdk.sdkConfiguration)

	sdk.Healthcheck = newHealthcheck(sdk.sdkConfiguration)

	sdk.Http3 = newHttp3(sdk.sdkConfiguration)

	sdk.LoggingAzureblob = newLoggingAzureblob(sdk.sdkConfiguration)

	sdk.LoggingBigquery = newLoggingBigquery(sdk.sdkConfiguration)

	sdk.LoggingCloudfiles = newLoggingCloudfiles(sdk.sdkConfiguration)

	sdk.LoggingDatadog = newLoggingDatadog(sdk.sdkConfiguration)

	sdk.LoggingDigitalocean = newLoggingDigitalocean(sdk.sdkConfiguration)

	sdk.LoggingElasticsearch = newLoggingElasticsearch(sdk.sdkConfiguration)

	sdk.LoggingFtp = newLoggingFtp(sdk.sdkConfiguration)

	sdk.LoggingGcs = newLoggingGcs(sdk.sdkConfiguration)

	sdk.LoggingHeroku = newLoggingHeroku(sdk.sdkConfiguration)

	sdk.LoggingHoneycomb = newLoggingHoneycomb(sdk.sdkConfiguration)

	sdk.LoggingHTTPS = newLoggingHTTPS(sdk.sdkConfiguration)

	sdk.LoggingKafka = newLoggingKafka(sdk.sdkConfiguration)

	sdk.LoggingKinesis = newLoggingKinesis(sdk.sdkConfiguration)

	sdk.LoggingLogentries = newLoggingLogentries(sdk.sdkConfiguration)

	sdk.LoggingLoggly = newLoggingLoggly(sdk.sdkConfiguration)

	sdk.LoggingLogshuttle = newLoggingLogshuttle(sdk.sdkConfiguration)

	sdk.LoggingNewrelic = newLoggingNewrelic(sdk.sdkConfiguration)

	sdk.LoggingOpenstack = newLoggingOpenstack(sdk.sdkConfiguration)

	sdk.LoggingPapertrail = newLoggingPapertrail(sdk.sdkConfiguration)

	sdk.LoggingPubsub = newLoggingPubsub(sdk.sdkConfiguration)

	sdk.LoggingS3 = newLoggingS3(sdk.sdkConfiguration)

	sdk.LoggingScalyr = newLoggingScalyr(sdk.sdkConfiguration)

	sdk.LoggingSftp = newLoggingSftp(sdk.sdkConfiguration)

	sdk.LoggingSplunk = newLoggingSplunk(sdk.sdkConfiguration)

	sdk.LoggingSumologic = newLoggingSumologic(sdk.sdkConfiguration)

	sdk.LoggingSyslog = newLoggingSyslog(sdk.sdkConfiguration)

	sdk.Package = newPackage(sdk.sdkConfiguration)

	sdk.Pool = newPool(sdk.sdkConfiguration)

	sdk.RequestSettings = newRequestSettings(sdk.sdkConfiguration)

	sdk.Resource = newResource(sdk.sdkConfiguration)

	sdk.ResponseObject = newResponseObject(sdk.sdkConfiguration)

	sdk.Settings = newSettings(sdk.sdkConfiguration)

	sdk.Star = newStar(sdk.sdkConfiguration)

	sdk.Historical = newHistorical(sdk.sdkConfiguration)

	sdk.TLSActivations = newTLSActivations(sdk.sdkConfiguration)

	sdk.TLSBulkCertificates = newTLSBulkCertificates(sdk.sdkConfiguration)

	sdk.TLSCertificates = newTLSCertificates(sdk.sdkConfiguration)

	sdk.TLSConfigurations = newTLSConfigurations(sdk.sdkConfiguration)

	sdk.TLSDomains = newTLSDomains(sdk.sdkConfiguration)

	sdk.MutualAuthentication = newMutualAuthentication(sdk.sdkConfiguration)

	sdk.TLSPrivateKeys = newTLSPrivateKeys(sdk.sdkConfiguration)

	sdk.TLSSubscriptions = newTLSSubscriptions(sdk.sdkConfiguration)

	sdk.IamUserGroups = newIamUserGroups(sdk.sdkConfiguration)

	sdk.Realtime = newRealtime(sdk.sdkConfiguration)

	sdk.WafFirewalls = newWafFirewalls(sdk.sdkConfiguration)

	sdk.WafFirewallVersions = newWafFirewallVersions(sdk.sdkConfiguration)

	sdk.WafExclusions = newWafExclusions(sdk.sdkConfiguration)

	sdk.WafActiveRules = newWafActiveRules(sdk.sdkConfiguration)

	sdk.WafRules = newWafRules(sdk.sdkConfiguration)

	sdk.WafRuleRevisions = newWafRuleRevisions(sdk.sdkConfiguration)

	sdk.WafTags = newWafTags(sdk.sdkConfiguration)

	return sdk
}
