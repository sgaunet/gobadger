[![Go Report Card](https://goreportcard.com/badge/github.com/sgaunet/gobadger)](https://goreportcard.com/report/github.com/sgaunet/gobadger)
[![GitHub release](https://img.shields.io/github/release/sgaunet/gobadger.svg)](https://github.com/sgaunet/gobadger/releases/latest)
![GitHub Downloads](https://img.shields.io/github/downloads/sgaunet/gobadger/total)
![Coverage](https://raw.githubusercontent.com/wiki/sgaunet/gobadger/coverage-badge.svg)
[![linter](https://github.com/sgaunet/gobadger/actions/workflows/coverage.yml/badge.svg)](https://github.com/sgaunet/gobadger/actions/workflows/coverage.yml)
[![coverage](https://github.com/sgaunet/gobadger/actions/workflows/coverage.yml/badge.svg)](https://github.com/sgaunet/gobadger/actions/workflows/coverage.yml)
[![Snapshot Build](https://github.com/sgaunet/gobadger/actions/workflows/snapshot.yml/badge.svg)](https://github.com/sgaunet/gobadger/actions/workflows/snapshot.yml)
[![Release Build](https://github.com/sgaunet/gobadger/actions/workflows/release.yml/badge.svg)](https://github.com/sgaunet/gobadger/actions/workflows/release.yml)
![License](https://img.shields.io/github/license/sgaunet/gobadger.svg)

# gobadger

Tool to generate badge (svg format).
It has been created to be used for private gitlab repositories. The badge need to be an artifact. See below on how to do.

# Usage

```
$ gobadger -h
Usage of gobadger:
  -c string
        color of badge (default "#5272B4")
  -o string
        output file name (default "badge.svg")
  -t string
        title
  -v string
        Value for the title
```

# In gitlab ci :

Create a stage in the CI to create the badges. They have to be artifacts.

```
build_badges2:
  stage: build
  image: sgaunet/gobadger:latest   # <= replace with the last tag
  script:
      # https://gitlab.com/%{project_path}/-/jobs/artifacts/main/raw/ref.svg?job=build_badges2
      - gobadger -o ref.svg -t godoc -v reference
      # https://gitlab.com/%{project_path}/-/jobs/artifacts/main/raw/badge.svg?job=build_badges2
      - gobadger -o badge.svg -t title -v value -c "#00FF00"
  artifacts:
    name: badge.svg
    paths:
      - badge.svg
      - ref.svg
    expire_in: 2 days
```

And add the badges in the project (settings > General > Badges ):

* Name: Name of the badge
* Link: Link of the button ex: https://gitlab.com/%{project_path}/-/commits/%{default_branch}
* Badge: https://gitlab.com/%{project_path}/-/jobs/artifacts/main/raw/badge.svg?job=build_badges2


# Install

## In your docker image

```
FROM sgaunet/gobadger:latest AS build

FROM ...

COPY --from=build /usr/bin/gobadger /usr/bin/gobadger
```

## Install the binary

```
curl -L -o gobadger https://github.com/sgaunet/gobadger/releases/download/v0.2.0/gobadger_0.2.0_linux_amd64
```

## Project Status

🟨 **Maintenance Mode**: This project is in maintenance mode.

While we are committed to keeping the project's dependencies up-to-date and secure, please note the following:

- New features are unlikely to be added
- Bug fixes will be addressed, but not necessarily promptly
- Security updates will be prioritized

## Issues and Bug Reports

We still encourage you to use our issue tracker for:

- 🐛 Reporting critical bugs
- 🔒 Reporting security vulnerabilities
- 🔍 Asking questions about the project

Please check existing issues before creating a new one to avoid duplicates.

## Contributions

🤝 Limited contributions are still welcome.

While we're not actively developing new features, we appreciate contributions that:

- Fix bugs
- Update dependencies
- Improve documentation
- Enhance performance or security

If you're interested in contributing, please read our [CONTRIBUTING.md](link-to-contributing-file) guide for more information on how to get started.

## Support

As this project is in maintenance mode, support may be limited. We appreciate your understanding and patience.

Thank you for your interest in our project!