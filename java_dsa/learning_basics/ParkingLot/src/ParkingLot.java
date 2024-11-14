import java.util.HashMap;
import java.util.Map;
import java.util.Objects;

public class ParkingLot {
    private int totalFloors;
    private Floor[] floorList;
    private Map<String, Ticket> vehicleNumberToTicketMap;

    // ParkingLot : method to create parking lot for given floors and slots
    public ParkingLot(int floors, int bikeSlots, int carSlots, int truckSlots) {
        this.totalFloors = floors;
        this.floorList = new Floor[floors];
        this.vehicleNumberToTicketMap = new HashMap<>();
        for (int i = 0; i < floors; i++) {
            Floor floor = new Floor(i, bikeSlots, carSlots, truckSlots);
            this.floorList[i] = floor;
        }
    }

    // addFloorToParkingLot : method to increase floor count
    void addFloorToParkingLot(int bikeSlots, int carSlots, int truckSlots) {
        this.totalFloors++;
        Floor[] newFloorList = new Floor[this.totalFloors];
        if (this.totalFloors - 1 >= 0) System.arraycopy(this.floorList, 0, newFloorList, 0, this.totalFloors - 1);
        Floor floor = new Floor(this.totalFloors, bikeSlots, carSlots, truckSlots);
        newFloorList[this.totalFloors - 1] = floor;
        this.floorList = newFloorList;
    }

    // getFreeFloorForVehicle : method to get first free floor for parking
    public int getFreeFloorForVehicle(String type) {
        int floorNumber = -1;
        for (int i = 0; i < this.totalFloors; i++) {
            if ((Objects.equals(type, Vehicle.vehicleTypeBike) && this.floorList[i].freeBikeSlots() > 0) ||
                    (Objects.equals(type, Vehicle.vehicleTypeCar) && this.floorList[i].freeCarSlots() > 0) ||
                    (Objects.equals(type, Vehicle.vehicleTypeTruck) && this.floorList[i].freeTruckSlots() > 0)) {
                floorNumber = i;
                break;
            }
        }
        return floorNumber;
    }

    // occupySlotOnFloor: method to occupy slot.
    public int occupySlotInParkingLot(String vehicleType, int floorNumber) {
        if (Objects.equals(vehicleType, Vehicle.vehicleTypeBike)) {
            return this.floorList[floorNumber].occupyBikeSlot();
        } else if (Objects.equals(vehicleType, Vehicle.vehicleTypeCar)) {
            return this.floorList[floorNumber].occupyCarSlot();
        } else if (Objects.equals(vehicleType, Vehicle.vehicleTypeTruck)) {
            return this.floorList[floorNumber].occupyTruckSlot();
        }
        return -1;
    }

    public void getFreeSlotsForVehicle(String vehicleType) {
        for (int i = 0; i < this.totalFloors; i++) {
            if (Objects.equals(vehicleType, Vehicle.vehicleTypeBike)) {
                System.out.printf("Floor Number: %d, Free Slots: %d \n", i, this.floorList[i].freeBikeSlots());
            } else if (Objects.equals(vehicleType, Vehicle.vehicleTypeCar)) {
                System.out.printf("Floor Number: %d, Free Slots: %d \n", i, this.floorList[i].freeCarSlots());
            } else if (Objects.equals(vehicleType, Vehicle.vehicleTypeTruck)) {
                System.out.printf("Floor Number: %d, Free Slots: %d \n", i, this.floorList[i].freeTruckSlots());
            }
        }
    }

    public void updateParkingLotTicketMapping(Vehicle vehicle, Ticket ticket) {
        this.vehicleNumberToTicketMap.put(vehicle.getRegistrationNumber(), ticket);
    }

    public void printParkingLotTicketMapping() {
        for (Map.Entry<String, Ticket> entry : this.vehicleNumberToTicketMap.entrySet())
            System.out.println("Vehicle Number = " + entry.getKey() + ", Ticket Number = " + entry.getValue().getTicketNumber());
    }

    public void addSlotOnFloor(int floorNumber, String type){
        this.floorList[floorNumber].addParkingSlot(type);
    }

    public boolean checkIfVehicleParked(String vehicleNumber){
        return this.vehicleNumberToTicketMap.containsKey(vehicleNumber);
    }

    public void unOccupySlotInParkingLot(String vehicleNumber){
        Ticket ticket = this.vehicleNumberToTicketMap.get(vehicleNumber);
        int floor = ticket.getFloor();
        int slotNumber = ticket.getSlot();
        String vehicleType = ticket.getVehicleType();
        if (Objects.equals(vehicleType, Vehicle.vehicleTypeBike)) {
            this.floorList[floor].unOccupyBikeSlot(slotNumber);
        } else if (Objects.equals(vehicleType, Vehicle.vehicleTypeCar)) {
            this.floorList[floor].unOccupyCarSlot(slotNumber);
        } else if (Objects.equals(vehicleType, Vehicle.vehicleTypeTruck)) {
            this.floorList[floor].unOccupyTruckSlot(slotNumber);
        }
        this.vehicleNumberToTicketMap.remove(vehicleNumber);
    }

}
