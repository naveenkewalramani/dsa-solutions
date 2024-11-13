import java.util.UUID;

public class Ticket {
    private int floor;
    private int slot;
    private String ticketNumber;
    private String vehicleNumber;

    public Ticket(int floor, int slot, String vehicleNumber) {
        this.floor = floor;
        this.slot = slot;
        this.vehicleNumber = vehicleNumber;
        this.ticketNumber = UUID.randomUUID().toString();
        System.out.printf("Ticket created, Number: %s, Floor: %d, Slot: %d\n", this.ticketNumber, this.floor, this.slot);
    }

    public String getTicketNumber(){
        return this.ticketNumber;
    }
}
