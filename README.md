<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
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
[![Go][golang-shield]][golang-url]


<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/codelif/hyprnotify">
    <img src="assets/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Hyprnotify</h3>

  <p align="center">
    A DBus Implementation for 'hyprctl notify'
    <br />
    <br />
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
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#compiling">Compiling</a></li>
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

![](assets/demo.gif)

Hyprnotify is a [Freedesktop.org](https://specifications.freedesktop.org/notification-spec/notification-spec-latest.html) compliant notification daemon implementing `hyprctl notify` as its backend.


<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

To get hyprnotify you can either download the binary build by github actions. Or build it locally.

### Prerequisites

 - `go` compiler
 - `libnotify` to send notifications with `notify-send` (optional)

### Compiling

1. Clone the repo and cd into it
   ```sh
   git clone https://github.com/codelif/hyprnotify.git
   cd hyprnotify
   ```
2. Build 
   ```sh
   go build ./cmd/hyprnotify
   ```
3. Run the binary
   ```sh
   ./hyprnotify
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage
Execute the daemon:
```sh
./hyprnotify
```
Now execute any of the command in another shell:\

Send a notification:
```sh
notify-send "Hello, World!"
```
Add a font-size hint:
```sh
notify-send "This is very big!" -h string:x-hyprnotify-font-size:40
```
Add a urgency hint and last for 20 seconds:
```sh
notify-send "This is serious stuff!" -u critical -t 20000 
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Implement the DBus Specification
- [x] Replace shell command invocation with IPC
- [ ] Hints Support:
    - [x] urgency
    - [x] font-size
    - [ ] color
- [ ] Add support for sound
- [ ] Fix race condition in `CloseNotification` Signal
- [ ] Scrap the Project


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the Apache-2.0 License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Harsh Sharma - [@codelif_](https://x.com/codelif_) - goharsh007@google.com

Project Link: [https://github.com/codelif/hyprnotify](https://github.com/codelif/hyprnotify)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [hyprwm community](https://github.com/hyprwm/Hyprland) for Hyprland (special thanks to [vaxry](https://github.com/vaxerski))
* [go](https://go.dev) for go

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/codelif/hyprnotify.svg?style=for-the-badge
[contributors-url]: https://github.com/codelif/hyprnotify/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/codelif/hyprnotify.svg?style=for-the-badge
[forks-url]: https://github.com/codelif/hyprnotify/network/members
[stars-shield]: https://img.shields.io/github/stars/codelif/hyprnotify.svg?style=for-the-badge
[stars-url]: https://github.com/codelif/hyprnotify/stargazers
[issues-shield]: https://img.shields.io/github/issues/codelif/hyprnotify.svg?style=for-the-badge
[issues-url]: https://github.com/codelif/hyprnotify/issues
[license-shield]: https://img.shields.io/github/license/codelif/hyprnotify.svg?style=for-the-badge
[license-url]: https://github.com/codelif/hyprnotify/blob/master/LICENSE.txt
[product-screenshot]: images/screenshot.png
[golang-shield]: https://img.shields.io/badge/Golang-00ADD8?style=for-the-badge&logo=go&logoColor=FFFFFF
[golang-url]: https://go.dev
