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
![](https://github.com/codelif/hyprnotify/actions/workflows/go.yml/badge.svg)


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
        <li><a href="#arch-linux">Arch Linux</a></li>
        <li><a href="#other-linux-distributions">Other Linux Distributions</a></li>
        <li><a href="#building-from-source">Building from Source</a></li>
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

Hyprnotify is a [Freedesktop.org](https://specifications.freedesktop.org/notification-spec/notification-spec-latest.html) compliant notification daemon implementing `hyprctl notify` as its backend.


![](assets/demo.gif)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started
![](https://repology.org/badge/vertical-allrepos/hyprnotify.svg)

To get hyprnotify you can either download the binary build by github actions. Or build it locally.

### Arch Linux
`hyprnotify` is available on the AUR. You can install it with a AUR helper like `yay`.
```sh
yay -S hyprnotify
```

### NixOS
`hyprnotify` is available on NixOS unstable.
It is quite simple to add:
1. Add the unstable tarball:
```nix
let
  unstableTarball =
    fetchTarball
      https://github.com/NixOS/nixpkgs/archive/nixos-unstable.tar.gz;
in
# ...
```

2. Add an overlay to use unstable packages
```nix
nixpkgs.config = {
  packageOverrides = pkgs: {
    unstable = import unstableTarball {
      config = config.nixpkgs.config;
    };
  };
};
```

3. Add `unstable.hyprnotify` to your `environment.systemPackages` 

### Other Linux Distributions
You can download the release binaries directly from the [releases](https://github.com/codelif/hyprnotify/releases) page.

### Building from Source
#### Prerequisites

 - `go` compiler
 - `alsa-lib` or `libasound` for sound support
 - `libnotify` to send notifications with `notify-send` (optional)

#### Compiling

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
hyprnotify
```

### Examples

Send a notification:
```sh
notify-send "Hello, World!"
```
Add a font-size hint:
```sh
notify-send "This is very big!" -h int:x-hyprnotify-font-size:40
```
Add an urgency hint and last for 20 seconds:
```sh
notify-send "This is serious stuff!" -u critical -t 20000 
```
### Custom Hints
|          Hint           |              Example                |              Comment             |
|:-----------------------:|:-----------------------------------:|:---------------------------------|
| `x-hyprnotify-font-size`| `int:x-hyprnotify-font-size:30`     | font size for notification       |
| `x-hyprnotify-color`    | `string:x-hyprnotify-color:#ff30fa` | hex color code for notif. color  |
| `x-hyprnotify-icon`     | `int:x-hyprnotify-icon:3`           | icon identifier for notification |


#### `x-hyprnotify-icon`
| ID |   Icon | Preview |
|:--:|:-------|:-------:|
|`0` |WARNING |![WARNING](https://github.com/codelif/hyprnotify/assets/68972644/7bf5ff97-1d6a-45b0-9715-7e8d1535d866)|
|`1` |INFO    |![INFO](https://github.com/codelif/hyprnotify/assets/68972644/473e5752-42b3-44cf-bd07-bed64abc9660)|
|`2` |HINT    |![HINT](https://github.com/codelif/hyprnotify/assets/68972644/ffc6ff60-1058-4e5b-8fe3-611ba4b40206)|
|`3` |ERROR   |![ERROR](https://github.com/codelif/hyprnotify/assets/68972644/630b2979-382b-4fe3-902e-6ee3526fcfe4)|
|`4` |CONFUSED|![CONFUSED](https://github.com/codelif/hyprnotify/assets/68972644/64ae100b-dc2c-46dd-be8c-afdacb03042b)|
|`5` |OK      |![OK](https://github.com/codelif/hyprnotify/assets/68972644/2b66a258-5e07-4683-a798-fe6f47d67716)|

### Audio Playback
A notification sound is played along with a notification.

To disable this behaviour pass `--silent` flag when executing.
```sh
hyprnotify --silent
```
`--no-sound` and `-s` also works the same way.

### Note about `replace-id`:
When using `replace-id` with `notify-send`
```sh
notify-send --replace-id=10 "Hello"
```
All the notifications with IDs of more than `replace-id` will also be deleted. (11, 12, 13...) \
This is due to the inherent design of `hyprctl dismissnotify`. So, it is not fixable.\
\
Due to this, it is advisable to use it to replace only the latest notification.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Implement the DBus Specification
- [x] Replace shell command invocation with IPC
- [x] Hints Support:
    - [x] urgency
    - [x] font-size
    - [x] color
    - [x] icon
- [ ] Add support for sound
    - [x] Default sound support
    - [ ] sound hints
- [x] Fix race condition in `CloseNotification` Signal
- [ ] Scrap the Project


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

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

Harsh Sharma - [@codelif_](https://x.com/codelif_) - harsh@codelif.in

Project Link: [https://github.com/codelif/hyprnotify](https://github.com/codelif/hyprnotify)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [hyprwm community](https://github.com/hyprwm/Hyprland) for Hyprland (special thanks to [vaxry](https://github.com/vaxerski))
* [go](https://go.dev) for go
* [Freedesktop.org Desktop Notification Specification](https://specifications.freedesktop.org/notification-spec/notification-spec-latest.html)

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
