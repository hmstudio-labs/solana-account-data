package solanaaccountdata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gagliardetto/solana-go"
	sol "github.com/gagliardetto/solana-go"
)

type UseMethod uint8
type TokenStandard uint8

type Uses struct {
	UseMethod UseMethod
	Remaining uint64
	Total     uint64
}

type Creator struct {
	Address  sol.PublicKey
	Verified bool
	Share    uint8
}

type Collection struct {
	Verified bool
	Address  solana.PublicKey
}
type Data struct {
	Name                 string
	Symbol               string
	URI                  string
	SellerFeeBasisPoints uint16
	Creators             *[]Creator
}

type MetaplexAccountData struct {
	Key                 uint8
	UpdateAuthority     sol.PublicKey
	Mint                sol.PublicKey
	Data                Data
	PrimarySaleHappened bool
	IsMutable           bool
	EditionNonce        *uint8
	TokenStandard       *TokenStandard
	Collection          *Collection
	Uses                *Uses
}

type TokenMetadata struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Image       string `json:"image"`
	ShowName    bool   `json:"showName"`
	CreatedOn   string `json:"createdOn"`
	Twitter     string `json:"twitter"`
}

func (m *MetaplexAccountData) GetTokenMetadata() (*TokenMetadata, error) {
	url, err := sanitizeURL(m.Data.URI)
	if err != nil {
		return nil, err
	}
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("request failure, status code %d", response.StatusCode)
	}
	var tokenMetadata TokenMetadata
	err = json.Unmarshal(bodyBytes, &tokenMetadata)
	if err != nil {
		return nil, err
	}
	return &tokenMetadata, nil
}

func sanitizeURL(rawURL string) (string, error) {
	// 去除所有控制字符（包括\x00）
	clean := strings.Map(func(r rune) rune {
		if r >= 32 && r != 127 { // 保留可打印字符
			return r
		}
		return -1 // 删除不可见字符
	}, rawURL)

	// 验证URL格式
	_, err := url.Parse(clean)
	if err != nil {
		return "", fmt.Errorf("无效URL: %w", err)
	}
	return clean, nil
}
