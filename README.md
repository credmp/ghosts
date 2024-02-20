<div id="top"></div>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]


<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/credmp/ghosts">
    <img src="docs/icon-ghosts-128.png" alt="Logo" width="128" height="128">
  </a>

<h3 align="center">Go Hosts file editor</h3>

  <p align="center">
    A command-line tool to easily manage you hosts file.
    <br />
    <br />
    <a href="https://github.com/credmp/ghosts">View Demo</a>
    ·
    <a href="https://github.com/credmp/ghosts/issues">Report Bug</a>
    ·
    <a href="https://github.com/credmp/ghosts/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

*Warning*: this repository is not yet feature complete, do not use. Please us [hed](https://github.com/credmp/hed) instead.

`ghosts` allows you to manipulate your hosts file from the command-line. By providing safe and easy commands you can add new hosts and aliases to your environment. This is a rewrite of the tool [hed](https://github.com/credmp/hed) in the Go programming language.

This tool was inspired by my students to whom I teach a Basic Cyber Security class. In this class we utilize [Hack The Box as a learning platform](https://www.youtube.com/watch?v=3b2Xul3gu_8&t=3592s) and most students struggle with editing the `hosts` file when they get started. To make this easier for them I wrote a tool that gives them a safe means of adding and removing hosts in this file.

The tool is to be used as an user that can get elevated rights using `sudo`.

<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [Go](https://www.go.dev/)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Installation

#### Using `go`

If you have the `go` binary installed you can simply install `ghosts` by running

```sh
go get github.com/credmp/ghosts
```

It will pull in the latest version.

#### Manual installation

1. Create your local `bin` directory
   ```sh
   mkdir ~/.local/bin
   ```
2. Download the latest binary release
   ```sh
   wget https://github.com/credmp/ghosts/releases/latest/download/ghosts -O ~/.local/bin/ghosts
   ```
3. Make it executable
   ```sh
   chmod +x ~/.local/bin/ghosts
   ```
4. Ensure the `bin` directory is in your path

   ```sh
   echo export PATH=\$PATH:~/.local/bin >> ~/.zshrc # if you use zsh
   echo export PATH=\$PATH:~/.local/bin >> ~/.bashrc # if you use bash
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage


### View the current hostsfile

`ghosts show` will color print the current hosts file.

```sh
ghosts show
```

Output:

```
#                    This is a comment
127.0.0.1	localhost	
::1	localhost	
127.0.1.1	pop-os.localdomain	pop-os
```

### Add a new entry

```sh
ghosts add example.com 127.1.1.1
```

Will add the following line to the hosts file.

```
127.1.1.1	example.com
```

### Add a subdomain

```sh
ghosts add demo.example.com
```

Will update the hosts file to add the subdomain to the parent domain as an alias

```
127.1.1.1	example.com	demo.example.com
```

### Remove a hostname

```sh
ghosts delete demo.example.com
```

If it is the primary `name` , the shortest alias will be chosen as new `name` for the host entry. If there are no aliases, the entire record is deleted.

```sh
ghosts delete 127.1.1.1
```

Will remove the entire record even if there are many aliases defined.

### Add an alias 

``` sh
ghosts alias demo.example.com arjenwiersma.nl
```

Add a non-subdomain alias to a hostname. This is useful when a host (ip) has many different hostnames. Instead of adding an entry for every unique top level domain they can be added as aliasses.

### Testing

Use the `--file` parameter to test the features of `ghosts` on a file that is not your `hosts` file.

```sh
ghosts --file test.txt add example.com 127.0.0.1
```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/credmp/ghosts/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Write your beautiful code
4. Ensure test coverage did not decrease (`cargo tarpaulin --verbose --all-features --workspace --timeout 120 --out Lcov`)
5. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
6. Push to the Branch (`git push origin feature/AmazingFeature`)
7. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the GPLv3 License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Arjen Wiersma - [@credmp](https://twitter.com/credmp) - [My website](https://www.arjenwiersma.nl/)

Project Link: [https://github.com/credmp/ghosts](https://github.com/credmp/ghosts)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* My students for showing me that editing a `hosts` file is not that easy.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/credmp/ghosts.svg?style=for-the-badge
[contributors-url]: https://github.com/credmp/ghosts/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/credmp/ghosts.svg?style=for-the-badge
[forks-url]: https://github.com/credmp/ghosts/network/members
[stars-shield]: https://img.shields.io/github/stars/credmp/ghosts.svg?style=for-the-badge
[stars-url]: https://github.com/credmp/ghosts/stargazers
[issues-shield]: https://img.shields.io/github/issues/credmp/ghosts.svg?style=for-the-badge
[issues-url]: https://github.com/credmp/ghosts/issues
[license-shield]: https://img.shields.io/github/license/credmp/ghosts.svg?style=for-the-badge
[license-url]: https://github.com/credmp/ghosts/blob/master/LICENSE.txt
[product-screenshot]: images/cast.gif

