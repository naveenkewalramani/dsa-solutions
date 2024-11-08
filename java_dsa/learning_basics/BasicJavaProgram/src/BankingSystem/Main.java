package BankingSystem;

import java.util.UUID;

public class Main {
    public static void main(String[] args) {
        Customer c1 = createCustomer("John Doe");
        Customer c2 = createCustomer("Doe John");
        Customer c3 = createCustomer("John Wick");
        TransferMoney(c1, c2, 20);
        TransferMoney(c2, c3, 10);
        System.out.println(c1.account.getAccountBalance());
        System.out.println(c2.account.getAccountBalance());
        System.out.println(c3.account.getAccountBalance());
    }

    static Customer createCustomer(String name) {
        Customer customer = new Customer();
        customer.name = name;
        customer.customerID = UUID.randomUUID().toString();
        customer.account = createAccount(customer.customerID);
        return customer;
    }

    static Account createAccount(String customerID) {
        Account account = new Account();
        account.accountBalance = 0;
        account.accountStatus = "ACTIVE";
        account.customerID = customerID;
        return account;
    }

    static void TransferMoney(Customer source, Customer destination, float amount) {
        source.account.spendMoney(amount);
        destination.account.receiveMoney(amount);
    }
}
