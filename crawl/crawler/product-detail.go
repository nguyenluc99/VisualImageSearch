package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getProductDetail(pID int32) (*ProductDetail, error) {
	url := fmt.Sprintf("https://tiki.vn/api/v2/products/%d", pID)
	res, err := http.Get(url)
	if err != nil {
		return &ProductDetail{}, err
	}

	var productDetail ProductDetail
	if err := json.NewDecoder(res.Body).Decode(&productDetail); err != nil {
		log.Fatalf("JSON parse error: %v", err)
		return &ProductDetail{}, err
	}

	return &productDetail, nil
}

func downloadImage(imageUrl, filePath string) error {
	//Get the response bytes from the url
	response, err := http.Get(imageUrl)
	if err != nil {
		return fmt.Errorf("error requesting image: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
        return fmt.Errorf("received non 200 response code: %v", response.StatusCode)
	}
	//Create a empty file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file image: %v", err)
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("error copy image data to file: %v", err)
	}

	return nil
}
