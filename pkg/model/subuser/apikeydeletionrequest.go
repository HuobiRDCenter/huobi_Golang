package subuser

type ApiKeyDeletionRequest struct {
	SubUid    int64  `json:"subUid"`
	AccessKey string `json:"accessKey"`
}
