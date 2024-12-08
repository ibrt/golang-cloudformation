// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_elasticloadbalancingv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__Type is the CloudFormation type for AWS::ElasticLoadBalancingV2::Listener.AuthenticateOidcConfig.
	AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__Type = "AWS::ElasticLoadBalancingV2::Listener.AuthenticateOidcConfig"
)

var (
	// AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap reports all the CloudFormation properties for AWS::ElasticLoadBalancingV2::Listener.AuthenticateOidcConfig.
	AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap = struct {
		AuthenticationRequestExtraParams string
		AuthorizationEndpoint            string
		ClientId                         string
		ClientSecret                     string
		Issuer                           string
		OnUnauthenticatedRequest         string
		Scope                            string
		SessionCookieName                string
		SessionTimeout                   string
		TokenEndpoint                    string
		UseExistingClientSecret          string
		UserInfoEndpoint                 string
	}{
		AuthenticationRequestExtraParams: "AuthenticationRequestExtraParams",
		AuthorizationEndpoint:            "AuthorizationEndpoint",
		ClientId:                         "ClientId",
		ClientSecret:                     "ClientSecret",
		Issuer:                           "Issuer",
		OnUnauthenticatedRequest:         "OnUnauthenticatedRequest",
		Scope:                            "Scope",
		SessionCookieName:                "SessionCookieName",
		SessionTimeout:                   "SessionTimeout",
		TokenEndpoint:                    "TokenEndpoint",
		UseExistingClientSecret:          "UseExistingClientSecret",
		UserInfoEndpoint:                 "UserInfoEndpoint",
	}

	// AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesSlice reports all the CloudFormation properties for AWS::ElasticLoadBalancingV2::Listener.AuthenticateOidcConfig.
	AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesSlice = []string{
		AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap.AuthenticationRequestExtraParams,
		AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap.AuthorizationEndpoint,
		AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap.ClientId,
		AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap.ClientSecret,
		AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap.Issuer,
		AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap.OnUnauthenticatedRequest,
		AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap.Scope,
		AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap.SessionCookieName,
		AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap.SessionTimeout,
		AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap.TokenEndpoint,
		AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap.UseExistingClientSecret,
		AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__PropertiesMap.UserInfoEndpoint,
	}
)

// AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig is a binding for AWS::ElasticLoadBalancingV2::Listener.AuthenticateOidcConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html
type AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig struct {
	// AuthenticationRequestExtraParams is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listener-authenticateoidcconfig-authenticationrequestextraparams
	AuthenticationRequestExtraParams cfz.ExpressionMap[string] `json:"AuthenticationRequestExtraParams,omitempty"`

	// AuthorizationEndpoint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listener-authenticateoidcconfig-authorizationendpoint
	AuthorizationEndpoint cfz.Expression[string] `json:"AuthorizationEndpoint,omitempty"`

	// ClientId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listener-authenticateoidcconfig-clientid
	ClientId cfz.Expression[string] `json:"ClientId,omitempty"`

	// ClientSecret is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listener-authenticateoidcconfig-clientsecret
	ClientSecret cfz.Expression[string] `json:"ClientSecret,omitempty"`

	// Issuer is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listener-authenticateoidcconfig-issuer
	Issuer cfz.Expression[string] `json:"Issuer,omitempty"`

	// OnUnauthenticatedRequest is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listener-authenticateoidcconfig-onunauthenticatedrequest
	OnUnauthenticatedRequest cfz.Expression[string] `json:"OnUnauthenticatedRequest,omitempty"`

	// Scope is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listener-authenticateoidcconfig-scope
	Scope cfz.Expression[string] `json:"Scope,omitempty"`

	// SessionCookieName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listener-authenticateoidcconfig-sessioncookiename
	SessionCookieName cfz.Expression[string] `json:"SessionCookieName,omitempty"`

	// SessionTimeout is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listener-authenticateoidcconfig-sessiontimeout
	SessionTimeout cfz.Expression[string] `json:"SessionTimeout,omitempty"`

	// TokenEndpoint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listener-authenticateoidcconfig-tokenendpoint
	TokenEndpoint cfz.Expression[string] `json:"TokenEndpoint,omitempty"`

	// UseExistingClientSecret is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listener-authenticateoidcconfig-useexistingclientsecret
	UseExistingClientSecret cfz.Expression[bool] `json:"UseExistingClientSecret,omitempty"`

	// UserInfoEndpoint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listener-authenticateoidcconfig-userinfoendpoint
	UserInfoEndpoint cfz.Expression[string] `json:"UserInfoEndpoint,omitempty"`
}

// New__AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig initializes a new AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig.
func New__AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig() AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	return AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) GetType() string {
	return AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig__Type
}

