package BankingSystem;


public class Main {
    public static void main(String[] args) {
        Customer c1 = createCustomer("John Doe");
        Customer c2 = createCustomer("Doe John");
        Customer c3 = createCustomer("John Wick");
        c1.getCustomerDetails();
        c2.getCustomerDetails();
        c3.getCustomerDetails();
        TransferMoney(c1, c2, 20);
        TransferMoney(c2, c3, 10);
        System.out.println("Account No: " + c1.account.getAccountNumber() + " , Account Balance: " + c1.account.getAccountBalance());
        System.out.println("Account No: " + c2.account.getAccountNumber() + " , Account Balance: " + c2.account.getAccountBalance());
        System.out.println("Account No: " + c3.account.getAccountNumber() + " , Account Balance: " + c3.account.getAccountBalance());
    }

    static Customer createCustomer(String name) {
        Customer customer = new Customer(name);
        customer.account = createAccount(customer.customerID);
        return customer;
    }

    static Account createAccount(String customerID) {
        return new Account(customerID);
    }

    static void TransferMoney(Customer source, Customer destination, float amount) {
        source.account.spendMoney(amount);
        destination.account.receiveMoney(amount);
    }
}
