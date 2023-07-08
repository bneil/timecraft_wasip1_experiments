# Timecraft / Wasip1 Experiments

Mostly just doing this as a toy repo to play around with this technology since it came up
in the wasmtime newsletter.

To run, use the Makefile and make sure to follow the steps from timecraft in the initial setup.

I put this together mostly because I usually use gofiber as a webserver, but the wasip1 doesn't handle that yet which im assuming
is why they "timecraft" created their own implementation (or at least the errors I ran into, `socket.go:9:10: undefined: syscall.ForkLock` among others).

