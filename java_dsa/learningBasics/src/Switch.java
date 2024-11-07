import java.util.Scanner;

public class Switch {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println("Enter a number");
        int number = sc.nextInt();
        switch (number % 2) {
            case 0:
                System.out.println("Number is Even");
                break;
            case 1:
                System.out.println("Number is Odd");
                break;
        }
    }
}
