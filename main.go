package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	ParkingFloorCapacity = 20
	ParkingFloorCount    = 3
	FirstHourRate        = 4
	SecondHourRate       = 3
	SucceedingHourRate   = 2
)

type ParkingLotManager interface {
	AddFloor(floor ParkingFloor)
	GetFloors() []ParkingFloor
	AddAdmin(admin Admin)
	GetAdmins() []Admin
}

type ParkingFloorManager interface {
	AddParkingSpot(spot ParkingSpot)
	GetParkingSpots() []ParkingSpot
}

type ParkingSpotManager interface {
	ParkVehicle(vehicle Vehicle)
	IsAvailable() bool
}

// ParkingLot is the main struct that represents a parking lot
type ParkingLot struct {
	name           string
	address        string
	capacity       int
	floors         []ParkingFloor
	admins         []Admin
	entrancePanels []EntrancePanel
	exitPanels     []ExitPanel
	displayBoard   []DisplayBoard
}

type ParkingFloor struct {
	floorNumber  int
	spots        []ParkingSpot
	entryGate    Gate
	exitGate     Gate
	displayBoard DisplayBoard
}

type ParkingSpot struct {
	spotNumber  int
	floor       int
	vehicle     Vehicle
	isAvailable bool
}

type Vehicle struct {
	licensePlate string
	vehicleType  VehicleType
}

type VehicleType struct {
	wheelCount int
	size       string
}

type Gate struct {
	gateNumber int
}

type Ticket struct {
	ticketNumber int
	vehicle      Vehicle
	entryTime    time.Time
	exitTime     time.Time
	charge       int
}

type Customer struct {
	name        string
	phoneNumber string
	vehicle     Vehicle
	ticket      Ticket
}

type Admin struct {
	name string
}

type EntrancePanel struct {
	panelNumber int
}

type ExitPanel struct {
	panelNumber int
}

type DisplayBoard struct {
	boardNumber int
}

func NewParkingLot(name, address string, capacity int) *ParkingLot {
	return &ParkingLot{
		name:     name,
		address:  address,
		capacity: capacity,
	}
}

// AddFloor adds a floor to the parking lot
func (p *ParkingLot) AddFloor(floor ParkingFloor) {
	p.floors = append(p.floors, floor)
}

// GetFloors returns the floors in the parking lot
func (p *ParkingLot) GetFloors() []ParkingFloor {
	return p.floors
}

// AddAdmin adds an admin to the parking lot
func (p *ParkingLot) AddAdmin(admin Admin) {
	p.admins = append(p.admins, admin)
}

// GetAdmins returns the admins in the parking lot
func (p *ParkingLot) GetAdmins() []Admin {
	return p.admins
}

// GetCapacity returns the total capacity of the parking lot
func (p *ParkingLot) GetCapacity() int {
	return p.capacity
}

// GetAvailableSpots returns the number of available parking spots
func (p *ParkingLot) GetAvailableSpots() int {
	count := 0
	for _, floor := range p.floors {
		for _, spot := range floor.spots {
			if spot.IsAvailable() {
				count++
			}
		}
	}
	return count
}

// ParkVehicle parks a vehicle in the parking lot
func (p *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
	for _, floor := range p.floors {
		for i := range floor.spots {
			spot := &floor.spots[i]
			if spot.IsAvailable() {
				spot.ParkVehicle(vehicle)
				fmt.Printf("Vehicle parked on Floor %d, Spot %d\n", floor.floorNumber, spot.spotNumber)
				return true
			}
		}
	}
	return false
}

// AddParkingSpot adds a parking spot to a floor
func (p *ParkingFloor) AddParkingSpot(spot ParkingSpot) {
	p.spots = append(p.spots, spot)
}

// GetParkingSpots returns the parking spots on a floor
func (p *ParkingFloor) GetParkingSpots() []ParkingSpot {
	return p.spots
}

