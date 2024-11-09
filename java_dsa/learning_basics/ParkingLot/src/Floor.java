public class Floor {
    int floorNumber;
    int totalBikeSlot;
    int occupiedBikeSlots;
    int totalCarSlots;
    int occupiedCarSlots;
    int totalTruckSlot;
    int occupiedTruckSlots;

    public Floor(int floorNumber, int bikeSlot, int carSlot, int truckSlot) {
        this.floorNumber = floorNumber;
        this.totalBikeSlot = bikeSlot;
        this.occupiedBikeSlots = 0;
        this.totalCarSlots = carSlot;
        this.occupiedCarSlots = 0;
        this.totalTruckSlot = truckSlot;
        this.occupiedTruckSlots = 0;
    }

    public int freeCarSlots() {
        return this.totalCarSlots - this.occupiedCarSlots;
    }

    public int freeBikeSlots() {
        return this.totalBikeSlot - this.occupiedBikeSlots;
    }

    public int freeTruckSlots() {
        return this.totalTruckSlot - this.occupiedTruckSlots;
    }

    public void occupyCarSlot() {
        this.occupiedCarSlots++;
    }

    public void occupyBikeSlot() {
        this.occupiedBikeSlots++;
    }

    public void occupyTruckSlot() {
        this.occupiedTruckSlots++;
    }

    public void unOccupyCarSlot() {
        this.occupiedCarSlots--;
    }

    public void unOccupyBikeSlot() {
        this.occupiedBikeSlots--;
    }

    public void unOccupyTruckSlot() {
        this.occupiedTruckSlots--;
    }
}
