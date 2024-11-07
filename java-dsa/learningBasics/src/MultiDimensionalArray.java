import java.util.Arrays;

public class MultiDimensionalArray {
    public static void main(String[] args) {
        int[][] arr = new int[3][3];
        for (int i = 0; i < arr.length; i++) {
            for (int j = 0; j < arr[i].length; j++) {
                arr[i][j] = i + j;
            }
        }
            for (int[] ints : arr) {
            System.out.println(Arrays.toString(ints));
        }
    }
}
