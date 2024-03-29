# HyprNotify
A needless DBus implementation for `hyprctl notify` so I can use `notify-send`.

## Overview
This implements a pseudo notification daemon, which **tries to** manage the state of Hyprland's internal notification manager. This requires the centuries-old advanced technology called "YOLO".

Along with that, this also implements the `org.freedesktop.Notifications` DBus Name, so you can use utilities like `notify-send` to manage notifications.

## Demo
![hyprnotify_demo](https://github.com/codelif/hyprnotify/assets/68972644/d9985035-3c8e-43cf-97e1-7f25219039e3)

## TODO
 - [x] Implement the DBus Protocol
 - [ ] Replace shell command invocation with IPC
 - [ ] Scrap the Project

## Why?
No Reason in particular, just wanted to learn Go & DBus this ~~week~~ today.

## How?
By sacrificing everything I stood for. This thing is filled with race conditions.
