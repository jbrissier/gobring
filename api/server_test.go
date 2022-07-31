package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/jbrissier/gobring/model"
)

func init() {
	// switch context to find the static svelt project
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestGetFN(t *testing.T) {

	router := SetupRouter("test.db")
	db := model.NewBringDB("test.db")

	type getTest struct {
		name, url string
	}

	bring, _ := db.GetAllBrings()
	if len(bring) == 0 {
		t.Error("no brings found ")
		return
	}

	tests := []getTest{
		{"svelt", "/"},
		{"all brings", "/api/bring"},
		{"one bring", fmt.Sprintf("/api/bring/%d", uint64(bring[0].ID))},
		{"bing items", fmt.Sprintf("/api/bring/%d/items", uint64(bring[0].ID))},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s:%s", test.name, test.url), func(t *testing.T) {

			req, _ := http.NewRequest("GET", test.url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			if w.Code != 200 {
				t.Errorf("Expected status code 200, got %d", w.Code)

			}

		})

	}

}
