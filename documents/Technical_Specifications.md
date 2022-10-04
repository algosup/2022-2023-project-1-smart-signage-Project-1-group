# Technical Specifications

## Stakeolders

| Person/Organization | Role              |
| ------------------- | ----------------- |
| SignAll             | Customers         |
| Tanguy HERRMANN     | Tech Consultant   |
| Théo TROUVE         | Project Manager   |
| Alexandre BOBIS     | Program Manager   |
| Romain NICOLAON     | Tech Leader       |
| Nicolas MIDA        | Software Engineer |
| Audrey TELLIEZ      | Quality Assurance |

<hr>

<details><summary>Table of contents</summary>

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

</details>

<hr>

### 1. Hardware

#### 1.1. Main device

It is designed to be hosted on a LoRa-E5 development kit which is a compact and easy to use development toolkit. It consists of a LoRa-E5 development board, an antenna (EU868), a USB cable.

#### 1.2. Modules

The main device is composed of the following modules:
- LoRa-E5 development board
- LoRa-E5 STM32WLE5JC
  - LoRa RF chip
  - Microcontroller chip
  - ARM Cortex-M4 core
  - LoRa Semtech SX126X chip.
- 28 pins

### 2. Software

#### 2.1. Language

We are going to use Go language and the framework [TinyGo](https://tinygo.org/) to develop the project. TinyGo is a Go compiler for small places. It compiles Go programs to small executables that can be run on microcontrollers, WebAssembly, and command-line tools.

#### 2.2. Output

We are going to show the status of the signage on a web page. The web page will be hosted on a server. The server will be hosted on [The Things Network](https://www.thethingsnetwork.org) (TTN) and will be accessible from anywhere.


### 3. Data

#### 3.1. Data storage

The data will be stored on the TTN server. The data will be accessible from anywhere and will be used to display the status of the signage.

#### 3.2. Data composition

The data will be composed by the following fields:
- Device ID
- Device status
- Device consumption
- Device problem

