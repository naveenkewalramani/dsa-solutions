import java.util.Scanner;

public class Power {
    public static void main(String []args){
        Scanner input = new Scanner(System.in);
        System.out.println("Enter the number");
        int num1 = input.nextInt();
        System.out.println("Enter the power");
        int num2 = input.nextInt();
        long output = num1;
        for (int i = 1; i < num2; i++){
            output  *= num1;
        }
        System.out.println(output);
    }
}
