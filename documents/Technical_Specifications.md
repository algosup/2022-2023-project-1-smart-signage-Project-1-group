# Technical Specifications

## Stakeolders

| Person/Organization | Role              |
| :-----------------: | :---------------: |
| SignAll             | Customers         |
| Tanguy HERRMANN     | Tech Consultant   |
| Théo TROUVE         | Tech Leader & Project Manager   |
| Alexandre BOBIS     | Tech Leader & Program Manager   |
| Nicolas MIDA        | Tech Leader & Software Engineer |
| Audrey TELLIEZ      | Tech Leader & Quality Assurance |

---

<details close>

<summary>Table of contents</summary><blockquote>

- [Technical Specifications](#technical-specifications)
  - [Stakeolders](#stakeolders)
  - [1. Hardware](#1-hardware)
    - [1.1. Main device](#11-main-device)
    - [1.2. Modules](#12-modules)
  - [2. Software](#2-software)
    - [2.1. Language](#21-language)
    - [2.2. Output](#22-output)
  - [3. Data](#3-data)
    - [3.1. Data storage](#31-data-storage)
    - [3.2. Data composition](#32-data-composition)
  - [4. Working](#4-working)
    - [4.1 Hardware environment](#41-hardware-environment)
    - [4.2 Software environment](#42-software-environment)
    - [4.3 Tests](#43-tests)
  - [5. Glosary](#5-glosary)

</details>

---

## 1. Hardware

### 1.1. Main device

It is designed to be hosted on a LoRa-E5 development kit[^1] which is a compact and easy to use development toolkit.

### 1.2. Modules

The main device is composed of the following modules:

- LoRa-E5 development board
![LoRa-E5 development board](https://files.seeedstudio.com/wiki/LoRa-E5_Development_Kit/202003261_preview-07.png)
- LoRa-E5 STM32WLE5JC module
![LoRa-E5 STM32WLE5JC module](https://files.seeedstudio.com/products/317990687/image/lora-e5_Preview-07.png)
- ARM Cortex-M4 core
- LoRa Semtech SX126X

Here are other materials:

- Bluepill STM32F103C8T6
![Bluepill STM32F103C8T6](https://imgs.search.brave.com/LvhMMiYgDuNs6HVsy-dNzPMn5a0rI3RyypIw2D7t3YM/rs:fit:1200:1200:1/g:ce/aHR0cHM6Ly9za3Rl/Y2h3b3Jrcy5jYS93/cC1jb250ZW50L3Vw/bG9hZHMvMjAxNy8w/Ni9TVDMyRHVpbm8t/NC5qcGc)
- ST-Link V2
![ST-Link V2](https://imgs.search.brave.com/RM9DUVUjJ4VqXGis_IfS5uZmj_i_U0ynfEg8zw14vKo/rs:fit:600:600:1/g:ce/aHR0cDovL3d3dy5q/eXNvdXJjZS5jb20v/d2ViL3VzZXJmaWxl/cy9wcm9kdWN0L0lD/L2JpZy9TVExJTktT/VE04U1RNMzIuanBn)
- LED's
- Pins, to link the modules together
![Pin's](https://imgs.search.brave.com/g9FQkFHOMjFngiFO9Iw6zDc-pZhvx1HC-Hll0iT4hZI/rs:fit:1200:1200:1/g:ce/aHR0cHM6Ly9pbWcu/c3RhdGljYmcuY29t/L2ltYWdlcy9vYXVw/bG9hZC9iYW5nZ29v/ZC9pbWFnZXMvQjIv/NjMvMjEzZTI4ZDUt/YzYwYS00MjRlLWE2/MjUtNTRmNWFhMTgy/YTM1LmpwZw)
- XY MOS Arduino Module
![XY MOS Arduino Module](https://imgs.search.brave.com/YlK9JCSdovTfClTsIybDoKx2Zx0Vx4Z1dUk72L6Dhqc/rs:fit:502:531:1/g:ce/aHR0cHM6Ly9pbWFn/ZXMudWEucHJvbS5z/dC8yMzU4OTAwNTUw/X3c2NDBfaDY0MF9t/b2R1bC10cmFuemlz/dG9yLTE1YS5qcGc)
- ZMCT103C Module
![ZMCT103C Module](https://imgs.search.brave.com/EgnBCmh1qpjRkgURmjdW1ul4y5MJCGKbVQSyrU-gtTs/rs:fit:1200:1200:1/g:ce/aHR0cHM6Ly9zMy1h/cC1zb3V0aGVhc3Qt/MS5hbWF6b25hd3Mu/Y29tL2EyLmRhdGFj/YWNpcXVlcy5jb20v/d20vTVRVNU16QS8z/NzYyMTg1NDUyLzMx/OTUwMTQxNzcuanBn)

## 2. Software

### 2.1. Language

We are going to use Go language and the framework TinyGo[^2] to develop the project. TinyGo is a Go compiler for small places. It compiles Go programs to small executables that can be run on microcontrollers, WebAssembly, and command-line tools. It is based on LLVM[^3], so it supports a wide range of microcontrollers and other platforms.

### 2.2. Output

We are going to show the status of the signage on a web page. The web page will be hosted on a server. The server will be hosted on The Things Network[^4] (TTN) and will be accessible from anywhere.

## 3. Data

### 3.1. Data storage

The data will be stored on the TTN server. The data will be accessible from anywhere and will be used to display the status of the signage.

### 3.2. Data composition

The data will be composed by the following fields:

- Device ID
- Device status
- Device consumption
- Device problem

## 4. Working

### 4.1 Hardware environment

![Bluepill STM32F103C8T6](https://kevoster.files.wordpress.com/2017/05/stm32f103c8t6-bluepill.gif?w=665&zoom=2)
![XY MOS Arduino Module](https://makershop.ie/image/catalog/p/00304/xy-mos-1.jpg)
![ZMCT103C Module](https://components101.com/sites/default/files/inline-images/ZMCT103C-Sensor-with-MCU_0.jpg)

### 4.2 Software environment

// INSERT CODE HERE

### 4.3 Tests

Here is a link to the [board test](https://github.com/algosup/2022-2023-project-1-smart-signage-Project-1-group/blob/main/documents/Board_Test.md)

## 5. Glosary

[^1]: [LoRa-E5 Development Kit](https://www.st.com/en/partner-products-and-services/lora-e5-development-kit.html#overview) consists of a LoRa-E5 Dev Board, an antenna, a USB type C cable, and a 2*AA 3V Battery Holder. It's a module that supports LoRaWAN protocol on global frequency band.

[^2]: [TinyGo](https://tinygo.org/) is a Go compiler for small places. It compiles Go programs to small executables that can be run on microcontrollers, WebAssembly, and command-line tools. It is based on LLVM, so it supports a wide range of microcontrollers and other platforms.

[^3]: [LLVM](https://en.wikipedia.org/wiki/LLVM) is a compiler infrastructure project that includes a collection of modular and reusable compiler and toolchain technologies. The name LLVM originally seans *Low Level Virtual Machine*.

[^4]: [The Things Network](https://www.thethingsnetwork.org) provides a set of open tools and a global, open network to build IoT application at low cost, featuring maximum security and ready to scale.
