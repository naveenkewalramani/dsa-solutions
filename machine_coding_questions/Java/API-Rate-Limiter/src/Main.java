import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        System.out.println("Welcome to rate limiter");
        Scanner sc = new Scanner(System.in);
        System.out.println("Choose the rate limiter scheme number");
        System.out.print("1. Token Bucket\n");
        int scheme = sc.nextInt();
        RateLimiter rl = null;
        if (scheme == 1) {
            System.out.println("Enter the refill rate and refill token");
            int refillRate = sc.nextInt();
            int refillToken = sc.nextInt();
            rl = new TokenBucket(refillRate, refillToken);
        }
        while(true) {
            System.out.println("Enter the API request timestamp");
            int timestamp = sc.nextInt();
            assert rl != null;
            boolean isAllowed = rl.checkIfAllowed(timestamp);
            if (isAllowed) {
                System.out.println("Request can be processed");
            }else{
                System.out.println("Request cannot be processed");
            }
        }
    }
}