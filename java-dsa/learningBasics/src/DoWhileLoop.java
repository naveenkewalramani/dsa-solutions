public class DoWhileLoop {
    public static void main(String[] args) {
        System.out.println(loop(10));
        System.out.println(loop(20));
        System.out.println(loop(-1));
    }

    public static int loop(int num){
        int counter = 0;
        do {
            counter++;
        }while(counter  < num);
        return counter;
    }
}
