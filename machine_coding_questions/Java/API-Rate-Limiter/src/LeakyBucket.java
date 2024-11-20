public class LeakyBucket extends RateLimiter{

    @Override
    public boolean checkIfAllowed(int timeStamp) {
        return timeStamp > 0;
    }
}
