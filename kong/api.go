package kong

// API represents an API in Kong
// Read https://getkong.org/docs/latest/admin-api/#api-object
type API struct {
	CreatedAt              *int64    `json:"created_at,omitempty"` //TODO marshal to time.Time
	Hosts                  []*string `json:"hosts,omitempty"`
	Methods                []*string `json:"methods,omitempty"` //TODO move to a stricter data type
	URIs                   []*string `json:"uris,omitempty"`
	HTTPIfTerminated       *bool     `json:"http_if_terminated,omitempty"`
	HTTPSOnly              *bool     `json:"https_only,omitempty"`
	ID                     *string   `json:"id,omitempty"`
	Name                   *string   `json:"name"`
	PreserveHost           *bool     `json:"preserve_host,omitempty"`
	Retries                *int      `json:"retries,omitempty"`
	StripURI               *bool     `json:"strip_uri,omitempty"`
	UpstreamConnectTimeout *int      `json:"upstream_connect_timeout,omitempty"`
	UpstreamReadTimeout    *int      `json:"upstream_read_timeout,omitempty"`
	UpstreamSendTimeout    *int      `json:"upstream_send_timeout,omitempty"`
	UpstreamURL            *string   `json:"upstream_url"`
}

// Valid checks if all the fields in API are valid
func (api *API) Valid() bool {
	if isEmptyString(api.Name) || isEmptyString(api.UpstreamURL) {
		return false
	}
	if len(api.Hosts) == 0 && len(api.Methods) == 0 && len(api.URIs) == 0 {
		return false
	}
	// TODO
	// TODO name must only contain alphanumeric and '., -, _, ~' characters
	// TODO check upstreamurl by parsing
	// TODO check methods are valid http methods
	// TODO check URIs starts with /
	// TODO all timeouts must be an integer between 1 and 2147483647
	// TODO "retries": "must be an integer between 0 and 32767"
	// TODO strip all of them
	return true
}