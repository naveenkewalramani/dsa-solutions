package ParameterAccessExample;

public class Main {
    public static void main(String[] args) {
        BaseIncrement base = new BaseIncrement();
        DeriveIncrement derive = new DeriveIncrement();
        base.getCounterValues();
        derive.getCounterValues();
        base.incrementPublic();
        base.incrementProtected();
        base.incrementPublic();
        base.incrementProtected();
        base.getCounterValues();
        derive.incrementPublic(5);
        derive.incrementProtected(5);
        base.getCounterValues();
    }
}
