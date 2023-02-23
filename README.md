# Sea battle

Status: In progress
Tags: Go

# Sea battle

## Rules

[https://www.cs.nmsu.edu/~bdu/TA/487/brules.htm](https://www.cs.nmsu.edu/~bdu/TA/487/brules.htm)

# Introduction

This game will be a game for the terminals.

# 📦 Component (package) Description

## Ship, damage, empty cells

String representation of these variables on the ********map********.

| Ship | S |
| --- | --- |
| Damage | X |
| Empty (represents empty shell) | E |
| Miss | M |

## 🗺️ PlayerMap

The main map of the game. The whole map will be assembled (in the game) by 2 maps. 1 map for 1 player.

| Type | Struct |
| --- | --- |

| Value members | Description | Type |
| --- | --- | --- |
| theMap | It contains the key to the row and number as column. | map[string][]string |
| width | The width of the battle map. It is 10 by default game. | const int |
| height | A height is the same as width. | const int |
| currentPlayer | Represents the player, needed to check weather the player is hitting the own map. | int |

| Functions | Description | Return type |
| --- | --- | --- |
| Constructor [init] | Creates a map according to the width and height. By default all the values inside of the map is “E”. |  |
| Stringer [interface] | It will show the current state of the map. It is going to show the all variables inside the PlayerMap for now. | string |
| GetDamage(string) | Accepts string coordinates as the parameter and damages with “X”.  |  |

## 🚢 Ships

[Battleship (game)](https://en.wikipedia.org/wiki/Battleship_(game))

The ship types

| Ship class | Occupied cells | Quantity (in game) |
| --- | --- | --- |
| Carrier | 5 | 1 |
| Battleship | 4 | 1 |
| Destroyer | 3 | 1 |
| Submarine | 3 | 2 |
| Patrol Boat | 2 | 2 |

# UML

[https://miro.com/app/board/uXjVPnd31KE=/](https://miro.com/app/board/uXjVPnd31KE=/)

# 📂 Repository

https://github.com/StarLightNova/sea-battle
