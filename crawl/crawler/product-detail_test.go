package crawler

import "testing"

func Test_downloadImage(t *testing.T) {

	if err := downloadImage("https://salt.tikicdn.com/ts/product/e6/b8/07/f9696d3ef3c10a6bd370ec1afb31f152.jpg", "a.jpg"); err != nil {
		t.Errorf("downloadImage() error = %v", err)
	}
}
