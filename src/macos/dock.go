package macos

import ice "shylinux.com/x/icebergs"

const DOCK = "dock"

func init() { Index.MergeCommands(ice.Commands{DOCK: {Actions: CmdHashAction()}}) }
