import java.util.Arrays;

public class BasicImplementation {
    public static void main(String[] args) {
        int[] arr1 = new int[]{1,2,3,4,5};
        int[] arr2 = new int[5];
        int[] arr3 = {1,2,3,4,5};
        for (int i = 0; i < arr1.length; i++){
            System.out.println(arr1[i]);
        }
        for (int i = 0; i < arr2.length; i++){
            System.out.println(arr2[i]);
        }
        for (int i = 0; i < arr3.length; i++){
            System.out.println(arr3[i]);
        }

        int[] values = new int[10];
        System.out.println(values.length);
        System.out.println(Arrays.toString(values));
        for (int i =0 ; i < 10;i++){
            values[i] = i;
        }
        System.out.println(Arrays.toString(values));
    }
}