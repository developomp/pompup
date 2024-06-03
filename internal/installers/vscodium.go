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

func init() {
	register(&Installer{
		Name: "vscodium",
		Desc: "vscode but 100% FOSS",
		Setup: func() {
			// not using flatpak version due to permission issues such as lazydocker not working inside the integrated terminal
			wrapper.ParuOnce("vscodium-bin")
			wrapper.ParuOnce("vscodium-bin-marketplace")

			wrapper.WriteFile(
				wrapper.InHome(".config/VSCodium/User/settings.json"),
				vscodiumConfig,
			)

			// install vscode extensions
			out, err := exec.Command("codium", "--list-extensions").Output()
			if err != nil {
				pterm.Fatal.Println("Failed to list installed vscode extensions:", err)
			}
			// contains lowercase names
			installedExtensions := strings.Split(string(out), "\n")

			// disable vscode spell check
			// cSpell:disable

			// vscodeExtensions is a list of vscode extensions retrieved by using the following command: codium --list-extensions
			var vscodeExtensions = []string{
				// general vscde features
				"aaron-bond.better-comments",            // comment color highlighting
				"EditorConfig.EditorConfig",             // EditorConfig support
				"mikestead.dotenv",                      // .env file support
				"mkxml.vscode-filesize",                 // show file size on bottom bar
				"streetsidesoftware.code-spell-checker", // spell checking
				"WakaTime.vscode-wakatime",              // Wakatime support
				"usernamehw.errorlens",                  // show error on lines they originated from

				// general editing feature
				"christian-kohler.path-intellisense",      // autocomplete filenames
				"dbankier.vscode-quick-select",            // selecting quoted text
				"earshinov.sort-lines-by-selection",       // sort lines by selection
				"stkb.rewrap",                             // auto-wrap long text
				"plievone.vscode-template-literal-editor", // template string syntax highlight

				// general web tech
				"bradlc.vscode-tailwindcss",     // Tailwind support
				"dbaeumer.vscode-eslint",        // ESLint support
				"esbenp.prettier-vscode",        // Prettier support
				"formulahendry.auto-rename-tag", // HTML/XML tag auto rename
				"jock.svg",                      // SVG support
				"naumovs.color-highlight",       // Color stuff like #000000
				"stylelint.vscode-stylelint",    // CSS linter

				// JS & TS ecosystem
				"oven.bun-vscode",                     // bun support
				"denoland.vscode-deno",                // deno support
				"yoavbls.pretty-ts-errors",            // make TS error human readable
				"bierner.jsdoc-markdown-highlighting", // markdown syntax highlighting for jsdoc comments
				"vunguyentuan.vscode-postcss",         // postcss support

				// icons & themes
				"PKief.material-icon-theme",   // explorer icon theme
				"zhuangtongfa.material-theme", // One Dark Pro theme

				// git / github
				"eamodio.gitlens",              // basic git action as GUI
				"github.vscode-github-actions", // Github action features

				// reverse engineering
				"tintinweb.vscode-decompiler", // Decompile jar, apk, pyc, pyo, binary executable, etc
				"icsharpcode.ilspy-vscode",    // C# reverse engineering
				"ms-vscode.hexeditor",         // hex editor

				// Markdown
				"yzhang.markdown-all-in-one",                           // markdown swiss army knife
				"jeff-tian.markdown-katex",                             // markdown KaTeX support
				"bierner.markdown-mermaid",                             // markdown mermaid preview
				"bpruitt-goddard.mermaid-markdown-syntax-highlighting", // markdown mermaid color highlight

				// Rust
				"rust-lang.rust-analyzer",  // rust support
				"tamasfe.even-better-toml", // toml support
				"serayuzgur.crates",        // crates.io features (crate version, vscode palette actions, etc)

				// Python
				"charliermarsh.ruff",        // fast python formatter with drop-in parity with flake8, isort, and black
				"ms-python.black-formatter", // popular python formatter
				"ms-python.python",          // python support
				"ms-toolsai.jupyter",        // jupyter extension pack

				// Go
				"golang.go", // Go support

				// C++
				"ms-vscode.cpptools", // C/C++ support

				// C#
				"ms-dotnettools.csharp",                // C# support
				"ms-dotnettools.csdevkit",              // C# dev tools
				"ms-dotnettools.vscode-dotnet-runtime", // .NET install tools
				"ms-vscode.mono-debug",                 // mono Debugging

				// Unity
				"visualstudiotoolsforunity.vstuc", // unity-vscode integration

				// Svelte
				"svelte.svelte-vscode",            // svelte support
				"ardenivanov.svelte-intellisense", // svelte intellisense

				// Shell
				"foxundermoon.shell-format", // shell formatting
				"timonwong.shellcheck",      // shell linting

				// Docker
				"ms-azuretools.vscode-docker", // docker support

				// LaTeX
				"James-Yu.latex-workshop", // LaTeX support

				// vim
				"XadillaX.viml", // vim script support

				// CSV
				"mechatroner.rainbow-csv", // better csv coloring

				// Godot
				"alfish.godot-files",         // godot file support
				"geequlim.godot-tools",       // godot-vscode integratrion
				"neikeq.godot-csharp-vscode", // godot C# tools

				// nginx
				"ahmadalli.vscode-nginx-conf", // nginx config support

				// Linux
				"coolbear.systemd-unit-file",      // systemd unit file support
				"nico-castell.linux-desktop-file", // desktop entry file support

				// Tauri
				"tauri-apps.tauri-vscode", // tauri support

				// Terraform / OpenTofu
				"gamunu.opentofu", // OpenTofu support
			}

			// cSpell:enable
			// re-enable vscode spell check

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
