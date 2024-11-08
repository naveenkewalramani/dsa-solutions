package RoboMove;

public class Main {
    public static void main(String[] args) {
        Robot robot = new Robot();
        robot.move("N", 1);
        robot.move("S", 1);
        robot.move("E", 1);
        robot.move("W", 1);
        robot.jump(1);
        robot.jump(-1);
        robot.getLocation();
    }
}
