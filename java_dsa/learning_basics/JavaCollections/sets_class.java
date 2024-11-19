package JavaCollections;

import java.util.Set;
import java.util.HashSet;
import java.util.LinkedHashSet;
import java.util.TreeSet;

public class sets_class {

	public static void main(String[] args) {
		// Does not maintain order, and return with no particular order
		Set<String> set1 = new HashSet<String>();
		set1.add("Dog");
		set1.add("Cat");
		set1.add("Mouse");
		set1.add("Horse");
		set1.add("Cow");
		set1.add("Dog");
		System.out.println(set1);

		// maintain order or enqueue
		Set<String> set2 = new LinkedHashSet<String>();
		set2.add("Dog");
		set2.add("Cat");
		set2.add("Mouse");
		set2.add("Horse");
		set2.add("Cow");
		set2.add("Dog");
		System.out.println(set2);

		// maintain order of alphabetical/numeric
		Set<String> set3 = new TreeSet<String>();
		set3.add("Dog");
		set3.add("Cat");
		set3.add("Mouse");
		set3.add("Horse");
		set3.add("Cow");
		set3.add("Dog");
		System.out.println(set3);

		for (String element : set1) {
			System.out.print(element);
		}
		System.out.println();
		for (String element : set2) {
			System.out.print(element);
		}
		System.out.println();
		for (String element : set3) {
			System.out.print(element);
		}
		System.out.println();

		// Set containing
		if (set1.contains("human")) {
			System.out.println("Human is an animal");
		} else {
			System.out.println("Human is not an animal");
		}

		if (set1.isEmpty()) {
			System.out.println("set1 is empty");
		} else {
			System.out.println("set1 is not empty");
		}
	}
}
