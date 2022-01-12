package crawler

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Category struct {
	Name string
	ID   int
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// asd
func crawlManyPage(cat *Category, pageCount int) {
	for page := 1; page <= pageCount; page++ {
		crawlOnePage(cat, page)
	}
}

func crawlOnePage(cat *Category, page int) {
	// for easy argument passing, we ensure the existence of the directory here
	// if it does not exist, we create it
	dir := fmt.Sprintf("images/%s", cat.Name)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}
	pIDs, err := getProductIDs(fmt.Sprintf("https://tiki.vn/api/v2/products?limit=48&category=%d&page=%d", cat.ID, page))
	if err != nil {
		panic(err)
	}
	for _, pID := range pIDs {
		fmt.Println("pID: ", pID)
		productDetail, err := getProductDetail(pID)
		Must(err)
		for idx, image := range productDetail.Images {
			imageUrl := image.SmallURL
			// download image to images folder
			filePath := fmt.Sprintf("images/%s/%d-%d.jpg", cat.Name, pID, idx)
			err := downloadImage(imageUrl, filePath)
			if err != nil {
				log.Printf("Download image error: productID- %v %v %v", pID, err, imageUrl)
			}
			log.Printf("Download image success: productID- %v - %v", pID, filePath)
			time.Sleep(time.Millisecond * 300)
		}
		err = saveProductDetail(productDetail)
		if err != nil {
			log.Printf("Save product detail error: productID- %v %v", pID, err)
		}
		time.Sleep(time.Second * 3)
	}
}
func Run() {
    pageCount := 10
	crawlManyPage(&Category{"keo", 8276}, pageCount)
	crawlManyPage(&Category{"trai-cay", 44824}, pageCount)
	crawlManyPage(&Category{"bia", 53584}, pageCount)
	crawlManyPage(&Category{"man-hinh", 28930}, pageCount)
	crawlManyPage(&Category{"vi-nam", 959}, pageCount) // ví nam
	crawlManyPage(&Category{"den-hoc", 2016}, pageCount)
	crawlManyPage(&Category{"tai-nghe-gaming", 2971}, pageCount)
	crawlManyPage(&Category{"mi-pho-chao", 54390}, pageCount) // mì, phở, cháo
	crawlManyPage(&Category{"chuot-van-phong", 1838}, pageCount) // chuột vản phòng
	crawlManyPage(&Category{"chuot-gaming", 3428}, pageCount) // chuôt gaming
	crawlManyPage(&Category{"chao", 1984}, pageCount)
	crawlManyPage(&Category{"noi", 2136}, pageCount) // nồi
	crawlManyPage(&Category{"mu-bao-hiem", 11908}, pageCount)
	crawlManyPage(&Category{"binh-giu-nhiet", 1936}, pageCount)
	crawlManyPage(&Category{"vong-tay", 15318}, pageCount)
	// crawlManyPage(&Category{"keo", 8276}, pageCount)
	// crawlManyPage(&Category{"keo", 8276}, pageCount)
	// crawlManyPage(&Category{"keo", 8276}, pageCount)
	// crawlManyPage(&Category{"keo", 8276}, pageCount)
	// crawlManyPage(&Category{"keo", 8276}, pageCount)
	// crawlManyPage(&Category{"keo", 8276}, pageCount)

}
