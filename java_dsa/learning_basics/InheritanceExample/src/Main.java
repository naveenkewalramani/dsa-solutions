import java.util.UUID;
public class Main {
    public static void main(String[] args) {
        String customerID = UUID.randomUUID().toString();
        String status = "ACTIVE";
        DepositsAccount a1 = new DepositsAccount("8001", customerID,status);
        LoanAccount a2 = new LoanAccount("8002", customerID,status);
        CreditCardAccount a3 = new CreditCardAccount("8003", customerID,status);
        System.out.println(a1.getProductType());
        System.out.println(a2.getProductType());
        System.out.println(a3.getProductType());
        a2.disburseFunds(20.0F);
        a2.repayFunds(10.0F);
        System.out.println(a2.getAccountBalance());
    }
}