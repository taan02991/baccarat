package types

// baccarat module event types
const (
	// TODO: Create your event types
	EventTypeUpdateGame    		= "updateGame"
	EventTypeRevealResult    	= "revealResult"

	// TODO: Create keys fo your events, the values will be derivided from the msg
	AttributeKeyGameID  		= "gameID"
	AttributeKeyWinner  		= "winner"
	AttributeKeyAddress  		= "address"
	AttributeKeyReward  		= "reward"
	AttributeKeyBetSide			= "betSide"
	AttributeKeyCard			= "card"
	AttributeKeyResultHash		= "resultHash"

	// TODO: Some events may not have values for that reason you want to emit that something happened.
	// AttributeValueDoubleSign = "double_sign"

	AttributeValueCategory = ModuleName
)
