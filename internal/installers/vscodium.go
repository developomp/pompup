package installers

import (
	_ "embed"
	"os/exec"
	"slices"
	"strings"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/home/.config/VSCodium/User/settings.json
var vscodiumConfig []byte

//go:embed assets/home/.config/VSCodium/product.json
var vscodiumProduct []byte

// disable vscode spell check
// cSpell:disable

// vscodeExtensions is a list of vscode extensions retrieved by using the following command: codium --list-extensions
var vscodeExtensions = []string{
	// General
	"mikestead.dotenv",
	"mkxml.vscode-filesize",
	"WakaTime.vscode-wakatime",

	// Code Tests

	"hbenl.vscode-test-explorer",
	"ms-vscode.test-adapter-converter",

	// Editor
	"dbaeumer.vscode-eslint",
	"christian-kohler.path-intellisense",
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

	// GitHub
	"github.vscode-github-actions",

	// HTML, CSS, Design
	"jock.svg",
	"formulahendry.auto-rename-tag",
	"bradlc.vscode-tailwindcss",
	"naumovs.color-highlight",
	"stylelint.vscode-stylelint",

	// Reverse Engineering
	"icsharpcode.ilspy-vscode",
	"ms-vscode.hexeditor",
	"tintinweb.vscode-decompiler",

	// Markdown
	"jeff-tian.markdown-katex",
	"bierner.markdown-mermaid",
	"bpruitt-goddard.mermaid-markdown-syntax-highlighting",
	"yzhang.markdown-all-in-one",

	// JS & TS Ecosystem
	"oven.bun-vscode",
	"denoland.vscode-deno",
	"idered.npm",
	"DigitalBrainstem.javascript-ejs-support",
	"plievone.vscode-template-literal-editor",
	"dsznajder.es7-react-js-snippets",
	"yoavbls.pretty-ts-errors",
	"bierner.jsdoc-markdown-highlighting",
	"vunguyentuan.vscode-postcss",

	// Rust
	"tamasfe.even-better-toml",
	"bungcip.better-toml",
	"rust-lang.rust-analyzer",
	"serayuzgur.crates",

	// Python
	"donjayamanne.python-extension-pack",
	"ms-python.black-formatter",
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
	"ms-dotnettools.csdevkit",
	"ms-dotnettools.csharp",
	"ms-dotnettools.vscode-dotnet-runtime",
	"ms-vscode.mono-debug",

	// Svelte
	"svelte.svelte-vscode",
	"ardenivanov.svelte-intellisense",

	// Dart, Flutter
	"dart-code.dart-code",
	"dart-code.flutter",

	// Shell
	"foxundermoon.shell-format",
	"timonwong.shellcheck",

	// Docker
	"ms-azuretools.vscode-docker",

	// LaTeX
	"James-Yu.latex-workshop",

	// vim
	"XadillaX.viml",

	// CSV
	"mechatroner.rainbow-csv",

	// Godot
	"alfish.godot-files",
	"geequlim.godot-tools",
	"neikeq.godot-csharp-vscode",

	// nginx
	"ahmadalli.vscode-nginx-conf",

	// XML
	"redhat.vscode-xml",

	// Linux
	"coolbear.systemd-unit-file",
	"nico-castell.linux-desktop-file",

	// Unity
	"visualstudiotoolsforunity.vstuc",

	// Tauri
	"tauri-apps.tauri-vscode",
}

// cSpell:enable
// re-enable vscode spell check

func init() {
	register(&Installer{
		Name: "vscodium",
		Desc: "vscode but 100% FOSS",
		Setup: func() {
			// not using flatpak version due to permission issues such as lazydocker not working inside the integrated terminal
			wrapper.ParuOnce("vscodium-bin")

			wrapper.WriteFile(
				wrapper.InHome(".config/VSCodium/User/settings.json"),
				vscodiumConfig,
			)

			// enable vscode extension store
			wrapper.WriteFile(
				wrapper.InHome(".config/VSCodium/product.json"),
				vscodiumProduct,
			)

			// install vscode extensions
			out, err := exec.Command("codium", "--list-extensions").Output()
			if err != nil {
				pterm.Fatal.Println("Failed to list installed vscode extensions:", err)
			}
			// contains lowercase names
			installedExtensions := strings.Split(string(out), "\n")

			for _, extensionName := range vscodeExtensions {
				if slices.Contains(installedExtensions, strings.ToLower(extensionName)) {
					continue // skip installed
				}

				pterm.Debug.Println("Installing vscode extension", extensionName)
				err := wrapper.Run("codium", "--install-extension", extensionName)
				if err != nil {
					pterm.Fatal.Println("Failed to install vscode extension", extensionName, ":", err)
				}
			}
		},
	})
}
