public class ParkingLot {
    int totalFloors;
    Floor []floors;

    // ParkingLot : method to create parking lot for given floors and slots
    public ParkingLot (int floors, int bikeSlots, int carSlots,int truckSlots) {
        this.totalFloors = floors;
        this.floors = new Floor[floors];
        for (int i = 0;i<floors;i++) {
            Floor floor = new Floor(i,bikeSlots,carSlots,truckSlots);
            this.floors[i] = floor;
        }
    }

    // addFloorToParkingLot : method to increase floor count
    void addFloorToParkingLot( int bikeSlots, int carSlots,int truckSlots){
        this.totalFloors ++;
        Floor []newFloors = new Floor[this.totalFloors];
        for (int i = 0; i<this.totalFloors-1; i++){
            newFloors[i] = this.floors[i];
        }
        Floor floor = new Floor(this.totalFloors, bikeSlots,carSlots,truckSlots);
        newFloors[this.totalFloors-1] = floor;
        this.floors = newFloors;
    }
}
