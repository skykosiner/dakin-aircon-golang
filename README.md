# What is this?
I have air con around my house with the brand Dakin and you can use a local web
api to control them so I wanted a way to control my air con locally on my
computer.

This allows me to change all the most changed settings about my air con
straight from my terminal.

# Setup
You need to create the config file ~/.config/aircon/aircon.json. In this file you
need to add the ip of the main aircon you wish control, and the aircon ip of
the aircon that conflict with yours and the aircon you went to toggle into that
state if so. If you don't have this porblem or don't want this feature you can
just remove the code looking for that before you build. (All of that can be found in `utils.go`).

Aircon config file:
```json
{
    "mainIp": "10.0.0.14",
    "conflictAirconOne": "10.0.0.9",
    "conflictAirconTwo": "10.0.0.17"
}
```
## Build
To build all you need to do is run the build script in the root of the project.
You might want to make change the path from ~/.local/bin/aircon-stuff/aircon to
~/.local/bin/aircon.

Also make sure that ~/.local/bin is added to the path of whatever shell you use.

# Usage
```bash
# Toggle between on and off
aircon toggle

# To change the temputer
aircon (any number between 18 - 30)

# To change the current fan speed
aircon fan-(night or level 1 - 5)

# To get the current status of the aircon
aircon status

# To set hot or cold
aircon (hot or cold)

# Sometimes you'll have an aircon that can cause a conflict with the aircon you
# want to contcrol. If you run aircon conflict it will turn off that aircon so
# that your aircon turn on. You need to make sure to change the ip address of
# the aircon in the conflict function before Compiling

# If you need help or forget the commands you can always run
aircon help
```
