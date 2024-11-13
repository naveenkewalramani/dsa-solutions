import java.util.HashMap;
import java.util.Map;
import java.util.Objects;

public class ParkingLot {
    private int totalFloors;
    private Floor[] floors;
    private Map<String, Ticket> ticketMap;

    // ParkingLot : method to create parking lot for given floors and slots
    public ParkingLot(int floors, int bikeSlots, int carSlots, int truckSlots) {
        this.totalFloors = floors;
        this.floors = new Floor[floors];
        this.ticketMap = new HashMap<>();
        for (int i = 0; i < floors; i++) {
            Floor floor = new Floor(i, bikeSlots, carSlots, truckSlots);
            this.floors[i] = floor;
        }
    }

    // addFloorToParkingLot : method to increase floor count
    void addFloorToParkingLot(int bikeSlots, int carSlots, int truckSlots) {
        this.totalFloors++;
        Floor[] newFloors = new Floor[this.totalFloors];
        if (this.totalFloors - 1 >= 0) System.arraycopy(this.floors, 0, newFloors, 0, this.totalFloors - 1);
        Floor floor = new Floor(this.totalFloors, bikeSlots, carSlots, truckSlots);
        newFloors[this.totalFloors - 1] = floor;
        this.floors = newFloors;
    }

    public int getFreeFloorForVehicle(String type) {
        int floorNumber = -1;
        for (int i = 0; i < this.totalFloors; i++) {
            if ((Objects.equals(type, Vehicle.vehicleTypeBike) && this.floors[i].freeBikeSlots() > 0) ||
                    (Objects.equals(type, Vehicle.vehicleTypeCar) && this.floors[i].freeCarSlots() > 0) ||
                    (Objects.equals(type, Vehicle.vehicleTypeTruck) && this.floors[i].freeTruckSlots() > 0)) {
                floorNumber = i;
                break;
            }
        }
        return floorNumber;
    }

    public int occupySlotOnFloor(String type, int floorNumber) {
        if (Objects.equals(type, Vehicle.vehicleTypeBike)) {
            return this.floors[floorNumber].occupyBikeSlot();
        } else if (Objects.equals(type, Vehicle.vehicleTypeCar)) {
            return this.floors[floorNumber].occupyCarSlot();
        } else if (Objects.equals(type, Vehicle.vehicleTypeTruck)) {
            return this.floors[floorNumber].occupyTruckSlot();
        }
        return -1;
    }

    public void getFreeSlotsForVehicle(String type) {
        for (int i = 0; i < this.totalFloors; i++) {
            if (Objects.equals(type, Vehicle.vehicleTypeBike)) {
                System.out.printf("Floor Number: %d, Free Slots: %d \n", i, this.floors[i].freeBikeSlots());
            } else if (Objects.equals(type, Vehicle.vehicleTypeCar)) {
                System.out.printf("Floor Number: %d, Free Slots: %d \n", i, this.floors[i].freeCarSlots());
            } else if (Objects.equals(type, Vehicle.vehicleTypeTruck)) {
                System.out.printf("Floor Number: %d, Free Slots: %d \n", i, this.floors[i].freeTruckSlots());
            }
        }
    }

    public void updateParkingLotTicketMapping(Vehicle vehicle, Ticket ticket) {
        this.ticketMap.put(vehicle.getRegistrationNumber(), ticket);
    }

    public void printParkingLotTicketMapping() {
        for (Map.Entry<String, Ticket> entry : this.ticketMap.entrySet())
            System.out.println("Key = " + entry.getKey() + ", Value = " + entry.getValue().toString());
    }
}
