package matt

import (
	"fmt"

	"github.com/hashicorp/packer/packer"
	"github.com/mitchellh/multistep"
)

type stepHello struct {
	Greeting string
}

func (s *stepHello) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	ui.Say(fmt.Sprintf("Hello: %s", s.Greeting))

	return multistep.ActionContinue
}

func (s *stepHello) Cleanup(state multistep.StateBag) {
	ui := state.Get("ui").(packer.Ui)
	ui.Say(fmt.Sprintf("Goodbye: %s", s.Greeting))

	_, cancelled := state.GetOk(multistep.StateCancelled)
	_, halted := state.GetOk(multistep.StateHalted)
	if !cancelled && !halted {
		return
	}

}
