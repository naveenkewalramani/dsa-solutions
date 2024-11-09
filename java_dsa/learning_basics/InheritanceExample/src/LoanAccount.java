public class LoanAccount extends Account {
    float disbursedAmount;
    float repaymentAmount;

    public LoanAccount(String number, String customerID, String status){
        super(number, customerID, status, "LOAN");
    }

    public void disburseFunds(float amount){
        this.disbursedAmount += amount;
    }

    public void repayFunds(float amount){
        this.repaymentAmount += amount;
    }

    public float getAccountBalance(){
        return this.disbursedAmount - this.repaymentAmount;
    }
}
