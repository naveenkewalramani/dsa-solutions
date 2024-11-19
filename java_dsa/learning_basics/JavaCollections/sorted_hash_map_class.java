package JavaCollections;

import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.TreeMap;
import java.util.Map;

public class sorted_hash_map_class {
	public static void main(String[] args) {
		HashMap<Integer, String> map1 = new HashMap<Integer, String>();
		// Doubly linked list hash map
		LinkedHashMap<Integer, String> map2 = new LinkedHashMap<Integer, String>();
		// Sort automatically
		TreeMap<Integer, String> map3 = new TreeMap<Integer, String>();
		System.out.println("Map 1");
		testMap(map1);
		System.out.println("Map 2");
		testMap(map2);
		System.out.println("Map 3");
		testMap(map3);
	}

	public static void testMap(Map<Integer, String> map) {
		map.put(1, "a");
		map.put(3, "b");
		map.put(5, "c");
		map.put(7, "d");
		map.put(2, "e");
		map.put(4, "f");
		map.put(6, "g");
		for (Integer key : map.keySet()) {
			System.out.println("Key : " + key + " , Value : " + map.get(key));
		}
	}
}
