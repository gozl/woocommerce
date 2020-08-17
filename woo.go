package woocommerce

// WooCommerce is a woocommerce API client
type WooCommerce struct {
	URL string
	Username string
	Password string
}

// NewWooCommerce creates a new instance of WooCommerce
func NewWooCommerce(url, username, password string) WooCommerce {
	return WooCommerce{
		URL: url,
		Username: username,
		Password: password,
	}
}