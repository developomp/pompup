package workflows

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

// BD_PLUGINS is an array of betterdiscord plugin IDs
var BD_PLUGINS = [...]string{
	"https://tharki-god.github.io/BetterDiscordPlugins/FakeDeafen.plugin.js",
	"https://raw.githubusercontent.com/SyndiShanX/Better-Discord-Plugins/main/MemberCounter/MemberCounter.plugin.js",
	"https://tharki-god.github.io/BetterDiscordPlugins/MentionCacheFix.plugin.js",
	"https://tharki-god.github.io/BetterDiscordPlugins/PremiumScreenShare.plugin.js",
	"https://raw.githubusercontent.com/QWERTxD/BetterDiscordPlugins/main/TypingUsersAvatars/TypingUsersAvatars.plugin.js",
	"https://tharki-god.github.io/BetterDiscordPlugins/VoiceChatUtilities.plugin.js",
	"https://raw.githubusercontent.com/riolubruh/YABDP4Nitro/main/YABDP4Nitro.plugin.js",
	"https://betterdiscord.app/Download?id=509", // https://betterdiscord.app/plugin/BetterBotTags
	"https://betterdiscord.app/Download?id=61",  // https://betterdiscord.app/plugin/BetterFriendList
	"https://betterdiscord.app/Download?id=62",  // https://betterdiscord.app/plugin/BetterNsfwTag
	"https://betterdiscord.app/Download?id=63",  // https://betterdiscord.app/plugin/BetterSearchPage
	"https://betterdiscord.app/Download?id=28",  // https://betterdiscord.app/plugin/BlurNSFW
	"https://betterdiscord.app/Download?id=228", // https://betterdiscord.app/plugin/CallTimeCounter
	"https://betterdiscord.app/Download?id=64",  // https://betterdiscord.app/plugin/CharCounter
	"https://betterdiscord.app/Download?id=67",  // https://betterdiscord.app/plugin/CompleteTimestamps
	"https://betterdiscord.app/Download?id=176", // https://betterdiscord.app/plugin/Copier
	"https://betterdiscord.app/Download?id=351", // https://betterdiscord.app/plugin/DisableStickerSuggestions
	"https://betterdiscord.app/Download?id=186", // https://betterdiscord.app/plugin/DoNotTrack
	"https://betterdiscord.app/Download?id=878", // https://betterdiscord.app/plugin/enhancecodeblocks
	"https://betterdiscord.app/Download?id=356", // https://betterdiscord.app/plugin/Hide%20Chat%20Icons
	"https://betterdiscord.app/Download?id=83",  // https://betterdiscord.app/plugin/ImageUtilities
	"https://betterdiscord.app/Download?id=598", // https://betterdiscord.app/plugin/InMyVoice
	"https://betterdiscord.app/Download?id=295", // https://betterdiscord.app/plugin/InvisibleTyping
	"https://betterdiscord.app/Download?id=85",  // https://betterdiscord.app/plugin/LastMessageDate
	"https://betterdiscord.app/Download?id=29",  // https://betterdiscord.app/plugin/PermissionsViewer
	"https://betterdiscord.app/Download?id=421", // https://betterdiscord.app/plugin/PinIcon
	"https://betterdiscord.app/Download?id=158", // https://betterdiscord.app/plugin/PlatformIndicators
	"https://betterdiscord.app/Download?id=518", // https://betterdiscord.app/plugin/ProProfile
	"https://betterdiscord.app/Download?id=179", // https://betterdiscord.app/plugin/RedditMentions
	"https://betterdiscord.app/Download?id=693", // https://betterdiscord.app/plugin/RelativeTimestamps
	"https://betterdiscord.app/Download?id=699", // https://betterdiscord.app/plugin/ReplaceTimestamps
	"https://betterdiscord.app/Download?id=97",  // https://betterdiscord.app/plugin/RevealAllSpoilers
	"https://betterdiscord.app/Download?id=190", // https://betterdiscord.app/plugin/RoleMembers
	"https://betterdiscord.app/Download?id=599", // https://betterdiscord.app/plugin/RoleMentionIcons
	"https://betterdiscord.app/Download?id=159", // https://betterdiscord.app/plugin/ShowAllActivities
	"https://betterdiscord.app/Download?id=60",  // https://betterdiscord.app/plugin/ShowBadgesInChat
	"https://betterdiscord.app/Download?id=291", // https://betterdiscord.app/plugin/ShowConnections
	"https://betterdiscord.app/Download?id=306", // https://betterdiscord.app/plugin/ShutUpClyde
	"https://betterdiscord.app/Download?id=104", // https://betterdiscord.app/plugin/SpellCheck
	"https://betterdiscord.app/Download?id=98",  // https://betterdiscord.app/plugin/SplitLargeMessages
	"https://betterdiscord.app/Download?id=162", // https://betterdiscord.app/plugin/StaffTag
	"https://betterdiscord.app/Download?id=8",   // https://betterdiscord.app/plugin/SuppressReplyMentions
	"https://betterdiscord.app/Download?id=196", // https://betterdiscord.app/plugin/TypingIndicator
	"https://betterdiscord.app/Download?id=593", // https://betterdiscord.app/plugin/TypingUsersPopouts
	"https://betterdiscord.app/Download?id=381", // https://betterdiscord.app/plugin/UserVoiceShow
	"https://betterdiscord.app/Download?id=160", // https://betterdiscord.app/plugin/VoiceUsersCounter
}

func init() {
	register(&Workflow{
		Name: "Discord",
		Desc: "popular revolt alternative",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Flatpak("com.discordapp.Discord")
			install.Paru("betterdiscordctl-git") // betterdiscord installer

			for _, url := range BD_PLUGINS {
				pterm.Info.Printfln("Installing BetterDiscord plugin: %s", url)

				dirname, err := os.UserHomeDir()
				if err != nil {
					pterm.Fatal.Println(err)
				}

				// assumes that plugins is located in "~/.var/app/com.discordapp.Discord/config/BetterDiscord/plugins" because I'm using flatpak
				err = exec.Command(
					"wget",
					"--content-disposition", // use HTTP response header to decide the file name (so it's named to a proper js file)
					"--no-clobber",          // skip if file already exists
					"-P",                    // save to...
					fmt.Sprintf("%s/.var/app/com.discordapp.Discord/config/BetterDiscord/plugins", dirname),
					url,
				).Run()

				if err != nil {
					pterm.Fatal.Println("Error while downloading", url)
				}
			}

			exec.Command("betterdiscordctl", "-i", "flatpak", "install").Run()
		},
	})
}
