package macosx

import ice "shylinux.com/x/icebergs"

const SETTINGS = "settings"

func init() { Index.MergeCommands(ice.Commands{SETTINGS: {Actions: CmdHashAction()}}) }
