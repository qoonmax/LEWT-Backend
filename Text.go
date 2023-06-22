package text

type Data struct {
	inputText  string
	ResultText string `json:"text"`
}

func NewText() *Data {
	return &Data{inputText: "", ResultText: ""}
}

func (d *Data) ResetInputText() {
	d.inputText = ""
}

func (d *Data) ResetResultText() {
	d.ResultText = ""
}

func (d *Data) RemoveLastSymbol() {
	runes := []rune(d.inputText)
	if len(runes) > 0 {
		d.inputText = string(runes[:len(runes)-1])
	}
}

func (d *Data) GetInputText() string {
	return d.inputText
}

func (d *Data) AddInputSymbol(inputSymbol string) {
	d.inputText += inputSymbol
}
