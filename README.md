# Fastly

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/fastly-test-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
	"context"
	"log"
)

func main() {
	s := fastly.New(
		fastly.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.ApexRedirect.DeleteApexRedirect(ctx, operations.DeleteApexRedirectRequest{
		ApexRedirectID: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [ApexRedirect](docs/sdks/apexredirect/README.md)

* [DeleteApexRedirect](docs/sdks/apexredirect/README.md#deleteapexredirect) - Delete an apex redirect
* [GetApexRedirect](docs/sdks/apexredirect/README.md#getapexredirect) - Get an apex redirect
* [ListApexRedirects](docs/sdks/apexredirect/README.md#listapexredirects) - List apex redirects
* [UpdateApexRedirect](docs/sdks/apexredirect/README.md#updateapexredirect) - Update an apex redirect

### [Billing](docs/sdks/billing/README.md)

* [GetInvoice](docs/sdks/billing/README.md#getinvoice) - Get an invoice
* [GetInvoiceByID](docs/sdks/billing/README.md#getinvoicebyid) - Get an invoice
* [GetInvoiceMtd](docs/sdks/billing/README.md#getinvoicemtd) - Get month-to-date billing estimate

### [Content](docs/sdks/content/README.md)

* [ContentCheck](docs/sdks/content/README.md#contentcheck) - Check status of content in each POP's cache

### [Customer](docs/sdks/customer/README.md)

* [DeleteCustomer](docs/sdks/customer/README.md#deletecustomer) - Delete a customer
* [GetCustomer](docs/sdks/customer/README.md#getcustomer) - Get a customer
* [GetLoggedInCustomer](docs/sdks/customer/README.md#getloggedincustomer) - Get the logged in customer
* [ListUsers](docs/sdks/customer/README.md#listusers) - List users
* [UpdateCustomer](docs/sdks/customer/README.md#updatecustomer) - Update a customer

### [User](docs/sdks/user/README.md)

* [CreateUser](docs/sdks/user/README.md#createuser) - Create a user
* [DeleteUser](docs/sdks/user/README.md#deleteuser) - Delete a user
* [GetCurrentUser](docs/sdks/user/README.md#getcurrentuser) - Get the current user
* [GetUser](docs/sdks/user/README.md#getuser) - Get a user
* [RequestPasswordReset](docs/sdks/user/README.md#requestpasswordreset) - Request a password reset
* [UpdateUser](docs/sdks/user/README.md#updateuser) - Update a user
* [UpdateUserPassword](docs/sdks/user/README.md#updateuserpassword) - Update the user's password

### [BillingAddress](docs/sdks/billingaddress/README.md)

* [AddBillingAddr](docs/sdks/billingaddress/README.md#addbillingaddr) - Add a billing address to a customer
* [DeleteBillingAddr](docs/sdks/billingaddress/README.md#deletebillingaddr) - Delete a billing address
* [GetBillingAddr](docs/sdks/billingaddress/README.md#getbillingaddr) - Get a billing address
* [UpdateBillingAddr](docs/sdks/billingaddress/README.md#updatebillingaddr) - Update a billing address

### [Contact](docs/sdks/contact/README.md)

* [DeleteContact](docs/sdks/contact/README.md#deletecontact) - Delete a contact
* [ListContacts](docs/sdks/contact/README.md#listcontacts) - List contacts

### [Tokens](docs/sdks/tokens/README.md)

* [GetTokenCurrent](docs/sdks/tokens/README.md#gettokencurrent) - Get the current token
* [ListTokensCustomer](docs/sdks/tokens/README.md#listtokenscustomer) - List tokens for a customer
* [ListTokensUser](docs/sdks/tokens/README.md#listtokensuser) - List tokens for the authenticated user
* [RevokeToken](docs/sdks/tokens/README.md#revoketoken) - Revoke a token
* [RevokeTokenCurrent](docs/sdks/tokens/README.md#revoketokencurrent) - Revoke the current token

### [Pop](docs/sdks/pop/README.md)

* [ListPops](docs/sdks/pop/README.md#listpops) - List Fastly POPs

### [DomainOwnerships](docs/sdks/domainownerships/README.md)

* [ListDomainOwnerships](docs/sdks/domainownerships/README.md#listdomainownerships) - List domain-ownerships

### [EnabledProducts](docs/sdks/enabledproducts/README.md)

* [DisableProduct](docs/sdks/enabledproducts/README.md#disableproduct) - Disable a product
* [EnableProduct](docs/sdks/enabledproducts/README.md#enableproduct) - Enable a product
* [GetEnabledProduct](docs/sdks/enabledproducts/README.md#getenabledproduct) - Get enabled product

### [Events](docs/sdks/events/README.md)

* [GetEvent](docs/sdks/events/README.md#getevent) - Get an event
* [ListEvents](docs/sdks/events/README.md#listevents) - List events

### [Invitations](docs/sdks/invitations/README.md)

* [CreateInvitation](docs/sdks/invitations/README.md#createinvitation) - Create an invitation
* [DeleteInvitation](docs/sdks/invitations/README.md#deleteinvitation) - Delete an invitation
* [ListInvitations](docs/sdks/invitations/README.md#listinvitations) - List invitations

### [IamPermissions](docs/sdks/iampermissions/README.md)

* [ListPermissions](docs/sdks/iampermissions/README.md#listpermissions) - List permissions

### [PublicIPList](docs/sdks/publiciplist/README.md)

* [ListFastlyIps](docs/sdks/publiciplist/README.md#listfastlyips) - List Fastly's public IPs

### [Purge](docs/sdks/purge/README.md)

* [PurgeAll](docs/sdks/purge/README.md#purgeall) - Purge everything from a service
* [PurgeSingleURL](docs/sdks/purge/README.md#purgesingleurl) - Purge a URL
* [PurgeTag](docs/sdks/purge/README.md#purgetag) - Purge by surrogate key tag

### [RateLimiter](docs/sdks/ratelimiter/README.md)

* [DeleteRateLimiter](docs/sdks/ratelimiter/README.md#deleteratelimiter) - Delete a rate limiter
* [GetRateLimiter](docs/sdks/ratelimiter/README.md#getratelimiter) - Get a rate limiter
* [ListRateLimiters](docs/sdks/ratelimiter/README.md#listratelimiters) - List rate limiters

### [ConfigStore](docs/sdks/configstore/README.md)

* [CreateConfigStore](docs/sdks/configstore/README.md#createconfigstore) - Create a config store
* [DeleteConfigStore](docs/sdks/configstore/README.md#deleteconfigstore) - Delete a config store
* [GetConfigStore](docs/sdks/configstore/README.md#getconfigstore) - Describe a config store
* [GetConfigStoreInfo](docs/sdks/configstore/README.md#getconfigstoreinfo) - Get config store metadata
* [ListConfigStoreServices](docs/sdks/configstore/README.md#listconfigstoreservices) - List linked services
* [ListConfigStores](docs/sdks/configstore/README.md#listconfigstores) - List config stores
* [UpdateConfigStore](docs/sdks/configstore/README.md#updateconfigstore) - Update a config store

### [ConfigStoreItem](docs/sdks/configstoreitem/README.md)

* [BulkUpdateConfigStoreItem](docs/sdks/configstoreitem/README.md#bulkupdateconfigstoreitem) - Update multiple entries in a config store
* [CreateConfigStoreItem](docs/sdks/configstoreitem/README.md#createconfigstoreitem) - Create an entry in a config store
* [DeleteConfigStoreItem](docs/sdks/configstoreitem/README.md#deleteconfigstoreitem) - Delete an item from a config store
* [GetConfigStoreItem](docs/sdks/configstoreitem/README.md#getconfigstoreitem) - Get an item from a config store
* [ListConfigStoreItems](docs/sdks/configstoreitem/README.md#listconfigstoreitems) - List items in a config store
* [UpdateConfigStoreItem](docs/sdks/configstoreitem/README.md#updateconfigstoreitem) - Update an entry in a config store
* [UpsertConfigStoreItem](docs/sdks/configstoreitem/README.md#upsertconfigstoreitem) - Insert or update an entry in a config store

### [KvStore](docs/sdks/kvstore/README.md)

* [CreateStore](docs/sdks/kvstore/README.md#createstore) - Create an kv store.
* [DeleteStore](docs/sdks/kvstore/README.md#deletestore) - Delete an kv store.
* [GetStore](docs/sdks/kvstore/README.md#getstore) - Describe an kv store.
* [GetStores](docs/sdks/kvstore/README.md#getstores) - List kv stores.

### [KvStoreItem](docs/sdks/kvstoreitem/README.md)

* [DeleteKeyFromStore](docs/sdks/kvstoreitem/README.md#deletekeyfromstore) - Delete kv store item.
* [GetKeys](docs/sdks/kvstoreitem/README.md#getkeys) - List kv store keys.
* [GetValueForKey](docs/sdks/kvstoreitem/README.md#getvalueforkey) - Get the value of an kv store item
* [SetValueForKey](docs/sdks/kvstoreitem/README.md#setvalueforkey) - Insert an item into an kv store

### [IamRoles](docs/sdks/iamroles/README.md)

* [DeleteARole](docs/sdks/iamroles/README.md#deletearole) - Delete a role
* [GetARole](docs/sdks/iamroles/README.md#getarole) - Get a role
* [ListRolePermissions](docs/sdks/iamroles/README.md#listrolepermissions) - List permissions in a role
* [ListRoles](docs/sdks/iamroles/README.md#listroles) - List roles

### [Service](docs/sdks/service/README.md)

* [CreateService](docs/sdks/service/README.md#createservice) - Create a service
* [DeleteService](docs/sdks/service/README.md#deleteservice) - Delete a service
* [GetService](docs/sdks/service/README.md#getservice) - Get a service
* [GetServiceDetail](docs/sdks/service/README.md#getservicedetail) - Get service details
* [ListServiceDomains](docs/sdks/service/README.md#listservicedomains) - List the domains within a service
* [ListServices](docs/sdks/service/README.md#listservices) - List services
* [SearchService](docs/sdks/service/README.md#searchservice) - Search for a service by name
* [UpdateService](docs/sdks/service/README.md#updateservice) - Update a service

### [ServiceAuthorizations](docs/sdks/serviceauthorizations/README.md)

* [CreateServiceAuthorization](docs/sdks/serviceauthorizations/README.md#createserviceauthorization) - Create service authorization
* [DeleteServiceAuthorization](docs/sdks/serviceauthorizations/README.md#deleteserviceauthorization) - Delete service authorization
* [ListServiceAuthorization](docs/sdks/serviceauthorizations/README.md#listserviceauthorization) - List service authorizations
* [ShowServiceAuthorization](docs/sdks/serviceauthorizations/README.md#showserviceauthorization) - Show service authorization
* [UpdateServiceAuthorization](docs/sdks/serviceauthorizations/README.md#updateserviceauthorization) - Update service authorization

### [IamServiceGroups](docs/sdks/iamservicegroups/README.md)

* [DeleteAServiceGroup](docs/sdks/iamservicegroups/README.md#deleteaservicegroup) - Delete a service group
* [GetAServiceGroup](docs/sdks/iamservicegroups/README.md#getaservicegroup) - Get a service group
* [ListServiceGroupServices](docs/sdks/iamservicegroups/README.md#listservicegroupservices) - List services to a service group
* [ListServiceGroups](docs/sdks/iamservicegroups/README.md#listservicegroups) - List service groups

### [ACLEntry](docs/sdks/aclentry/README.md)

* [BulkUpdateACLEntries](docs/sdks/aclentry/README.md#bulkupdateaclentries) - Update multiple ACL entries
* [CreateACLEntry](docs/sdks/aclentry/README.md#createaclentry) - Create an ACL entry
* [DeleteACLEntry](docs/sdks/aclentry/README.md#deleteaclentry) - Delete an ACL entry
* [GetACLEntry](docs/sdks/aclentry/README.md#getaclentry) - Describe an ACL entry
* [ListACLEntries](docs/sdks/aclentry/README.md#listaclentries) - List ACL entries
* [UpdateACLEntry](docs/sdks/aclentry/README.md#updateaclentry) - Update an ACL entry

### [DictionaryItem](docs/sdks/dictionaryitem/README.md)

* [BulkUpdateDictionaryItem](docs/sdks/dictionaryitem/README.md#bulkupdatedictionaryitem) - Update multiple entries in an edge dictionary
* [CreateDictionaryItem](docs/sdks/dictionaryitem/README.md#createdictionaryitem) - Create an entry in an edge dictionary
* [DeleteDictionaryItem](docs/sdks/dictionaryitem/README.md#deletedictionaryitem) - Delete an item from an edge dictionary
* [GetDictionaryItem](docs/sdks/dictionaryitem/README.md#getdictionaryitem) - Get an item from an edge dictionary
* [ListDictionaryItems](docs/sdks/dictionaryitem/README.md#listdictionaryitems) - List items in an edge dictionary
* [UpdateDictionaryItem](docs/sdks/dictionaryitem/README.md#updatedictionaryitem) - Update an entry in an edge dictionary
* [UpsertDictionaryItem](docs/sdks/dictionaryitem/README.md#upsertdictionaryitem) - Insert or update an entry in an edge dictionary

### [Diff](docs/sdks/diff/README.md)

* [DiffServiceVersions](docs/sdks/diff/README.md#diffserviceversions) - Diff two service versions

### [Server](docs/sdks/server/README.md)

* [CreatePoolServer](docs/sdks/server/README.md#createpoolserver) - Add a server to a pool
* [DeletePoolServer](docs/sdks/server/README.md#deletepoolserver) - Delete a server from a pool
* [GetPoolServer](docs/sdks/server/README.md#getpoolserver) - Get a pool server
* [ListPoolServers](docs/sdks/server/README.md#listpoolservers) - List servers in a pool
* [UpdatePoolServer](docs/sdks/server/README.md#updatepoolserver) - Update a server

### [Publish](docs/sdks/publish/README.md)

* [Publish](docs/sdks/publish/README.md#publish) - Send messages to Fanout subscribers

### [Snippet](docs/sdks/snippet/README.md)

* [CreateSnippet](docs/sdks/snippet/README.md#createsnippet) - Create a snippet
* [DeleteSnippet](docs/sdks/snippet/README.md#deletesnippet) - Delete a snippet
* [GetSnippet](docs/sdks/snippet/README.md#getsnippet) - Get a versioned snippet
* [GetSnippetDynamic](docs/sdks/snippet/README.md#getsnippetdynamic) - Get a dynamic snippet
* [ListSnippets](docs/sdks/snippet/README.md#listsnippets) - List snippets
* [UpdateSnippetDynamic](docs/sdks/snippet/README.md#updatesnippetdynamic) - Update a dynamic snippet

### [Stats](docs/sdks/stats/README.md)

* [GetServiceStats](docs/sdks/stats/README.md#getservicestats) - Get stats for a service

### [VclDiff](docs/sdks/vcldiff/README.md)

* [VclDiffServiceVersions](docs/sdks/vcldiff/README.md#vcldiffserviceversions) - Get a comparison of the VCL changes between two service versions

### [Version](docs/sdks/version/README.md)

* [ActivateServiceVersion](docs/sdks/version/README.md#activateserviceversion) - Activate a service version
* [CloneServiceVersion](docs/sdks/version/README.md#cloneserviceversion) - Clone a service version
* [CreateServiceVersion](docs/sdks/version/README.md#createserviceversion) - Create a service version
* [DeactivateServiceVersion](docs/sdks/version/README.md#deactivateserviceversion) - Deactivate a service version
* [GetServiceVersion](docs/sdks/version/README.md#getserviceversion) - Get a version of a service
* [ListServiceVersions](docs/sdks/version/README.md#listserviceversions) - List versions of a service
* [LockServiceVersion](docs/sdks/version/README.md#lockserviceversion) - Lock a service version
* [UpdateServiceVersion](docs/sdks/version/README.md#updateserviceversion) - Update a service version
* [ValidateServiceVersion](docs/sdks/version/README.md#validateserviceversion) - Validate a service version

### [ACL](docs/sdks/acl/README.md)

* [CreateACL](docs/sdks/acl/README.md#createacl) - Create a new ACL
* [DeleteACL](docs/sdks/acl/README.md#deleteacl) - Delete an ACL
* [GetACL](docs/sdks/acl/README.md#getacl) - Describe an ACL
* [ListAcls](docs/sdks/acl/README.md#listacls) - List ACLs
* [UpdateACL](docs/sdks/acl/README.md#updateacl) - Update an ACL

### [Backend](docs/sdks/backend/README.md)

* [CreateBackend](docs/sdks/backend/README.md#createbackend) - Create a backend
* [DeleteBackend](docs/sdks/backend/README.md#deletebackend) - Delete a backend
* [GetBackend](docs/sdks/backend/README.md#getbackend) - Describe a backend
* [ListBackends](docs/sdks/backend/README.md#listbackends) - List backends
* [UpdateBackend](docs/sdks/backend/README.md#updatebackend) - Update a backend

### [Vcl](docs/sdks/vcl/README.md)

* [CreateCustomVcl](docs/sdks/vcl/README.md#createcustomvcl) - Create a custom VCL file
* [DeleteCustomVcl](docs/sdks/vcl/README.md#deletecustomvcl) - Delete a custom VCL file
* [GetCustomVcl](docs/sdks/vcl/README.md#getcustomvcl) - Get a custom VCL file
* [GetCustomVclBoilerplate](docs/sdks/vcl/README.md#getcustomvclboilerplate) - Get boilerplate VCL
* [GetCustomVclGenerated](docs/sdks/vcl/README.md#getcustomvclgenerated) - Get the generated VCL for a service
* [GetCustomVclRaw](docs/sdks/vcl/README.md#getcustomvclraw) - Download a custom VCL file
* [ListCustomVcl](docs/sdks/vcl/README.md#listcustomvcl) - List custom VCL files
* [SetCustomVclMain](docs/sdks/vcl/README.md#setcustomvclmain) - Set a custom VCL file as main
* [UpdateCustomVcl](docs/sdks/vcl/README.md#updatecustomvcl) - Update a custom VCL file

### [CacheSettings](docs/sdks/cachesettings/README.md)

* [CreateCacheSettings](docs/sdks/cachesettings/README.md#createcachesettings) - Create a cache settings object
* [DeleteCacheSettings](docs/sdks/cachesettings/README.md#deletecachesettings) - Delete a cache settings object
* [GetCacheSettings](docs/sdks/cachesettings/README.md#getcachesettings) - Get a cache settings object
* [ListCacheSettings](docs/sdks/cachesettings/README.md#listcachesettings) - List cache settings objects
* [UpdateCacheSettings](docs/sdks/cachesettings/README.md#updatecachesettings) - Update a cache settings object

### [Condition](docs/sdks/condition/README.md)

* [CreateCondition](docs/sdks/condition/README.md#createcondition) - Create a condition
* [DeleteCondition](docs/sdks/condition/README.md#deletecondition) - Delete a condition
* [GetCondition](docs/sdks/condition/README.md#getcondition) - Describe a condition
* [ListConditions](docs/sdks/condition/README.md#listconditions) - List conditions
* [UpdateCondition](docs/sdks/condition/README.md#updatecondition) - Update a condition

### [Dictionary](docs/sdks/dictionary/README.md)

* [CreateDictionary](docs/sdks/dictionary/README.md#createdictionary) - Create an edge dictionary
* [DeleteDictionary](docs/sdks/dictionary/README.md#deletedictionary) - Delete an edge dictionary
* [GetDictionary](docs/sdks/dictionary/README.md#getdictionary) - Get an edge dictionary
* [ListDictionaries](docs/sdks/dictionary/README.md#listdictionaries) - List edge dictionaries
* [UpdateDictionary](docs/sdks/dictionary/README.md#updatedictionary) - Update an edge dictionary

### [DictionaryInfo](docs/sdks/dictionaryinfo/README.md)

* [GetDictionaryInfo](docs/sdks/dictionaryinfo/README.md#getdictionaryinfo) - Get edge dictionary metadata

### [Director](docs/sdks/director/README.md)

* [CreateDirector](docs/sdks/director/README.md#createdirector) - Create a director
* [DeleteDirector](docs/sdks/director/README.md#deletedirector) - Delete a director
* [GetDirector](docs/sdks/director/README.md#getdirector) - Get a director
* [ListDirectors](docs/sdks/director/README.md#listdirectors) - List directors

### [DirectorBackend](docs/sdks/directorbackend/README.md)

* [CreateDirectorBackend](docs/sdks/directorbackend/README.md#createdirectorbackend) - Create a director-backend relationship
* [DeleteDirectorBackend](docs/sdks/directorbackend/README.md#deletedirectorbackend) - Delete a director-backend relationship
* [GetDirectorBackend](docs/sdks/directorbackend/README.md#getdirectorbackend) - Get a director-backend relationship

### [Domain](docs/sdks/domain/README.md)

* [CheckDomain](docs/sdks/domain/README.md#checkdomain) - Validate DNS configuration for a single domain on a service
* [CheckDomains](docs/sdks/domain/README.md#checkdomains) - Validate DNS configuration for all domains on a service
* [CreateDomain](docs/sdks/domain/README.md#createdomain) - Add a domain name to a service
* [DeleteDomain](docs/sdks/domain/README.md#deletedomain) - Remove a domain from a service
* [GetDomain](docs/sdks/domain/README.md#getdomain) - Describe a domain
* [ListDomains](docs/sdks/domain/README.md#listdomains) - List domains
* [UpdateDomain](docs/sdks/domain/README.md#updatedomain) - Update a domain

### [Gzip](docs/sdks/gzip/README.md)

* [CreateGzipConfig](docs/sdks/gzip/README.md#creategzipconfig) - Create a gzip configuration
* [DeleteGzipConfig](docs/sdks/gzip/README.md#deletegzipconfig) - Delete a gzip configuration
* [GetGzipConfigs](docs/sdks/gzip/README.md#getgzipconfigs) - Get a gzip configuration
* [ListGzipConfigs](docs/sdks/gzip/README.md#listgzipconfigs) - List gzip configurations
* [UpdateGzipConfig](docs/sdks/gzip/README.md#updategzipconfig) - Update a gzip configuration

### [Header](docs/sdks/header/README.md)

* [CreateHeaderObject](docs/sdks/header/README.md#createheaderobject) - Create a Header object
* [DeleteHeaderObject](docs/sdks/header/README.md#deleteheaderobject) - Delete a Header object
* [GetHeaderObject](docs/sdks/header/README.md#getheaderobject) - Get a Header object
* [ListHeaderObjects](docs/sdks/header/README.md#listheaderobjects) - List Header objects
* [UpdateHeaderObject](docs/sdks/header/README.md#updateheaderobject) - Update a Header object

### [Healthcheck](docs/sdks/healthcheck/README.md)

* [CreateHealthcheck](docs/sdks/healthcheck/README.md#createhealthcheck) - Create a health check
* [DeleteHealthcheck](docs/sdks/healthcheck/README.md#deletehealthcheck) - Delete a health check
* [GetHealthcheck](docs/sdks/healthcheck/README.md#gethealthcheck) - Get a health check
* [ListHealthchecks](docs/sdks/healthcheck/README.md#listhealthchecks) - List health checks
* [UpdateHealthcheck](docs/sdks/healthcheck/README.md#updatehealthcheck) - Update a health check

### [Http3](docs/sdks/http3/README.md)

* [CreateHttp3](docs/sdks/http3/README.md#createhttp3) - Enable support for HTTP/3
* [DeleteHttp3](docs/sdks/http3/README.md#deletehttp3) - Disable support for HTTP/3
* [GetHttp3](docs/sdks/http3/README.md#gethttp3) - Get HTTP/3 status

### [LoggingAzureblob](docs/sdks/loggingazureblob/README.md)

* [CreateLogAzure](docs/sdks/loggingazureblob/README.md#createlogazure) - Create an Azure Blob Storage log endpoint
* [DeleteLogAzure](docs/sdks/loggingazureblob/README.md#deletelogazure) - Delete the Azure Blob Storage log endpoint
* [GetLogAzure](docs/sdks/loggingazureblob/README.md#getlogazure) - Get an Azure Blob Storage log endpoint
* [ListLogAzure](docs/sdks/loggingazureblob/README.md#listlogazure) - List Azure Blob Storage log endpoints
* [UpdateLogAzure](docs/sdks/loggingazureblob/README.md#updatelogazure) - Update an Azure Blob Storage log endpoint

### [LoggingBigquery](docs/sdks/loggingbigquery/README.md)

* [CreateLogBigquery](docs/sdks/loggingbigquery/README.md#createlogbigquery) - Create a BigQuery log endpoint
* [DeleteLogBigquery](docs/sdks/loggingbigquery/README.md#deletelogbigquery) - Delete a BigQuery log endpoint
* [GetLogBigquery](docs/sdks/loggingbigquery/README.md#getlogbigquery) - Get a BigQuery log endpoint
* [ListLogBigquery](docs/sdks/loggingbigquery/README.md#listlogbigquery) - List BigQuery log endpoints
* [UpdateLogBigquery](docs/sdks/loggingbigquery/README.md#updatelogbigquery) - Update a BigQuery log endpoint

### [LoggingCloudfiles](docs/sdks/loggingcloudfiles/README.md)

* [CreateLogCloudfiles](docs/sdks/loggingcloudfiles/README.md#createlogcloudfiles) - Create a Cloud Files log endpoint
* [DeleteLogCloudfiles](docs/sdks/loggingcloudfiles/README.md#deletelogcloudfiles) - Delete the Cloud Files log endpoint
* [GetLogCloudfiles](docs/sdks/loggingcloudfiles/README.md#getlogcloudfiles) - Get a Cloud Files log endpoint
* [ListLogCloudfiles](docs/sdks/loggingcloudfiles/README.md#listlogcloudfiles) - List Cloud Files log endpoints
* [UpdateLogCloudfiles](docs/sdks/loggingcloudfiles/README.md#updatelogcloudfiles) - Update the Cloud Files log endpoint

### [LoggingDatadog](docs/sdks/loggingdatadog/README.md)

* [CreateLogDatadog](docs/sdks/loggingdatadog/README.md#createlogdatadog) - Create a Datadog log endpoint
* [DeleteLogDatadog](docs/sdks/loggingdatadog/README.md#deletelogdatadog) - Delete a Datadog log endpoint
* [GetLogDatadog](docs/sdks/loggingdatadog/README.md#getlogdatadog) - Get a Datadog log endpoint
* [ListLogDatadog](docs/sdks/loggingdatadog/README.md#listlogdatadog) - List Datadog log endpoints
* [UpdateLogDatadog](docs/sdks/loggingdatadog/README.md#updatelogdatadog) - Update a Datadog log endpoint

### [LoggingDigitalocean](docs/sdks/loggingdigitalocean/README.md)

* [CreateLogDigocean](docs/sdks/loggingdigitalocean/README.md#createlogdigocean) - Create a DigitalOcean Spaces log endpoint
* [DeleteLogDigocean](docs/sdks/loggingdigitalocean/README.md#deletelogdigocean) - Delete a DigitalOcean Spaces log endpoint
* [GetLogDigocean](docs/sdks/loggingdigitalocean/README.md#getlogdigocean) - Get a DigitalOcean Spaces log endpoint
* [ListLogDigocean](docs/sdks/loggingdigitalocean/README.md#listlogdigocean) - List DigitalOcean Spaces log endpoints
* [UpdateLogDigocean](docs/sdks/loggingdigitalocean/README.md#updatelogdigocean) - Update a DigitalOcean Spaces log endpoint

### [LoggingElasticsearch](docs/sdks/loggingelasticsearch/README.md)

* [CreateLogElasticsearch](docs/sdks/loggingelasticsearch/README.md#createlogelasticsearch) - Create an Elasticsearch log endpoint
* [DeleteLogElasticsearch](docs/sdks/loggingelasticsearch/README.md#deletelogelasticsearch) - Delete an Elasticsearch log endpoint
* [GetLogElasticsearch](docs/sdks/loggingelasticsearch/README.md#getlogelasticsearch) - Get an Elasticsearch log endpoint
* [ListLogElasticsearch](docs/sdks/loggingelasticsearch/README.md#listlogelasticsearch) - List Elasticsearch log endpoints
* [UpdateLogElasticsearch](docs/sdks/loggingelasticsearch/README.md#updatelogelasticsearch) - Update an Elasticsearch log endpoint

### [LoggingFtp](docs/sdks/loggingftp/README.md)

* [CreateLogFtp](docs/sdks/loggingftp/README.md#createlogftp) - Create an FTP log endpoint
* [DeleteLogFtp](docs/sdks/loggingftp/README.md#deletelogftp) - Delete an FTP log endpoint
* [GetLogFtp](docs/sdks/loggingftp/README.md#getlogftp) - Get an FTP log endpoint
* [ListLogFtp](docs/sdks/loggingftp/README.md#listlogftp) - List FTP log endpoints
* [UpdateLogFtp](docs/sdks/loggingftp/README.md#updatelogftp) - Update an FTP log endpoint

### [LoggingGcs](docs/sdks/logginggcs/README.md)

* [CreateLogGcs](docs/sdks/logginggcs/README.md#createloggcs) - Create a GCS log endpoint
* [DeleteLogGcs](docs/sdks/logginggcs/README.md#deleteloggcs) - Delete a GCS log endpoint
* [GetLogGcs](docs/sdks/logginggcs/README.md#getloggcs) - Get a GCS log endpoint
* [ListLogGcs](docs/sdks/logginggcs/README.md#listloggcs) - List GCS log endpoints
* [UpdateLogGcs](docs/sdks/logginggcs/README.md#updateloggcs) - Update a GCS log endpoint

### [LoggingHeroku](docs/sdks/loggingheroku/README.md)

* [CreateLogHeroku](docs/sdks/loggingheroku/README.md#createlogheroku) - Create a Heroku log endpoint
* [DeleteLogHeroku](docs/sdks/loggingheroku/README.md#deletelogheroku) - Delete the Heroku log endpoint
* [GetLogHeroku](docs/sdks/loggingheroku/README.md#getlogheroku) - Get a Heroku log endpoint
* [ListLogHeroku](docs/sdks/loggingheroku/README.md#listlogheroku) - List Heroku log endpoints
* [UpdateLogHeroku](docs/sdks/loggingheroku/README.md#updatelogheroku) - Update the Heroku log endpoint

### [LoggingHoneycomb](docs/sdks/logginghoneycomb/README.md)

* [CreateLogHoneycomb](docs/sdks/logginghoneycomb/README.md#createloghoneycomb) - Create a Honeycomb log endpoint
* [DeleteLogHoneycomb](docs/sdks/logginghoneycomb/README.md#deleteloghoneycomb) - Delete the Honeycomb log endpoint
* [GetLogHoneycomb](docs/sdks/logginghoneycomb/README.md#getloghoneycomb) - Get a Honeycomb log endpoint
* [ListLogHoneycomb](docs/sdks/logginghoneycomb/README.md#listloghoneycomb) - List Honeycomb log endpoints
* [UpdateLogHoneycomb](docs/sdks/logginghoneycomb/README.md#updateloghoneycomb) - Update a Honeycomb log endpoint

### [LoggingHTTPS](docs/sdks/logginghttps/README.md)

* [CreateLogHTTPS](docs/sdks/logginghttps/README.md#createloghttps) - Create an HTTPS log endpoint
* [DeleteLogHTTPS](docs/sdks/logginghttps/README.md#deleteloghttps) - Delete an HTTPS log endpoint
* [GetLogHTTPS](docs/sdks/logginghttps/README.md#getloghttps) - Get an HTTPS log endpoint
* [ListLogHTTPS](docs/sdks/logginghttps/README.md#listloghttps) - List HTTPS log endpoints
* [UpdateLogHTTPS](docs/sdks/logginghttps/README.md#updateloghttps) - Update an HTTPS log endpoint

### [LoggingKafka](docs/sdks/loggingkafka/README.md)

* [CreateLogKafka](docs/sdks/loggingkafka/README.md#createlogkafka) - Create a Kafka log endpoint
* [DeleteLogKafka](docs/sdks/loggingkafka/README.md#deletelogkafka) - Delete the Kafka log endpoint
* [GetLogKafka](docs/sdks/loggingkafka/README.md#getlogkafka) - Get a Kafka log endpoint
* [ListLogKafka](docs/sdks/loggingkafka/README.md#listlogkafka) - List Kafka log endpoints

### [LoggingKinesis](docs/sdks/loggingkinesis/README.md)

* [CreateLogKinesis](docs/sdks/loggingkinesis/README.md#createlogkinesis) - Create  an Amazon Kinesis log endpoint
* [DeleteLogKinesis](docs/sdks/loggingkinesis/README.md#deletelogkinesis) - Delete the Amazon Kinesis log endpoint
* [GetLogKinesis](docs/sdks/loggingkinesis/README.md#getlogkinesis) - Get an Amazon Kinesis log endpoint
* [ListLogKinesis](docs/sdks/loggingkinesis/README.md#listlogkinesis) - List Amazon Kinesis log endpoints

### [LoggingLogentries](docs/sdks/logginglogentries/README.md)

* [~~CreateLogLogentries~~](docs/sdks/logginglogentries/README.md#createloglogentries) - Create a Logentries log endpoint :warning: **Deprecated**
* [~~DeleteLogLogentries~~](docs/sdks/logginglogentries/README.md#deleteloglogentries) - Delete a Logentries log endpoint :warning: **Deprecated**
* [~~GetLogLogentries~~](docs/sdks/logginglogentries/README.md#getloglogentries) - Get a Logentries log endpoint :warning: **Deprecated**
* [~~ListLogLogentries~~](docs/sdks/logginglogentries/README.md#listloglogentries) - List Logentries log endpoints :warning: **Deprecated**
* [~~UpdateLogLogentries~~](docs/sdks/logginglogentries/README.md#updateloglogentries) - Update a Logentries log endpoint :warning: **Deprecated**

### [LoggingLoggly](docs/sdks/loggingloggly/README.md)

* [CreateLogLoggly](docs/sdks/loggingloggly/README.md#createlogloggly) - Create a Loggly log endpoint
* [DeleteLogLoggly](docs/sdks/loggingloggly/README.md#deletelogloggly) - Delete a Loggly log endpoint
* [GetLogLoggly](docs/sdks/loggingloggly/README.md#getlogloggly) - Get a Loggly log endpoint
* [ListLogLoggly](docs/sdks/loggingloggly/README.md#listlogloggly) - List Loggly log endpoints
* [UpdateLogLoggly](docs/sdks/loggingloggly/README.md#updatelogloggly) - Update a Loggly log endpoint

### [LoggingLogshuttle](docs/sdks/logginglogshuttle/README.md)

* [CreateLogLogshuttle](docs/sdks/logginglogshuttle/README.md#createloglogshuttle) - Create a Log Shuttle log endpoint
* [DeleteLogLogshuttle](docs/sdks/logginglogshuttle/README.md#deleteloglogshuttle) - Delete a Log Shuttle log endpoint
* [GetLogLogshuttle](docs/sdks/logginglogshuttle/README.md#getloglogshuttle) - Get a Log Shuttle log endpoint
* [ListLogLogshuttle](docs/sdks/logginglogshuttle/README.md#listloglogshuttle) - List Log Shuttle log endpoints
* [UpdateLogLogshuttle](docs/sdks/logginglogshuttle/README.md#updateloglogshuttle) - Update a Log Shuttle log endpoint

### [LoggingNewrelic](docs/sdks/loggingnewrelic/README.md)

* [CreateLogNewrelic](docs/sdks/loggingnewrelic/README.md#createlognewrelic) - Create a New Relic log endpoint
* [DeleteLogNewrelic](docs/sdks/loggingnewrelic/README.md#deletelognewrelic) - Delete a New Relic log endpoint
* [GetLogNewrelic](docs/sdks/loggingnewrelic/README.md#getlognewrelic) - Get a New Relic log endpoint
* [ListLogNewrelic](docs/sdks/loggingnewrelic/README.md#listlognewrelic) - List New Relic log endpoints
* [UpdateLogNewrelic](docs/sdks/loggingnewrelic/README.md#updatelognewrelic) - Update a New Relic log endpoint

### [LoggingOpenstack](docs/sdks/loggingopenstack/README.md)

* [CreateLogOpenstack](docs/sdks/loggingopenstack/README.md#createlogopenstack) - Create an OpenStack log endpoint
* [DeleteLogOpenstack](docs/sdks/loggingopenstack/README.md#deletelogopenstack) - Delete an OpenStack log endpoint
* [GetLogOpenstack](docs/sdks/loggingopenstack/README.md#getlogopenstack) - Get an OpenStack log endpoint
* [ListLogOpenstack](docs/sdks/loggingopenstack/README.md#listlogopenstack) - List OpenStack log endpoints
* [UpdateLogOpenstack](docs/sdks/loggingopenstack/README.md#updatelogopenstack) - Update an OpenStack log endpoint

### [LoggingPapertrail](docs/sdks/loggingpapertrail/README.md)

* [CreateLogPapertrail](docs/sdks/loggingpapertrail/README.md#createlogpapertrail) - Create a Papertrail log endpoint
* [DeleteLogPapertrail](docs/sdks/loggingpapertrail/README.md#deletelogpapertrail) - Delete a Papertrail log endpoint
* [GetLogPapertrail](docs/sdks/loggingpapertrail/README.md#getlogpapertrail) - Get a Papertrail log endpoint
* [ListLogPapertrail](docs/sdks/loggingpapertrail/README.md#listlogpapertrail) - List Papertrail log endpoints
* [UpdateLogPapertrail](docs/sdks/loggingpapertrail/README.md#updatelogpapertrail) - Update a Papertrail log endpoint

### [LoggingPubsub](docs/sdks/loggingpubsub/README.md)

* [CreateLogGcpPubsub](docs/sdks/loggingpubsub/README.md#createloggcppubsub) - Create a GCP Cloud Pub/Sub log endpoint
* [DeleteLogGcpPubsub](docs/sdks/loggingpubsub/README.md#deleteloggcppubsub) - Delete a GCP Cloud Pub/Sub log endpoint
* [GetLogGcpPubsub](docs/sdks/loggingpubsub/README.md#getloggcppubsub) - Get a GCP Cloud Pub/Sub log endpoint
* [ListLogGcpPubsub](docs/sdks/loggingpubsub/README.md#listloggcppubsub) - List GCP Cloud Pub/Sub log endpoints
* [UpdateLogGcpPubsub](docs/sdks/loggingpubsub/README.md#updateloggcppubsub) - Update a GCP Cloud Pub/Sub log endpoint

### [LoggingS3](docs/sdks/loggings3/README.md)

* [CreateLogAwsS3](docs/sdks/loggings3/README.md#createlogawss3) - Create an AWS S3 log endpoint
* [DeleteLogAwsS3](docs/sdks/loggings3/README.md#deletelogawss3) - Delete an AWS S3 log endpoint
* [GetLogAwsS3](docs/sdks/loggings3/README.md#getlogawss3) - Get an AWS S3 log endpoint
* [ListLogAwsS3](docs/sdks/loggings3/README.md#listlogawss3) - List AWS S3 log endpoints
* [UpdateLogAwsS3](docs/sdks/loggings3/README.md#updatelogawss3) - Update an AWS S3 log endpoint

### [LoggingScalyr](docs/sdks/loggingscalyr/README.md)

* [CreateLogScalyr](docs/sdks/loggingscalyr/README.md#createlogscalyr) - Create a Scalyr log endpoint
* [DeleteLogScalyr](docs/sdks/loggingscalyr/README.md#deletelogscalyr) - Delete the Scalyr log endpoint
* [GetLogScalyr](docs/sdks/loggingscalyr/README.md#getlogscalyr) - Get a Scalyr log endpoint
* [ListLogScalyr](docs/sdks/loggingscalyr/README.md#listlogscalyr) - List Scalyr log endpoints
* [UpdateLogScalyr](docs/sdks/loggingscalyr/README.md#updatelogscalyr) - Update the Scalyr log endpoint

### [LoggingSftp](docs/sdks/loggingsftp/README.md)

* [CreateLogSftp](docs/sdks/loggingsftp/README.md#createlogsftp) - Create an SFTP log endpoint
* [DeleteLogSftp](docs/sdks/loggingsftp/README.md#deletelogsftp) - Delete an SFTP log endpoint
* [GetLogSftp](docs/sdks/loggingsftp/README.md#getlogsftp) - Get an SFTP log endpoint
* [ListLogSftp](docs/sdks/loggingsftp/README.md#listlogsftp) - List SFTP log endpoints
* [UpdateLogSftp](docs/sdks/loggingsftp/README.md#updatelogsftp) - Update an SFTP log endpoint

### [LoggingSplunk](docs/sdks/loggingsplunk/README.md)

* [CreateLogSplunk](docs/sdks/loggingsplunk/README.md#createlogsplunk) - Create a Splunk log endpoint
* [DeleteLogSplunk](docs/sdks/loggingsplunk/README.md#deletelogsplunk) - Delete a Splunk log endpoint
* [GetLogSplunk](docs/sdks/loggingsplunk/README.md#getlogsplunk) - Get a Splunk log endpoint
* [ListLogSplunk](docs/sdks/loggingsplunk/README.md#listlogsplunk) - List Splunk log endpoints
* [UpdateLogSplunk](docs/sdks/loggingsplunk/README.md#updatelogsplunk) - Update a Splunk log endpoint

### [LoggingSumologic](docs/sdks/loggingsumologic/README.md)

* [CreateLogSumologic](docs/sdks/loggingsumologic/README.md#createlogsumologic) - Create a Sumologic log endpoint
* [DeleteLogSumologic](docs/sdks/loggingsumologic/README.md#deletelogsumologic) - Delete a Sumologic log endpoint
* [GetLogSumologic](docs/sdks/loggingsumologic/README.md#getlogsumologic) - Get a Sumologic log endpoint
* [ListLogSumologic](docs/sdks/loggingsumologic/README.md#listlogsumologic) - List Sumologic log endpoints
* [UpdateLogSumologic](docs/sdks/loggingsumologic/README.md#updatelogsumologic) - Update a Sumologic log endpoint

### [LoggingSyslog](docs/sdks/loggingsyslog/README.md)

* [CreateLogSyslog](docs/sdks/loggingsyslog/README.md#createlogsyslog) - Create a syslog log endpoint
* [DeleteLogSyslog](docs/sdks/loggingsyslog/README.md#deletelogsyslog) - Delete a syslog log endpoint
* [GetLogSyslog](docs/sdks/loggingsyslog/README.md#getlogsyslog) - Get a syslog log endpoint
* [ListLogSyslog](docs/sdks/loggingsyslog/README.md#listlogsyslog) - List Syslog log endpoints
* [UpdateLogSyslog](docs/sdks/loggingsyslog/README.md#updatelogsyslog) - Update a syslog log endpoint

### [Package](docs/sdks/package/README.md)

* [GetPackage](docs/sdks/package/README.md#getpackage) - Get details of the service's Compute@Edge package.
* [PutPackage](docs/sdks/package/README.md#putpackage) - Upload a Compute@Edge package.

### [Pool](docs/sdks/pool/README.md)

* [CreateServerPool](docs/sdks/pool/README.md#createserverpool) - Create a server pool
* [DeleteServerPool](docs/sdks/pool/README.md#deleteserverpool) - Delete a server pool
* [GetServerPool](docs/sdks/pool/README.md#getserverpool) - Get a server pool
* [ListServerPools](docs/sdks/pool/README.md#listserverpools) - List server pools
* [UpdateServerPool](docs/sdks/pool/README.md#updateserverpool) - Update a server pool

### [RequestSettings](docs/sdks/requestsettings/README.md)

* [DeleteRequestSettings](docs/sdks/requestsettings/README.md#deleterequestsettings) - Delete a Request Settings object
* [GetRequestSettings](docs/sdks/requestsettings/README.md#getrequestsettings) - Get a Request Settings object
* [ListRequestSettings](docs/sdks/requestsettings/README.md#listrequestsettings) - List Request Settings objects
* [UpdateRequestSettings](docs/sdks/requestsettings/README.md#updaterequestsettings) - Update a Request Settings object

### [Resource](docs/sdks/resource/README.md)

* [CreateResource](docs/sdks/resource/README.md#createresource) - Create a resource link
* [DeleteResource](docs/sdks/resource/README.md#deleteresource) - Delete a resource link
* [GetResource](docs/sdks/resource/README.md#getresource) - Display a resource link
* [ListResources](docs/sdks/resource/README.md#listresources) - List resource links
* [UpdateResource](docs/sdks/resource/README.md#updateresource) - Update a resource link

### [ResponseObject](docs/sdks/responseobject/README.md)

* [DeleteResponseObject](docs/sdks/responseobject/README.md#deleteresponseobject) - Delete a Response Object
* [GetResponseObject](docs/sdks/responseobject/README.md#getresponseobject) - Get a Response object
* [ListResponseObjects](docs/sdks/responseobject/README.md#listresponseobjects) - List Response objects

### [Settings](docs/sdks/settings/README.md)

* [GetServiceSettings](docs/sdks/settings/README.md#getservicesettings) - Get service settings
* [UpdateServiceSettings](docs/sdks/settings/README.md#updateservicesettings) - Update service settings

### [Star](docs/sdks/star/README.md)

* [CreateServiceStar](docs/sdks/star/README.md#createservicestar) - Create a star
* [DeleteServiceStar](docs/sdks/star/README.md#deleteservicestar) - Delete a star
* [GetServiceStar](docs/sdks/star/README.md#getservicestar) - Get a star
* [ListServiceStars](docs/sdks/star/README.md#listservicestars) - List stars

### [Historical](docs/sdks/historical/README.md)

* [GetHistStats](docs/sdks/historical/README.md#gethiststats) - Get historical stats
* [GetHistStatsAggregated](docs/sdks/historical/README.md#gethiststatsaggregated) - Get aggregated historical stats
* [GetHistStatsField](docs/sdks/historical/README.md#gethiststatsfield) - Get historical stats for a single field
* [GetHistStatsService](docs/sdks/historical/README.md#gethiststatsservice) - Get historical stats for a single service
* [GetHistStatsServiceField](docs/sdks/historical/README.md#gethiststatsservicefield) - Get historical stats for a single service/field combination
* [GetRegions](docs/sdks/historical/README.md#getregions) - Get region codes
* [GetUsage](docs/sdks/historical/README.md#getusage) - Get usage statistics
* [GetUsageMonth](docs/sdks/historical/README.md#getusagemonth) - Get month-to-date usage statistics
* [GetUsageService](docs/sdks/historical/README.md#getusageservice) - Get usage statistics per service

### [TLSActivations](docs/sdks/tlsactivations/README.md)

* [CreateTLSActivation](docs/sdks/tlsactivations/README.md#createtlsactivation) - Enable TLS for a domain using a custom certificate
* [DeleteTLSActivation](docs/sdks/tlsactivations/README.md#deletetlsactivation) - Disable TLS on a domain
* [GetTLSActivation](docs/sdks/tlsactivations/README.md#gettlsactivation) - Get a TLS activation
* [ListTLSActivations](docs/sdks/tlsactivations/README.md#listtlsactivations) - List TLS activations
* [UpdateTLSActivation](docs/sdks/tlsactivations/README.md#updatetlsactivation) - Update a certificate

### [TLSBulkCertificates](docs/sdks/tlsbulkcertificates/README.md)

* [DeleteBulkTLSCert](docs/sdks/tlsbulkcertificates/README.md#deletebulktlscert) - Delete a certificate
* [GetTLSBulkCert](docs/sdks/tlsbulkcertificates/README.md#gettlsbulkcert) - Get a certificate
* [ListTLSBulkCerts](docs/sdks/tlsbulkcertificates/README.md#listtlsbulkcerts) - List certificates
* [UpdateBulkTLSCert](docs/sdks/tlsbulkcertificates/README.md#updatebulktlscert) - Update a certificate
* [UploadTLSBulkCert](docs/sdks/tlsbulkcertificates/README.md#uploadtlsbulkcert) - Upload a certificate

### [TLSCertificates](docs/sdks/tlscertificates/README.md)

* [CreateTLSCert](docs/sdks/tlscertificates/README.md#createtlscert) - Create a TLS certificate
* [DeleteTLSCert](docs/sdks/tlscertificates/README.md#deletetlscert) - Delete a TLS certificate
* [GetTLSCert](docs/sdks/tlscertificates/README.md#gettlscert) - Get a TLS certificate
* [ListTLSCerts](docs/sdks/tlscertificates/README.md#listtlscerts) - List TLS certificates
* [UpdateTLSCert](docs/sdks/tlscertificates/README.md#updatetlscert) - Update a TLS certificate

### [TLSConfigurations](docs/sdks/tlsconfigurations/README.md)

* [GetTLSConfig](docs/sdks/tlsconfigurations/README.md#gettlsconfig) - Get a TLS configuration
* [ListTLSConfigs](docs/sdks/tlsconfigurations/README.md#listtlsconfigs) - List TLS configurations
* [UpdateTLSConfig](docs/sdks/tlsconfigurations/README.md#updatetlsconfig) - Update a TLS configuration

### [TLSDomains](docs/sdks/tlsdomains/README.md)

* [ListTLSDomains](docs/sdks/tlsdomains/README.md#listtlsdomains) - List TLS domains

### [MutualAuthentication](docs/sdks/mutualauthentication/README.md)

* [CreateMutualTLSAuthentication](docs/sdks/mutualauthentication/README.md#createmutualtlsauthentication) - Create a Mutual Authentication
* [DeleteMutualTLS](docs/sdks/mutualauthentication/README.md#deletemutualtls) - Delete a Mutual TLS
* [GetMutualAuthentication](docs/sdks/mutualauthentication/README.md#getmutualauthentication) - Get a Mutual Authentication
* [ListMutualAuthentications](docs/sdks/mutualauthentication/README.md#listmutualauthentications) - List Mutual Authentications
* [PatchMutualAuthentication](docs/sdks/mutualauthentication/README.md#patchmutualauthentication) - Update a Mutual Authentication

### [TLSPrivateKeys](docs/sdks/tlsprivatekeys/README.md)

* [CreateTLSKey](docs/sdks/tlsprivatekeys/README.md#createtlskey) - Create a TLS private key
* [DeleteTLSKey](docs/sdks/tlsprivatekeys/README.md#deletetlskey) - Delete a TLS private key
* [GetTLSKey](docs/sdks/tlsprivatekeys/README.md#gettlskey) - Get a TLS private key
* [ListTLSKeys](docs/sdks/tlsprivatekeys/README.md#listtlskeys) - List TLS private keys

### [TLSSubscriptions](docs/sdks/tlssubscriptions/README.md)

* [CreateGlobalsignEmailChallenge](docs/sdks/tlssubscriptions/README.md#createglobalsignemailchallenge) - Creates a GlobalSign email challenge.
* [CreateTLSSub](docs/sdks/tlssubscriptions/README.md#createtlssub) - Create a TLS subscription
* [DeleteGlobalsignEmailChallenge](docs/sdks/tlssubscriptions/README.md#deleteglobalsignemailchallenge) - Delete a GlobalSign email challenge
* [DeleteTLSSub](docs/sdks/tlssubscriptions/README.md#deletetlssub) - Delete a TLS subscription
* [GetTLSSub](docs/sdks/tlssubscriptions/README.md#gettlssub) - Get a TLS subscription
* [ListTLSSubs](docs/sdks/tlssubscriptions/README.md#listtlssubs) - List TLS subscriptions
* [PatchTLSSub](docs/sdks/tlssubscriptions/README.md#patchtlssub) - Update a TLS subscription

### [IamUserGroups](docs/sdks/iamusergroups/README.md)

* [DeleteAUserGroup](docs/sdks/iamusergroups/README.md#deleteausergroup) - Delete a user group
* [GetAUserGroup](docs/sdks/iamusergroups/README.md#getausergroup) - Get a user group
* [ListUserGroupMembers](docs/sdks/iamusergroups/README.md#listusergroupmembers) - List members of a user group
* [ListUserGroupRoles](docs/sdks/iamusergroups/README.md#listusergrouproles) - List roles in a user group
* [ListUserGroupServiceGroups](docs/sdks/iamusergroups/README.md#listusergroupservicegroups) - List service groups in a user group
* [ListUserGroups](docs/sdks/iamusergroups/README.md#listusergroups) - List user groups

### [Realtime](docs/sdks/realtime/README.md)

* [GetStatsLast120Seconds](docs/sdks/realtime/README.md#getstatslast120seconds) - Get real-time data for the last 120 seconds
* [GetStatsLast120SecondsLimitEntries](docs/sdks/realtime/README.md#getstatslast120secondslimitentries) - Get a limited number of real-time data entries
* [GetStatsLastSecond](docs/sdks/realtime/README.md#getstatslastsecond) - Get real-time data from specified time

### [WafFirewalls](docs/sdks/waffirewalls/README.md)

* [~~CreateWafFirewall~~](docs/sdks/waffirewalls/README.md#createwaffirewall) - Create a firewall :warning: **Deprecated**
* [~~DeleteWafFirewall~~](docs/sdks/waffirewalls/README.md#deletewaffirewall) - Delete a firewall :warning: **Deprecated**
* [~~GetWafFirewall~~](docs/sdks/waffirewalls/README.md#getwaffirewall) - Get a firewall :warning: **Deprecated**
* [~~ListWafFirewalls~~](docs/sdks/waffirewalls/README.md#listwaffirewalls) - List firewalls :warning: **Deprecated**
* [~~UpdateWafFirewall~~](docs/sdks/waffirewalls/README.md#updatewaffirewall) - Update a firewall :warning: **Deprecated**

### [WafFirewallVersions](docs/sdks/waffirewallversions/README.md)

* [~~CloneWafFirewallVersion~~](docs/sdks/waffirewallversions/README.md#clonewaffirewallversion) - Clone a firewall version :warning: **Deprecated**
* [~~CreateWafFirewallVersion~~](docs/sdks/waffirewallversions/README.md#createwaffirewallversion) - Create a firewall version :warning: **Deprecated**
* [~~DeployActivateWafFirewallVersion~~](docs/sdks/waffirewallversions/README.md#deployactivatewaffirewallversion) - Deploy or activate a firewall version :warning: **Deprecated**
* [~~GetWafFirewallVersion~~](docs/sdks/waffirewallversions/README.md#getwaffirewallversion) - Get a firewall version :warning: **Deprecated**
* [~~ListWafFirewallVersions~~](docs/sdks/waffirewallversions/README.md#listwaffirewallversions) - List firewall versions :warning: **Deprecated**
* [~~UpdateWafFirewallVersion~~](docs/sdks/waffirewallversions/README.md#updatewaffirewallversion) - Update a firewall version :warning: **Deprecated**

### [WafExclusions](docs/sdks/wafexclusions/README.md)

* [~~CreateWafRuleExclusion~~](docs/sdks/wafexclusions/README.md#createwafruleexclusion) - Create a WAF rule exclusion :warning: **Deprecated**
* [~~DeleteWafRuleExclusion~~](docs/sdks/wafexclusions/README.md#deletewafruleexclusion) - Delete a WAF rule exclusion :warning: **Deprecated**
* [~~GetWafRuleExclusion~~](docs/sdks/wafexclusions/README.md#getwafruleexclusion) - Get a WAF rule exclusion :warning: **Deprecated**
* [~~ListWafRuleExclusions~~](docs/sdks/wafexclusions/README.md#listwafruleexclusions) - List WAF rule exclusions :warning: **Deprecated**
* [~~UpdateWafRuleExclusion~~](docs/sdks/wafexclusions/README.md#updatewafruleexclusion) - Update a WAF rule exclusion :warning: **Deprecated**

### [WafActiveRules](docs/sdks/wafactiverules/README.md)

* [~~BulkUpdateWafActiveRules~~](docs/sdks/wafactiverules/README.md#bulkupdatewafactiverules) - Update multiple active rules :warning: **Deprecated**
* [~~CreateWafActiveRule~~](docs/sdks/wafactiverules/README.md#createwafactiverule) - Add a rule to a WAF as an active rule :warning: **Deprecated**
* [~~CreateWafActiveRulesTag~~](docs/sdks/wafactiverules/README.md#createwafactiverulestag) - Create active rules by tag :warning: **Deprecated**
* [~~DeleteWafActiveRule~~](docs/sdks/wafactiverules/README.md#deletewafactiverule) - Delete an active rule :warning: **Deprecated**
* [~~GetWafActiveRule~~](docs/sdks/wafactiverules/README.md#getwafactiverule) - Get an active WAF rule object :warning: **Deprecated**
* [~~ListWafActiveRules~~](docs/sdks/wafactiverules/README.md#listwafactiverules) - List active rules on a WAF :warning: **Deprecated**
* [~~UpdateWafActiveRule~~](docs/sdks/wafactiverules/README.md#updatewafactiverule) - Update an active rule :warning: **Deprecated**

### [WafRules](docs/sdks/wafrules/README.md)

* [~~GetWafRule~~](docs/sdks/wafrules/README.md#getwafrule) - Get a rule :warning: **Deprecated**
* [~~ListWafRules~~](docs/sdks/wafrules/README.md#listwafrules) - List available WAF rules :warning: **Deprecated**

### [WafRuleRevisions](docs/sdks/wafrulerevisions/README.md)

* [~~GetWafRuleRevision~~](docs/sdks/wafrulerevisions/README.md#getwafrulerevision) - Get a revision of a rule :warning: **Deprecated**
* [~~ListWafRuleRevisions~~](docs/sdks/wafrulerevisions/README.md#listwafrulerevisions) - List revisions for a rule :warning: **Deprecated**

### [WafTags](docs/sdks/waftags/README.md)

* [~~ListWafTags~~](docs/sdks/waftags/README.md#listwaftags) - List tags :warning: **Deprecated**
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->



<!-- End Dev Containers -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                                      | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| sdkerrors.BillingAddressVerificationErrorResponse | 400                                               | application/vnd.api+json                          |
| sdkerrors.SDKError                                | 400-600                                           | */*                                               |


## Example

```go
package main

import (
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
	"context"
	"log"
)

func main() {
	s := fastly.New(
		fastly.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.BillingAddress.AddBillingAddr(ctx, operations.AddBillingAddrRequest{
		BillingAddressRequest: &components.BillingAddressRequest{
			Data: &components.BillingAddressRequestData{
				Attributes: &components.BillingAddressAttributesInput{
					Address1:   fastly.String("80719 Dorothea Mountain"),
					Address2:   fastly.String("Apt. 652"),
					City:       fastly.String("New Rasheedville"),
					Country:    fastly.String("US"),
					Locality:   fastly.String("New Castle"),
					PostalCode: fastly.String("53538-5902"),
					State:      fastly.String("DE"),
				},
			},
		},
		CustomerID: "x4xCwxxJxGCx123Rx5xTx",
	})
	if err != nil {

		var e *sdkerrors.BillingAddressVerificationErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.fastly.com` | None |
| 1 | `https://rt.fastly.com` | None |

For example:

```go
package main

import (
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
	"context"
	"log"
)

func main() {
	s := fastly.New(
		fastly.WithServerIndex(1),
		fastly.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.ApexRedirect.DeleteApexRedirect(ctx, operations.DeleteApexRedirectRequest{
		ApexRedirectID: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Object != nil {
		// handle response
	}
}

```


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:

```go
package main

import (
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
	"context"
	"log"
)

func main() {
	s := fastly.New(
		fastly.WithServerURL("https://api.fastly.com"),
		fastly.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.ApexRedirect.DeleteApexRedirect(ctx, operations.DeleteApexRedirectRequest{
		ApexRedirectID: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->



<!-- Start Authentication -->

# Authentication

## Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name    | Type    | Scheme  |
| ------- | ------- | ------- |
| `Token` | apiKey  | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:

```go
package main

import (
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
	"context"
	"log"
)

func main() {
	s := fastly.New(
		fastly.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.ApexRedirect.DeleteApexRedirect(ctx, operations.DeleteApexRedirectRequest{
		ApexRedirectID: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Object != nil {
		// handle response
	}
}

```

## Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:

```go
package main

import (
	fastly "Fastly"
	"Fastly/models/components"
	"Fastly/models/operations"
	"context"
	"log"
)

func main() {
	s := fastly.New()

	operationSecurity := operations.UpdateUserPasswordSecurity{
		Password: "",
		Username: "",
	}

	ctx := context.Background()
	res, err := s.User.UpdateUserPassword(ctx, &components.PasswordChange{}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.UserResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
