# ws

A CLI tool to serve contents of your local directories quickly

    ws

will serve the contents of the **pwd(1)** as **http://localhost:8765**.

While

    ws -d /var/tmp

will serve the contents of **/var/tmp**.

And

    ws -p 3456

will serve the contents as **http://localhost:3456**.

## Installation

Requirements: [Go][go] lang.

And then

    make all
    make install

[go]: http://golang.org/
