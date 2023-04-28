package macosx

import ice "shylinux.com/x/icebergs"

const CONTROLS = "controls"

func init() { Index.MergeCommands(ice.Commands{CONTROLS: {Actions: CmdHashAction()}}) }
