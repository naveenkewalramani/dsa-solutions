import java.util.UUID;

public class Ticket {
    int floor;
    int slot;
    String ticketNumber;
    String vehicleNumber;

    public Ticket(int floor, int slot, String vehicleNumber) {
        this.floor = floor;
        this.slot = slot;
        this.vehicleNumber = vehicleNumber;
        this.ticketNumber = UUID.randomUUID().toString();
    }
}
