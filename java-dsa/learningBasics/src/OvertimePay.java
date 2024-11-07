// Write a program to calculate overtime pay of 10 employees. Overtime is paid at the rate of
// Rs 12.00 per hour foe every hour worked above 40 hours. Assume that employees do not work for fractional
// part of an hour

import java.util.Scanner;

public class OvertimePay {
    public static void main(String[] args){
        int rate = 12;
        int numberOfEmployees = 10 ;
        Scanner input =  new Scanner(System.in);
        System.out.println("Enter the number of hours worked");
        String workingHoursInStr  = input.nextLine();
        int workingHours = Integer.parseInt(workingHoursInStr);
        int overpaidAmount = 0;
        if (workingHours > 40) {
            overpaidAmount = (workingHours - 40)*rate*numberOfEmployees;
        }
        System.out.println("Overpaid amount will be: " + overpaidAmount);
    }
}
