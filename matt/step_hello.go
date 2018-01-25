package matt

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
)

type stepHello struct {
	Greeting string
}

func (s *stepHello) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
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
