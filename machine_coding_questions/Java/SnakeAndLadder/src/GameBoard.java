public class GameBoard {
    private Snake[] snakes;
    private Ladder[] ladders;
    private Player[] players;

    GameBoard(int snakeCount, int ladderCount, int playerCount) {
        snakes = new Snake[snakeCount];
        ladders = new Ladder[ladderCount];
        players = new Player[playerCount];
    }
}
