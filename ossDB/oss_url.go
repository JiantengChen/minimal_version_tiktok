package oss

func GenerateCoverUrl(VideoUrl string) string {
	// VideoUrl + "?x-oss-process=video/snapshot,t_1000,m_fast"
	CoverUrl := VideoUrl + "?x-oss-process=video/snapshot,t_1000,m_fast"
	return CoverUrl
}

func GeneratePlayUrl(playName string) string {
	// https://BucketName.Endpoint/ObjectName
	Url := "https://" + MyBucket + "." + Endpoint + "/" + playName
	return Url
}