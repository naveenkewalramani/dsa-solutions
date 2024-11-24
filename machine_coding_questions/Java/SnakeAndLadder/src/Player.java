public class Player {
    private String name;
    private int currentPosition;
    private int previousPosition;

    Player(String name) {
        this.name = name;
        this.currentPosition = 0;
        this.previousPosition = 0;
    }

    public void setPosition(int position, int diceValue) {
        this.previousPosition = this.currentPosition;
        this.currentPosition = position;
        displayPosition(diceValue);
    }

    public void displayPosition(int diceValue) {
        System.out.printf("%s rolled a %d and moved from %d to %d", this.name, diceValue, this.previousPosition, this.currentPosition);
    }
}
