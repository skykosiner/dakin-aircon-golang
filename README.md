# What is this?
I have air con around my house with the brand Dakin and you can use a local web
# If
api to control them so I wanted a way to control my air con locally on my
computer.

This allows me to change all the most changed settings about my air con
straight from my terminal.

# Setup
The first thing you need to do is find the ip of the air con you want to control
and change it in the code. My ip is set to `10.0.0.24`, but yours could be
`192.168.1.69` and you'd want to change everywhere in the code `10.0.0.24` is
mentioned to the ip of your air con.

Once you make the changes to the code simply save your changes run `go build -o
air con ./cmd/main.go` and move the air con executable file to a `~/.local/bin`
path


# Usage
```
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
```
