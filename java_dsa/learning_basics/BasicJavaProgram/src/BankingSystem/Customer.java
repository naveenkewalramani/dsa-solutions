package BankingSystem;

import java.util.UUID;

class Customer {
     Account account;
     String name;
     String customerID;

     public Customer(String name) {
          this.name = name;
          this.customerID =  UUID.randomUUID().toString();;
     }

     void getCustomerDetails() {
          System.out.println("CustomerID: " +  this.customerID  + " , Name: " + this.name  + " , AccountID: " +  this.account.getAccountNumber());
     }
}

