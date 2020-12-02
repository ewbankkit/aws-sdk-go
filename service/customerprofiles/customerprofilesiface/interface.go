// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package customerprofilesiface provides an interface to enable mocking the Amazon Connect Customer Profiles service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package customerprofilesiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/customerprofiles"
)

// CustomerProfilesAPI provides an interface to enable mocking the
// customerprofiles.CustomerProfiles service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Connect Customer Profiles.
//    func myFunc(svc customerprofilesiface.CustomerProfilesAPI) bool {
//        // Make svc.AddProfileKey request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := customerprofiles.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCustomerProfilesClient struct {
//        customerprofilesiface.CustomerProfilesAPI
//    }
//    func (m *mockCustomerProfilesClient) AddProfileKey(input *customerprofiles.AddProfileKeyInput) (*customerprofiles.AddProfileKeyOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCustomerProfilesClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CustomerProfilesAPI interface {
	AddProfileKey(*customerprofiles.AddProfileKeyInput) (*customerprofiles.AddProfileKeyOutput, error)
	AddProfileKeyWithContext(aws.Context, *customerprofiles.AddProfileKeyInput, ...request.Option) (*customerprofiles.AddProfileKeyOutput, error)
	AddProfileKeyRequest(*customerprofiles.AddProfileKeyInput) (*request.Request, *customerprofiles.AddProfileKeyOutput)

	CreateDomain(*customerprofiles.CreateDomainInput) (*customerprofiles.CreateDomainOutput, error)
	CreateDomainWithContext(aws.Context, *customerprofiles.CreateDomainInput, ...request.Option) (*customerprofiles.CreateDomainOutput, error)
	CreateDomainRequest(*customerprofiles.CreateDomainInput) (*request.Request, *customerprofiles.CreateDomainOutput)

	CreateProfile(*customerprofiles.CreateProfileInput) (*customerprofiles.CreateProfileOutput, error)
	CreateProfileWithContext(aws.Context, *customerprofiles.CreateProfileInput, ...request.Option) (*customerprofiles.CreateProfileOutput, error)
	CreateProfileRequest(*customerprofiles.CreateProfileInput) (*request.Request, *customerprofiles.CreateProfileOutput)

	DeleteDomain(*customerprofiles.DeleteDomainInput) (*customerprofiles.DeleteDomainOutput, error)
	DeleteDomainWithContext(aws.Context, *customerprofiles.DeleteDomainInput, ...request.Option) (*customerprofiles.DeleteDomainOutput, error)
	DeleteDomainRequest(*customerprofiles.DeleteDomainInput) (*request.Request, *customerprofiles.DeleteDomainOutput)

	DeleteIntegration(*customerprofiles.DeleteIntegrationInput) (*customerprofiles.DeleteIntegrationOutput, error)
	DeleteIntegrationWithContext(aws.Context, *customerprofiles.DeleteIntegrationInput, ...request.Option) (*customerprofiles.DeleteIntegrationOutput, error)
	DeleteIntegrationRequest(*customerprofiles.DeleteIntegrationInput) (*request.Request, *customerprofiles.DeleteIntegrationOutput)

	DeleteProfile(*customerprofiles.DeleteProfileInput) (*customerprofiles.DeleteProfileOutput, error)
	DeleteProfileWithContext(aws.Context, *customerprofiles.DeleteProfileInput, ...request.Option) (*customerprofiles.DeleteProfileOutput, error)
	DeleteProfileRequest(*customerprofiles.DeleteProfileInput) (*request.Request, *customerprofiles.DeleteProfileOutput)

	DeleteProfileKey(*customerprofiles.DeleteProfileKeyInput) (*customerprofiles.DeleteProfileKeyOutput, error)
	DeleteProfileKeyWithContext(aws.Context, *customerprofiles.DeleteProfileKeyInput, ...request.Option) (*customerprofiles.DeleteProfileKeyOutput, error)
	DeleteProfileKeyRequest(*customerprofiles.DeleteProfileKeyInput) (*request.Request, *customerprofiles.DeleteProfileKeyOutput)

	DeleteProfileObject(*customerprofiles.DeleteProfileObjectInput) (*customerprofiles.DeleteProfileObjectOutput, error)
	DeleteProfileObjectWithContext(aws.Context, *customerprofiles.DeleteProfileObjectInput, ...request.Option) (*customerprofiles.DeleteProfileObjectOutput, error)
	DeleteProfileObjectRequest(*customerprofiles.DeleteProfileObjectInput) (*request.Request, *customerprofiles.DeleteProfileObjectOutput)

	DeleteProfileObjectType(*customerprofiles.DeleteProfileObjectTypeInput) (*customerprofiles.DeleteProfileObjectTypeOutput, error)
	DeleteProfileObjectTypeWithContext(aws.Context, *customerprofiles.DeleteProfileObjectTypeInput, ...request.Option) (*customerprofiles.DeleteProfileObjectTypeOutput, error)
	DeleteProfileObjectTypeRequest(*customerprofiles.DeleteProfileObjectTypeInput) (*request.Request, *customerprofiles.DeleteProfileObjectTypeOutput)

	GetDomain(*customerprofiles.GetDomainInput) (*customerprofiles.GetDomainOutput, error)
	GetDomainWithContext(aws.Context, *customerprofiles.GetDomainInput, ...request.Option) (*customerprofiles.GetDomainOutput, error)
	GetDomainRequest(*customerprofiles.GetDomainInput) (*request.Request, *customerprofiles.GetDomainOutput)

	GetIntegration(*customerprofiles.GetIntegrationInput) (*customerprofiles.GetIntegrationOutput, error)
	GetIntegrationWithContext(aws.Context, *customerprofiles.GetIntegrationInput, ...request.Option) (*customerprofiles.GetIntegrationOutput, error)
	GetIntegrationRequest(*customerprofiles.GetIntegrationInput) (*request.Request, *customerprofiles.GetIntegrationOutput)

	GetProfileObjectType(*customerprofiles.GetProfileObjectTypeInput) (*customerprofiles.GetProfileObjectTypeOutput, error)
	GetProfileObjectTypeWithContext(aws.Context, *customerprofiles.GetProfileObjectTypeInput, ...request.Option) (*customerprofiles.GetProfileObjectTypeOutput, error)
	GetProfileObjectTypeRequest(*customerprofiles.GetProfileObjectTypeInput) (*request.Request, *customerprofiles.GetProfileObjectTypeOutput)

	GetProfileObjectTypeTemplate(*customerprofiles.GetProfileObjectTypeTemplateInput) (*customerprofiles.GetProfileObjectTypeTemplateOutput, error)
	GetProfileObjectTypeTemplateWithContext(aws.Context, *customerprofiles.GetProfileObjectTypeTemplateInput, ...request.Option) (*customerprofiles.GetProfileObjectTypeTemplateOutput, error)
	GetProfileObjectTypeTemplateRequest(*customerprofiles.GetProfileObjectTypeTemplateInput) (*request.Request, *customerprofiles.GetProfileObjectTypeTemplateOutput)

	ListAccountIntegrations(*customerprofiles.ListAccountIntegrationsInput) (*customerprofiles.ListAccountIntegrationsOutput, error)
	ListAccountIntegrationsWithContext(aws.Context, *customerprofiles.ListAccountIntegrationsInput, ...request.Option) (*customerprofiles.ListAccountIntegrationsOutput, error)
	ListAccountIntegrationsRequest(*customerprofiles.ListAccountIntegrationsInput) (*request.Request, *customerprofiles.ListAccountIntegrationsOutput)

	ListDomains(*customerprofiles.ListDomainsInput) (*customerprofiles.ListDomainsOutput, error)
	ListDomainsWithContext(aws.Context, *customerprofiles.ListDomainsInput, ...request.Option) (*customerprofiles.ListDomainsOutput, error)
	ListDomainsRequest(*customerprofiles.ListDomainsInput) (*request.Request, *customerprofiles.ListDomainsOutput)

	ListIntegrations(*customerprofiles.ListIntegrationsInput) (*customerprofiles.ListIntegrationsOutput, error)
	ListIntegrationsWithContext(aws.Context, *customerprofiles.ListIntegrationsInput, ...request.Option) (*customerprofiles.ListIntegrationsOutput, error)
	ListIntegrationsRequest(*customerprofiles.ListIntegrationsInput) (*request.Request, *customerprofiles.ListIntegrationsOutput)

	ListProfileObjectTypeTemplates(*customerprofiles.ListProfileObjectTypeTemplatesInput) (*customerprofiles.ListProfileObjectTypeTemplatesOutput, error)
	ListProfileObjectTypeTemplatesWithContext(aws.Context, *customerprofiles.ListProfileObjectTypeTemplatesInput, ...request.Option) (*customerprofiles.ListProfileObjectTypeTemplatesOutput, error)
	ListProfileObjectTypeTemplatesRequest(*customerprofiles.ListProfileObjectTypeTemplatesInput) (*request.Request, *customerprofiles.ListProfileObjectTypeTemplatesOutput)

	ListProfileObjectTypes(*customerprofiles.ListProfileObjectTypesInput) (*customerprofiles.ListProfileObjectTypesOutput, error)
	ListProfileObjectTypesWithContext(aws.Context, *customerprofiles.ListProfileObjectTypesInput, ...request.Option) (*customerprofiles.ListProfileObjectTypesOutput, error)
	ListProfileObjectTypesRequest(*customerprofiles.ListProfileObjectTypesInput) (*request.Request, *customerprofiles.ListProfileObjectTypesOutput)

	ListProfileObjects(*customerprofiles.ListProfileObjectsInput) (*customerprofiles.ListProfileObjectsOutput, error)
	ListProfileObjectsWithContext(aws.Context, *customerprofiles.ListProfileObjectsInput, ...request.Option) (*customerprofiles.ListProfileObjectsOutput, error)
	ListProfileObjectsRequest(*customerprofiles.ListProfileObjectsInput) (*request.Request, *customerprofiles.ListProfileObjectsOutput)

	ListTagsForResource(*customerprofiles.ListTagsForResourceInput) (*customerprofiles.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *customerprofiles.ListTagsForResourceInput, ...request.Option) (*customerprofiles.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*customerprofiles.ListTagsForResourceInput) (*request.Request, *customerprofiles.ListTagsForResourceOutput)

	PutIntegration(*customerprofiles.PutIntegrationInput) (*customerprofiles.PutIntegrationOutput, error)
	PutIntegrationWithContext(aws.Context, *customerprofiles.PutIntegrationInput, ...request.Option) (*customerprofiles.PutIntegrationOutput, error)
	PutIntegrationRequest(*customerprofiles.PutIntegrationInput) (*request.Request, *customerprofiles.PutIntegrationOutput)

	PutProfileObject(*customerprofiles.PutProfileObjectInput) (*customerprofiles.PutProfileObjectOutput, error)
	PutProfileObjectWithContext(aws.Context, *customerprofiles.PutProfileObjectInput, ...request.Option) (*customerprofiles.PutProfileObjectOutput, error)
	PutProfileObjectRequest(*customerprofiles.PutProfileObjectInput) (*request.Request, *customerprofiles.PutProfileObjectOutput)

	PutProfileObjectType(*customerprofiles.PutProfileObjectTypeInput) (*customerprofiles.PutProfileObjectTypeOutput, error)
	PutProfileObjectTypeWithContext(aws.Context, *customerprofiles.PutProfileObjectTypeInput, ...request.Option) (*customerprofiles.PutProfileObjectTypeOutput, error)
	PutProfileObjectTypeRequest(*customerprofiles.PutProfileObjectTypeInput) (*request.Request, *customerprofiles.PutProfileObjectTypeOutput)

	SearchProfiles(*customerprofiles.SearchProfilesInput) (*customerprofiles.SearchProfilesOutput, error)
	SearchProfilesWithContext(aws.Context, *customerprofiles.SearchProfilesInput, ...request.Option) (*customerprofiles.SearchProfilesOutput, error)
	SearchProfilesRequest(*customerprofiles.SearchProfilesInput) (*request.Request, *customerprofiles.SearchProfilesOutput)

	TagResource(*customerprofiles.TagResourceInput) (*customerprofiles.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *customerprofiles.TagResourceInput, ...request.Option) (*customerprofiles.TagResourceOutput, error)
	TagResourceRequest(*customerprofiles.TagResourceInput) (*request.Request, *customerprofiles.TagResourceOutput)

	UntagResource(*customerprofiles.UntagResourceInput) (*customerprofiles.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *customerprofiles.UntagResourceInput, ...request.Option) (*customerprofiles.UntagResourceOutput, error)
	UntagResourceRequest(*customerprofiles.UntagResourceInput) (*request.Request, *customerprofiles.UntagResourceOutput)

	UpdateDomain(*customerprofiles.UpdateDomainInput) (*customerprofiles.UpdateDomainOutput, error)
	UpdateDomainWithContext(aws.Context, *customerprofiles.UpdateDomainInput, ...request.Option) (*customerprofiles.UpdateDomainOutput, error)
	UpdateDomainRequest(*customerprofiles.UpdateDomainInput) (*request.Request, *customerprofiles.UpdateDomainOutput)

	UpdateProfile(*customerprofiles.UpdateProfileInput) (*customerprofiles.UpdateProfileOutput, error)
	UpdateProfileWithContext(aws.Context, *customerprofiles.UpdateProfileInput, ...request.Option) (*customerprofiles.UpdateProfileOutput, error)
	UpdateProfileRequest(*customerprofiles.UpdateProfileInput) (*request.Request, *customerprofiles.UpdateProfileOutput)
}

var _ CustomerProfilesAPI = (*customerprofiles.CustomerProfiles)(nil)
