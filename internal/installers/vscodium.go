package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/home/.config/VSCodium/User/settings.json
var vscodiumConfig []byte

//go:embed assets/home/.config/VSCodium/product.json
var vscodiumProduct []byte

// vscodeExtensions is a list of vscode extensions retrieved by using the following command: codium --list-extensions
var vscodeExtensions = []string{
	// General
	"mikestead.dotenv",
	"mkxml.vscode-filesize",
	"WakaTime.vscode-wakatime",

	// Editor
	"dbaeumer.vscode-eslint",
	"usernamehw.errorlens",
	"earshinov.sort-lines-by-selection",
	"EditorConfig.EditorConfig",
	"dbankier.vscode-quick-select",
	"hajdaini.select-until-pattern",
	"esbenp.prettier-vscode",
	"aaron-bond.better-comments",
	"streetsidesoftware.code-spell-checker",
	"stkb.rewrap",

	// icons & themes
	"PKief.material-icon-theme",
	"zhuangtongfa.material-theme", // One Dark Pro

	// Git
	"eamodio.gitlens",
	"mhutchie.git-graph",
	"vivaxy.vscode-conventional-commits",

	// Design
	"jock.svg",
	"formulahendry.auto-rename-tag",
	"bradlc.vscode-tailwindcss",
	"styled-components.vscode-styled-components",
	"figma.figma-vscode-extension",
	"naumovs.color-highlight",

	// Reverse Engineering
	"icsharpcode.ilspy-vscode",
	"ms-vscode.hexeditor",
	"tintinweb.vscode-decompiler",

	// Markdown
	"jeff-tian.markdown-katex",
	"bierner.markdown-mermaid",
	"bpruitt-goddard.mermaid-markdown-syntax-highlighting",
	"sebsojeda.vscode-svx",
	"unifiedjs.vscode-mdx",
	"yzhang.markdown-all-in-one",

	// JS & TS
	"denoland.vscode-deno",
	"DigitalBrainstem.javascript-ejs-support",
	"plievone.vscode-template-literal-editor",
	"dsznajder.es7-react-js-snippets",
	"yoavbls.pretty-ts-errors",
	"bierner.jsdoc-markdown-highlighting",

	// Rust
	"eww-yuck.yuck",
	"tamasfe.even-better-toml",
	"bungcip.better-toml",
	"rust-lang.rust-analyzer",
	"serayuzgur.crates",

	// Python
	"ms-python.isort",
	"ms-python.python",
	"ms-python.vscode-pylance",
	"ms-toolsai.jupyter-keymap",
	"ms-toolsai.jupyter-renderers",
	"ms-toolsai.jupyter",
	"ms-toolsai.vscode-jupyter-cell-tags",
	"ms-toolsai.vscode-jupyter-slideshow",

	// Go
	"golang.go",

	// C++
	"ms-vscode.cpptools",

	// C#
	"ms-dotnettools.csharp",
	"ms-dotnettools.vscode-dotnet-runtime",
	"ms-vscode.mono-debug",

	// Svelte
	"svelte.svelte-vscode",

	// Shell
	"foxundermoon.shell-format",

	// Docker
	"ms-azuretools.vscode-docker",

	// LaTeX
	"James-Yu.latex-workshop",

	// vim
	"XadillaX.viml",

	// CSV
	"mechatroner.rainbow-csv",

	// Godot
	"geequlim.godot-tools",
	"Razoric.gdscript-toolkit-formatter",
	"neikeq.godot-csharp-vscode",

	// nginx
	"ahmadalli.vscode-nginx-conf",

	// XML
	"redhat.vscode-xml",

	// GraphQL
	"GraphQL.vscode-graphql-syntax",
	"GraphQL.vscode-graphql",

	// Linux
	"coolbear.systemd-unit-file",
	"nico-castell.linux-desktop-file",
}

func init() {
	register(&Installer{
		Name: "vscodium",
		Desc: "vscode but 100% FOSS",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			// not using flatpak version due to permission issues such as lazydocker not working inside the integrated terminal
			wrapper.ParuOnce("vscodium-bin")

			restoreVscodeSettings()
			enableVscodeExtensionStore()

			for _, extensionName := range vscodeExtensions {
				installVscodeExtension(extensionName)
			}
		},
	})
}

func restoreVscodeSettings() {
	err := wrapper.WriteFile(
		wrapper.InHome(".config/VSCodium/User/settings.json"),
		vscodiumConfig,
	)
	if err != nil {
		pterm.Fatal.Println("Failed to restore vscodium settings:", err)
	}
}

func enableVscodeExtensionStore() {
	err := wrapper.WriteFile(
		wrapper.InHome(".config/VSCodium/product.json"),
		vscodiumProduct,
	)
	if err != nil {
		pterm.Fatal.Println("Failed to enable vscode extension store:", err)
	}
}

func installVscodeExtension(extensionName string) {
	err := wrapper.Run("codium", "--install-extension", extensionName, "--force")
	if err != nil {
		pterm.Fatal.Println("Failed to install vscode extension", extensionName, ":", err)
	}
}
