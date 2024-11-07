from collections import defaultdict, deque

class Twitter:

    def __init__(self):
        self.follows = defaultdict(set)
        self.feed = deque()

    def postTweet(self, userId: int, tweetId: int) -> None:
        self.feed.appendleft((userId, tweetId))

    def getNewsFeed(self, userId: int) -> List[int]:
        return [tweetId for user, tweetId in self.feed if userId == user or user in self.follows[userId]][:10]

    def follow(self, followerId: int, followeeId: int) -> None:
        self.follows[followerId].add(followeeId)

    def unfollow(self, followerId: int, followeeId: int) -> None:
        self.follows[followerId].discard(followeeId)

