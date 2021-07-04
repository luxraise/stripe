package stripe

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testAPIKey = os.Getenv("STRIPE_TEST_API_KEY")

func TestClient_customer_cycle(t *testing.T) {
	var (
		c   *Client
		err error
	)

	if c, err = New(testAPIKey); err != nil {
		t.Fatal(err)
	}

	var customer Customer
	name := fmt.Sprintf("Test %d", time.Now().Unix())
	customer.Name = &name

	var created Customer
	if created, err = c.CreateCustomer(customer); err != nil {
		t.Fatal(err)
	}

	switch {
	case len(created.ID) == 0:
		t.Fatal("empty stripe ID encountered")
	case *created.Name != *customer.Name:
		t.Fatalf("invalid name, expected <%s> and received <%s>", *customer.Name, *created.Name)
	}

	var retrieved Customer
	if retrieved, err = c.GetCustomer(created.ID); err != nil {
		t.Fatal(err)
	}

	switch {
	case retrieved.ID != created.ID:
		t.Fatalf("invalid ID encounterd, expected <%s> and received <%s>", created.ID, retrieved.ID)
	case *retrieved.Name != *customer.Name:
		t.Fatalf("invalid name, expected <%s> and received <%s>", *customer.Name, *retrieved.Name)
	}

	// Create copy of retrieved valued
	edited := retrieved
	editedName := fmt.Sprintf("Test %d", time.Now().Unix())
	edited.Name = &editedName

	var updated Customer
	if updated, err = c.UpdateCustomer(created.ID, edited); err != nil {
		t.Fatal(err)
	}

	switch {
	case updated.ID != created.ID:
		t.Fatalf("invalid ID encounterd, expected <%s> and received <%s>", created.ID, retrieved.ID)
	case *updated.Name != *edited.Name:
		t.Fatalf("invalid name, expected <%s> and received <%s>", *edited.Name, *updated.Name)
	}

}

func TestClient_credit_card_cycle(t *testing.T) {
	//var (
	//	c   *Client
	//	err error
	//)
	//
	//if c, err = New(testAPIKey); err != nil {
	//	t.Fatal(err)
	//}
	//
	//var customer Customer
	//name := fmt.Sprintf("Test %d", time.Now().Unix())
	//customer.Name = &name

	//	var created Customer
	//	if created, err = c.CreateCustomer(customer); err != nil {
	//		t.Fatal(err)
	//	}

}

func TestClient_CreateCardToken(t *testing.T) {
	var (
		c   *Client
		err error
	)

	if c, err = New(testAPIKey); err != nil {
		t.Fatal(err)
	}

	var card Card
	card.CardNumber = "4242424242424242"
	card.CVC = String("123")
	card.ExpirationMonth = 11
	card.ExpirationYear = 2026

	var created Token
	if created, err = c.CreateCardToken(card); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Created: %+v\n", created)
}
