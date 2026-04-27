package main

import (
	"fmt"
	"net/http"
	"time"
)

// Simulasi database di memory untuk test jika belum konek DB sungguhan
var trackingEvents []TrackingEvent
var nextTrackingID = 1

// INI BUAT TES LEWAT DOCKER, FUNCTIONAL TES
var OrderServiceURL = "http://order-service:8080"

// INI BUAT UNIT TES, LEWAT LOCAL
// var OrderServiceURL = "http://localhost:8080"

// ResiValidator adalah interface untuk mempermudah unit test (mocking)
// Digunakan untuk mengecek apakah resi valid/ada di Order Service
type ResiValidator interface {
	CheckResi(resi string) bool
}

type RealResiValidator struct{}

func (v RealResiValidator) CheckResi(resi string) bool {
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/order?resi=%s", OrderServiceURL, resi),
		nil,
	)

	// Set timeout agar tidak gantung jika service mati
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("ORDER SERVICE ERROR:", err)
		return false
	}

	fmt.Println("ORDER SERVICE STATUS:", resp.StatusCode)

	// Asumsi return 200 OK berarti resi valid
	return resp.StatusCode == 200
}

// --- Tracking Functions ---

// GetTrackingStatus mengambil timeline dari database (belum diimplementasikan)
func GetTrackingStatus(resi string) *TrackingResponse {
	return nil
}

// InsertTrackingEvent memasukkan event baru dengan memvalidasi resi terlebih dahulu
func InsertTrackingEvent(req TrackingEvent, v ResiValidator) (TrackingEvent, error) {
	return TrackingEvent{}, nil
}

// --- Map/Location Functions ---

// CalculateDistance memanggil API Maps/OSM untuk hitung jarak (belum diimplementasikan)
func CalculateDistance(req DistanceRequest) *DistanceResponse {
	return nil
}

// CalculateRoute optimasi rute dari origin ke destination melewati waypoints
func CalculateRoute(req RouteRequest) *DistanceResponse {
	return nil
}

// GetCourierLocation mengambil lokasi kurir saat ini dari cache/Redis
func GetCourierLocation(courierID string) *CourierLocation {
	return nil
}