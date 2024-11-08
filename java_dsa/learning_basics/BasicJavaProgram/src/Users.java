class User {
    String name;
    int age;
    String gender;
    String email;
}

public class Users {
    public static void main(String[] args) {
        User u = new User();
        u.name = "naveen";
        u.age = 18;
        u.gender = "male";
        u.email = "naveen@gmail.com";
        System.out.println(u.name);
    }
}
