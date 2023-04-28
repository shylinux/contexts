package macosx

import ice "shylinux.com/x/icebergs"

const NOTIFICATIONS = "notifications"

func init() { Index.MergeCommands(ice.Commands{NOTIFICATIONS: {Actions: CmdHashAction()}}) }
