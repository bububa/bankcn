package bankcn

import (
	"encoding/json"
	"errors"
	"strings"
	"sync"

	_ "embed"
)

//go:embed card_bins.json
var binData []byte

var (
	onceBinInit sync.Once
	binList     []CardBin
)

func setupCardBin() {
	onceInit.Do(func() {
		json.Unmarshal(binData, &binList)
	})
}

func CardBinList() []CardBin {
	setupCardBin()
	return binList
}

type CardBin struct {
	Bin  string   `json:"bin,omitempty"`
	Bank string   `json:"bank,omitempty"`
	Type CardType `json:"type,omitempty"`
	Len  int      `json:"len,omitempty"`
}

// IsBankCard 检测是否是银行卡
func IsBankCard(bankCardNo string) bool {
	return regBankCard.MatchString(bankCardNo)
}

// SearchCardBin 根据传入的数据源，搜索银行卡信息
// 存在部分不同银行的卡Bin相同，但卡号长度不同。如622303和622305，16位是南京银行，18位是中国工商银行。所以cardBin查询时，如果卡号输入不完整，只给第一个结果。
// 当前有以下卡Bin存在重复：690755,622442,622425,622302,622308,622309,622510,622162,622307,622303,622305,621260
func SearchCardBin(bankCardNo string) []CardBin {
	binList := CardBinList()
	var ret []CardBin
	for _, cardBin := range binList {
		if !strings.HasPrefix(bankCardNo, cardBin.Bin) {
			continue
		}
		ret = append(ret, cardBin)
	}
	return ret
}

// GetBankByCardBin 根据卡号获取银行信息
func GetBankByCardBin(bankCardNo string, bank *Bank) error {
	if !IsBankCard(bankCardNo) {
		return errors.New("无效的银行卡号")
	}
	binList := SearchCardBin(bankCardNo)
	if len(binList) == 0 {
		return errors.New("没有找到该卡号的银行信息")
	}
	l := len(bankCardNo)
	for _, bin := range binList {
		if bin.Len == l {
			bank.Bank = bin.Bank
			bank.CardType = bin.Type
			bank.GetName()
		}
	}
	return nil
}
