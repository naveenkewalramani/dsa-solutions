import java.util.Scanner;

class Robot {
	private String name;
	private int current_x, current_y;

	// Method to set name
	public void setName(String name) {
		this.name = name;
	}

	// Method to set X coordinate
	public void setXCoordinate(Integer coordinate_x) {
		this.current_x = coordinate_x;
	}

	// Method to set Y coordinate
	public void setYCoordinate(Integer coordinate_y) {
		this.current_y = coordinate_y;
	}

	// Method to move robot to left
	public void moveLeft(Integer steps) {
		this.current_x -= steps;
	}

	// Method to move robot to right
	public void moveRight(Integer steps) {
		this.current_x += steps;
	}

	// Method to move robot to up
	public void moveUp(Integer steps) {
		this.current_y += steps;
	}

	// Method to move robot to down
	public void moveDown(Integer steps) {
		this.current_y -= steps;
	}

	// Method to show current location
	public void currentLocation() {
		System.out.println("Current X coordinate : " + this.current_x + " ,Current Y coordinate : " + this.current_y);
	}
}

public class RobotMotion {

	public static void main(String[] args) {
		Robot robot1 = new Robot();
		Scanner input = new Scanner(System.in);
		System.out.println("Hi to robot program. Please enter the robot name, current X,Y coordinates in new lines");
		robot1.setName(input.nextLine());
		robot1.setXCoordinate(input.nextInt());
		robot1.setYCoordinate(input.nextInt());
		System.out.println("Enter 1 to move left,Enter 2 to move right, Enter 3 to move up, "
				+ "Enter 4 to move down, Enter 5 to know current location, Enter 6 to exit");
		int to_do = input.nextInt();
		int steps;
		while (to_do != 6) {
			switch (to_do) {
			case (1):
				steps = input.nextInt();
				robot1.moveLeft(steps);
				break;
			case (2):
				steps = input.nextInt();
				robot1.moveRight(steps);
				break;
			case (3):
				steps = input.nextInt();
				robot1.moveUp(steps);
				break;
			case (4):
				steps = input.nextInt();
				robot1.moveDown(steps);
				break;
			case (5):
				robot1.currentLocation();
				break;

			}
			System.out.println("Next command please");
			to_do = input.nextInt();

		}
		System.out.println("I am finished with movement. ");
		robot1.currentLocation();
	}
}
