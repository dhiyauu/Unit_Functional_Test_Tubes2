package main

import "testing"

// ===============================
// MOCK VALIDATOR
// ===============================
type MockResiValidator struct{}

func (m MockResiValidator) Validate(resi string) bool {
	return true
}

// ===============================
// MOCK REPOSITORY
// ===============================
type MockRepository struct{}

func (m MockRepository) Insert(event TrackingEvent) error {
	return nil
}

// ===============================
// UNIT TEST INSERT TRACKING
// ===============================
func TestInsertTrackingEvent(t *testing.T) {

	mockValidator := MockResiValidator{}
	mockRepo := MockRepository{}

	req := TrackingEvent{
		Resi:      "RESI123",
		Lokasi:    "Gudang Jakarta",
		Event:     "Paket diterima di gudang",
		Timestamp: "2026-04-27 10:00:00",
	}

	event, err := InsertTrackingEvent(
		req,
		mockValidator,
		mockRepo,
	)

	if err != nil {
		t.Fatal(err)
	}

	// Test akan FAILED kalau service belum selesai
	if event.Resi != "RESI123" {
		t.Fail()
	}
}

// ===============================
// UNIT TEST GET TRACKING
// ===============================
func TestGetTrackingStatus(t *testing.T) {

	resi := "RESI123"

	resp := GetTrackingStatus(resi)

	// Akan FAILED karena service masih return nil
	if resp == nil {
		t.Fatal("Expected response but got nil")
	}

	if resp.Resi != resi {
		t.Fail()
	}
}

// ===============================
// UNIT TEST DISTANCE
// ===============================
func TestCalculateDistance(t *testing.T) {

	req := DistanceRequest{
		OriginAddress:      "Bandung",
		DestinationAddress: "Jakarta",
	}

	resp := CalculateDistance(req)

	// Akan FAILED karena service masih return nil
	if resp == nil {
		t.Fatal("Expected response but got nil")
	}

	if resp.DistanceKm <= 0 {
		t.Fail()
	}
}