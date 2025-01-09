# Super Simple Inventory Database - SSID

SSID is exacly that; super simple, but still functional. And, Its free!

I wrote this to inventory my office and server cabinet using a $60 bar-code scanner.

## Functions

- Register your items by serialnr, price and description
- Sort by location, item type, manufacturer and status
- You add your own locations, types, manufacts and statuses in the app.

## Getting started

[Download binaries](bin/)

### Build from code

- Clone ssid reposetory and run 'make build'.
  - You find binary in the /bin folder ex: ssid_darwin.arm64 on a mac.
  
### Runing SSID

- Create a folder for the application and put the binary in it.
  - SSID will create sub-folders and files at this location and need the right privileges.
- After starting SSID you can access the start page in your browser at <https://ipaddress:5005>.
  - SSID generates it's own TLS-crtificates.
  - You can edit the port number in the srv.yaml file in the app folder

## Using SSID
  
- When you see the  Login' page, click on the 'Signup' link at the bottom.
  - First user will be administrator (admin only work with users, not inventory).
  - After adding aditional users, login as admin and authenticate them.
- Login as a user and go to the 'Tools' menu.
  - Populate the locations, type, manufact and status properties.
  - The 'New' status is a default property, and is already added.
  - You can rename it to your prefered language by selecting it, and click 'Add/Update'.
- SSID is npw ready for use.

## License

Copyright (c) 2025 Raadig AS

This project is licensed under the MIT License - see the [license](LICENSE) file for details.