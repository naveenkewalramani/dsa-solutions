package ParameterAccessExample;


public class BaseIncrement {

    public int CounterPublic;
    protected int CounterProtected;

    BaseIncrement() {
        CounterPublic = 0;
        CounterProtected = 0;
    }

    void incrementPublic() {
        CounterPublic++;
    }

    void incrementProtected() {
        CounterProtected++;
    }

    void getCounterValues() {
        System.out.println("counterPublic: " + CounterPublic + ", counterProtected: " + CounterProtected);
    }
}
