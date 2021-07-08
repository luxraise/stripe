# Stripe
Stripe is a Client SDK for the Stripe API

## Usage
### New
```go
func ExampleNew() {
	var err error
	if testClient, err = New("[Stripe API Key]"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Stripe Client has been initialized! %v\n", testClient)
}
```

### Client.CreateCustomer
```go
func ExampleClient_CreateCustomer() {
	var (
		customer Customer
		created  Customer
		err      error
	)

	customer.Name = String("Leeroy Jenkins")

	if created, err = testClient.CreateCustomer(customer); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Stripe Customer has been created! %v\n", created)
}
```

### Client.GetCustomer
```go
func ExampleClient_GetCustomer() {
	var (
		customer Customer
		err      error
	)

	if customer, err = testClient.GetCustomer("[Stripe Customer ID]"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Stripe Customer has been retrieved! %v\n", customer)
}
```

### Client.UpdateCustomer
```go
func ExampleClient_UpdateCustomer() {
	var (
		customer Customer
		updated  Customer
		err      error
	)

	customer.Name = String("Leeroy Jenkins (Legend)")

	if updated, err = testClient.UpdateCustomer("[Stripe Customer ID]", customer); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Stripe Customer has been updated! %v\n", updated)
}
```

### Client.RemoveCustomer
```go
func ExampleClient_RemoveCustomer() {
	var err error
	if err = testClient.RemoveCustomer("[Stripe Customer ID]"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Stripe Customer has been removed!\n")
}
```
### Client.AddCreditCard
```go
func ExampleClient_AddCreditCard() {
	var (
		card    Card
		created Card
		err     error
	)

	card.CardNumber = "4242424242424242"
	card.CVC = String("123")
	card.ExpirationMonth = 11
	card.ExpirationYear = 2026

	if created, err = testClient.AddCreditCard("[Stripe Customer ID]", card); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Stripe Customer has had Credit Card added! %v\n", created)
}
```

### Client.RemoveCreditCard
```go
func ExampleClient_RemoveCreditCard() {
	var err error
	if err = testClient.RemoveCreditCard("[Stripe Customer ID]", "[Stripe Card ID]"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Stripe Customer has had Credit Card remove!\n")
}
```

### Client.CreateCharge
```go
func ExampleClient_CreateCharge() {
	var (
		charge  Charge
		created Charge
		err     error
	)

	charge.Amount = 1337
	charge.Currency = "usd"
	charge.Source = "[Stripe source ID]"

	if created, err = testClient.CreateCharge("[Stripe source ID]", charge); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Stripe Charge has been created! %v\n", created)
}
```