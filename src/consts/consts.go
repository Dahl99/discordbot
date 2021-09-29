package consts

//									global
//-----------------------------------------------------------------------------

// DecodingFailed contains string to be sent if decoding fails
const DecodingFailed string = "Something wrong happened when decoding data"

//								  handlers.go
//-----------------------------------------------------------------------------

//Help contains all the commands available
const Help string = "```Current commands are:\n\tping\n\tcard <card name>\n\tdice <die sides>\n\tinsult\n\tadvice\n\tkanye"
const MusicHelp string = "\n\nMusic commands:\n\tjoin\n\tleave\n\tplay <youtube url/query>\n\tskip\n\tstop```"

//									youtube.go
//-----------------------------------------------------------------------------

const YoutubeSearchEndpoint string = "https://www.googleapis.com/youtube/v3/search?part=snippet&type=video&key="
const YoutubeFindEndpoint string = "https://www.googleapis.com/youtube/v3/videos?part=snippet&key="

const YtVideoUrl string = "https://www.youtube.com/watch?v="
