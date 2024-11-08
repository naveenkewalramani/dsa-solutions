package RoboMove;

public class Main {
    public static void main(String[] args) {
        Robot robot1 = new Robot();
        robot1.getLocation();
        robot1.move("N", 1);
        robot1.move("S", 2);
        robot1.move("E", 3);
        robot1.move("W", 4);
        robot1.jump(2);
        robot1.jump(-1);
        robot1.getLocation();

        Robot robot2 = new Robot(1,1,1);
        robot2.getLocation();
        robot2.move("N", 1);
        robot2.move("S", 2);
        robot2.move("E", 3);
        robot2.move("W", 4);
        robot2.jump(2);
        robot2.jump(-1);
        robot2.getLocation();
    }
}
