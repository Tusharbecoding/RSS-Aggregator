package main 

import (
  "time"
  "log"
  "github.com/Tusharbecoding/RSS-Aggregator/internal/database"
  "sync"
  "context"
)

func startScraping (db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
  log.Println("Starting scraping on %v goroutines in %v duration", concurrency, timeBetweenRequest)
  
  ticker := time.NewTicker(timeBetweenRequest)

  for ; ; <-ticker.C {
    feeds, err := db.GetNextFeedsToFetch(
      context.Background(),
      int32(concurrency),
      ) 
    if err != nil {
      log.Println("Error getting feeds: ", err)
      continue
    }

    wg := &sync.WaitGroup{}
    for _, feed := range feeds {
      wg.Add(1)

      go scrapeFeed(db, wg, feed)
    }
    wg.Wait()
  }

}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
  defer wg.Done()

  _, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
  if err != nil {
    log.Println("Error marking feed as fetched: ", err)
    return
  }

  rssFeed, err := urlToFeed(feed.Url)
  if err != nil {
    log.Println("Error fetching feed: ", err)
    return
  }

  for _, item := range rssFeed.Channel.Item {
    description := sql.NullString{}
    if item.Description != "" {
      description.String = item.Description
      description.Valid = true
    }

    pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
    if err != nil {
      log.Printf("Could not parse data %v with error %v", item.PubDate, err)
      continue
    }

    _, err = db.CreatePost(context.Background(), 
      database.CreatePostParams{
        ID: uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        Title: item.Title,
        Description: description,
        PublishedAt: pubAt,
        Url: item.Link,
        FeedID: feed.ID,
      }
      )
    if err != nil {
      log.Printf("Error creating post: %v", err)
    }
  }


  log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

}
