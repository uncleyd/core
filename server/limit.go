package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	libredis "github.com/go-redis/redis/v8"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"github.com/uncleyd/core/config"
	"github.com/uncleyd/core/logger"
	rdmodels "github.com/uncleyd/core/redis"
	"net/http"
	"strconv"
)

type Middleware mgin.Middleware

func RedisLimit() gin.HandlerFunc {
	// Define a limit rate to 4 requests per hour.
	rate, err := limiter.NewRateFromFormatted(config.Get().Limit.Num + "-S")
	if err != nil {
		logger.Sugar.Error(err)
		return nil
	}

	// Create a redis client.
	des := fmt.Sprintf("redis://%s:%s/%d", config.Get().Redis[0].Host, config.Get().Redis[0].Port, config.Get().Redis[0].Index)
	option, err := libredis.ParseURL(des)
	if err != nil {
		logger.Sugar.Error(err)
		return nil
	}
	client := libredis.NewClient(option)

	// Create a store with the redis client.
	store, err := sredis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix: "limiter",
	})
	if err != nil {
		logger.Sugar.Error(err)
		return nil
	}

	// Create a new middleware with the limiter instance.
	middleware := NewMiddleware(limiter.New(store, rate))
	return middleware
}

func NewMiddleware(limiter *limiter.Limiter) gin.HandlerFunc {
	middleware := Middleware{
		Limiter:        limiter,
		OnError:        mgin.DefaultErrorHandler,
		OnLimitReached: DefaultLimitReachedHandler,
		KeyGetter:      DefaultKeyGetter,
		ExcludedKey:    nil,
	}

	return func(ctx *gin.Context) {
		middleware.Handle(ctx)
	}
}

func DefaultKeyGetter(c *gin.Context) string {
	return c.Query("guid")
}

func DefaultLimitReachedHandler(c *gin.Context) {
	// 添加黑名单
	guid := c.Query("guid")
	bid := c.Query("bid")
	times := config.Get().Limit.Clean
	if times == 0 {
		times = 43200
	}
	opt := limiter.Options{
		IPv4Mask:           limiter.DefaultIPv4Mask,
		IPv6Mask:           limiter.DefaultIPv6Mask,
		TrustForwardHeader: false,
	}
	req := c.Request
	ip := limiter.GetIP(req, opt)

	sugar.Errorf("Too Many Requests from '%s' on '%s',bid:'%s', guid:'%s' ", ip, c.Request.URL.Path, bid, guid)
	// 加小黑屋
	rdmodels.SetBlackList(guid, times)

	c.String(http.StatusTooManyRequests, "Too Many Requests")
}

// Handle gin request.
func (middleware *Middleware) Handle(c *gin.Context) {
	if c.Query("guid") == "" {
		c.Next()
		return
	}

	if err := middleware.InBlack(c); err != nil {
		c.Abort()
		return
	}

	key := middleware.KeyGetter(c)
	if middleware.ExcludedKey != nil && middleware.ExcludedKey(key) {
		c.Next()
		return
	}

	context, err := middleware.Limiter.Get(c, key)
	if err != nil {
		middleware.OnError(c, err)
		c.Abort()
		return
	}

	c.Header("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
	c.Header("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
	c.Header("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))

	if context.Reached {
		middleware.OnLimitReached(c)
		c.Abort()
		return
	}

	c.Next()
}

func (middleware *Middleware) InBlack(ctx *gin.Context) error {
	guid := ctx.Query("guid")
	if guid == "" {
		return nil
	}

	c := GinContext{
		Context: ctx,
	}

	// 查询在不在小黑屋，在直接返回
	if rdmodels.ExistsBlackList(guid) {
		c.Error("Too Many Requests")
		return fmt.Errorf("Too Many Requests")
	}

	return nil
}
