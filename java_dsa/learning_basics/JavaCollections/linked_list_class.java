// Compare timing of ArrayList and LinkedList

// Appending in ArrayList is fast
// Getting in ArrayList is fast
package JavaCollections;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

public class linked_list_class {
	public static void main(String[] args) {
		// Declaring as List don't make any difference, initialization make difference
		List<Integer> array_list = new ArrayList<Integer>();
		List<Integer> linked_list = new LinkedList<Integer>();
		doTimingsAdd(1, array_list);
		doTimingsAdd(2, linked_list);
		doTimingsGet(array_list);
		doTimingsGet(linked_list);
	}

	private static void doTimingsAdd(Integer type, List<Integer> list) {
		long startTime = System.currentTimeMillis();
		for (int i = 0; i < 1E5; i++) {
			list.add(type);
		}
		for (int i = 0; i < 1E5; i++) {
			list.add(type);
		}
		long endTime = System.currentTimeMillis();
		System.out.println("Time Taken : " + (endTime - startTime) + " ms");
	}

	private static void doTimingsGet(List<Integer> list) {
		long startTime = System.currentTimeMillis();
		for (int i = 0; i < 1E5; i++) {
			list.get(i);
		}
		long endTime = System.currentTimeMillis();
		System.out.println("Time Taken : " + (endTime - startTime) + " ms");
	}
}
