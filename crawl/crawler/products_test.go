package crawler

// import (
// 	"fmt"
// 	"net/http"
// 	"testing"
// )

// func TestMain(m *testing.M) {
// 	// simple http server that serves the testdata directory
// 	go func() {
// 		http.Handle("/", http.FileServer(http.Dir("testdata")))
// 		http.ListenAndServe(":8080", nil)
// 	}()
// 	runTests := m.Run()
// 	fmt.Println("DONE", runTests)
// }

// func Test_getProductIDs(t *testing.T) {
// 	t.Run("Product list file", func(t *testing.T) {
// 		pIDs := getProductIDs("http://localhost:8080/products-list-api.json")
// 		result := []int32{72917624, 54622585, 42465518, 93154028, 128574812, 13158001, 145397453, 113075152, 136227595, 105812707, 141562457, 135716509, 105812772, 90072284, 105812767, 105812802, 5490513, 87381949, 60184289, 7614336, 54457494, 52676472, 140204791, 100130030, 134832143, 52469160, 43179817, 134757953, 71985140, 42115137, 134660934, 71984872, 134760917, 136037597, 72753179, 58679495, 134840309, 109544149, 104807183, 129507016, 87381030, 55995066, 53039077, 106661189, 135715996, 142792662, 135651314, 136040326}
// 		if len(pIDs) != len(result) {
// 			t.Errorf("Expected %d products, got %d", len(result), len(pIDs))
// 		}
// 	})
// }
