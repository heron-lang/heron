<div align="center">

<img alt='heron' src="https://raw.githubusercontent.com/PoseidonCoder/heron/main/small_logo.png" />

![GoV1.15](https://img.shields.io/github/go-mod/go-version/PoseidonCoder/heron?style=for-the-badge)
![license: GPL-3.0](https://img.shields.io/github/license/PoseidonCoder/heron?style=for-the-badge)
![](https://img.shields.io/github/commit-activity/m/PoseidonCoder/heron.svg?style=for-the-badge)
![](https://img.shields.io/github/last-commit/PoseidonCoder/heron.svg?style=for-the-badge)
![codacy grade: B](https://img.shields.io/codacy/grade/1ee92d90d52b4b6e822345c7d5462be2?style=for-the-badge)
![contributions are welcome](https://img.shields.io/badge/contributions-welcome-orange.svg?style=for-the-badge)

# a *powerful* CSS preprocessor

</div>

## Table of Contents

*    [Advantages](#advantages)
*    [Planned features](#planned-features)
*    [Installation for Windows](#if-you-have-windows-you-can-just-download-the-installer-from-github-releaseshttpsgithubcomposeidoncoderheronreleases)
*    [Installation for other operating systems](#if-you-dont-use-windows)
*    [Installation with Go CLI](#or-if-you-have-go)
*    [Usage](#usage)

## Advantages

*    creates small and compact CSS
*    supports nested selectors
*    single-line ***and*** multi-line comments

## Planned Features

*    variables
*    mixins
*    standard library

## Installation

### If you have windows, you can just download the installer from [Github releases](https://github.com/PoseidonCoder/heron/releases)

### If you don't use windows

*    [install the raw executable from the `dist` folder](https://github.com/PoseidonCoder/heron/tree/main/dist)
*    To use the executable anywhere from the command line, [add it to your PATH](https://katiek2.github.io/path-doc/)

### Or, if you have Go

Run `go install github.com/PoseidonCoder/heron` in the command line. This will fetch the package, compile the project,
and add it to your [PATH](https://katiek2.github.io/path-doc/)

### *PATH is a system variable that contains programs which are allowed to run anywhere from the command line*

## Usage

```yaml
input.he
```

<pre><code>
ul {
    li {
        color: purple; //the text of the list items within this list will be purple
    }

    background-color: blue; //the unordered list will have a blue background
}
</code></pre>

To compile this bit of Heron code run the following in your terminal:

```yaml
command-line
```

```bash
heron input.he output.css
```

Heron will automatically create `output.css` if the output file is not found

```yaml
output.css
```

```css
ul li {
    color: purple;
}

ul {
    background-color: blue;
}
```


## Contributing to Heron *(I can't do this all by myself)*

1. Fork this repository.
2. Create a branch: `git checkout -b <branch_name>`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`
4. Push to the original branch: `git push origin <project_name>/<location>`
5. Create the pull request.

Alternatively see the GitHub documentation on [creating a pull request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request).

### *Note: this project is still a work in progress and therefore still has many flaws*
