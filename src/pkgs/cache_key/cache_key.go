package cache_key

// TokenBlackListCacheKey ...
func TokenBlackListCacheKey(key string) string {
	return ":token:BlackList:" + key
}
