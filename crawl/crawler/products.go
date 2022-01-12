package crawler

import (
	"encoding/json"
	"log"
	"net/http"
)

// getProductIDs push the productIDs to the given channel get product IDs from the listing API
func getProductIDs(url string) ([]int32, error) {
	pIDs := []int32{}
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var productsListRes ProductListResponse

	if err := json.NewDecoder(resp.Body).Decode(&productsListRes); err != nil {
		log.Println("parsing error: ", err)
		return nil, err
	}
	for _, product := range productsListRes.ProductShortList {
		pIDs = append(pIDs, product.ID)
	}
	defer resp.Body.Close()
	return pIDs, nil
}
