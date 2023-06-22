package keyboard

import (
	text "LEWT_Backend"
	"LEWT_Backend/translateService"
	hook "github.com/robotn/gohook"
	"time"
)

func Listen(data *text.Data) {
	evChan := hook.Start()
	defer hook.End()

	translateTimer := time.NewTimer(2 * time.Second)
	reset := make(chan bool)

	go func() {
		for {
			select {
			case <-translateTimer.C:
				translateInputText(data)
			case <-reset:
				translateTimer.Reset(2 * time.Second)
			}
		}
	}()

	for ev := range evChan {
		if ev.Kind == hook.KeyDown {
			reset <- true

			if ev.Keychar == Space {
				translateInputText(data)
			} else if ev.Keychar == Dot || ev.Keychar == QuestionMark {
				addInputSymbol(data, ev.Keychar)
				translateInputText(data)
				data.ResetInputText()
				continue
			} else if ev.Keychar == Enter {
				translateInputText(data)
				data.ResetInputText()
				continue
			}
			addInputSymbol(data, ev.Keychar)
		}
	}
}

func translateInputText(data *text.Data) {
	translateService.Translate(data)
}

func addInputSymbol(data *text.Data, inputRune rune) {
	runes := []rune(data.GetInputText())

	if inputRune == Backspace {
		data.RemoveLastSymbol()
	} else if len(runes) == 0 && inputRune == Space {
		return
	} else {
		data.AddInputSymbol(string(inputRune))
	}
}
