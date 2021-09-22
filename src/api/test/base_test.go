package test

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/omniaw/golang-testing/src/api/app"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	fmt.Println("about to start the application...")
	//in order not to block running the testcases use go routine
	go app.StartApp()
	fmt.Println("application started. about to start the test cases...")
	os.Exit(m.Run())
}
