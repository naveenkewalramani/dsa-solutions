import java.util.Scanner;

public class SwitchWithDefaultStatement {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println("Choose the gender between M/F/T");
        String gender = sc.nextLine();
        switch (gender) {
            case "M":
                System.out.println("Gender is male");
                break;
            case ("F"):
                System.out.println("Gender is female");
                break;
            case ("T"):
                System.out.println("Gender is transgender");
                break;
            default:
                System.out.println("Invalid gender");
        }
    }
}
