package BankingSystem;

import java.util.UUID;

class Account {
    float accountBalance;
    String accountNumber;
    String accountStatus;
    String customerID;

    public Account(String customerID){
        this.accountBalance = 0;
        this.accountNumber = UUID.randomUUID().toString();
        this.accountStatus = "ACTIVE";
        this.customerID = customerID;
    }

    float getAccountBalance(){
        return accountBalance;
    }

    void spendMoney(float amount){
        accountBalance -= amount;
    }

    void receiveMoney(float amount){
        accountBalance += amount;
    }

    String getAccountNumber(){
        return accountNumber;
    }

    String getAccountStatus(){
        return accountStatus;
    }

    void updateStatus(String status){
        accountStatus = status;
    }
}