// Set__AuthenticationRequestExtraParams updates property "AuthenticationRequestExtraParams".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) Set__AuthenticationRequestExtraParams(v cfz.ExpressionMap[string]) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.AuthenticationRequestExtraParams = v
	return t
}

// SetM__AuthenticationRequestExtraParams updates property "AuthenticationRequestExtraParams".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) SetM__AuthenticationRequestExtraParams(v ...map[string]cfz.Expression[string]) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.AuthenticationRequestExtraParams = cfz.M(v...)
	return t
}

// SetMV__AuthenticationRequestExtraParams updates property "AuthenticationRequestExtraParams".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) SetMV__AuthenticationRequestExtraParams(v ...map[string]string) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.AuthenticationRequestExtraParams = cfz.MV(v...)
	return t
}

// Set__AuthorizationEndpoint updates property "AuthorizationEndpoint".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) Set__AuthorizationEndpoint(v cfz.Expression[string]) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.AuthorizationEndpoint = v
	return t
}

// SetV__AuthorizationEndpoint updates property "AuthorizationEndpoint".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) SetV__AuthorizationEndpoint(v string) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.AuthorizationEndpoint = cfz.V(v)
	return t
}

// Set__ClientId updates property "ClientId".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) Set__ClientId(v cfz.Expression[string]) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.ClientId = v
	return t
}

// SetV__ClientId updates property "ClientId".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) SetV__ClientId(v string) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.ClientId = cfz.V(v)
	return t
}

// Set__ClientSecret updates property "ClientSecret".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) Set__ClientSecret(v cfz.Expression[string]) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.ClientSecret = v
	return t
}

// SetV__ClientSecret updates property "ClientSecret".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) SetV__ClientSecret(v string) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.ClientSecret = cfz.V(v)
	return t
}

// Set__Issuer updates property "Issuer".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) Set__Issuer(v cfz.Expression[string]) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.Issuer = v
	return t
}

// SetV__Issuer updates property "Issuer".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) SetV__Issuer(v string) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.Issuer = cfz.V(v)
	return t
}

// Set__OnUnauthenticatedRequest updates property "OnUnauthenticatedRequest".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) Set__OnUnauthenticatedRequest(v cfz.Expression[string]) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.OnUnauthenticatedRequest = v
	return t
}

// SetV__OnUnauthenticatedRequest updates property "OnUnauthenticatedRequest".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) SetV__OnUnauthenticatedRequest(v string) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.OnUnauthenticatedRequest = cfz.V(v)
	return t
}

// Set__Scope updates property "Scope".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) Set__Scope(v cfz.Expression[string]) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.Scope = v
	return t
}

// SetV__Scope updates property "Scope".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) SetV__Scope(v string) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.Scope = cfz.V(v)
	return t
}

// Set__SessionCookieName updates property "SessionCookieName".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) Set__SessionCookieName(v cfz.Expression[string]) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.SessionCookieName = v
	return t
}

// SetV__SessionCookieName updates property "SessionCookieName".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) SetV__SessionCookieName(v string) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.SessionCookieName = cfz.V(v)
	return t
}

// Set__SessionTimeout updates property "SessionTimeout".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) Set__SessionTimeout(v cfz.Expression[string]) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.SessionTimeout = v
	return t
}

// SetV__SessionTimeout updates property "SessionTimeout".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) SetV__SessionTimeout(v string) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.SessionTimeout = cfz.V(v)
	return t
}

// Set__TokenEndpoint updates property "TokenEndpoint".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) Set__TokenEndpoint(v cfz.Expression[string]) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.TokenEndpoint = v
	return t
}

// SetV__TokenEndpoint updates property "TokenEndpoint".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) SetV__TokenEndpoint(v string) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.TokenEndpoint = cfz.V(v)
	return t
}

// Set__UseExistingClientSecret updates property "UseExistingClientSecret".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) Set__UseExistingClientSecret(v cfz.Expression[bool]) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.UseExistingClientSecret = v
	return t
}

// SetV__UseExistingClientSecret updates property "UseExistingClientSecret".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) SetV__UseExistingClientSecret(v bool) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.UseExistingClientSecret = cfz.V(v)
	return t
}

// Set__UserInfoEndpoint updates property "UserInfoEndpoint".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) Set__UserInfoEndpoint(v cfz.Expression[string]) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.UserInfoEndpoint = v
	return t
}

// SetV__UserInfoEndpoint updates property "UserInfoEndpoint".
func (t AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig) SetV__UserInfoEndpoint(v string) AWS_ElasticLoadBalancingV2_Listener_AuthenticateOidcConfig {
	t.UserInfoEndpoint = cfz.V(v)
	return t
}
