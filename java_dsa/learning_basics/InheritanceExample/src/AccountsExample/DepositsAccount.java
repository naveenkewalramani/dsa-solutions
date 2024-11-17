package AccountsExample;

public class DepositsAccount extends Account {
    float balance;

    public DepositsAccount(String number, String customerID, String status) {
        super(number, customerID, status, "DEPOSITS");
    }

    public float getAccountBalance() {
        return balance;
    }
}
