package JavaCollections;

import java.util.HashMap;
import java.util.Map;

public class hash_map_class {
	// HashMap is not ordered. !Reminder
	public static void main(String[] args) {
		// <key_type, value_type>
		HashMap<Integer, String> map = new HashMap<Integer, String>();

		long startTime = System.currentTimeMillis();

		// Adding key value pair
		for (int i = 0; i < 10; i++) {
			map.put(i, Integer.toString(i));
		}

		long endTime = System.currentTimeMillis();

		System.out.println("Adding to hash time :" + (endTime - startTime) + "ms");

		// Getting key value pair
		for (int i = 0; i < 10; i++) {
			System.out.println("Key : " + i + " , Value : " + map.get(i));
		}

		// Getting value for no key
		System.out.println("Key : 12 , Value : " + map.get(12));

		// For each loop
		for (Map.Entry<Integer, String> entry : map.entrySet()) {
			int key = entry.getKey();
			String value = entry.getValue();
			System.out.println("Key : " + key + " , Value : " + value);
		}
	}

}
