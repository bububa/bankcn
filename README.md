# 国内银行支行的联行号和地区信息, 银行图标, 校验银行卡

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/bankcn.svg)](https://pkg.go.dev/github.com/bububa/bankcn)
[![Go](https://github.com/bububa/bankcn/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/bankcn/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/bankcn/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/bankcn/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/bankcn.svg)](https://github.com/bububa/bankcn)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/bankcn)](https://goreportcard.com/report/github.com/bububa/bankcn)
[![GitHub license](https://img.shields.io/github/license/bububa/bankcn.svg)](https://github.com/bububa/bankcn/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/bankcn.svg)](https://GitHub.com/bububa/bankcn/releases/)


## API List

- 获取所有支行 [ BankBranchList() []Bank ]
- 根据 areaID 获取当前区域下所有支行 [ BankListByArea(bankID string, areaID string) []*Bank ]
- 检测是否是银行卡 [ IsBankCard(bankCardNo string) bool ]
- 根据卡号获取银行信息 [ GetBankByCardBin(bankCardNo string, bank *Bank) error ]
- 使用阿里接口查询银行卡信息 [ GetBankByCardOnline(cardNo string, bankInfo *Bank) error ]


## Usage
```golang

package main

import (
	"fmt"
	"log"

	"github.com/bububa/bankcn"
)

func main() {
    cardNo := "XXXXXX"
	var bank bankcn.Bank
	if err := bankcn.GetBankByCardOnline(cardNo, &bank); err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf("%+v\n", bank)
	if err := bankcn.GetBankByCardBin(cardNo, &bank); err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf("%+v\n", bank)
}

```

