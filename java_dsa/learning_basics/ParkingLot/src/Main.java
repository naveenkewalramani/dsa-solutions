import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        System.out.println("Welcome to parking lot");
        System.out.println("Enter the number of floor and slots per floor");
        Scanner scanner = new Scanner(System.in);
        int floor = scanner.nextInt();
        int slots = scanner.nextInt();

        ParkingLot parkingLot = new ParkingLot(floor, slots);

    }
}