import java.util.ArrayList;
import java.util.Scanner;

public class GenericArrayList {
	public static void main(String[] args) {

		ArrayList<String> strings = new ArrayList<String>();
		ArrayList<Integer> numbers = new ArrayList<Integer>();
		try (Scanner input = new Scanner(System.in)) {
			for (int i = 0; i < 5; i++) {
				strings.add(input.nextLine());
			}
			for (int i = 0; i < 5; i++) {
				numbers.add(input.nextInt());
			}
		}
		System.out.println("Array List Strings");
		for (int i = 0; i < 5; i++) {
			System.out.println(strings.get(i));
		}
		System.out.println("Array List Integer");
		for (int i = 0; i < 5; i++) {
			System.out.println(numbers.get(i));
		}
	}
}
