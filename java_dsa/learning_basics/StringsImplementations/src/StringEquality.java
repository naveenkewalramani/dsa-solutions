
public class StringEquality {
    public static void main(String[] args) {
       String s1 = "naveen";
       String s2  = "naveen";
       String s3 = "Naveen";

       if (s1 == s2){
           System.out.println("s1 is equal to s2");
       }else{
           System.out.println("s1 is not equal to s2");
       }

       if (s1 == s3){
           System.out.println("s1 is equal to s3");
       }else{
           System.out.println("s1 is not equal to s3");
       }

       String s4 = new String("naveen");
       if (s1 == s4){
           System.out.println("s1 is equal to s4");
       }else{
           System.out.println("s1 is not equal to s4");
       }

       String s5 = new String("naveen");
       if (s1.equals(s5)){
           System.out.println("s1 is equal to s5");
       }else{
           System.out.println("s1 is not equal to s5");
       }
    }
}