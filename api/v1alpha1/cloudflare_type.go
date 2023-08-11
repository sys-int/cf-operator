package v1alpha1

type CloudflareDetails struct {
	//+kubebuilder:validation:Required
	// Cloudflare Domain to which this tunnel belongs to
	Domain string `json:"domain,omitempty"`

	//+kubebuilder:validation:Required
	// Secret containing Cloudflare API key/token
	Secret string `json:"secret,omitempty"`

	//+kubebuilder:validation:Optional
	// Account Name in Cloudflare. AccountName and AccountId cannot be both empty. If both are provided, Account ID is used if valid, else falls back to Account Name.
	AccountName string `json:"accountName,omitempty"`

	//+kubebuilder:validation:Optional
	// Account ID in Cloudflare. AccountId and AccountName cannot be both empty. If both are provided, Account ID is used if valid, else falls back to Account Name.
	AccountId string `json:"accountId,omitempty"`

	//+kubebuilder:validation:Optional
	// Email to use along with API Key for Delete operations for new tunnels only, or as an alternate to API Token
	Email string `json:"email,omitempty"`

	//+kubebuilder:validation:Optional
	//+kubebuilder:default:=CLOUDFLARE_API_KEY
	// Key in the secret to use for Cloudflare API Key, defaults to CLOUDFLARE_API_KEY. Needs Email also to be provided.
	// For Delete operations for new tunnels only, or as an alternate to API Token
	CLOUDFLARE_API_KEY string `json:"CLOUDFLARE_API_KEY,omitempty"`

	//+kubebuilder:validation:Optional
	//+kubebuilder:default:=CLOUDFLARE_API_TOKEN
	// Key in the secret to use for Cloudflare API token, defaults to CLOUDFLARE_API_TOKEN
	CLOUDFLARE_API_TOKEN string `json:"CLOUDFLARE_API_TOKEN,omitempty"`

	//+kubebuilder:validation:Optional
	//+kubebuilder:default:=CLOUDFLARE_TUNNEL_CREDENTIAL_FILE
	// Key in the secret to use as credentials.json for an existing tunnel, defaults to CLOUDFLARE_TUNNEL_CREDENTIAL_FILE
	CLOUDFLARE_TUNNEL_CREDENTIAL_FILE string `json:"CLOUDFLARE_TUNNEL_CREDENTIAL_FILE,omitempty"`

	//+kubebuilder:validation:Optional
	//+kubebuilder:default:=CLOUDFLARE_TUNNEL_CREDENTIAL_SECRET
	// Key in the secret to use as tunnel secret for an existing tunnel, defaults to CLOUDFLARE_TUNNEL_CREDENTIAL_SECRET
	CLOUDFLARE_TUNNEL_CREDENTIAL_SECRET string `json:"CLOUDFLARE_TUNNEL_CREDENTIAL_SECRET,omitempty"`
}
