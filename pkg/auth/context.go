package auth

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/openshift-online/ocm-sdk-go/authentication"
)

// Context key type defined to avoid collisions in other pkgs using context
// See https://golang.org/pkg/context/#WithValue
type contextKey string

const (
	// Context Keys
	// FilterByOrganisation is used to determine whether resources are filtered by a user's organisation or as an individual owner
	contextFilterByOrganisation contextKey = "filter-by-organisation"

	// ocm token claim keys
	ocmUsernameKey string = "username"
	ocmOrgIdKey    string = "org_id"
	isOrgAdmin     string = "is_org_admin"

	// sso.redhat.com token claim keys
	ssoRHUsernameKey  string = "preferred_username"
	ssoRhAccountIdKey string = "account_id"
	// mas sso token claim keys
	masSsoOrgIdKey string = "rh-org-id"
)

func GetUsernameFromClaims(claims jwt.MapClaims) string {
	if claims[ocmUsernameKey] != nil {
		return claims[ocmUsernameKey].(string)
	}

	if claims[ssoRHUsernameKey] != nil {
		return claims[ssoRHUsernameKey].(string)
	}

	return ""
}

func GetAccountIdFromClaims(claims jwt.MapClaims) string {
	if claims[ssoRhAccountIdKey] != nil {
		return claims[ssoRhAccountIdKey].(string)
	}
	return ""
}

func GetOrgIdFromClaims(claims jwt.MapClaims) string {
	if claims[ocmOrgIdKey] != nil {
		return claims[ocmOrgIdKey].(string)
	}

	if claims[masSsoOrgIdKey] != nil {
		return claims[masSsoOrgIdKey].(string)
	}

	return ""
}

func GetIsOrgAdminFromClaims(claims jwt.MapClaims) bool {
	if claims[isOrgAdmin] != nil {
		return claims[isOrgAdmin].(bool)
	}
	return false
}

func SetFilterByOrganisationContext(ctx context.Context, filterByOrganisation bool) context.Context {
	return context.WithValue(ctx, contextFilterByOrganisation, filterByOrganisation)
}

func GetFilterByOrganisationFromContext(ctx context.Context) bool {
	filterByOrganisation := ctx.Value(contextFilterByOrganisation)
	if filterByOrganisation == nil {
		return false
	}
	return filterByOrganisation.(bool)
}

func SetTokenInContext(ctx context.Context, token *jwt.Token) context.Context {
	return authentication.ContextWithToken(ctx, token)
}

func GetClaimsFromContext(ctx context.Context) (jwt.MapClaims, error) {
	var claims jwt.MapClaims
	token, err := authentication.TokenFromContext(ctx)
	if err != nil {
		return claims, fmt.Errorf("failed to get jwt token from context: %v", err)
	}

	if token != nil && token.Claims != nil {
		claims = token.Claims.(jwt.MapClaims)
	}
	return claims, nil
}
