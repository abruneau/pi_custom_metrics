[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <!-- <a href="https://github.com/abruneau/pi_custom_metrics">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

  <h3 align="center">Raspberry Pi Custom Metrics</h3>

  <p align="center">
    Collect additional system metric of a Raspberry Pi for Datadog
    <br />
    <a href="https://github.com/abruneau/pi_custom_metrics"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/abruneau/pi_custom_metrics">View Demo</a>
    ·
    <a href="https://github.com/abruneau/pi_custom_metrics/issues">Report Bug</a>
    ·
    <a href="https://github.com/abruneau/pi_custom_metrics/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
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
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Datadog is great for monitoring infrastructure among other things, and the IoT Agent is very light.
But on small devices like a Raspberry Pi, cpu and gpu temperature, as well number of process are important to monitor.

This project collects metrics every 15 seconds and send them to the Datadog agent via [DogStatsD](https://docs.datadoghq.com/developers/dogstatsd/?tab=hostagent#pagetitle).

Collected metrics are:

* GPU Temperature
* CPU Temperature 
* Number of threads
* Number of processes

### Built With

* [DogStatsD](https://docs.datadoghq.com/developers/dogstatsd/?tab=hostagent#pagetitle)
* [GoReleaser](https://goreleaser.com/)
* [Logrus](https://github.com/sirupsen/logrus)
* [Viper](github.com/spf13/viper)



<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Installation

1. Clone the repo
   ```sh
   wget https://github.com/abruneau/pi_custom_metrics/releases/download/1.0.1/pi_custom_metrics_1.0.1_linux_armv7.deb
   ```
2. Install deb packages
   ```sh
   sudo apt install pi_custom_metrics_v1.0.0-next_linux_armv7.deb
   ```



<!-- USAGE EXAMPLES -->
## Usage

To start/stop datadog-custom-metrics:

```sh
sudo systemctl start datadog-custom-metrics
sudo systemctl stop datadog-custom-metrics
```

To enable/disable datadog-custom-metrics starting automatically on boot:

```sh
sudo systemctl enable datadog-custom-metrics
sudo systemctl disable datadog-custom-metrics
```

To reload datadog-custom-metrics:

```sh
sudo systemctl restart datadog-custom-metrics
```

To view datadog-custom-metrics logs:

```sh
journalctl -f -u datadog-custom-metrics
```


<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/abruneau/pi_custom_metrics/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Your Name - [@BruneauAntonin](https://twitter.com/BruneauAntonin) - antonin.bruneau@gmail.com

Project Link: [https://github.com/abruneau/pi_custom_metrics](https://github.com/abruneau/pi_custom_metrics)



<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

* [go-to-deb](https://github.com/alexhowarth/go-to-deb)
<!-- * []()
* []() -->





<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/abruneau/pi_custom_metrics.svg?style=for-the-badge
[contributors-url]: https://github.com/abruneau/pi_custom_metrics/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/abruneau/pi_custom_metrics.svg?style=for-the-badge
[forks-url]: https://github.com/abruneau/pi_custom_metrics/network/members
[stars-shield]: https://img.shields.io/github/stars/abruneau/pi_custom_metrics.svg?style=for-the-badge
[stars-url]: https://github.com/abruneau/pi_custom_metrics/stargazers
[issues-shield]: https://img.shields.io/github/issues/abruneau/pi_custom_metrics.svg?style=for-the-badge
[issues-url]: https://github.com/abruneau/pi_custom_metrics/issues
[license-shield]: https://img.shields.io/github/license/abruneau/pi_custom_metrics.svg?style=for-the-badge
[license-url]: https://github.com/abruneau/pi_custom_metrics/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/antoninbruneau