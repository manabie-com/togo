package dao

import "github.com/patrickmn/go-cache"

func (dao *DAO) GetTokenInCache(username string) int {
	if tokenInCache, found := dao.cache.Get("auth:" + username); found {
		token := tokenInCache.(int)
		return token
	}
	return -1
}

func (dao *DAO) SetTokenToCache(token string, userID int) {
	dao.cache.Set("auth:"+token, userID, cache.DefaultExpiration)
}
