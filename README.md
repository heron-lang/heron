<div style="text-align: center">
<img alt='heron' src="https://raw.githubusercontent.com/PoseidonCoder/heron/main/small_logo.png" />
<h1>a <em>powerful</em> CSS preprocessor</h1>

![GoV1.15](https://img.shields.io/github/go-mod/go-version/PoseidonCoder/heron?style=for-the-badge)
![license: GPL-3.0](https://img.shields.io/github/license/PoseidonCoder/heron?style=for-the-badge)
![](https://img.shields.io/github/commit-activity/m/PoseidonCoder/heron.svg?style=for-the-badge)
![](https://img.shields.io/github/last-commit/PoseidonCoder/heron.svg?style=for-the-badge)
![contributions are welcome](https://img.shields.io/badge/contributions-welcome-orange.svg?style=for-the-badge)
</div>

_____


## Features

* nested selectors

## Planned Features

* comments
* variables
* mixins
* standard library

## Installation

#### If you have windows, you can just download the installer from [Github releases](https://github.com/PoseidonCoder/heron/releases)

#### If you don't use windows:
* [install the raw executable from the `dist` folder](https://github.com/PoseidonCoder/heron/tree/main/dist)
* To use the executable anywhere from the command line, [add it to your PATH](https://katiek2.github.io/path-doc/)

#### Or, if you have Go:
Run `go install github.com/PoseidonCoder/heron` in the command line. This will fetch the package, compile the project, and add it to your [PATH](https://katiek2.github.io/path-doc/)

#### _PATH is a system variable that contains programs which are allowed to run anywhere from the command line_

## Usage

```yaml
input.he
```
```
ul {
    li {
        color: purple;
    }

    background-color: blue;
}
```

___

To compile this bit of Heron code run the following in your terminal:
```yaml
command-line
```
```bash
heron input.he output.css
```
Heron will automatically create `output.css` if the output file is not found

___

```yaml
output.css
```
```css
ul li{color:purple;}ul{background-color:blue;}
```
___

### *Note: this project is still a work in progress and therefore still has many flaws*
