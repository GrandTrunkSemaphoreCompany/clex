# Clex

> The Clacks is the informal nickname for the semaphore system that is the fastest means of non-magical communication on the Discworld. [1](https://wiki.lspace.org/mediawiki/Clacks)

> Hex is the Unseen University's organic/inorganic/magical super-computer [2](https://wiki.lspace.org/mediawiki/Hex)

Clex is the name for a computer controlled Clacks system, literally *Clacks* + *Hex*. It's a portmanteau because neologism's are cool.

The current idea involves a multi-component webserver, including:

- Encoders from byte data into
  - serial data for writing to a physical clacks 
  - image data for writing to a virtual clacks
- Decoders from
  - camera video for reading from a physical clacks
  - image files for reading from a virtual clacks  
- GET handlers
    - provides a way to read messages sent to this "tower"
- POST handlers
    - provides a way to add messages onto a queue to be sent out by this tower
- Internal queues to hold the sending, receiving and relaying of messages

## CLI usage

`go run main.go --id 151 --camera internal,200 --camera usb1,204 --shutter usb3,200`

This will run a Clex "tower" as Tower number 151. It uses a camera on the internal conneciton, pointing at tower 200 and another one on a usb1 interface pointing at tower 204. It has a single shutter on usb3 pointed at tower 200.
