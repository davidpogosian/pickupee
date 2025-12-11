package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidpogosian/pickupee/platform/initialization"
)

func setupTestScenario1(t *testing.T, db *sql.DB) {
	_, err := db.Exec("INSERT INTO items (name) VALUES (?), (?), (?)", "Eggs", "Milk", "Bread")
	if err != nil {
		t.Fatalf("Failed to insert items: %v", err)
	}
}

func setupTestScenario2(t *testing.T, db *sql.DB) {
	_, err := db.Exec("INSERT INTO items (name) VALUES (?), (?), (?)", "Eggs", "Milk", "Bread")
	if err != nil {
		t.Fatalf("Failed to insert Items: %v", err)
	}

	_, err = db.Exec("INSERT INTO orders (user_id) VALUES (?)", 42)
	if err != nil {
		t.Fatalf("Failed to insert Order: %v", err)
	}

	_, err = db.Exec("INSERT INTO order_items (order_id, item_id) VALUES (?, ?), (?, ?)", 1, 1, 1, 3)
	if err != nil {
		t.Fatalf("Failed to insert OrderItems: %v", err)
	}
}

func TestPlaceOrderHTTP(t *testing.T) {
	db, r := initialization.CreateServer(":memory:")
	defer db.Close()

	setupTestScenario1(t, db)

	placeOrderReq := map[string]any{
		"user_id":  42,
		"item_ids": []int{1, 3},
	}
	reqBody, err := json.Marshal(placeOrderReq)
	if err != nil {
		t.Fatalf("json marshal failed: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/placeOrder", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("/placeOrder returned status %d", w.Code)
	}

	var resp map[string]int
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check order ID
	orderID, ok := resp["order_id"]
	if !ok || orderID <= 0 {
		t.Fatalf("Invalid order_id in response: %v", resp)
	} else {
		t.Logf("Order ID: %d", orderID)
	}
}

func TestListOrdersForUser(t *testing.T) {
	db, r := initialization.CreateServer(":memory:")
	defer db.Close()

	setupTestScenario2(t, db)

	req := httptest.NewRequest(http.MethodGet, "/orders?user_id=42", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("/placeOrder returned status %d", w.Code)
	}

	var resp map[string][]int
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check order IDs
	orderIDs, ok := resp["order_ids"]
	if !ok {
		t.Fatalf("Invalid order_id in response: %v", resp)
	} else if len(orderIDs) != 1 {
		t.Fatalf("Expected one order_id, got none")
	} else {
		t.Log("Order IDs:")
		for _, orderID := range orderIDs {
			t.Logf("Order ID: %d \n", orderID)
		}
	}
}
