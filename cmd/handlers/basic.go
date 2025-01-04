package handlers

type BasicCard struct {
	Front string `json:"Front"`
	Back  string `json:"Back"`
}

func (b BasicCard) Field() {
	return
}
