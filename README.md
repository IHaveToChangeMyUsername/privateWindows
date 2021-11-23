# private-room
Hide Private stuff, when the door opens. 

# Overview
todo

# Index

todo

# Hardware requirements
## 1. esp8266

Search for [esp8266 d1 mini][] to find where you can buy one.\
Other esp8266's will work too.

[esp8266 d1 mini]: https://duckduckgo.com/?q=buy+esp8266+d1+mini

<details>
  <summary>Picture of the one I have</summary>

TODO...
</details>

&nbsp;
## 2. Sensor

You will need a sensor that can detect, when the door opens.\
I am using a [reed switch][] (magnet sensor) but glueing something together with wires and aluminium foil will work too (I think)

[reed switch]: https://en.wikipedia.org/wiki/Reed_switch

&nbsp;
## 3. Some wires
You will need to connect the sensor to the esp8266. I strongly recommand a soldering iron here. If you find another solution, feel free to tell me. I will add your solution here.

&nbsp;
# Software requirements
Currently **Windows** and **Linux** are supported.
<details>
  <summary>For the Linux guys</summary>

  I am personally use i3 as a window manager.
  Please open an issue and tell me your window manger, so I can add support for it.
</details> 

## Flash the esp8266
todo

## Download the program for windows
Click [here][]

[here]: room/archive/refs/heads/main.zip

# Tips & Tricks
<details>
  <summary>Disable Window minimize animation to hide windows faster</summary>

  1. Open the start menu and search for *"Advanced System Settings"* and click on the first result
  2. Under Performance, click Settings
  3. Uncheck *"Animate windows when minimizing or maximizing option"*
  4. Click Ok.
</details>