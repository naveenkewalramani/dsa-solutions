package AccountsExample;

public class Account {
    String number;
    String customerID;
    String status;
    long OpeningTimestamp;
    String productType;

    public Account(String number, String customerID, String status, String productType) {
        this.number = number;
        this.customerID = customerID;
        this.status = status;
        this.OpeningTimestamp = System.currentTimeMillis();
        this.productType = productType;
    }

    public String getProductType() {
        return this.productType;
    }

    public String getAccountNumber() {
        return this.number;
    }
}
