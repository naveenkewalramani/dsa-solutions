
public class StringBuilderExample {
    public static void main(String[] args) {
        StringBuilder sb = new StringBuilder();
        String[] arr = new String[]{"a", "b", "c", "d", "e", "f", "g", "h"};
        for (int i = 0; i < arr.length; i++) {
            sb.append(arr[i]);
        }
        System.out.println(sb.toString());
    }
}
