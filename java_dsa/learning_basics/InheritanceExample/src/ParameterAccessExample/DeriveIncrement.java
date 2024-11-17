package ParameterAccessExample;

public class DeriveIncrement extends BaseIncrement {

    void incrementPublic(int num){
        System.out.println(num + " " + CounterPublic);
        this.CounterPublic += num;
    }

    void incrementProtected(int num) {
        this.CounterProtected += num;
    }
    void getCounterValues() {
        System.out.println("counterPublic: " + this.CounterPublic + ", counterProtected: " + this.CounterProtected);
    }
}
