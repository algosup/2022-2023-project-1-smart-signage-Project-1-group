# Functional Specifications

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

---

### Table of contents

- [Functional Specifications](#functional-specifications)
  - [Stakeolders](#stakeolders)
    - [Table of contents](#table-of-contents)
    - [1. Overview](#1-overview)
    - [2. Project Target](#2-project-target)
    - [3. Requirement Specifications](#3-requirement-specifications)
      - [3.1. Device Overview](#31-device-overview)
    - [4. Development Environement and Requirements](#4-development-environement-and-requirements)
    - [5. Glosary](#5-glosary)

### 1. Overview

The goal of the project is to provide to SignAll[^1] a new product for them, that will make their signage smart by being connected. The name of this project is Appsolu.
Nowadays, their existing products are not connected therefore users must be on-site to know if the signage is on, functional, or out of order. Also, users cannot switch the signage on/off remotely.

### 2. Project Target

For this product, we want to receive a message to see remotely the signage status in real time.
The main features of the project are:

- The monitoring of the signage like on or off, if their is a failure or an overheating of the product.
- A remote control of the signage like switching on/off, using a dimmable light bulb.
- A consumption reduction and compliance with environmental laws like switching off at given time, adjust intensity based on ambient lighting.

### 3. Requirement Specifications

#### 3.1. Device Overview

As written in the project target, the application should be able to provide many informations like if the signage is on, functional, or out of order and switch the signage on/off remotely.

The device will have a power switch control board, LEDs[^2], ambient light sensor, temperature sensor, inductive current sensor, current sensor

### 4. Development Environement and Requirements

- Go[^3]
  - TinyGo[^4]
- MacOSx/Windows on development
- Arduino[^5] on production environment

### 5. Glosary

[^1]: SignAll
It's a French company that has been manufacturing large luminous signage since 1962. They supply a large number of customers such as McDonald’s, Burger King, La Poste, Orange, AXA, Crédit Agricole, Total, etc. to name a few.

[^2]: LED
A light-emitting diode (LED) is a semiconductor light source that emits light when current flows through it.

[^3]: Go
Go is a compiled and concurrent programming language inspired by other languages (C and Pascal). This language was developed by Google.

[^4]: TinyGo
It brings the Go programming language to embedded systems and to the modern web by creating a new compiler. You can compile and run TinyGo programs on a wide variety of platforms, including microcontrollers, web browsers, and servers.

[^5]: Arduino
Arduino is an open-source hardware and software company, project, and user community that designs and manufactures single-board microcontrollers and microcontroller kits for building digital devices.
