package main

import (
	"time"
	"uuid"

	"github.com/mdarkanurl/GoApp/internal/database"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	APIKey   string    `json:"api_key"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:       uuid.UUID(dbUser.ID),
		Name:     dbUser.Name,
		APIKey:   dbUser.ApiKey,
		CreateAt: dbUser.CreateAt,
		UpdateAt: dbUser.UpdateAt,
	}
}

type Feed struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	Url      string    `json:"url"`
	UserID   uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:       uuid.UUID(dbFeed.ID),
		Name:     dbFeed.Name,
		CreateAt: dbFeed.CreateAt,
		UpdateAt: dbFeed.UpdateAt,
		Url:      dbFeed.Url,
		UserID:   uuid.UUID(dbFeed.UserID),
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID       uuid.UUID `json:"id"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	UserID   uuid.UUID `json:"user_id"`
	FeedID   uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:       uuid.UUID(dbFeedFollow.ID),
		CreateAt: dbFeedFollow.CreateAt,
		UpdateAt: dbFeedFollow.UpdateAt,
		UserID:   uuid.UUID(dbFeedFollow.UserID),
		FeedID:   uuid.UUID(dbFeedFollow.FeedID),
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFoolows := []FeedFollow{}

	for _, dbFeedFollow := range dbFeedFollows {
		feedFoolows = append(feedFoolows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFoolows
}
