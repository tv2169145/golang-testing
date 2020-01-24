package test

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/tv2169145/golang-testing/src/api/app"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	go app.StartApp()
	os.Exit(m.Run())
}
