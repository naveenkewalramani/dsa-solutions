public class CreditCardAccount  extends Account {
    float offeredLimit;
    float utilizedLimit;

    public CreditCardAccount(String number, String customerID, String status){
        super(number, customerID, status, "CREDIT_CARD");
    }

    public float getAccountBalance(){
        return this.offeredLimit - this.utilizedLimit;
    }
}
