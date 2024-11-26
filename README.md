# What is this?
I have air con around my house with the brand Dakin and you can use a local web
api to control them so I wanted a way to control my air con locally on my
computer.

This allows me to change all the most changed settings about my air con
straight from my terminal.

# Setup
You need to create a config file at `~/.config/aircon/aircon.json`. In this file you
need to add the ip of the air con you wish control. You can extend the code to
work with more then one units if you want, I'm more then happy to accept the PR
:).
Air con config file:
```json
{
    "aircon_ip": "10.0.0.5"
}
```
## Build
To build all you need to do is run the build script in the root of the project.
This will build the project and put it at your `~/.local/bin`. Make sure that
`~/.local/bin` is added to your PATH.

# Usage
```bash
# Get the current air con status
aircon status
23.0 Cold Night Off

# Toggle power of the air con
aircon toggle

# Change between modes on the air con
aircon mode Heat
aircon mode Cold

# Set the temperature to what you desrise
aircon temp 28

# Set to any of the fan modes listed below
# 1,2,3,4,5,Night
aircon fan Night
```

You can also check out [`scripts/control`](/scripts/control), to
give you some examples on how you can write your own scripts for
the air con.
