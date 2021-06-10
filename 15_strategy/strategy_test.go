package strategy

func ExamplePayByCash() {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()
	// Output:
	// Pay $123 to Ada by cash
}

func ExamplePayByBank() {
	payment := NewPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()
	// Output:
	// Pay $888 to Bob by bank account 0002
}

func ExampleOperationAdd_DoOperation() {
	operation := NewCalculation(&OperationAdd{})
	operation.ExecuteStrategy(1, 2)
}

func ExampleOperationSubtract_DoOperation() {
	operation := NewCalculation(&OperationSubtract{})
	operation.ExecuteStrategy(2, 1)
}

func ExampleOperationMultiply_DoOperation() {
	operation := NewCalculation(&OperationMultiply{})
	operation.ExecuteStrategy(2, 1)
}