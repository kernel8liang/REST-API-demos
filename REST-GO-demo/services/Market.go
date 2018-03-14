package services

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/MsloveDl/HuobiProAPI/config"
	"github.com/MsloveDl/HuobiProAPI/models"
	"github.com/MsloveDl/HuobiProAPI/untils"
)

// 批量操作的API下个版本再封装

//------------------------------------------------------------------------------------------
// 交易API

// 获取K线数据
// strSymbol: 交易对, btcusdt, bccbtc......
// strPeriod: K线类型, 1min, 5min, 15min......
// nSize: 获取数量, [1-2000]
// return: KLineReturn 对象
func GetKLine(strSymbol, strPeriod string, nSize int) models.KLineReturn, bool {
	kLineReturn := models.KLineReturn{}

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	mapParams["period"] = strPeriod
	mapParams["size"] = strconv.Itoa(nSize)

	strRequestUrl := "/market/history/kline"
	strUrl := config.MARKET_URL + strRequestUrl

	jsonKLineReturn, ok := untils.HttpGetRequest(strUrl, mapParams)
    if !ok {
        return kLineReturn, ok
    }
	json.Unmarshal([]byte(jsonKLineReturn), &kLineReturn)

	return kLineReturn, ok
}

// 获取聚合行情
// strSymbol: 交易对, btcusdt, bccbtc......
// return: TickReturn对象
func GetTicker(strSymbol string) models.TickerReturn, bool {
	tickerReturn := models.TickerReturn{}

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol

	strRequestUrl := "/market/detail/merged"
	strUrl := config.MARKET_URL + strRequestUrl

	jsonTickReturn, ok := untils.HttpGetRequest(strUrl, mapParams)
    if !ok {
        return tickerReturn, ok
    }
	json.Unmarshal([]byte(jsonTickReturn), &tickerReturn)

	return tickerReturn, ok
}

// 获取交易深度信息
// strSymbol: 交易对, btcusdt, bccbtc......
// strType: Depth类型, step0、step1......stpe5 (合并深度0-5, 0时不合并)
// return: MarketDepthReturn对象
func GetMarketDepth(strSymbol, strType string) models.MarketDepthReturn, bool {
	marketDepthReturn := models.MarketDepthReturn{}

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	mapParams["type"] = strType

	strRequestUrl := "/market/depth"
	strUrl := config.MARKET_URL + strRequestUrl

	jsonMarketDepthReturn, ok := untils.HttpGetRequest(strUrl, mapParams)
    if !ok {
	    return marketDepthReturn, ok
    }
	json.Unmarshal([]byte(jsonMarketDepthReturn), &marketDepthReturn)

	return marketDepthReturn, ok
}

// 获取交易细节信息
// strSymbol: 交易对, btcusdt, bccbtc......
// return: TradeDetailReturn对象
func GetTradeDetail(strSymbol string) models.TradeDetailReturn, bool {
	tradeDetailReturn := models.TradeDetailReturn{}

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol

	strRequestUrl := "/market/trade"
	strUrl := config.MARKET_URL + strRequestUrl

	jsonTradeDetailReturn, ok := untils.HttpGetRequest(strUrl, mapParams)
    if !ok {
        return tradeDetailReturn, ok
    }
	json.Unmarshal([]byte(jsonTradeDetailReturn), &tradeDetailReturn)

	return tradeDetailReturn, ok
}

// 批量获取最近的交易记录
// strSymbol: 交易对, btcusdt, bccbtc......
// nSize: 获取交易记录的数量, 范围1-2000
// return: TradeReturn对象
func GetTrade(strSymbol string, nSize int) models.TradeReturn, bool {
	tradeReturn := models.TradeReturn{}

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	mapParams["size"] = strconv.Itoa(nSize)

	strRequestUrl := "/market/history/trade"
	strUrl := config.MARKET_URL + strRequestUrl

	jsonTradeReturn, ok := untils.HttpGetRequest(strUrl, mapParams)
    if !ok {
        return tradeReturn, ok
    }
	json.Unmarshal([]byte(jsonTradeReturn), &tradeReturn)

	return tradeReturn, ok
}

// 获取Market Detail 24小时成交量数据
// strSymbol: 交易对, btcusdt, bccbtc......
// return: MarketDetailReturn对象
func GetMarketDetail(strSymbol string) models.MarketDetailRetur, bool {
	marketDetailReturn := models.MarketDetailReturn{}

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol

	strRequestUrl := "/market/detail"
	strUrl := config.MARKET_URL + strRequestUrl

	jsonMarketDetailReturn, ok := untils.HttpGetRequest(strUrl, mapParams)
    if !ok {
        return marketDetailReturn, ok
    }
	json.Unmarshal([]byte(jsonMarketDetailReturn), &marketDetailReturn)

	return marketDetailReturn, ok
}

