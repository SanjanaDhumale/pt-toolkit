# System Layer

## Purpose

The System Layer is responsible for interacting with the operating system.

It executes commands such as:

- go version
- docker --version
- docker compose version
- java --version
- py --version
- git --version

The layer is reusable and independent of the CLI.

Other modules like Install, Doctor and Update can use the same functions.