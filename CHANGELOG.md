# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/), and this project adheres to [Semantic Versioning](https://semver.org/).

## [Unreleased]

### Added

* Changelog
* Lots of help dialogs
* Guard detection
* Authority detection
* Self-assignable roles
* `info` command to get information in the form of a Rich Embed about a member

### Changed

* Now use Inoxydable's icon
* Rewritten everything to use the new database wrappers

### Deprecated

### Removed

* `Love` and `Points` are to be re-introduced in a while.

### Fixed

* Members with more than one role will no longuer surpass her intellectual capabilities
* Fixed a bug where she wouldn't ask for guard if the welcome channel was set by Discord's welcome message

### Security

* Ban new members with `discord.gg` in their username
* New permission management system for internal commands

## [2.0.0] - 2018-09-01

This marks the end of the [GitLab](https://gitlab.com/NatoBoram/Go-Miiko) migration from [GitHub](https://github.com/NatoBoram/Go-Miiko) and the introduction of this changelog.

### Added

* Database wrappers. A lot of them
* More keywords for guardless members
* Presentation channel for people to introduce themselves
* Message when the guard wasn't understood
* GitLab CI

### Changed

* Changed from GitHub to GitLab
* Put everything in the same package
* "Welcome" channel is now "channel-welcome"

### Removed

* Fall in love

## Types of changes

* `Added` for new features.
* `Changed` for changes in existing functionality.
* `Deprecated` for soon-to-be removed features.
* `Removed` for now removed features.
* `Fixed` for any bug fixes.
* `Security` in case of vulnerabilities.
