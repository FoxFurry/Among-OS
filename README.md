<div id="top"></div>

[![LinkedIn][linkedin-shield]][linkedin-url]


<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h3 align="center">
    SOMMIP
  </h3>

  <a href="https://github.com/FoxFurry/Among-OS">
    <img src="_assets/amongos.gif" alt="Logo">
  </a>

  <p align="center">
    Lab 1
    <br>
    Isac Arthur
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#vision-of-os">Vision of OS</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#supported-commands">Supported commands</a></li>

  </ol>
</details>


<!-- ABOUT THE PROJECT -->
## About The Project

This very simple and sussy project is an implementation of SOMMIP Lab 1 written in Golang. Among OS does not support any filesystem related functionality
nor supports users and sessions. It's a simple printer with very basic user interaction

**Task description:**  
_For the task/lab "OS Simulator" you need to crate an simple application (in any programming language) which will simulate an simple Command Line Operating System. It should include short "booting part of the PC or of the OS and after that it should include "simulation" of an CLI OS with a few text commands (3-6 different commands).
This should be your individual VISION of an simple CLI OS!_

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- VISION OF OS -->

## Vision of OS
OS is an isolated infrastructure which gives user a higher level interface. This interface should manage the hardware and provide an easy to-use set of tools, which end user may require.
Different OS may have different levels of this interface, for example: Windows 10 is designed for generic user and provides GUI with very easy interaction, Arch linux on the other hand is a geek-oriented OS.
Arch provides no GUI, just a terminal and very little preinstalled packages. The level of interface which OS provides depends on target audience.

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

* Install Go 1.17 (or at least 1.15)  
  [Go Install Guide](https://golang.org/doc/install)

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/FoxFurry/Among-OS.git
   ```
2. Install the dependicies
   ```shell
    $ go mod download
    ```

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Supported commands

### Echo

Simply redirects seconds parameter to stdout

**Example:** 
```shell
> echo Sus
Sus
```

### Help

Prints help information for command specified as second parameter  
If no parameter specified - prints help for itself

**Example:**
```shell
> help echo
Echo is a command tool used for displaying lines of text or string which are passed as arguments on the command line
```
```shell
> help
Help is a simple command which show description of another command specified as parameter.
        Example:
        help echo

```

### Shutdown

Closes Among os with exit code 0 

**Example:**
```shell
> shutdown
Shutting down Among OS

Process finished with the exit code 0

```

### Uname

Prints kernel-related information  
Supports next parameters:
- _-r_ -- print kernel release version
- _-p_ -- print CPU information
- _-s_ -- print kernel name
- _-a_ -- print Bruh  
If no parameters is specified - kernel name will be shown (e.g. uname -s)

**Example:**
```shell
> uname -r
kernel release 5.69 k-among 
```
```shell
> uname -rrrraaaaaaa
kernel release 5.69 k-among Bruh 
```
```shell
> uname -rpsa
kernel release 5.69 k-among Intel(R) Core(TM) i9-6969K CPU @ 6.969Ghz Among OS Bruh 
```

<p align="right">(<a href="#top">back to top</a>)</p>


[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/arthur-isac-412a6519b/
