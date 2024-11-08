package BankingSystem;

class Account {
    float accountBalance;
    String accountNumber;
    String accountStatus;
    String customerID;

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
