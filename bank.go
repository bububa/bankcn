package bankcn

import (
	"encoding/json"
	"sort"
	"sync"

	_ "embed"
)

//go:embed banks.json
var bankData []byte

// Bank 银行信息
type Bank struct {
	// 银行标识符
	Bank string `json:"bank,omitempty"`
	// 银行名字
	Name string `json:"name,omitempty"`
	// 银行联行号
	BankUnionID string `json:"bank_union_id,omitempty"`
	// 联系地址
	Address string `json:"address,omitempty"`
	// 联系电话
	Phone string `json:"phone,omitempty"`
	// 所在区域代号
	AreaID string `json:"area_id,omitempty"`
	// CardType 银行卡类型
	CardType CardType `json:"card_type"`
}

var (
	onceInit       sync.Once
	bankBranchList []Bank
	areaMapBank    map[string]map[string][]*Bank
	bankList       []Bank
)

func setupBank() {
	onceInit.Do(func() {
		json.Unmarshal(bankData, &bankBranchList)
		areaMapBank := make(map[string]map[string][]*Bank, len(bankBranchList))
		for _, bank := range bankBranchList {
			if areaMapBank[bank.Bank] == nil {
				areaMapBank[bank.Bank] = make(map[string][]*Bank)
			}
			areaID := bank.AreaID
			switch len(areaID) {
			case 4:
				areaMapBank[bank.Bank][areaID[:2]] = append(areaMapBank[bank.Bank][areaID[:2]], &bank)
				fallthrough
			case 2:
				areaMapBank[bank.Bank][areaID] = append(areaMapBank[bank.Bank][areaID], &bank)
			}
		}
		bi := make([]Bank, 0, len(bankMap))
		for k, v := range bankMap {
			bi = append(bi, Bank{
				Bank: k,
				Name: v,
			})
		}
		sort.Slice(bi, func(i, j int) bool {
			return bi[i].Bank < bi[j].Bank
		})
		bankList = bi
	})
}

// BankList 获取所有支行
func BankBranchList() []Bank {
	setupBank()
	return bankBranchList
}

func AreaMapBank() map[string]map[string][]*Bank {
	setupBank()
	return areaMapBank
}

// BankListByArea 根据 areaID 获取当前区域下所有支行
func BankListByArea(bankID string, areaID string) []*Bank {
	mp := AreaMapBank()
	b := mp[bankID]
	if b == nil {
		return nil
	}
	return b[areaID]
}

// GetName 根据bank 获取中文名
func (bank *Bank) GetName() string {
	bank.Name = bankMap[bank.Bank]
	return bank.Name
}