//------------------------------------------------------------------------------------------
// 公共API

// 查询系统支持的所有交易及精度
// return: SymbolsReturn对象
func GetSymbols() models.SymbolsReturn, bool {
	symbolsReturn := models.SymbolsReturn{}

	strRequestUrl := "/v1/common/symbols"
	strUrl := config.TRADE_URL + strRequestUrl

	jsonSymbolsReturn, ok := untils.HttpGetRequest(strUrl, nil)
    if !ok {
        return symbolsReturn, ok
    }
	json.Unmarshal([]byte(jsonSymbolsReturn), &symbolsReturn)

	return symbolsReturn, ok
}

// 查询系统支持的所有币种
// return: CurrencysReturn对象
func GetCurrencys() models.CurrencysReturn, bool {
	currencysReturn := models.CurrencysReturn{}

	strRequestUrl := "/v1/common/currencys"
	strUrl := config.TRADE_URL + strRequestUrl

	jsonCurrencysReturn, ok := untils.HttpGetRequest(strUrl, nil)
    if !ok {
        return currencysReturn, ok
    }
	json.Unmarshal([]byte(jsonCurrencysReturn), &currencysReturn)

	return currencysReturn, ok
}

// 查询系统当前时间戳
// return: TimestampReturn对象
func GetTimestamp() models.TimestampReturn, bool {
	timestampReturn := models.TimestampReturn{}

	strRequest := "/v1/common/timestamp"
	strUrl := config.TRADE_URL + strRequest

	jsonTimestampReturn, ok := untils.HttpGetRequest(strUrl, nil)
    if !ok {
        return timestampReturn, ok
    }
	json.Unmarshal([]byte(jsonTimestampReturn), &timestampReturn)

	return timestampReturn, ok
}

//------------------------------------------------------------------------------------------
// 用户资产API

// 查询当前用户的所有账户, 根据包含的私钥查询
// return: AccountsReturn对象
func GetAccounts() models.AccountsReturn, bool {
	accountsReturn := models.AccountsReturn{}

	strRequest := "/v1/account/accounts"
	jsonAccountsReturn, ok := untils.ApiKeyGet(make(map[string]string), strRequest)
    if !ok {
        return accountsReturn, ok
    }
	json.Unmarshal([]byte(jsonAccountsReturn), &accountsReturn)

	return accountsReturn, ok
}

// 根据账户ID查询账户余额
// nAccountID: 账户ID, 不知道的话可以通过GetAccounts()获取, 可以只现货账户, C2C账户, 期货账户
// return: BalanceReturn对象
func GetAccountBalance(strAccountID string) models.BalanceReturn, bool {
	balanceReturn := models.BalanceReturn{}

	strRequest := fmt.Sprintf("/v1/account/accounts/%s/balance", strAccountID)
	jsonBanlanceReturn, ok := untils.ApiKeyGet(make(map[string]string), strRequest)
    if !ok {
        return balanceReturn, ok
    }
	json.Unmarshal([]byte(jsonBanlanceReturn), &balanceReturn)

	return balanceReturn, ok
}

//------------------------------------------------------------------------------------------
// 交易API

// 下单
// placeRequestParams: 下单信息
// return: PlaceReturn对象
func Place(placeRequestParams models.PlaceRequestParams) models.PlaceReturn, bool {
	placeReturn := models.PlaceReturn{}

	mapParams := make(map[string]string)
	mapParams["account-id"] = placeRequestParams.AccountID
	mapParams["amount"] = placeRequestParams.Amount
	if 0 < len(placeRequestParams.Price) {
		mapParams["price"] = placeRequestParams.Price
	}
	if 0 < len(placeRequestParams.Source) {
		mapParams["source"] = placeRequestParams.Source
	}
	mapParams["symbol"] = placeRequestParams.Symbol
	mapParams["type"] = placeRequestParams.Type

	strRequest := "/v1/order/orders/place"
	jsonPlaceReturn, ok := untils.ApiKeyPost(mapParams, strRequest)
    if !ok {
	    return placeReturn, ok
    }
	json.Unmarshal([]byte(jsonPlaceReturn), &placeReturn)

	return placeReturn, ok
}

// 申请撤销一个订单请求
// strOrderID: 订单ID
// return: PlaceReturn对象
func SubmitCancel(strOrderID string) models.PlaceReturn {
	placeReturn := models.PlaceReturn{}

	strRequest := fmt.Sprintf("/v1/order/orders/%s/submitcancel", strOrderID)
	jsonPlaceReturn := untils.ApiKeyPost(make(map[string]string), strRequest)
	json.Unmarshal([]byte(jsonPlaceReturn), &placeReturn)

	return placeReturn
}
