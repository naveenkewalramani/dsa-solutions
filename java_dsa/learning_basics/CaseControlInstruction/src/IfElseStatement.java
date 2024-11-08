import java.util.Scanner;

public class IfElseStatement {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        System.out.println("Enter the number");
        int number = in.nextInt();
        if (number > 0){
            System.out.println(number + " is positive");
        }else if (number < 0){
            System.out.println(number + " is negative");
        }else{
            System.out.println(number + " is zero");
        }
    }
}
