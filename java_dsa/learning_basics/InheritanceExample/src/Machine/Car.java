package Machine;

public class Car extends Machine{
	
	public Car() {
		this.type = "Car";
	}
	public Car(String name, String model_name, int uuid) {
		this.name = name;
		this.model_name = model_name;
		this.uuid = uuid;
		this.type = "Car";
	}
	

}
