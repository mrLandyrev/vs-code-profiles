# Create profiles for Visual Studio Code

A little automation to create and run different sets of extensions and interface settings for VS Code
Used `--extension-dir` and `--user-data-dir` launch parameters for change directory on every startup

## Build:

- `git clone https://github.com/mrLandyrev/vs-code-profiles.git`
- `cd vs-code-profiles`
- `go build`

## Using:
`code-profiles PROIFILE_ALIAS PATH` open editor with predifined directories

`code-profiles --create` open dialog to create new profile

`code-profiles --list` show all saved profiles

`code-profiles --delete PROFILE_NAME` delete profile

`code-profiles --update PROFILE_NAME` open update dialog