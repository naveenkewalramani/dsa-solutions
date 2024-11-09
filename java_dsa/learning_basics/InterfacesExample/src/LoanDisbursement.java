public class LoanDisbursement implements Transfers{
    public void Transfer(String source, String destination, float amount){
        StringBuilder sb = new StringBuilder();
        sb.append("Loan Disbursement: Debit ")
                .append(source)
                .append(" and credit ")
                .append(destination)
                .append(" of amount ")
                .append(amount);
        System.out.println(sb.toString());
    }
}
