# net-cat

This code allows you to create a group chat using the nc (netcat) utility.

### Usage

-  clone https://github.com/vavasimioki/net-cat.git

- " go run ./cmd . " or " go run ./cmd 2525 "

### Connection

Run in a different terminal to connect:

 - nc localhost 8989

Enter the correct nickname and you will be in chat.

Note: The maximum name and message length is 1024 characters and only ASCII characters 32 through 126 are accepted.