// ParkVehicle parks a vehicle in a parking spot
func (p *ParkingSpot) ParkVehicle(vehicle Vehicle) {
	if p.isAvailable {
		p.vehicle = vehicle
		p.isAvailable = false
	}
}

// IsAvailable returns whether the parking spot is available
func (p *ParkingSpot) IsAvailable() bool {
	return p.isAvailable
}

// CreateTicket creates a new parking ticket
func (p *ParkingLot) CreateTicket(vehicle Vehicle) *Ticket {
	ticket := &Ticket{
		ticketNumber: len(p.floors) + 1, // Simplified ticket number
		vehicle:      vehicle,
		entryTime:    time.Now(),
	}
	fmt.Printf("Ticket %d created for vehicle %s\n", ticket.ticketNumber, vehicle.licensePlate)
	return ticket
}

// DisplayParkingSpots displays the parking spots on each floor
func (p *ParkingLot) DisplayParkingSpots() {
	for _, floor := range p.floors {
		fmt.Printf("Floor %d:\n", floor.floorNumber)
		for _, spot := range floor.spots {
			fmt.Printf("  Spot %d: Available: %t\n", spot.spotNumber, spot.isAvailable)
		}
	}
}

// CalculateCharge calculates the charge based on entry and exit times
func (t *Ticket) CalculateCharge() int {
	duration := time.Since(t.entryTime).Hours()
	var charge int
	if duration <= 1 {
		charge = FirstHourRate
	} else if duration <= 3 {
		charge = FirstHourRate + SecondHourRate
	} else {
		charge = FirstHourRate + SecondHourRate + SucceedingHourRate*int(duration-3)
	}
	return charge
}

// ProcessPayment processes the payment for a ticket
func (t *Ticket) ProcessPayment(paymentMethod string) {
	charge := t.CalculateCharge()
	t.charge = charge
	fmt.Printf("Payment of $%.2f processed for Ticket %d using %s\n", charge, t.ticketNumber, paymentMethod)
}

func main() {
	// Create a new parking lot
	parkingLot := NewParkingLot("Parking Lot 1", "123 Main St", ParkingFloorCapacity*ParkingFloorCount)

	// Create all parking floors
	for i := 0; i < ParkingFloorCount; i++ {
		parkingFloor := ParkingFloor{
			floorNumber: i + 1,
		}
		parkingLot.AddFloor(parkingFloor)
	}

	// Create all parking spots
	for i := 0; i < ParkingFloorCount; i++ {
		for j := 0; j < ParkingFloorCapacity; j++ {
			parkingSpot := ParkingSpot{
				spotNumber:  j + 1,
				floor:       i + 1,
				isAvailable: true,
			}
			parkingLot.floors[i].AddParkingSpot(parkingSpot)
		}
	}

	// Every parking floor has 1 admin
	for i := 0; i < ParkingFloorCount; i++ {
		admin := Admin{
			name: "Admin " + strconv.Itoa(i+1),
		}
		parkingLot.AddAdmin(admin)
	}

	// New Customer arrives with a vehicle
	customer := Customer{
		name:        "John Doe",
		phoneNumber: "123-456-7890",
		vehicle: Vehicle{
			licensePlate: "ABC123",
			vehicleType: VehicleType{
				wheelCount: 4,
				size:       "medium",
			},
		},
	}

	// Check if there are available spots
	if parkingLot.GetAvailableSpots() == 0 {
		fmt.Println("No available spots")
		return
	}

	// Park the vehicle for the customer
	if !parkingLot.ParkVehicle(customer.vehicle) {
		fmt.Println("Failed to park the vehicle")
		return
	}

	// Create a ticket for the customer
	ticket := parkingLot.CreateTicket(customer.vehicle)

	// assign ticket to customer
	customer.ticket = *ticket

	// some time passed and customer is ready to leave
	// process payment
	// give ticket to exit panel

	// Display all parking spots
	parkingLot.DisplayParkingSpots()
}
