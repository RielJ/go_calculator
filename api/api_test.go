package api

import (
	"testing"
)

func TestStore_Init(t *testing.T) {
	store := Store{}
	store.Init()
	if len(store) != 7 {
		t.Errorf("Expected 7 items, got %d", len(store))
	}
}

func TestStore_AddOrder(t *testing.T) {
	store := Store{}
	store.Init()
	store.AddOrder(RED)
	store.AddOrder(ORANGE)
	store.AddOrder(PINK)

	if store[RED].Orders != 1 {
		t.Errorf("Expected 1 order for RED, got %d", store[RED].Orders)
	}
	if store[ORANGE].Orders != 1 {
		t.Errorf("Expected 1 order for ORANGE, got %d", store[ORANGE].Orders)
	}
	if store[PINK].Orders != 1 {
		t.Errorf("Expected 1 order for PINK, got %d", store[PINK].Orders)
	}
	if store[GREEN].Orders != 0 {
		t.Errorf("Expected 0 order for GREEN, got %d", store[GREEN].Orders)
	}
}

func TestStore_CustomerMember(t *testing.T) {
	store := Store{}
	store.Init()
	// Add 1 order for each item
	// 50 + 120 + 80 + 40 + 30 + 50 + 90 = 460
	store.AddOrder(RED)
	store.AddOrder(ORANGE)
	store.AddOrder(PINK)
	store.AddOrder(GREEN)
	store.AddOrder(BLUE)
	store.AddOrder(YELLOW)
	store.AddOrder(PURPLE)

	// Calculate total
	// With membership, 460 * 0.9 = 414
	total := store.CalculateTotalCustomerMember()
	if total != 414 {
		t.Errorf("Expected 0, got %.2f", total)
	}
}

func TestStore_CalculateTotal(t *testing.T) {
	store := Store{}
	store.Init()

	store.AddOrder(RED)
	// 50

	store.AddOrder(ORANGE)
	store.AddOrder(ORANGE)
	// 120 * 0.95 + 120 * 0.95 = 228

	store.AddOrder(ORANGE)
	// 120

	store.AddOrder(PINK)
	store.AddOrder(PINK)
	// 80 * 0.95 + 80 * 0.95 = 152

	store.AddOrder(ORANGE)
	store.AddOrder(ORANGE)
	// 120 * 0.95 + 120 * 0.95 = 228

	// 50 + 228 + 120 + 152 + 228 = 778
	total := store.CalculateTotal()
	if total != 778 {
		t.Errorf("Expected 778, got %.2f", total)
	}
}

func TestStore_CalculateTotalWithMembership(t *testing.T) {
	store := Store{}
	store.Init()
	store.AddOrder(RED)

	store.AddOrder(ORANGE)
	store.AddOrder(ORANGE)

	store.AddOrder(ORANGE)
	store.AddOrder(PINK)
	store.AddOrder(PINK)

	store.AddOrder(ORANGE)
	store.AddOrder(ORANGE)
	total := store.CalculateTotalCustomerMember()
	if total != 700.20 {
		t.Errorf("Expected 700.20, got %.2f", total)
	}
}
