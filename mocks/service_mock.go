package main

type TrackingService interface {
	AddEvent(event TrackingEvent) error
	GetTracking(resi string) (TrackingResponse, error)

	GetDistance(origin, destination string) (DistanceResponse, error)
}