package bybit

import (
	"context"
	"log"

	bybitApi "github.com/wuhewuhe/bybit.go.api"
	"github.com/wuhewuhe/bybit.go.api/models"
)

type OrderBookRequest struct {
	Category string
	Symbol   string
	Limit    int
}

type OrderBookItem struct {
	Price       string
	Size        string
	IsBid       bool
	IsBestPrice bool
}

func GetOrderBook(client *bybitApi.Client, req *OrderBookRequest) ([]OrderBookItem, error) {
	orderBookInfo, err := client.
		NewOrderBookService().
		Category(models.Category(req.Category)).
		Symbol(req.Symbol).
		Limit(req.Limit).
		Do(context.Background())

	if err != nil {
		log.Println(err)
		return nil, err
	}

	//fmt.Println("orderBookInfo", bybitApi.PrettyPrint(orderBookInfo))

	return transformResponse(orderBookInfo), nil
}

func transformResponse(orderBookInfo *models.OrderBookInfo) []OrderBookItem {
	bidLen := len(orderBookInfo.Bids)
	askLen := len(orderBookInfo.Asks)
	isBestPrice := false
	orderBookList := make([]OrderBookItem, bidLen+askLen)

	for i, bid := range orderBookInfo.Bids {
		if i+1 == bidLen {
			isBestPrice = true
		}

		orderBookList[i] = OrderBookItem{
			Price:       bid[0],
			Size:        bid[1],
			IsBid:       true,
			IsBestPrice: isBestPrice,
		}
	}

	for i, ask := range orderBookInfo.Asks {
		if i != 0 {
			isBestPrice = false
		}

		orderBookList[i+bidLen] = OrderBookItem{
			Price:       ask[0],
			Size:        ask[1],
			IsBid:       false,
			IsBestPrice: isBestPrice,
		}
	}

	return orderBookList

}
