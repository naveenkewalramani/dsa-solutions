import java.util.Objects;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        System.out.println("Welcome to parking lot");
        System.out.println("Enter the number of floor, bike slots, car slots and truck slots per floor");
        Scanner scanner = new Scanner(System.in);
        int floor = scanner.nextInt();
        int bikeSlots = scanner.nextInt();
        int carSlots = scanner.nextInt();
        int truckSlots = scanner.nextInt();
        ParkingLot parkingLot = new ParkingLot(floor, bikeSlots, carSlots, truckSlots);
        System.out.println("Parking lot created");
        System.out.println(parkingLot.totalFloors);
        System.out.println(parkingLot.floors[0].totalBikeSlot);

        while (true){
            System.out.println("Enter the command");
            String command = scanner.next();
            if (Objects.equals(command, "ADD_FLOOR")){
                System.out.println("Enter the number of bike slots, car slots and truck slots");
                bikeSlots = scanner.nextInt();
                carSlots = scanner.nextInt();
                truckSlots = scanner.nextInt();
                parkingLot.addFloorToParkingLot(bikeSlots, carSlots, truckSlots);
            }else if (Objects.equals(command, "PARK_VEHICLE")){
                System.out.println("Enter the vehicle type, number, color");
                String vehicleType = scanner.next();
                String color = scanner.next();
                String vehicleNumber = scanner.next();
                Vehicle vehicle = new Vehicle(vehicleType, color, vehicleNumber);
            } else if (Objects.equals(command, "EXIT")){
                break;
            }
        }
        System.out.println("Terminating Parking lot program");
    }
}