package handlers

type Card struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  Params `json:"params"`
}

type Params struct {
	Note Note `json:"note"`
	// Options Options `json:"options"`
	// Tags Tags `json:"tags"`
	// Audio Audio `json:"audio"`
	// Video Video `json:"video"`
	// Picture Picture `json:"picture"`
}
type Note struct {
	DeckName  string `json:"deckName"`
	ModelName string `json:"modelName"`
	Fields    Fields `json:"fields"`
}

type Fields interface {
	Field()
}

func CreateCard(action, deckName, modelName string, fields Fields) Card {
	card := Card{}
	card.Action = action
	card.Version = 6
	card.Params = CreateParams(deckName, modelName, fields)
	return card
}

func CreateParams(deckName, modelName string, fields Fields) Params {
	params := Params{}
	params.Note = Note{}
	params.Note.DeckName = deckName
	params.Note.ModelName = modelName
	params.Note.Fields = fields
	return params
}
