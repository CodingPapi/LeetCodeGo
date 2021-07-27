package strToInt

type Machine struct {
	State  MState
	Symbol int
	Result int
	Inputs []rune
}

type MState int

const (
	Exit MState = iota
	Start
	Space
	Symbol
	Number
)

func isSpace(r rune) bool {
	return r == ' '
}

func isSymbol(r rune) bool {
	return r == '-' || r == '+'
}

func isDiget(r rune) bool {
	return int(r) >= '0' && int(r) <= '9'
}

func (machine *Machine) next(r rune) bool {
	switch machine.State {
	case Start, Space:
		if isSpace(r) {
			machine.State = Space
			return true
		} else if isSymbol(r) {
			machine.State = Symbol
			if r == '-' {
				machine.Symbol = -1
			} else {
				machine.Symbol = 1
			}
			return true
		} else if isDiget(r) {
			machine.State = Number
			machine.Result = machine.Result*10 + int(r) - '0'
			return true
		} else {
			machine.State = Exit
			return false
		}
	case Symbol, Number:
		if isDiget(r) {
			machine.State = Number
			machine.Result = machine.Result*10 + int(r) - '0'
			return true
		} else {
			machine.State = Exit
			return false
		}
	case Exit:
		return false
	}
	return true
}

func myAtoi(s string) int {
	machine := Machine{
		State:  Start,
		Symbol: 1,
		Result: 0,
		Inputs: []rune(s),
	}
	for _, r := range machine.Inputs {
		if !machine.next(r) {
			break
		}
	}
	return machine.Result * machine.Symbol
}
