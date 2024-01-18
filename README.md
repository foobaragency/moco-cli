# Moco CLI

A nice simple command line utility for interacting with moco time tracking.

## Installation 📦

First you need to tap the `foobaragency/foobaragency` homebrew tap.

```shell
brew tap foobaragency/foobaragency
```

Then you can install normally

```shell
brew install foobaragency/foobaragency/moco
```

## Authentication 🔑

You'll need your API key from moco in order to authenticate. It's available in your profile page under the "integrations" tab.
Once you have that key, simply execute `moco login`, enter your first name, last name, and your API key.

## Usage 📚

Moco CLI comes with a few basic commands built in.

| Command    | Description                                            |
| ---------- | ------------------------------------------------------ |
| `ls`       | List existing moco items (projects, tasks, activities) |
| `activity` | Actions for activities (new, edit, delete, etc.)       |
| `help`     | Help about any command                                 |
| `login`    | Log in to Moco                                         |
| `logout`   | Log out from Moco                                      |
| `stop`     | Stop time tracking for a given project                 |
