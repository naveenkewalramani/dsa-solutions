public class Plant {

	protected String type;
	protected Boolean has_leaves;
	protected Boolean has_flowers;

	public Plant(String type, Boolean leaves, Boolean flowers) {
		this.type = type;
		this.has_leaves = leaves;
		this.has_flowers = flowers;
	}

	public void plant_type() {
		System.out.println(this.type);
	}

	public void state() {
		System.out.println("Plant is growing");
	}
}
