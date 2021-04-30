// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

const (

	// ErrCodeAccessDenied for service response error code
	// "AccessDenied".
	//
	// Access denied.
	ErrCodeAccessDenied = "AccessDenied"

	// ErrCodeBatchTooLarge for service response error code
	// "BatchTooLarge".
	//
	// Invalidation batch specified is too large.
	ErrCodeBatchTooLarge = "BatchTooLarge"

	// ErrCodeCNAMEAlreadyExists for service response error code
	// "CNAMEAlreadyExists".
	//
	// The CNAME specified is already defined for CloudFront.
	ErrCodeCNAMEAlreadyExists = "CNAMEAlreadyExists"

	// ErrCodeCachePolicyAlreadyExists for service response error code
	// "CachePolicyAlreadyExists".
	//
	// A cache policy with this name already exists. You must provide a unique name.
	// To modify an existing cache policy, use UpdateCachePolicy.
	ErrCodeCachePolicyAlreadyExists = "CachePolicyAlreadyExists"

	// ErrCodeCachePolicyInUse for service response error code
	// "CachePolicyInUse".
	//
	// Cannot delete the cache policy because it is attached to one or more cache
	// behaviors.
	ErrCodeCachePolicyInUse = "CachePolicyInUse"

	// ErrCodeCannotChangeImmutablePublicKeyFields for service response error code
	// "CannotChangeImmutablePublicKeyFields".
	//
	// You can't change the value of a public key.
	ErrCodeCannotChangeImmutablePublicKeyFields = "CannotChangeImmutablePublicKeyFields"

	// ErrCodeDistributionAlreadyExists for service response error code
	// "DistributionAlreadyExists".
	//
	// The caller reference you attempted to create the distribution with is associated
	// with another distribution.
	ErrCodeDistributionAlreadyExists = "DistributionAlreadyExists"

	// ErrCodeDistributionNotDisabled for service response error code
	// "DistributionNotDisabled".
	//
	// The specified CloudFront distribution is not disabled. You must disable the
	// distribution before you can delete it.
	ErrCodeDistributionNotDisabled = "DistributionNotDisabled"

	// ErrCodeFieldLevelEncryptionConfigAlreadyExists for service response error code
	// "FieldLevelEncryptionConfigAlreadyExists".
	//
	// The specified configuration for field-level encryption already exists.
	ErrCodeFieldLevelEncryptionConfigAlreadyExists = "FieldLevelEncryptionConfigAlreadyExists"

	// ErrCodeFieldLevelEncryptionConfigInUse for service response error code
	// "FieldLevelEncryptionConfigInUse".
	//
	// The specified configuration for field-level encryption is in use.
	ErrCodeFieldLevelEncryptionConfigInUse = "FieldLevelEncryptionConfigInUse"

	// ErrCodeFieldLevelEncryptionProfileAlreadyExists for service response error code
	// "FieldLevelEncryptionProfileAlreadyExists".
	//
	// The specified profile for field-level encryption already exists.
	ErrCodeFieldLevelEncryptionProfileAlreadyExists = "FieldLevelEncryptionProfileAlreadyExists"

	// ErrCodeFieldLevelEncryptionProfileInUse for service response error code
	// "FieldLevelEncryptionProfileInUse".
	//
	// The specified profile for field-level encryption is in use.
	ErrCodeFieldLevelEncryptionProfileInUse = "FieldLevelEncryptionProfileInUse"

	// ErrCodeFieldLevelEncryptionProfileSizeExceeded for service response error code
	// "FieldLevelEncryptionProfileSizeExceeded".
	//
	// The maximum size of a profile for field-level encryption was exceeded.
	ErrCodeFieldLevelEncryptionProfileSizeExceeded = "FieldLevelEncryptionProfileSizeExceeded"

	// ErrCodeFunctionAlreadyExists for service response error code
	// "FunctionAlreadyExists".
	//
	// A function with the same name already exists in this AWS account. To create
	// a function, you must provide a unique name. To update an existing function,
	// use UpdateFunction.
	ErrCodeFunctionAlreadyExists = "FunctionAlreadyExists"

	// ErrCodeFunctionInUse for service response error code
	// "FunctionInUse".
	//
	// Cannot delete the function because it’s attached to one or more cache behaviors.
	ErrCodeFunctionInUse = "FunctionInUse"

	// ErrCodeFunctionSizeLimitExceeded for service response error code
	// "FunctionSizeLimitExceeded".
	//
	// The function is too large. For more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeFunctionSizeLimitExceeded = "FunctionSizeLimitExceeded"

	// ErrCodeIllegalDelete for service response error code
	// "IllegalDelete".
	//
	// You cannot delete a managed policy.
	ErrCodeIllegalDelete = "IllegalDelete"

	// ErrCodeIllegalFieldLevelEncryptionConfigAssociationWithCacheBehavior for service response error code
	// "IllegalFieldLevelEncryptionConfigAssociationWithCacheBehavior".
	//
	// The specified configuration for field-level encryption can't be associated
	// with the specified cache behavior.
	ErrCodeIllegalFieldLevelEncryptionConfigAssociationWithCacheBehavior = "IllegalFieldLevelEncryptionConfigAssociationWithCacheBehavior"

	// ErrCodeIllegalUpdate for service response error code
	// "IllegalUpdate".
	//
	// The update contains modifications that are not allowed.
	ErrCodeIllegalUpdate = "IllegalUpdate"

	// ErrCodeInconsistentQuantities for service response error code
	// "InconsistentQuantities".
	//
	// The value of Quantity and the size of Items don't match.
	ErrCodeInconsistentQuantities = "InconsistentQuantities"

	// ErrCodeInvalidArgument for service response error code
	// "InvalidArgument".
	//
	// An argument is invalid.
	ErrCodeInvalidArgument = "InvalidArgument"

	// ErrCodeInvalidDefaultRootObject for service response error code
	// "InvalidDefaultRootObject".
	//
	// The default root object file name is too big or contains an invalid character.
	ErrCodeInvalidDefaultRootObject = "InvalidDefaultRootObject"

	// ErrCodeInvalidErrorCode for service response error code
	// "InvalidErrorCode".
	//
	// An invalid error code was specified.
	ErrCodeInvalidErrorCode = "InvalidErrorCode"

	// ErrCodeInvalidForwardCookies for service response error code
	// "InvalidForwardCookies".
	//
	// Your request contains forward cookies option which doesn't match with the
	// expectation for the whitelisted list of cookie names. Either list of cookie
	// names has been specified when not allowed or list of cookie names is missing
	// when expected.
	ErrCodeInvalidForwardCookies = "InvalidForwardCookies"

	// ErrCodeInvalidFunctionAssociation for service response error code
	// "InvalidFunctionAssociation".
	//
	// A CloudFront function association is invalid.
	ErrCodeInvalidFunctionAssociation = "InvalidFunctionAssociation"

	// ErrCodeInvalidGeoRestrictionParameter for service response error code
	// "InvalidGeoRestrictionParameter".
	//
	// The specified geo restriction parameter is not valid.
	ErrCodeInvalidGeoRestrictionParameter = "InvalidGeoRestrictionParameter"

	// ErrCodeInvalidHeadersForS3Origin for service response error code
	// "InvalidHeadersForS3Origin".
	//
	// The headers specified are not valid for an Amazon S3 origin.
	ErrCodeInvalidHeadersForS3Origin = "InvalidHeadersForS3Origin"

	// ErrCodeInvalidIfMatchVersion for service response error code
	// "InvalidIfMatchVersion".
	//
	// The If-Match version is missing or not valid.
	ErrCodeInvalidIfMatchVersion = "InvalidIfMatchVersion"

	// ErrCodeInvalidLambdaFunctionAssociation for service response error code
	// "InvalidLambdaFunctionAssociation".
	//
	// The specified Lambda function association is invalid.
	ErrCodeInvalidLambdaFunctionAssociation = "InvalidLambdaFunctionAssociation"

	// ErrCodeInvalidLocationCode for service response error code
	// "InvalidLocationCode".
	//
	// The location code specified is not valid.
	ErrCodeInvalidLocationCode = "InvalidLocationCode"

	// ErrCodeInvalidMinimumProtocolVersion for service response error code
	// "InvalidMinimumProtocolVersion".
	//
	// The minimum protocol version specified is not valid.
	ErrCodeInvalidMinimumProtocolVersion = "InvalidMinimumProtocolVersion"

	// ErrCodeInvalidOrigin for service response error code
	// "InvalidOrigin".
	//
	// The Amazon S3 origin server specified does not refer to a valid Amazon S3
	// bucket.
	ErrCodeInvalidOrigin = "InvalidOrigin"

	// ErrCodeInvalidOriginAccessIdentity for service response error code
	// "InvalidOriginAccessIdentity".
	//
	// The origin access identity is not valid or doesn't exist.
	ErrCodeInvalidOriginAccessIdentity = "InvalidOriginAccessIdentity"

	// ErrCodeInvalidOriginKeepaliveTimeout for service response error code
	// "InvalidOriginKeepaliveTimeout".
	//
	// The keep alive timeout specified for the origin is not valid.
	ErrCodeInvalidOriginKeepaliveTimeout = "InvalidOriginKeepaliveTimeout"

	// ErrCodeInvalidOriginReadTimeout for service response error code
	// "InvalidOriginReadTimeout".
	//
	// The read timeout specified for the origin is not valid.
	ErrCodeInvalidOriginReadTimeout = "InvalidOriginReadTimeout"

	// ErrCodeInvalidProtocolSettings for service response error code
	// "InvalidProtocolSettings".
	//
	// You cannot specify SSLv3 as the minimum protocol version if you only want
	// to support only clients that support Server Name Indication (SNI).
	ErrCodeInvalidProtocolSettings = "InvalidProtocolSettings"

	// ErrCodeInvalidQueryStringParameters for service response error code
	// "InvalidQueryStringParameters".
	//
	// The query string parameters specified are not valid.
	ErrCodeInvalidQueryStringParameters = "InvalidQueryStringParameters"

	// ErrCodeInvalidRelativePath for service response error code
	// "InvalidRelativePath".
	//
	// The relative path is too big, is not URL-encoded, or does not begin with
	// a slash (/).
	ErrCodeInvalidRelativePath = "InvalidRelativePath"

	// ErrCodeInvalidRequiredProtocol for service response error code
	// "InvalidRequiredProtocol".
	//
	// This operation requires the HTTPS protocol. Ensure that you specify the HTTPS
	// protocol in your request, or omit the RequiredProtocols element from your
	// distribution configuration.
	ErrCodeInvalidRequiredProtocol = "InvalidRequiredProtocol"

	// ErrCodeInvalidResponseCode for service response error code
	// "InvalidResponseCode".
	//
	// A response code is not valid.
	ErrCodeInvalidResponseCode = "InvalidResponseCode"

	// ErrCodeInvalidTTLOrder for service response error code
	// "InvalidTTLOrder".
	//
	// The TTL order specified is not valid.
	ErrCodeInvalidTTLOrder = "InvalidTTLOrder"

	// ErrCodeInvalidTagging for service response error code
	// "InvalidTagging".
	//
	// The tagging specified is not valid.
	ErrCodeInvalidTagging = "InvalidTagging"

	// ErrCodeInvalidViewerCertificate for service response error code
	// "InvalidViewerCertificate".
	//
	// A viewer certificate specified is not valid.
	ErrCodeInvalidViewerCertificate = "InvalidViewerCertificate"

	// ErrCodeInvalidWebACLId for service response error code
	// "InvalidWebACLId".
	//
	// A web ACL ID specified is not valid. To specify a web ACL created using the
	// latest version of AWS WAF, use the ACL ARN, for example arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a.
	// To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example
	// 473e64fd-f30b-4765-81a0-62ad96dd167a.
	ErrCodeInvalidWebACLId = "InvalidWebACLId"

	// ErrCodeKeyGroupAlreadyExists for service response error code
	// "KeyGroupAlreadyExists".
	//
	// A key group with this name already exists. You must provide a unique name.
	// To modify an existing key group, use UpdateKeyGroup.
	ErrCodeKeyGroupAlreadyExists = "KeyGroupAlreadyExists"

	// ErrCodeMissingBody for service response error code
	// "MissingBody".
	//
	// This operation requires a body. Ensure that the body is present and the Content-Type
	// header is set.
	ErrCodeMissingBody = "MissingBody"

	// ErrCodeNoSuchCachePolicy for service response error code
	// "NoSuchCachePolicy".
	//
	// The cache policy does not exist.
	ErrCodeNoSuchCachePolicy = "NoSuchCachePolicy"

	// ErrCodeNoSuchCloudFrontOriginAccessIdentity for service response error code
	// "NoSuchCloudFrontOriginAccessIdentity".
	//
	// The specified origin access identity does not exist.
	ErrCodeNoSuchCloudFrontOriginAccessIdentity = "NoSuchCloudFrontOriginAccessIdentity"

	// ErrCodeNoSuchDistribution for service response error code
	// "NoSuchDistribution".
	//
	// The specified distribution does not exist.
	ErrCodeNoSuchDistribution = "NoSuchDistribution"

	// ErrCodeNoSuchFieldLevelEncryptionConfig for service response error code
	// "NoSuchFieldLevelEncryptionConfig".
	//
	// The specified configuration for field-level encryption doesn't exist.
	ErrCodeNoSuchFieldLevelEncryptionConfig = "NoSuchFieldLevelEncryptionConfig"

	// ErrCodeNoSuchFieldLevelEncryptionProfile for service response error code
	// "NoSuchFieldLevelEncryptionProfile".
	//
	// The specified profile for field-level encryption doesn't exist.
	ErrCodeNoSuchFieldLevelEncryptionProfile = "NoSuchFieldLevelEncryptionProfile"

	// ErrCodeNoSuchFunctionExists for service response error code
	// "NoSuchFunctionExists".
	//
	// The function does not exist.
	ErrCodeNoSuchFunctionExists = "NoSuchFunctionExists"

	// ErrCodeNoSuchInvalidation for service response error code
	// "NoSuchInvalidation".
	//
	// The specified invalidation does not exist.
	ErrCodeNoSuchInvalidation = "NoSuchInvalidation"

	// ErrCodeNoSuchOrigin for service response error code
	// "NoSuchOrigin".
	//
	// No origin exists with the specified Origin Id.
	ErrCodeNoSuchOrigin = "NoSuchOrigin"

	// ErrCodeNoSuchOriginRequestPolicy for service response error code
	// "NoSuchOriginRequestPolicy".
	//
	// The origin request policy does not exist.
	ErrCodeNoSuchOriginRequestPolicy = "NoSuchOriginRequestPolicy"

	// ErrCodeNoSuchPublicKey for service response error code
	// "NoSuchPublicKey".
	//
	// The specified public key doesn't exist.
	ErrCodeNoSuchPublicKey = "NoSuchPublicKey"

	// ErrCodeNoSuchRealtimeLogConfig for service response error code
	// "NoSuchRealtimeLogConfig".
	//
	// The real-time log configuration does not exist.
	ErrCodeNoSuchRealtimeLogConfig = "NoSuchRealtimeLogConfig"

	// ErrCodeNoSuchResource for service response error code
	// "NoSuchResource".
	//
	// A resource that was specified is not valid.
	ErrCodeNoSuchResource = "NoSuchResource"

	// ErrCodeNoSuchStreamingDistribution for service response error code
	// "NoSuchStreamingDistribution".
	//
	// The specified streaming distribution does not exist.
	ErrCodeNoSuchStreamingDistribution = "NoSuchStreamingDistribution"

	// ErrCodeOriginAccessIdentityAlreadyExists for service response error code
	// "CloudFrontOriginAccessIdentityAlreadyExists".
	//
	// If the CallerReference is a value you already sent in a previous request
	// to create an identity but the content of the CloudFrontOriginAccessIdentityConfig
	// is different from the original request, CloudFront returns a CloudFrontOriginAccessIdentityAlreadyExists
	// error.
	ErrCodeOriginAccessIdentityAlreadyExists = "CloudFrontOriginAccessIdentityAlreadyExists"

	// ErrCodeOriginAccessIdentityInUse for service response error code
	// "CloudFrontOriginAccessIdentityInUse".
	//
	// The Origin Access Identity specified is already in use.
	ErrCodeOriginAccessIdentityInUse = "CloudFrontOriginAccessIdentityInUse"

	// ErrCodeOriginRequestPolicyAlreadyExists for service response error code
	// "OriginRequestPolicyAlreadyExists".
	//
	// An origin request policy with this name already exists. You must provide
	// a unique name. To modify an existing origin request policy, use UpdateOriginRequestPolicy.
	ErrCodeOriginRequestPolicyAlreadyExists = "OriginRequestPolicyAlreadyExists"

	// ErrCodeOriginRequestPolicyInUse for service response error code
	// "OriginRequestPolicyInUse".
	//
	// Cannot delete the origin request policy because it is attached to one or
	// more cache behaviors.
	ErrCodeOriginRequestPolicyInUse = "OriginRequestPolicyInUse"

	// ErrCodePreconditionFailed for service response error code
	// "PreconditionFailed".
	//
	// The precondition in one or more of the request fields evaluated to false.
	ErrCodePreconditionFailed = "PreconditionFailed"

	// ErrCodePublicKeyAlreadyExists for service response error code
	// "PublicKeyAlreadyExists".
	//
	// The specified public key already exists.
	ErrCodePublicKeyAlreadyExists = "PublicKeyAlreadyExists"

	// ErrCodePublicKeyInUse for service response error code
	// "PublicKeyInUse".
	//
	// The specified public key is in use.
	ErrCodePublicKeyInUse = "PublicKeyInUse"

	// ErrCodeQueryArgProfileEmpty for service response error code
	// "QueryArgProfileEmpty".
	//
	// No profile specified for the field-level encryption query argument.
	ErrCodeQueryArgProfileEmpty = "QueryArgProfileEmpty"

	// ErrCodeRealtimeLogConfigAlreadyExists for service response error code
	// "RealtimeLogConfigAlreadyExists".
	//
	// A real-time log configuration with this name already exists. You must provide
	// a unique name. To modify an existing real-time log configuration, use UpdateRealtimeLogConfig.
	ErrCodeRealtimeLogConfigAlreadyExists = "RealtimeLogConfigAlreadyExists"

	// ErrCodeRealtimeLogConfigInUse for service response error code
	// "RealtimeLogConfigInUse".
	//
	// Cannot delete the real-time log configuration because it is attached to one
	// or more cache behaviors.
	ErrCodeRealtimeLogConfigInUse = "RealtimeLogConfigInUse"

	// ErrCodeRealtimeLogConfigOwnerMismatch for service response error code
	// "RealtimeLogConfigOwnerMismatch".
	//
	// The specified real-time log configuration belongs to a different AWS account.
	ErrCodeRealtimeLogConfigOwnerMismatch = "RealtimeLogConfigOwnerMismatch"

	// ErrCodeResourceInUse for service response error code
	// "ResourceInUse".
	//
	// Cannot delete this resource because it is in use.
	ErrCodeResourceInUse = "ResourceInUse"

	// ErrCodeStreamingDistributionAlreadyExists for service response error code
	// "StreamingDistributionAlreadyExists".
	//
	// The caller reference you attempted to create the streaming distribution with
	// is associated with another distribution
	ErrCodeStreamingDistributionAlreadyExists = "StreamingDistributionAlreadyExists"

	// ErrCodeStreamingDistributionNotDisabled for service response error code
	// "StreamingDistributionNotDisabled".
	//
	// The specified CloudFront distribution is not disabled. You must disable the
	// distribution before you can delete it.
	ErrCodeStreamingDistributionNotDisabled = "StreamingDistributionNotDisabled"

	// ErrCodeTestFunctionFailed for service response error code
	// "TestFunctionFailed".
	//
	// The CloudFront function failed.
	ErrCodeTestFunctionFailed = "TestFunctionFailed"

	// ErrCodeTooManyCacheBehaviors for service response error code
	// "TooManyCacheBehaviors".
	//
	// You cannot create more cache behaviors for the distribution.
	ErrCodeTooManyCacheBehaviors = "TooManyCacheBehaviors"

	// ErrCodeTooManyCachePolicies for service response error code
	// "TooManyCachePolicies".
	//
	// You have reached the maximum number of cache policies for this AWS account.
	// For more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyCachePolicies = "TooManyCachePolicies"

	// ErrCodeTooManyCertificates for service response error code
	// "TooManyCertificates".
	//
	// You cannot create anymore custom SSL/TLS certificates.
	ErrCodeTooManyCertificates = "TooManyCertificates"

	// ErrCodeTooManyCloudFrontOriginAccessIdentities for service response error code
	// "TooManyCloudFrontOriginAccessIdentities".
	//
	// Processing your request would cause you to exceed the maximum number of origin
	// access identities allowed.
	ErrCodeTooManyCloudFrontOriginAccessIdentities = "TooManyCloudFrontOriginAccessIdentities"

	// ErrCodeTooManyCookieNamesInWhiteList for service response error code
	// "TooManyCookieNamesInWhiteList".
	//
	// Your request contains more cookie names in the whitelist than are allowed
	// per cache behavior.
	ErrCodeTooManyCookieNamesInWhiteList = "TooManyCookieNamesInWhiteList"

	// ErrCodeTooManyCookiesInCachePolicy for service response error code
	// "TooManyCookiesInCachePolicy".
	//
	// The number of cookies in the cache policy exceeds the maximum. For more information,
	// see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyCookiesInCachePolicy = "TooManyCookiesInCachePolicy"

	// ErrCodeTooManyCookiesInOriginRequestPolicy for service response error code
	// "TooManyCookiesInOriginRequestPolicy".
	//
	// The number of cookies in the origin request policy exceeds the maximum. For
	// more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyCookiesInOriginRequestPolicy = "TooManyCookiesInOriginRequestPolicy"

	// ErrCodeTooManyDistributionCNAMEs for service response error code
	// "TooManyDistributionCNAMEs".
	//
	// Your request contains more CNAMEs than are allowed per distribution.
	ErrCodeTooManyDistributionCNAMEs = "TooManyDistributionCNAMEs"

	// ErrCodeTooManyDistributions for service response error code
	// "TooManyDistributions".
	//
	// Processing your request would cause you to exceed the maximum number of distributions
	// allowed.
	ErrCodeTooManyDistributions = "TooManyDistributions"

	// ErrCodeTooManyDistributionsAssociatedToCachePolicy for service response error code
	// "TooManyDistributionsAssociatedToCachePolicy".
	//
	// The maximum number of distributions have been associated with the specified
	// cache policy. For more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyDistributionsAssociatedToCachePolicy = "TooManyDistributionsAssociatedToCachePolicy"

	// ErrCodeTooManyDistributionsAssociatedToFieldLevelEncryptionConfig for service response error code
	// "TooManyDistributionsAssociatedToFieldLevelEncryptionConfig".
	//
	// The maximum number of distributions have been associated with the specified
	// configuration for field-level encryption.
	ErrCodeTooManyDistributionsAssociatedToFieldLevelEncryptionConfig = "TooManyDistributionsAssociatedToFieldLevelEncryptionConfig"

	// ErrCodeTooManyDistributionsAssociatedToKeyGroup for service response error code
	// "TooManyDistributionsAssociatedToKeyGroup".
	//
	// The number of distributions that reference this key group is more than the
	// maximum allowed. For more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyDistributionsAssociatedToKeyGroup = "TooManyDistributionsAssociatedToKeyGroup"

	// ErrCodeTooManyDistributionsAssociatedToOriginRequestPolicy for service response error code
	// "TooManyDistributionsAssociatedToOriginRequestPolicy".
	//
	// The maximum number of distributions have been associated with the specified
	// origin request policy. For more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyDistributionsAssociatedToOriginRequestPolicy = "TooManyDistributionsAssociatedToOriginRequestPolicy"

	// ErrCodeTooManyDistributionsWithFunctionAssociations for service response error code
	// "TooManyDistributionsWithFunctionAssociations".
	//
	// You have reached the maximum number of distributions that are associated
	// with a CloudFront function. For more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyDistributionsWithFunctionAssociations = "TooManyDistributionsWithFunctionAssociations"

	// ErrCodeTooManyDistributionsWithLambdaAssociations for service response error code
	// "TooManyDistributionsWithLambdaAssociations".
	//
	// Processing your request would cause the maximum number of distributions with
	// Lambda function associations per owner to be exceeded.
	ErrCodeTooManyDistributionsWithLambdaAssociations = "TooManyDistributionsWithLambdaAssociations"

	// ErrCodeTooManyDistributionsWithSingleFunctionARN for service response error code
	// "TooManyDistributionsWithSingleFunctionARN".
	//
	// The maximum number of distributions have been associated with the specified
	// Lambda function.
	ErrCodeTooManyDistributionsWithSingleFunctionARN = "TooManyDistributionsWithSingleFunctionARN"

	// ErrCodeTooManyFieldLevelEncryptionConfigs for service response error code
	// "TooManyFieldLevelEncryptionConfigs".
	//
	// The maximum number of configurations for field-level encryption have been
	// created.
	ErrCodeTooManyFieldLevelEncryptionConfigs = "TooManyFieldLevelEncryptionConfigs"

	// ErrCodeTooManyFieldLevelEncryptionContentTypeProfiles for service response error code
	// "TooManyFieldLevelEncryptionContentTypeProfiles".
	//
	// The maximum number of content type profiles for field-level encryption have
	// been created.
	ErrCodeTooManyFieldLevelEncryptionContentTypeProfiles = "TooManyFieldLevelEncryptionContentTypeProfiles"

	// ErrCodeTooManyFieldLevelEncryptionEncryptionEntities for service response error code
	// "TooManyFieldLevelEncryptionEncryptionEntities".
	//
	// The maximum number of encryption entities for field-level encryption have
	// been created.
	ErrCodeTooManyFieldLevelEncryptionEncryptionEntities = "TooManyFieldLevelEncryptionEncryptionEntities"

	// ErrCodeTooManyFieldLevelEncryptionFieldPatterns for service response error code
	// "TooManyFieldLevelEncryptionFieldPatterns".
	//
	// The maximum number of field patterns for field-level encryption have been
	// created.
	ErrCodeTooManyFieldLevelEncryptionFieldPatterns = "TooManyFieldLevelEncryptionFieldPatterns"

	// ErrCodeTooManyFieldLevelEncryptionProfiles for service response error code
	// "TooManyFieldLevelEncryptionProfiles".
	//
	// The maximum number of profiles for field-level encryption have been created.
	ErrCodeTooManyFieldLevelEncryptionProfiles = "TooManyFieldLevelEncryptionProfiles"

	// ErrCodeTooManyFieldLevelEncryptionQueryArgProfiles for service response error code
	// "TooManyFieldLevelEncryptionQueryArgProfiles".
	//
	// The maximum number of query arg profiles for field-level encryption have
	// been created.
	ErrCodeTooManyFieldLevelEncryptionQueryArgProfiles = "TooManyFieldLevelEncryptionQueryArgProfiles"

	// ErrCodeTooManyFunctionAssociations for service response error code
	// "TooManyFunctionAssociations".
	//
	// You have reached the maximum number of CloudFront function associations for
	// this distribution. For more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyFunctionAssociations = "TooManyFunctionAssociations"

	// ErrCodeTooManyFunctions for service response error code
	// "TooManyFunctions".
	//
	// You have reached the maximum number of CloudFront functions for this AWS
	// account. For more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyFunctions = "TooManyFunctions"

	// ErrCodeTooManyHeadersInCachePolicy for service response error code
	// "TooManyHeadersInCachePolicy".
	//
	// The number of headers in the cache policy exceeds the maximum. For more information,
	// see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyHeadersInCachePolicy = "TooManyHeadersInCachePolicy"

	// ErrCodeTooManyHeadersInForwardedValues for service response error code
	// "TooManyHeadersInForwardedValues".
	//
	// Your request contains too many headers in forwarded values.
	ErrCodeTooManyHeadersInForwardedValues = "TooManyHeadersInForwardedValues"

	// ErrCodeTooManyHeadersInOriginRequestPolicy for service response error code
	// "TooManyHeadersInOriginRequestPolicy".
	//
	// The number of headers in the origin request policy exceeds the maximum. For
	// more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyHeadersInOriginRequestPolicy = "TooManyHeadersInOriginRequestPolicy"

	// ErrCodeTooManyInvalidationsInProgress for service response error code
	// "TooManyInvalidationsInProgress".
	//
	// You have exceeded the maximum number of allowable InProgress invalidation
	// batch requests, or invalidation objects.
	ErrCodeTooManyInvalidationsInProgress = "TooManyInvalidationsInProgress"

	// ErrCodeTooManyKeyGroups for service response error code
	// "TooManyKeyGroups".
	//
	// You have reached the maximum number of key groups for this AWS account. For
	// more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyKeyGroups = "TooManyKeyGroups"

	// ErrCodeTooManyKeyGroupsAssociatedToDistribution for service response error code
	// "TooManyKeyGroupsAssociatedToDistribution".
	//
	// The number of key groups referenced by this distribution is more than the
	// maximum allowed. For more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyKeyGroupsAssociatedToDistribution = "TooManyKeyGroupsAssociatedToDistribution"

	// ErrCodeTooManyLambdaFunctionAssociations for service response error code
	// "TooManyLambdaFunctionAssociations".
	//
	// Your request contains more Lambda function associations than are allowed
	// per distribution.
	ErrCodeTooManyLambdaFunctionAssociations = "TooManyLambdaFunctionAssociations"

	// ErrCodeTooManyOriginCustomHeaders for service response error code
	// "TooManyOriginCustomHeaders".
	//
	// Your request contains too many origin custom headers.
	ErrCodeTooManyOriginCustomHeaders = "TooManyOriginCustomHeaders"

	// ErrCodeTooManyOriginGroupsPerDistribution for service response error code
	// "TooManyOriginGroupsPerDistribution".
	//
	// Processing your request would cause you to exceed the maximum number of origin
	// groups allowed.
	ErrCodeTooManyOriginGroupsPerDistribution = "TooManyOriginGroupsPerDistribution"

	// ErrCodeTooManyOriginRequestPolicies for service response error code
	// "TooManyOriginRequestPolicies".
	//
	// You have reached the maximum number of origin request policies for this AWS
	// account. For more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyOriginRequestPolicies = "TooManyOriginRequestPolicies"

	// ErrCodeTooManyOrigins for service response error code
	// "TooManyOrigins".
	//
	// You cannot create more origins for the distribution.
	ErrCodeTooManyOrigins = "TooManyOrigins"

	// ErrCodeTooManyPublicKeys for service response error code
	// "TooManyPublicKeys".
	//
	// The maximum number of public keys for field-level encryption have been created.
	// To create a new public key, delete one of the existing keys.
	ErrCodeTooManyPublicKeys = "TooManyPublicKeys"

	// ErrCodeTooManyPublicKeysInKeyGroup for service response error code
	// "TooManyPublicKeysInKeyGroup".
	//
	// The number of public keys in this key group is more than the maximum allowed.
	// For more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyPublicKeysInKeyGroup = "TooManyPublicKeysInKeyGroup"

	// ErrCodeTooManyQueryStringParameters for service response error code
	// "TooManyQueryStringParameters".
	//
	// Your request contains too many query string parameters.
	ErrCodeTooManyQueryStringParameters = "TooManyQueryStringParameters"

	// ErrCodeTooManyQueryStringsInCachePolicy for service response error code
	// "TooManyQueryStringsInCachePolicy".
	//
	// The number of query strings in the cache policy exceeds the maximum. For
	// more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyQueryStringsInCachePolicy = "TooManyQueryStringsInCachePolicy"

	// ErrCodeTooManyQueryStringsInOriginRequestPolicy for service response error code
	// "TooManyQueryStringsInOriginRequestPolicy".
	//
	// The number of query strings in the origin request policy exceeds the maximum.
	// For more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyQueryStringsInOriginRequestPolicy = "TooManyQueryStringsInOriginRequestPolicy"

	// ErrCodeTooManyRealtimeLogConfigs for service response error code
	// "TooManyRealtimeLogConfigs".
	//
	// You have reached the maximum number of real-time log configurations for this
	// AWS account. For more information, see Quotas (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html)
	// (formerly known as limits) in the Amazon CloudFront Developer Guide.
	ErrCodeTooManyRealtimeLogConfigs = "TooManyRealtimeLogConfigs"

	// ErrCodeTooManyStreamingDistributionCNAMEs for service response error code
	// "TooManyStreamingDistributionCNAMEs".
	//
	// Your request contains more CNAMEs than are allowed per distribution.
	ErrCodeTooManyStreamingDistributionCNAMEs = "TooManyStreamingDistributionCNAMEs"

	// ErrCodeTooManyStreamingDistributions for service response error code
	// "TooManyStreamingDistributions".
	//
	// Processing your request would cause you to exceed the maximum number of streaming
	// distributions allowed.
	ErrCodeTooManyStreamingDistributions = "TooManyStreamingDistributions"

	// ErrCodeTooManyTrustedSigners for service response error code
	// "TooManyTrustedSigners".
	//
	// Your request contains more trusted signers than are allowed per distribution.
	ErrCodeTooManyTrustedSigners = "TooManyTrustedSigners"

	// ErrCodeTrustedKeyGroupDoesNotExist for service response error code
	// "TrustedKeyGroupDoesNotExist".
	//
	// The specified key group does not exist.
	ErrCodeTrustedKeyGroupDoesNotExist = "TrustedKeyGroupDoesNotExist"

	// ErrCodeTrustedSignerDoesNotExist for service response error code
	// "TrustedSignerDoesNotExist".
	//
	// One or more of your trusted signers don't exist.
	ErrCodeTrustedSignerDoesNotExist = "TrustedSignerDoesNotExist"

	// ErrCodeUnsupportedOperation for service response error code
	// "UnsupportedOperation".
	//
	// This operation is not supported in this region.
	ErrCodeUnsupportedOperation = "UnsupportedOperation"
)
