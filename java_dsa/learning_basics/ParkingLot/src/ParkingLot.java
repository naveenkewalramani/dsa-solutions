public class ParkingLot {
    int floors;
    int parkingLotsPerFloor;
    int[][]parkingLots;

    // ParkingLot : method to create parking lot for given floors and slots
    public ParkingLot (int floors, int parkingLotsPerFloor) {
        this.floors = floors;
        this.parkingLotsPerFloor = parkingLotsPerFloor;
        this.parkingLots = new int[floors][parkingLotsPerFloor];
    }

    // addFloorToParkingLot : method to increase floor count
    void addFloorToParkingLot(){
        this.floors ++;
        int[][] newParkingLots = new int[this.floors][this.parkingLotsPerFloor];
        for (int i = 0; i<this.floors; i++){
            System.arraycopy(this.parkingLots[i], 0, newParkingLots[i], 0, this.parkingLotsPerFloor);
        }
        this.parkingLots = newParkingLots;
    }
}
