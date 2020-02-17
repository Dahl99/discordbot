package discordbot

//									globals
//-----------------------------------------------------------------------------

//Const containing string to be sent if decoding fails
const decodingFailed string = "Something wrong happened when decoding data"

//								messageCreate.go
//-----------------------------------------------------------------------------

//Const containing the prefix needed to use bot commands
const prefix string = "+"

//Help contains all the commands available
const help string = "Current commands are:\n\tping\n\tcard <card name>\n\tdice <die sides>\n\tinsult\n\tadvice"

//									scryfall.go
//-----------------------------------------------------------------------------

//Const containing the root of the url
const scryfallBaseURL string = "https://api.scryfall.com/cards/named?fuzzy="

//Const containing string to be sent if scryfall API is unavailable
const scryfallNotAvailable string = "Scryfall API not available at the moment."

//									advice.go
//-----------------------------------------------------------------------------

//contains url to adviceslip API
const adviceSlipURL string = "https://api.adviceslip.com/advice"

//Const containing string to be sent if adviceslip API is unavailable
const adviceslipNotAvailable string = "Adviceslip API not available at the moment."

//									insults.go
//-----------------------------------------------------------------------------

//insultURL contains the url for the API generating insults
const insultURL string = "https://evilinsult.com/generate_insult.php?lang=en&type=json"

//String to be sent if Evil Insult API isn't available
const evilInsultNotAvailable string = "Evil Insult API not available at the moment. Please try again later."
