package stripe

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

var testAPIKey = os.Getenv("STRIPE_TEST_API_KEY")

func TestClient_customer_cycle(t *testing.T) {
	var (
		c   *Client
		err error
	)

	if _, err = New(""); err != ErrEmptyAPIKey {
		t.Fatalf("invalid error, expected %v and recieved %v", ErrEmptyAPIKey, err)
	}

	if c, err = New(testAPIKey); err != nil {
		t.Fatal(err)
	}

	var customer Customer
	name := fmt.Sprintf("Test %d", time.Now().Unix())
	customer.Name = &name
	customer.Metadata = Dictionary{"foo": "bar"}
	customer.PreferredLocales = []string{"en-US"}

	var created Customer
	if created, err = c.CreateCustomer(customer); err != nil {
		t.Fatal(err)
	}

	switch {
	case len(created.ID) == 0:
		t.Fatal("empty stripe ID encountered")
	case *created.Name != *customer.Name:
		t.Fatalf("invalid name, expected <%s> and received <%s>", *customer.Name, *created.Name)
	case !reflect.DeepEqual(customer.Metadata, created.Metadata):
		t.Fatalf("invalid metadata, expected %v and received %v", customer.Metadata, created.Metadata)
	case !reflect.DeepEqual(customer.PreferredLocales, created.PreferredLocales):
		t.Fatalf("invalid metadata, expected %v and received %v", customer.PreferredLocales, created.PreferredLocales)
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

	if err = c.RemoveCustomer(created.ID); err != nil {
		t.Fatal(err)
	}

	var deletedCustomer Customer
	if deletedCustomer, err = c.GetCustomer(created.ID); err != nil {
		t.Fatal(err)
	}

	if deletedCustomer.Name != nil {
		t.Fatalf("invalid name, expected <nil> and received <%s>", *deletedCustomer.Name)
	}
}

func TestClient_credit_card_cycle(t *testing.T) {
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
	defer func() { _ = c.RemoveCustomer(created.ID) }()

	var card Card
	// Attempt to add empty card
	if _, err = c.AddCreditCard(created.ID, card); err == nil || err.Error() != "error creating card token: Missing required param: card[number]." {
		t.Fatalf("invalid error, expected <%s> and received <%v>", "error creating card token: Missing required param: card[number].", err)
	}

	card.CardNumber = "4242424242424242"
	card.CVC = String("123")
	card.ExpirationMonth = 11
	card.ExpirationYear = 2026

	var createdCard Card
	if createdCard, err = c.AddCreditCard(created.ID, card); err != nil {
		t.Fatal(err)
	}

	var cards []Card
	if cards, err = c.ListCards(created.ID); err != nil {
		t.Fatal(err)
	}

	if len(cards) != 1 {
		t.Fatalf("invalid number of cards, expected %d and received %d", 1, len(cards))
	}

	listedCard := cards[0]

	switch {
	case listedCard.ID != createdCard.ID:
		t.Fatalf("invalid tokenized card ID, expected <%s> and received <%s>", createdCard.ID, listedCard.ID)
	case listedCard.ExpirationMonth != createdCard.ExpirationMonth:
		t.Fatalf("invalid expiration month, expected <%d> and received <%d>", createdCard.ExpirationMonth, listedCard.ExpirationMonth)
	case listedCard.ExpirationYear != createdCard.ExpirationYear:
		t.Fatalf("invalid expiration month, expected <%d> and received <%d>", createdCard.ExpirationYear, listedCard.ExpirationYear)
	}

	if err = c.RemoveCreditCard(created.ID, listedCard.ID); err != nil {
		t.Fatal(err)
	}

	var updatedCards []Card
	if updatedCards, err = c.ListCards(created.ID); err != nil {
		t.Fatal(err)
	}

	if len(updatedCards) != 0 {
		t.Fatalf("invalid number of cards, expected %d and received %d", 0, len(updatedCards))
	}
}
