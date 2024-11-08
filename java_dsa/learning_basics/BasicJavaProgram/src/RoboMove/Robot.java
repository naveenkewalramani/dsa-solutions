package RoboMove;

import java.util.Objects;

public class Robot {
    private int x;
    private int y;
    private int z;

    public Robot(){
        this.x = 0;
        this.y = 0;
        this.z = 0;
    }

    public Robot(int x, int y, int z){
        this.x = x;
        this.y = y;
        this.z = z;
    }

    void move(String direction, int distance){
        if (!validate(direction)){
            System.out.println("Invalid direction");
        }
        if (Objects.equals(direction, "N")){
            this.y += distance;
        }else if (Objects.equals(direction, "S")){
            this.y -= distance;
        }else if (Objects.equals(direction, "E")){
            this.x += distance;
        }else if (Objects.equals(direction, "W")){
            this.x -= distance;
        }
    }

    void jump(int distance){
        this.z += distance;
    }

    void getLocation(){
        System.out.println("X: " + this.x+" , Y: " + this.y + " , Z: " + this.z);
    }

    private boolean validate(String direction){
        if (Objects.equals(direction, "N") || Objects.equals(direction, "S") || Objects.equals(direction, "E") || Objects.equals(direction, "W")){
            return true;
        }
        return false;
    }
}
