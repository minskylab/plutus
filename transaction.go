package plutus

import "time"

// FlowType is how the flow of resources (money in this case) is propagated
type FlowType string

// Charge represents a money flow from your customer to you
var Charge FlowType = "charge"

// TransactionState is the state of a transaction
type TransactionState string

// Created is when the transaction eas early created
var Created TransactionState = "created"

// Transaction represents a transaction of money
type Transaction struct {
	ID        string
	Type      FlowType
	State     TransactionState
	Snapshots []*TransactionSnapshot
	Give      []*Product
	Expected  []*Product
}

// TransactionSnapshot is a snapshot of one time transaction
type TransactionSnapshot struct {
	At       time.Time
	Snapshot Transaction
}
