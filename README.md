# gaspass
Store passwords without actually storing them

## How does it work?
Use one password for access all other passwords, but unlike password managers (like Keepass family) this tool does not store them on disk or in memory at all.
Gaspass is more a password generator than password manager. Every run you will get the password and this password will always be the same if you use the same parameters such as length, character set, resource and private key.
Working principle is the same as [lesspass](https://github.com/lesspass/lesspass), but gaspass uses modern [argon2id](https://en.wikipedia.org/wiki/Argon2) KDF (key derivation function) instead of PBKDF2-SHA1.

## Is it secure?
Generally yes, but it highly dependent on private key quality and protection

## ToDo
[] More tests
[] Better README.md
[] Travis CI
[] SECURITY.md
[] Resource management
[] GUI
