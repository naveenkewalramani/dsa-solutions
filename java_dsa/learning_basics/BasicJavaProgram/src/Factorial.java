import java.util.Scanner;

public class Factorial {
    public static void main(String []args){
        Scanner input = new Scanner(System.in);
        System.out.println("Enter the number for factorial computation");
        int number = input.nextInt();
        int output = factorial(number);
        System.out.println("Output: " + output);

    }

    public static int factorial(int number){
        if (number == 0 || number == 1){
            return 1;
        }
        return number * factorial(number - 1);
    }
}
