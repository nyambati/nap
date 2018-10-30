package nap

// API struct
type API struct {
	baseURL       string
	Resources     map[string]RestResource
	DefaultRouter *CBRouter
	Client        *Client
}
