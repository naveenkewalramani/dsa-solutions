public class Vehicle {
    String type;
    String registrationNumber;
    String color;
    Ticket ticket;

    public final static String vehicleTypeCar = "CAR";
    public final static String vehicleTypeBike = "BIKE";
    public final static String vehicleTypeTruck = "TRUCK";

    public Vehicle(String type, String registrationNumber, String color) {
        this.type = type;
        this.registrationNumber = registrationNumber;
        this.color = color;
    }

    public String getRegistrationNumber() {
        return this.registrationNumber;
    }
}
