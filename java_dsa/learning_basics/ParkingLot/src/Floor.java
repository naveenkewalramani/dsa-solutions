public class Floor {
    private int floorNumber;
    private  int[] bikeSlots;
    private int occupiedBikeSlots;
    private int[] carSlots;
    private int occupiedCarSlots;
    private int[] truckSlots;
    private int occupiedTruckSlots;

    public Floor(int floorNumber, int bikeSlot, int carSlot, int truckSlot) {
        this.floorNumber = floorNumber;
        this.bikeSlots = new int[bikeSlot];
        this.occupiedBikeSlots = 0;
        this.carSlots = new int[carSlot];
        this.occupiedCarSlots = 0;
        this.truckSlots = new int[truckSlot];
        this.occupiedTruckSlots = 0;
    }

    public int freeCarSlots() {
        return this.carSlots.length - this.occupiedCarSlots;
    }

    public int freeBikeSlots() {
        return this.bikeSlots.length - this.occupiedBikeSlots;
    }

    public int freeTruckSlots() {
        return this.truckSlots.length - this.occupiedTruckSlots;
    }

    public int occupyCarSlot() {
        int slotNumber = -1;
        boolean slotFound = false;
        for (int i = 0; i < this.carSlots.length; i++) {
            if (this.carSlots[i] == 0) {
                slotNumber = i;
                slotFound = true;
                break;
            }
        }
        if (slotFound){
            this.carSlots[slotNumber] = -1;
            this.occupiedCarSlots++;
            return slotNumber;
        }
        return -1;
    }

    public int occupyBikeSlot() {
        int slotNumber = -1;
        boolean slotFound = false;
        for (int i = 0; i < this.bikeSlots.length; i++) {
            if (this.bikeSlots[i] == 0) {
                slotNumber = i;
                slotFound = true;
                break;
            }
        }
        if (slotFound){
            this.bikeSlots[slotNumber] = -1;
            this.occupiedBikeSlots++;
            return slotNumber;
        }
        return -1;
    }

    public int occupyTruckSlot() {
        int slotNumber = -1;
        boolean slotFound = false;
        for (int i = 0; i < this.truckSlots.length; i++) {
            if (this.truckSlots[i] == 0) {
                slotNumber = i;
                slotFound = true;
                break;
            }
        }
        if (slotFound){
            this.truckSlots[slotNumber] = -1;
            this.occupiedTruckSlots++;
            return slotNumber;
        }
        return -1;
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
