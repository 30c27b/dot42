<div align="center">
	<h3><code>dot42</code></h3>
	<sub>Created by <a href="https://github.com/30c27b">Antoine Coulon</a></sub>
</div>

---

# Description
The aim of this application is to reliably manage settings and brew packages for students from the [42 Network](https://www.42.fr/42-network/) on school computers.

# Usage

## Installing dot42

To install **dot42**, type the following command in your terminal

**note:** *it it heavily recommanded to run this command on a clean session where brew is not yet installed locally.*
```sh
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/30c27b/dot42/master/scripts/install.sh)"
```
## What does it do ?

When you run the installation script, the following actions will be executed:

- this repository will be cloned to `~/.42`
- a LaunchAgent daemon will be added to `~/Library/LaunchAgents` and run the ``login.sh`` script at each login
- the `dot42` binary will install [Homebrew](https://github.com/Homebrew/brew) in your home directory and the packages in `configs/Brewfile`
- the path to brew will be added to your `.zshrc`.

By default, the script will install Homebrew according to the following tree:
```
📦 ~
 ┣ 📜 .zshrc
 ┣ 📂 usr
 ┃ ┣ 📂 Homebrew
 ┃ ┃ ┗ 📜 (Brew repository)
 ┃ ┗ 📂 bin
 ┃   ┗ 📜 (your binaries)
 ┗ 📂 Applications
   ┗ 📜 (your applications)
```

## Settings

Settings can be modified at `~/.42/configs/config.json`.

1. Homebrew will be installed according to this tree (with `[example]` corresponding to the `"example"` entry in `config.json`)

**Installation tree**
```
📦 ~
 ┣ 📜 [zshrcPath]
 ┣ 📂 [brewPath]
 ┃ ┣ 📂 [brewName]
 ┃ ┃ ┗ 📜 [brewRepository] (is cloned here)
 ┃ ┗ 📂 bin
 ┃   ┗ 📜 (your binaries)
 ┗ 📂[brewCaskPath]
   ┗ 📜 (your applications)
```
**Correspondig settings**
```json
{
	"brewPath": "usr",
	"brewCaskPath": "Applications",
	"brewName": "Homebrew",
	"brewRepository": "https://github.com/Homebrew/brew.git",
	"zshrcPath": ".zshrc"
}
```

2. The dark theme can be set here:
```json
{
	"darkTheme": true
}
```
3. You can change [macOS defaults settings](https://ss64.com/osx/defaults.html) by adding them to the `defaults` parameter in this format:
```json
{
	"defaults": [
		{
			"domain": "Apple Global Domain",
			"key": "com.apple.swipescrolldirection",
			"valueType": "bool",
			"value": "false"
		}
	]
}
```
**note:** *To update your settings without delogging, type `~/.42/bin/dot42 load`.*


# Copyright
**dot42** is distributed under the [ISC license](/LICENSE).
