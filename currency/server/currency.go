package server

import (
	"context"
	"log"
	"time"

	
	protos "github.com/cyberbono3/apis/currency/protos/currency"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Currency struct {
	log *log.Logger
}

// NewCurrency is concstructor for currency server
func NewCurrency(l *log.Logger) *Currency {
	return &Currency{l}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Println("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
//SubscribeRates sends a standart message back to client
func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {
	
	go func() {
		for {
			rr, err := src.Recv()
			if err == io.EOF {
				c.log.Info("Client has closed connection")
				break
			}
			if err != nil {
			c.log.Error("Unable to read from", err)
			}
			c.log.info("Handle client request", "request", err)
		}
	}()

	for {
		if err := src.Send(&protos.RateResponse{Rate: 12.1}); err != nil {
			return err
		}

		time.Sleep(5*time.Second)




	}
}
