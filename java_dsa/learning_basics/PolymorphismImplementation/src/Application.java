public class Application {

	public static void main(String[] args) {
		Plant p1 = new Plant("plant", true, false);
		Tree t1 = new Tree("tree", true, false);
		p1.plant_type();
		p1.state();

		t1.tree_type();
		t1.state();

		Plant p2 = t1;
		p2.plant_type();
		p2.state();
	}
}
