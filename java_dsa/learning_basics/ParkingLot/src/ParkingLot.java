import java.util.Objects;

public class ParkingLot {
    int totalFloors;
    Floor[] floors;

    // ParkingLot : method to create parking lot for given floors and slots
    public ParkingLot(int floors, int bikeSlots, int carSlots, int truckSlots) {
        this.totalFloors = floors;
        this.floors = new Floor[floors];
        for (int i = 0; i < floors; i++) {
            Floor floor = new Floor(i, bikeSlots, carSlots, truckSlots);
            this.floors[i] = floor;
        }
    }

    // addFloorToParkingLot : method to increase floor count
    void addFloorToParkingLot(int bikeSlots, int carSlots, int truckSlots) {
        this.totalFloors++;
        Floor[] newFloors = new Floor[this.totalFloors];
        for (int i = 0; i < this.totalFloors - 1; i++) {
            newFloors[i] = this.floors[i];
        }
        Floor floor = new Floor(this.totalFloors, bikeSlots, carSlots, truckSlots);
        newFloors[this.totalFloors - 1] = floor;
        this.floors = newFloors;
    }

    public int getFreeFloorForVehicle(String type) {
        int floorNumber = -1;
        for (int i = 0; i < this.totalFloors; i++) {
            if ((Objects.equals(type, "BIKE") && this.floors[i].freeBikeSlots() > 0) ||
                    (Objects.equals(type, "CAR") && this.floors[i].freeCarSlots() > 0) ||
                    (Objects.equals(type, "TRUCK") && this.floors[i].freeTruckSlots() > 0)) {
                floorNumber = i;
                break;
            }
        }
        return floorNumber;
    }

    public int occupySlotOnFloor(String type, int floorNumber) {
        if (Objects.equals(type, "BIKE")) {
            return this.floors[floorNumber].occupyBikeSlot();
        } else if (Objects.equals(type, "CAR")) {
            return this.floors[floorNumber].occupyCarSlot();
        } else if (Objects.equals(type, "TRUCK")) {
            return this.floors[floorNumber].occupyTruckSlot();
        }
        return -1;
    }

    public void getFreeSlotsForVehicle(String type) {
        for (int i = 0; i < this.totalFloors; i++) {
            if (Objects.equals(type, "BIKE")) {
                System.out.printf("Floor Number: %d, Free Slots: %d \n", i, this.floors[i].freeBikeSlots());
            } else if (Objects.equals(type, "CAR")) {
                System.out.printf("Floor Number: %d, Free Slots: %d \n", i, this.floors[i].freeCarSlots());
            } else if (Objects.equals(type, "TRUCK")) {
                System.out.printf("Floor Number: %d, Free Slots: %d \n", i, this.floors[i].freeTruckSlots());
            }
        }
    }
}
