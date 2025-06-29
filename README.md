# Pokedex-CLI

**Pokedex-CLI** is a Go-based REPL app that simplifies interaction with the powerful [PokéAPI](https://pokeapi.co/).

Explore the world of Pokémon with interactive map! Browse specific locations, check for Pokémon presence, and try your luck at catching them.

## Examples

### Pokémon Inspection

```bash
Pokedex > inspect onix
Name: onix
Height: 88
Weight: 2100
Stats:  - hp: 35
  - attack: 45
  - defense: 160
  - special-attack: 30
  - special-defense: 45
  - speed: 70
Types:
  - rock
  - ground
```

### Pokédex Interaction

```bash
Pokedex > pokedex
Your Pokedex:
  - geodude
  - whiscash
  - unown
  - onix
  - zubat
```

### Location Exploration

```bash
Pokedex > explore ravaged-path-area
Exploring ravaged-path-area...
Found Pokemon:
  - zubat
  - golbat
  - psyduck
  - golduck
  - geodude
  - magikarp
  - gyarados
  - barboach
  - whiscash
```

### Map browsing

```bash
Pokedex > map
solaceon-ruins-b3f-d
solaceon-ruins-b3f-e
solaceon-ruins-b4f-a
solaceon-ruins-b4f-b
solaceon-ruins-b4f-c
solaceon-ruins-b4f-d
solaceon-ruins-b5f
sinnoh-victory-road-1f
sinnoh-victory-road-2f
sinnoh-victory-road-b1f
sinnoh-victory-road-inside-b1f
sinnoh-victory-road-inside
sinnoh-victory-road-inside-exit
ravaged-path-area
oreburgh-gate-1f
oreburgh-gate-b1f
stark-mountain-area
stark-mountain-entrance
stark-mountain-inside
sendoff-spring-area
```

## Table of contents

- [Requirements](#requirements)
- [Installation Guide](#installation-guide)
- [Usage Guide](#usage-guide)
  - [General Info](#general-info)
  - [Command List](#command-list)

## Requirements

Any unix-based OS: Linux (or WSL for Windows) / MacOS.
To build/install the app you need to have **Go** version 1.23+ on your computer.

## Installation Guide

You can either build or install Pokedex-CLI on your computer:

- to build the app, navigate to the Pokedex-CLI root folder and use `go build` command. This will compile the app to a single executable file.
- to install the app use `go install` command from within the gator root folder. This will compile and install the app globally on your system. Now it will be accessible by `pokedex-cli` command in your CLI.

## Usage Guide

### General Info

You can browse the map for specific locations and explore them for Pokémon presence. Map sections are cached for 5 minutes. Additionally, you can attempt to catch a Pokémon and inspect any Pokémon you've caught. The chance of catching a Pokémon depends on its base experience — the higher the experience, the lower the chance.

### Command List

- `help`: Displays a help message
- `map`: Displays the next 20 locations on the map
- `mapb`: Displays the previous 20 locations on the map
- `explore <location_name>`: Displays all the Pokémon located in the provided area
- `catch <pokemon_name>`: Attempt to catch a Pokémon
- `pokedex`: Shows all the Pokémon you caught
- `inspect <pokemon_name>`: Shows the name, height, weight, stats and type(s) of the Pokémon
- `exit`: Exit the Pokédex
