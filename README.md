# PROBLEM STATEMENT: Create a Low-Level Design for a Parking Lot

## We will focus on the following set of requirements while designing the parking lot:

- The parking lot should have multiple floors where customers can park their cars.

- The parking lot should have multiple entry and exit points.

- The parking lot should book the nearest parking spot for the customers.

- Customers can collect a parking ticket from the entry points and can pay the parking fee at the exit points on their way out.

- The system should not allow more vehicles than the maximum capacity of the parking lot. If the parking is full, the system should be able to show a message at the entrance panel and on the parking display board on the ground floor. 

- Each parking floor will have many parking spots. The system should support multiple types of parking spots such as Small, Medium, Large etc.

- Each parking floor should have a display board showing any free parking spot for each spot type.

- Customers can pay the tickets at the automated exit panel or to the parking attendant. 

- Customers can pay via both cash and credit cards. 

- The system should support a per-hour parking fee model. For example, customers have to pay $4 for the first hour, $3 for the second and third hours, and $2 for all the remaining hours.

## Use Cases: Based on the above requirements, we can identify the following use cases:
1. Add/Remove/Edit Parking Floor: To add, remove or modify a parking floor from the system.
2. Add/Remove/Edit Parking Spot: To add, remove or modify a parking spot on a parking floor.
3. Add/Remove/Edit Parking Attendant: To add, remove or modify a parking attendant.
4. Take Ticket: To provide customers with a new parking ticket.
5. Scan Ticket: To scan a ticket at the exit panel.
6. Process Payment: To process the payment for the ticket.
7. Display Parking Floors: To display all the parking floors on the parking display board.
8. Display Free Spots: To display all the free parking spots on each floor on the parking display board.
9. Monitor Parking Lot: To monitor the parking lot for any discrepancies and issues.
10. Admin Login: To allow admins to log in to the system.
11. Customer Payment: To allow customers to pay for their ticket.

To create a class diagram, we need to identify the classes (entities) that will be used to implement the parking lot system. Based on the above requirements and use cases, we can have the following classes:
Remember, we need to keep our design simple and only include the main classes and their relationships. We can always add more classes later if required.

Step 1: Identify Classes/Objects
1. ParkingLot: Main class to manage the parking lot.
2. ParkingFloor: To manage each parking floor.
3. ParkingSpot: To manage each parking spot in the parking floor.
4. EntrancePanel: To manage the entrance panel.
5. ExitPanel: To manage the exit panel.
6. ParkingTicket: To manage the parking tickets.
7. Vehicle: To manage the vehicles.
8. Customer: To manage the customers.
9. Admin: To manage the admins.
10. Payment: To manage the payments.
11. DisplayBoard: To manage the display boards.
12. ParkingAttendant: To manage the parking attendants.

Step 2: Define attributes and functions for each class
1. ParkingLot: 
   - attributes: name, address, totalCapacity, parkingFloors, entrancePanels, exitPanels, displayBoards ....
   - functions: getParkingFloors, getEntrancePanels, getExitPanels, findAvailableParkingSpot() ....

2. ParkingFloor:
   - attributes: level, parkingSpots[], displayBoard, entrancePanelCount, exitPanelCount ....
   - functions: addParkingSpot(), removeParkingSpot(), updateDisplayBoard() ....

3. ParkingSpot:
   - attributes: spotNumber, spotType, isAvailable, vehicle ....
   - functions: assignVehicle(), removeVehicle() ....

4. ParkingTicket:
   - attributes: ticketNumber, issuedAt, paidAt, isLost, amount ....
   - functions: isPaid(), getAmount() ....

5. Vehicle:
   - attributes: licenseNumber, type, ticket ....
   - functions: assignTicket() ....

6. Customer:
   - attributes: name, vehicle, parkingTicket ....
   - functions: requestTicket() ....

7. Admin:
   - attributes: name, employeeId ....
   - functions: login(), addParkingFloor(), addParkingSpot() ....

	
Step 3: Define relationships between classes
- ParkingLot can have multiple ParkingFloors, EntrancePanels, ExitPanels, and DisplayBoards.
- Each ParkingFloor can have multiple ParkingSpots.
- Each ParkingSpot can have one Vehicle parked in it.
- ParkingTicket will be used by one Vehicle at a time.
- Customer can have one Vehicle.
- Admin can manage ParkingLot, ParkingFloor, ParkingSpot, etc.
- Payment will be associated with a ParkingTicket.
- DisplayBoard will be associated with ParkingFloor to display free spots.
- Customer makes Payment for a ParkingTicket.


Implementing the above design will give us a basic structure for our parking lot system. We can further enhance this design by adding more features and classes. 

