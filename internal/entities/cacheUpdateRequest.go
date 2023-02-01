package entities

type UpdateCacheRequest struct {
	RequestId     string `json:"request_id"`
	CacheNewValue string `json:"cache_new_value"`
}

type UpdateCacheResponse struct {
	RequestId     string `json:"request_id"`
	CacheOldValue string `json:"cache_old_value"`
	CacheNewValue string `json:"cache_new_value"`
}
