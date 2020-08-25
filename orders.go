package woocommerce

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gozl/woocommerce/models"
)

// GetOrders returns a list of existing orders
func (wc *WooCommerce) GetOrders(page int) ([]models.Order, error) {
	body, err := wc.GetOrdersJSON(page)
	if err != nil {
		return nil, err
	}

	var result []models.Order
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetOrdersJSON returns the raw JSON data from WooCommerce endpoint
func (wc *WooCommerce) GetOrdersJSON(page int) ([]byte, error) {
	if page < 1 {
		page = 1
	}
	
	apiURL := fmt.Sprintf("%s/wp-json/wc/v3/orders?page=%d", wc.URL, page)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(wc.Username, wc.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	_ = resp.Body.Close()

	return body, nil
}