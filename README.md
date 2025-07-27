# tinyrdev

`tinyrdev` is a lightweight, dependency-free CLI tool written in Go that scaffolds an R package structure. It allows full customization of the `DESCRIPTION` file via command-line flags, making it ideal for R developers who want to quickly bootstrap new packages.

---

## Features

- Creates a standard R package directory structure:
- Fully configurable `DESCRIPTION` file fields via CLI flags
- Configure image for build
- Configure Posit Public Package Manager URL for reproducibility
- Include github action
- No external dependencies — uses only Go's standard library

## Usage

| Flag           | Default Value                                                   | Description                                      |
| -------------- | --------------------------------------------------------------- | ------------------------------------------------ |
| `-name`        | `""`                                                            | Name of the R package (**required**)             |
| `-title`       | `"What the Package Does (One Line, Title Case)"`                | Title of the package                             |
| `-version`     | `"0.1.0"`                                                       | Version of the package                           |
| `-author`      | `"Your Name"`                                                   | Author of the package                            |
| `-maintainer`  | `"Your Name"`                                                   | Maintainer of the package                        |
| `-email`       | `"your@email.com"`                                              | Maintainer email                                 |
| `-description` | `"A short description of the package."`                         | Description of the package                       |
| `-license`     | `"MIT"`                                                         | License type                                     |
| `-dir`         | `"."`                                                           | Target directory to create the package in        |
| `-createdir`   | `true`                                                          | Should a folder with the package name be created |
| `-image`       | `"ghcr.io/rocker-org/devcontainer/r-ver:4"`                     | Devcontainer image                               |
| `-pppm`        | `"https://packagemanager.posit.co/cran/__linux__/noble/latest"` | Posit Public Package Manager URL                 |
| `-v`           | `false`                                                         | Print version                                    |


## Example

```
./tinyrdev-<version>-<GOOS>-<GOARCH> \
  -name=myStatsPkg \
  -title="Statistical Tools for Data Analysis" \
  -version=0.2.1 \
  -author="Jane Doe" \
  -maintainer="Jane Doe" \
  -email="jane.doe@example.com" \
  -description="Provides statistical functions for regression and hypothesis testing." \
  -license="MIT" \
  -dir=~/projects
```

## Acknowledgments

Microsoft 365 Copilot assisted throughout the development of this project. Its contributions helped improve code quality, automate repetitive tasks, and accelerate problem-solving.