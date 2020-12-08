package auth

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
)

/*
Helper class to retrieve component credentials.

If credentials are configured to be retrieved from ICredentialStore, it automatically locates ICredentialStore in component references and retrieve credentials from there using store_key parameter.

Configuration parameters
  credential:
  
    store_key: (optional) a key to retrieve the credentials from ICredentialStore
      ... other credential parameters
    credentials: alternative to credential
    
      [credential params 1]: first credential parameters
      ... credential parameters for key 1
      ...
      [credential params N]: Nth credential parameters
      ... credential parameters for key N
References
*:credential-store:*:*:1.0 (optional) Credential stores to resolve credentials
see
CredentialParams

see
ICredentialStore

Example
  config := NewConfigParamsFromTuples(
      "credential.user", "jdoe",
      "credential.pass",  "pass123"
  );
  
  credentialResolver := NewCredentialResolver();
  credentialResolver.Configure(config);
  credentialResolver.SetReferences(references);
  
  credentialResolver.Lookup("123", (err, credential) => {
      // Now use credential...
  });
*/
type CredentialResolver struct {
	credentials []*CredentialParams
	references  refer.IReferences
}

// Creates a new instance of credentials resolver.
// Returns *CredentialResolver
func NewEmptyCredentialResolver() *CredentialResolver {
	return &CredentialResolver{
		credentials: []*CredentialParams{},
		references:  nil,
	}
}

// Creates a new instance of credentials resolver.
// Parameters:
//   - config *config.ConfigParams
//   component configuration parameters
//   - references refer.IReferences
//   component references
// Returns *CredentialResolver
func NewCredentialResolver(config *config.ConfigParams,
	references refer.IReferences) *CredentialResolver {
	c := &CredentialResolver{
		credentials: []*CredentialParams{},
		references:  references,
	}

	if config != nil {
		c.Configure(config)
	}

	return c
}

// Configures component by passing configuration parameters.
// Parameters:
//   - config *config.ConfigParams
//   configuration parameters to be set.
func (c *CredentialResolver) Configure(config *config.ConfigParams) {
	credentials := NewManyCredentialParamsFromConfig(config)

	for _, credential := range credentials {
		c.credentials = append(c.credentials, credential)
	}
}

// Sets references to dependent components.
// Parameters:
//   - references refer.IReferences
//   references to locate the component dependencies.
func (c *CredentialResolver) SetReferences(references refer.IReferences) {
	c.references = references
}

// Gets all credentials configured in component configuration.
// Redirect to CredentialStores is not done at this point. If you need fully fleshed credential use lookup method instead.
// Returns []*CredentialParams
// a list with credential parameters
func (c *CredentialResolver) GetAll() []*CredentialParams {
	return c.credentials
}

// Adds a new credential to component credentials
// Parameters:
//   - credential *CredentialParams
//   new credential parameters to be added
func (c *CredentialResolver) Add(credential *CredentialParams) {
	c.credentials = append(c.credentials, credential)
}

func (c *CredentialResolver) lookupInStores(correlationId string,
	credential *CredentialParams) (result *CredentialParams, err error) {

	if !credential.UseCredentialStore() {
		return credential, nil
	}

	key := credential.StoreKey()
	if c.references == nil {
		return nil, nil
	}

	storeDescriptor := refer.NewDescriptor("*", "credential_store", "*", "*", "*")
	components := c.references.GetOptional(storeDescriptor)
	if len(components) == 0 {
		err := refer.NewReferenceError(correlationId, storeDescriptor)
		return nil, err
	}

	for _, component := range components {
		store, _ := component.(ICredentialStore)
		if store != nil {
			credential, err = store.Lookup(correlationId, key)
			if credential != nil || err != nil {
				return credential, err
			}
		}
	}

	return nil, nil
}

// Looks up component credential parameters. If credentials are configured to be retrieved from Credential store it finds a ICredentialStore and lookups credentials there.
// Parameters:
//   - correlationId string
// (optional) transaction id to trace execution through call chain.
// Returns *CredentialParams? error
func (c *CredentialResolver) Lookup(correlationId string) (*CredentialParams, error) {
	if len(c.credentials) == 0 {
		return nil, nil
	}

	lookupCredentials := []*CredentialParams{}

	for _, credential := range c.credentials {
		if !credential.UseCredentialStore() {
			return credential, nil
		}

		lookupCredentials = append(lookupCredentials, credential)
	}

	for _, credential := range lookupCredentials {
		c, err := c.lookupInStores(correlationId, credential)
		if c != nil || err != nil {
			return c, err
		}
	}

	return nil, nil
}
