public class Tree extends Plant {

	public Tree(String type, Boolean leaves, Boolean flowers) {
		super(type, flowers, flowers);
	}

	public void tree_type() {
		System.out.println(this.type);
	}

	public void state() {
		System.out.println("Tree is growing");
	}
}
