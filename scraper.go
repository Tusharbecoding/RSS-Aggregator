package main 

import (
  "time"
  "log"
  "github.com/Tusharbecoding/RSS-Aggregator/internal/database"
)

func startScraping (db *database.Queries, concurrecny int, timeBetweenRequest time.Duration) {
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
    for_, feed := range feeds {
      wg.Add(1)
    }
  }

}
