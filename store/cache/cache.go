package cache

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"manabie.com/togo"
	"strconv"
	"time"
)

type Config struct {
	HostAddr     string
	Pass         string
	DB           int
	DialTimeout  time.Duration
	MaxRetries   int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

const (
	UserTaskSetting                      = "USER_TASK_SETTING"
	NoExpiration           time.Duration = 0
	GetSetDoTaskWithScript               = `
		local value = redis.call('get', KEYS[1])
		if value == false then
			error("not setting key")
		else
			redis.call('set', KEYS[1], tonumber(value) + 1)
		end
		return tonumber(value) + 1
	`

	RollbackDoTaskWithScript = `
		local reply = {}
		local value = redis.call('get', KEYS[1])
		if value == false then
			error("not setting key")
		else
			redis.call('set', KEYS[1], tonumber(value) - 1)
		end
		table.insert(reply, tonumber(value) + 1)
		return reply
	`
)

type cacheImpl struct {
	redisClient *redis.Client
}

func New(cfg Config) (togo.Cache, error) {
	option := redis.Options{
		Addr:         cfg.HostAddr,
		Password:     cfg.Pass,
		DB:           cfg.DB,
		DialTimeout:  cfg.DialTimeout,
		MaxRetries:   cfg.MaxRetries,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	client := redis.NewClient(&option)
	if _, err := client.Ping().Result(); err != nil {
		return nil, errors.New(fmt.Sprintf("Error %s when ping redis", err))
	}

	return &cacheImpl{redisClient: client}, nil
}

func (c *cacheImpl) ResetTaskForNewDay(limitSettings []togo.UserTaskLimitEntity) error {
	params := make([]interface{}, 0)
	mLimitSetting := make(map[string]interface{})
	for _, setting := range limitSettings {
		key := buildKeyFromTaskLimit(setting.UserId, setting.TaskId)
		params = append(params, key)
		params = append(params, 0)
		mLimitSetting[key] = setting.Limit
	}

	if _, err := c.redisClient.HMSet(UserTaskSetting, mLimitSetting).Result(); err != nil {
		return errors.Wrap(err, "fail to reset task limit setting")
	}
	if _, err := c.redisClient.MSet(params...).Result(); err != nil {
		return errors.Wrap(err, "fail to initialize task current status")
	}
	return nil
}

func (c *cacheImpl) SetTaskLimit(userId int, taskId int, limit int) error {
	key := buildKeyFromTaskLimit(userId, taskId)

	mLimitSetting := map[string]interface{}{key: limit}

	if _, err := c.redisClient.HMSet(UserTaskSetting, mLimitSetting).Result(); err != nil {
		return errors.Wrap(err, "fail to set task limit setting")
	}

	if _, err := c.redisClient.SetNX(key, 0, NoExpiration).Result(); err != nil {
		return errors.Wrap(err, "fail to initialize task current status")
	}

	return nil
}

func (c *cacheImpl) CheckIfCanDoTask(userId int, taskId int) (int, error) {
	key := buildKeyFromTaskLimit(userId, taskId)
	setting, err := c.redisClient.HMGet(UserTaskSetting, key).Result()
	if len(setting) == 0 || setting[0] == nil {
		return -1, errors.New("can not find setting task setting")
	}

	settingValue, err := strconv.Atoi(setting[0].(string))
	if err != nil {
		return 0, errors.Wrap(err, "setting wrong value")
	}

	res, err := c.redisClient.Eval(GetSetDoTaskWithScript, []string{key}).Result()
	if err != nil {
		return 0, errors.Wrap(err, "fail to get user current status")
	}

	numDo := int(res.(int64))

	if numDo > settingValue {
		log.Errorf("user %d do task %d for the %d times and it exceed the limit: %d", userId, taskId, numDo, settingValue)
		return numDo, errors.New("do task over the limit setting value")
	} else {
		log.Infof("user %d do task %d for the %d times. Limit: %d", userId, taskId, numDo, settingValue)
	}

	return numDo, err
}

func (c *cacheImpl) RollbackIfCheckFail(userId int, taskId int) error {
	key := buildKeyFromTaskLimit(userId, taskId)

	_, err := c.redisClient.Eval(RollbackDoTaskWithScript, []string{key}).Result()
	if err != nil {
		return errors.Wrap(err, "fail to rollback user current status")
	}
	return nil
}

func buildKeyFromTaskLimit(userId int, taskId int) string {
	return fmt.Sprintf("%dU.%dT", userId, taskId)
}
