package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	statshouse "github.com/vkcom/statshouse-go"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wishlist/configs"
	wishlist3 "wishlist/internal/wishlist"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	appConfig := configs.LoadConfig()

	stathouseClient := statshouse.NewClient(
		log.Printf,
		"udp",
		appConfig.StathouseHost(),
		"",
	)

	defer stathouseClient.Close()

	stathouseClient.Metric("started", statshouse.Tags{1: "main"}).Count(1)

	config, err := pgxpool.ParseConfig(appConfig.DatabaseDsn())

	var dbpool *pgxpool.Pool

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		stathouseClient.MetricNamed(
			"pgxpoolCounter",
			statshouse.NamedTags{{"trigger", "AfterConnect"}},
		).Count(float64(dbpool.Stat().TotalConns()))

		return nil
	}

	config.BeforeClose = func(*pgx.Conn) {
		stathouseClient.MetricNamed(
			"pgxpoolCounter",
			statshouse.NamedTags{{"trigger", "AfterConnect"}},
		).Count(float64(dbpool.Stat().TotalConns()))
	}

	dbpool, err = pgxpool.NewWithConfig(ctx, config)

	fmt.Printf("\n\nMax size of pool %d\n\n", dbpool.Stat().MaxConns())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)

		stathouseClient.MetricNamed(
			"apilatencyValue",
			statshouse.NamedTags{{"route", c.FullPath()}, {"path", c.Request.URL.Path}},
		).Value(float64(latency.Microseconds()))
	})

	wishlist3.InitWishlistModule(dbpool, r)

	srv := &http.Server{
		Addr:    appConfig.Addr(),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
