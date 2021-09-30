package consts

//									global
//-----------------------------------------------------------------------------

// DecodingFailed contains string to be sent if decoding fails
const DecodingFailed string = "Something wrong happened when decoding data"

//									youtube.go
//-----------------------------------------------------------------------------

const YoutubeSearchEndpoint string = "https://www.googleapis.com/youtube/v3/search?part=snippet&type=video&key="
const YoutubeFindEndpoint string = "https://www.googleapis.com/youtube/v3/videos?part=snippet&key="

const YtVideoUrl string = "https://www.youtube.com/watch?v="
