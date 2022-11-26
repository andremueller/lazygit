package patch_panel

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var ExtractPatch = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Extract a patch into commit",
	ExtraCmdArgs: "",
	Skip:         false,
	SetupRepo: func(shell *Shell) {
		shell.
			CreateFile("test.txt", "").
			GitAdd("test.txt").
			GitAddAll().Commit("1").
			UpdateFile("test.txt", "one").
			GitAddAll().Commit("2").
			UpdateFile("test.txt", "one\n\ntwo").
			GitAddAll().Commit("3")
	},
	SetupConfig: func(cfg *config.AppConfig) {},
	Run: func(
		shell *Shell,
		input *Input,
		assert *Assert,
		keys config.KeybindingConfig,
	) {
		input.SwitchToCommitsWindow()
		input.Enter()
		input.Enter()
		input.NextItem()
		input.NextItem()
		input.NextItem()
		input.PressKeys(keys.Universal.Select)
		input.PressKeys(keys.Universal.CreatePatchOptionsMenu)
		input.PressKeys(keys.Universal.New)
	},
})
