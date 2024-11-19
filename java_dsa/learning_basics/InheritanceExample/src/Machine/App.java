package Machine;

public class App {

	public static void main(String[] args) {
		Machine mach1 = new Machine();
		Machine mach2 = new Machine("Second Machine", "Normal", 182);
		Car car1 = new Car();
		Car car2 = new Car("Second Car", "Awesome", 192);

		System.out.println("Machine 1 Data");
		mach1.start();
		mach1.class_type();
		mach1.stop();

		System.out.println("Machine 2 Data");
		mach2.start();
		mach2.class_type();
		mach2.stop();

		System.out.println("Car 1 Data");
		car1.start();
		car1.class_type();
		car1.stop();

		System.out.println("Car 2 Data");
		car2.start();
		car2.class_type();
		car2.stop();
	}
}