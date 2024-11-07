import java.util.Arrays;

public class Array {
    public static void main(String[] args) {
        int[] values = new int[10];
        System.out.println(values.length);
        System.out.println(Arrays.toString(values));
        for (int i =0 ; i < 10;i++){
            values[i] = i;
        }
        System.out.println(Arrays.toString(values));
    }


}
