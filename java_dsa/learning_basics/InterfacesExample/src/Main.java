public class Main {
    public static void main(String[] args) {
        LoanDisbursement loanDisbursement = new LoanDisbursement();
        LoanRepayment loanRepayment = new LoanRepayment();
        loanDisbursement.Transfer("8001", "8002", 200.33F);
        loanRepayment.Transfer("8002", "8001", 100.33F);
        loanRepayment.Transfer("8002", "8001", 100.00F);
    }
}