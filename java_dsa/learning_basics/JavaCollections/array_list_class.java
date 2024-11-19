package JavaCollections;

import java.util.ArrayList;
import java.util.Scanner;

public class array_list_class {
	public static void main(String[] args) {
		ArrayList<Integer> numbers = new ArrayList<Integer>();
		try (Scanner input = new Scanner(System.in)) {
			System.out.println("Enter the list size");
			int listSize = input.nextInt();
			for (int i = 0; i < listSize; i++) {
				numbers.add(input.nextInt());
			}
			System.out.println("Entered values are : ");
			for (int i = 0; i < listSize; i++) {
				System.out.println(numbers.get(i));
			}
		}
	}
}
