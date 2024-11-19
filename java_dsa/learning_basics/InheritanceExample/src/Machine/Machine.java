package Machine;

public class Machine {

	// protected means can be accessible from any child class within package
	protected String name;
	protected String type; // Important note first created as protected static but was set by child class.
	protected String model_name;
	protected int uuid;

	public Machine(String name, String model_name, int uuid) {
		this.name = name;
		this.model_name = model_name;
		this.uuid = uuid;
		this.type = "Machine";
	}

	public Machine() {
		this.type = "Machine";
	}

	public void start() {
		System.out.println("Machine name :" + this.name + " started");
	}

	public void stop() {
		System.out.println("Machine name :" + this.name + " stopped");
	}

	public void class_type() {
		System.out.println("Class type is " + this.type);
	}
}
