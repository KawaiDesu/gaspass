# gaspass
Store passwords without actually storing them

## How does it work?
Ose one password for access all other passwords, but в отличии от does not store them on disk or in memory at all.
Gaspass is more a password generator than password manager or store tool. Every run you will get the password and this password will be the same if you use the same parameters like length, character set, resource and private key.
Work scheme is very similar to [lesspass](https://github.com/lesspass/lesspass), but uses modern [argon2id](https://en.wikipedia.org/wiki/Argon2) KDF (key derivation function) instead of PBKDF2-SHA1.

## Is it secure?
Generally yes, but it depends on private key quality and "защиты ключа"



## ToDo
[] Tests
[] SECURITY.md
[] Resource management
[] GUI



Это шобы версии текстов основного файла

Compare this version with version of localized file to make sure toy read an actual information.
README.md version 0