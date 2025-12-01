package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidpogosian/pickupee/platform/initialization"
	"github.com/davidpogosian/pickupee/platform/repository"
	"github.com/davidpogosian/pickupee/platform/router"
	"github.com/davidpogosian/pickupee/service"
	_ "github.com/mattn/go-sqlite3"
)

func setupTestServer(t *testing.T) http.Handler {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open DB: %v", err)
	}

	if err := initialization.InitTables(db); err != nil {
		t.Fatalf("Failed to init tables: %v", err)
	}

	// Insert some items to use in orders
	_, err = db.Exec("INSERT INTO items (name) VALUES (?), (?), (?)", "item1", "item2", "item3")
	if err != nil {
		t.Fatalf("Failed to insert items: %v", err)
	}

	// Repositories & services
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)

	// Router
	r := router.Create(orderService)
	return r
}

func TestPlaceAndListOrdersHTTP(t *testing.T) {
	r := setupTestServer(t)

	// ----------------------
	// 1️⃣ Test PlaceOrder
	// ----------------------
	placeOrderReq := map[string]interface{}{
		"user_id":  42,
		"item_ids": []int{1, 2, 3},
	}
	reqBody, _ := json.Marshal(placeOrderReq)

	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("PlaceOrder returned status %d", w.Code)
	}

	var resp map[string]int
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	orderID, ok := resp["order_id"]
	if !ok || orderID <= 0 {
		t.Fatalf("Invalid order_id in response: %v", resp)
	}

	// ----------------------
	// 2️⃣ Test ListOrdersForUser
	// ----------------------
	listReq := map[string]interface{}{
		"user_id": 42,
	}
	listBody, _ := json.Marshal(listReq)

	req = httptest.NewRequest(http.MethodPost, "/orders/list", bytes.NewReader(listBody))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ListOrdersForUser returned status %d", w.Code)
	}

	var listResp map[string][]int
	if err := json.NewDecoder(w.Body).Decode(&listResp); err != nil {
		t.Fatalf("Failed to decode list response: %v", err)
	}

	orderIDs, ok := listResp["order_ids"]
	if !ok || len(orderIDs) != 1 || orderIDs[0] != orderID {
		t.Fatalf("Expected order_ids to contain %d, got %v", orderID, orderIDs)
	}
}
