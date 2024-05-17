package bankcn

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	onceClient sync.Once
	httpClient *http.Client
)

func defaultHttpClient() *http.Client {
	onceClient.Do(func() {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 100
		transport.MaxConnsPerHost = 100
		transport.MaxIdleConnsPerHost = 100
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		}
	})
	return httpClient
}

// GetBankByCardOnline 使用阿里接口查询银行卡信息
func GetBankByCardOnline(cardNo string, bankInfo *Bank) error {
	values := make(url.Values)
	values.Set("_input_charset", "utf-8")
	values.Set("cardBinCheck", "true")
	values.Set("cardNo", cardNo)
	uri := fmt.Sprintf("%s?%s", VALIDATE_GATEWAY, values.Encode())

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return err
	}
	resp, err := defaultHttpClient().Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var vv validResponse
	if err := json.NewDecoder(resp.Body).Decode(&vv); err != nil {
		return err
	}
	if !vv.Validated {
		return errors.New("这是无效的银行卡")
	}
	bankInfo.CardType = vv.CardType
	bankInfo.Bank = vv.Bank
	bankInfo.GetName()
	return nil
}

type validResponse struct {
	CardType  CardType `json:"cardType"`
	Bank      string   `json:"bank"`
	Validated bool     `json:"validated"`
}
