package main



import (
	"testing"
	"fmt"
	
	"github.com/cyberbono3/apis/product-api/sdk/client"
	"github.com/cyberbono3/apis/product-api/sdk/client/products"
)

func TestClient(t *testing.T)  {

	cfg := client.DefaultTransportConfig().WithHost("localhost:8080")
	c := client.NewHTTPClientWithConfig(nil,cfg)
	
	//c := client.Default
	
	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)
	
	if err != nil {
		t.Fatal(err)

	fmt.Printf("%#v", prod.GetPayload()[0])


	}
}