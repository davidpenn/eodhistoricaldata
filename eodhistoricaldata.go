package eodhistoricaldata

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

var (
	hostURL = "https://eodhistoricaldata.com"
)

// Client wraps eodhistoricaldata.com REST endpoints.
type Client struct {
	resty *resty.Client
	token string
}

// New constructs a new Client.
func New(token string) *Client {
	return &Client{
		resty: resty.New().SetHostURL(hostURL),
		token: token,
	}
}

type dividendData []*Dividend

func (m *dividendData) UnmarshalJSON(input []byte) error {
	var data map[string]json.RawMessage
	if err := json.Unmarshal(input, &data); err != nil {
		return err
	}
	var out []*Dividend
	for i := len(data) - 1; i >= 0; i-- {
		var cur *Dividend
		if err := json.Unmarshal(data[fmt.Sprintf("%d", i)], &cur); err != nil {
			return err
		}
		out = append(out, cur)
	}
	*m = dividendData(out)
	return nil
}

// GetHistoricalDividends for symbol on exchange.
// https://eodhistoricaldata.com/knowledgebase/api-splits-dividends/
func (m *Client) GetHistoricalDividends(symbol, exchange string, opts map[string]string) ([]*Dividend, error) {
	path := fmt.Sprintf("/api/div/%s.%s", symbol, exchange)
	params := make(map[string]string, len(opts)+2)
	for k, v := range opts {
		params[k] = v
	}
	params["api_token"] = m.token
	params["fmt"] = "json"
	var data dividendData
	return data, m.get(path, params, &data)
}

// GetHistoricalSplits for symbol on exchange.
// https://eodhistoricaldata.com/knowledgebase/api-splits-dividends/
func (m *Client) GetHistoricalSplits(symbol, exchange string, opts map[string]string) ([]*Split, error) {
	path := fmt.Sprintf("/api/splits/%s.%s", symbol, exchange)
	params := make(map[string]string, len(opts)+2)
	for k, v := range opts {
		params[k] = v
	}
	params["api_token"] = m.token
	params["fmt"] = "json"
	var data []*Split
	return data, m.get(path, params, &data)
}

// GetHistoricalShortInterest for symbol on exchange.
// https://eodhistoricaldata.com/knowledgebase/api-splits-dividends/
func (m *Client) GetHistoricalShortInterest(symbol, exchange string, opts map[string]string) ([]*ShortInterest, error) {
	path := fmt.Sprintf("/api/shorts/%s.%s", symbol, exchange)
	params := make(map[string]string, len(opts)+2)
	for k, v := range opts {
		params[k] = v
	}
	params["api_token"] = m.token
	params["fmt"] = "json"
	var data []*ShortInterest
	return data, m.get(path, params, &data)
}

// GetIntradayData for symbol on exchange.
// https://eodhistoricaldata.com/knowledgebase/intraday-historical-data-api/
func (m *Client) GetIntradayData(symbol, exchange string, opts map[string]string) ([]*IntradayData, error) {
	path := fmt.Sprintf("/api/intraday/%s.%s", symbol, exchange)
	params := make(map[string]string, len(opts)+2)
	for k, v := range opts {
		params[k] = v
	}
	params["api_token"] = m.token
	params["fmt"] = "json"
	var data []*IntradayData
	return data, m.get(path, params, &data)
}

// GetOptionData gets option data for symbol on exchange.
// https://eodhistoricaldata.com/knowledgebase/stock-options-data/
func (m *Client) GetOptionData(symbol, exchange string, opts map[string]string) (*OptionData, error) {
	path := fmt.Sprintf("/api/options/%s.%s", symbol, exchange)
	params := make(map[string]string, len(opts)+1)
	for k, v := range opts {
		params[k] = v
	}
	params["api_token"] = m.token
	var data *OptionData
	return data, m.get(path, params, &data)
}

// GetPrice of symbol on exchange. (20 minute delay)
// https://eodhistoricaldata.com/knowledgebase/live-realtime-stocks-api/
func (m *Client) GetPrice(symbol, exchange string) (*Price, error) {
	path := fmt.Sprintf("/api/real-time/%s.%s", symbol, exchange)
	params := map[string]string{
		"api_token": m.token,
		"fmt":       "json",
	}
	var data *Price
	return data, m.get(path, params, &data)
}

// GetPriceEOD get end of day prices for symbol on exchange.
// https://eodhistoricaldata.com/knowledgebase/api-for-historical-data-and-volumes/
func (m *Client) GetPriceEOD(symbol, exchange string, opts map[string]string) ([]*PriceEOD, error) {
	path := fmt.Sprintf("/api/eod/%s.%s", symbol, exchange)
	params := make(map[string]string, len(opts)+2)
	for k, v := range opts {
		params[k] = v
	}
	params["api_token"] = m.token
	params["fmt"] = "json"
	var data []*PriceEOD
	return data, m.get(path, params, &data)
}

// GetSymbolsForExchange gets a list of symbols for exchange.
// https://eodhistoricaldata.com/knowledgebase/list-symbols-exchange/
func (m *Client) GetSymbolsForExchange(exchange string) ([]*Symbol, error) {
	path := fmt.Sprintf("/api/exchanges/%s", exchange)
	params := map[string]string{
		"api_token": m.token,
		"fmt":       "json",
	}
	var data []*Symbol
	return data, m.get(path, params, &data)
}

// GetTechnicalIndicator for symbol on exchange.
// https://eodhistoricaldata.com/knowledgebase/technical-indicators-api/
func (m *Client) GetTechnicalIndicator(symbol, exchange, function string, period int, opts map[string]string) ([]*TechnicalIndicator, error) {
	path := fmt.Sprintf("/api/technical/%s.%s", symbol, exchange)
	params := make(map[string]string, len(opts)+3)
	for k, v := range opts {
		params[k] = v
	}
	params["api_token"] = m.token
	params["function"] = function
	params["period"] = fmt.Sprintf("%v", period)
	var data []*TechnicalIndicator
	return data, m.get(path, params, &data)
}

func (m *Client) get(path string, params map[string]string, dest interface{}) error {
	req := m.resty.R()
	if params != nil {
		req.SetQueryParams(params)
	}
	if dest != nil {
		req.SetResult(dest)
	}
	resp, err := req.Get(path)
	if resp.StatusCode() >= 400 && err == nil {
		return fmt.Errorf("%d - %s", resp.StatusCode(), resp.String())
	}
	return err
}
