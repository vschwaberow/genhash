# genhash

genhash is a tool to generate hashes on the commandline from stdio.

It can be used to generate single or multiple hashes for usage in password databases or even in penetration testing scenarios where you want to test password cracking tools.

## Install

genhash is written in Go. You can install the tool with your go installation using following command:

>
> go get -u github.com/vschwaberow/genhash
>

Or you've got the choice to download a binary distribution package under Releases.

## Usage

You can either provide a text string as argument, to be hashed by one of the program supported hash algorithms.

>
> genhash <HASH_ALGORITHM> <text string>
>

Or you can provide a list of text strings to be hashed over the stdin

>
> cat <LIST_OF_STRINGS> | genhash <HASH_ALGORITHM> -s
>

You can list all algorithms over the help function.

Supported are:

* Argon2 hash
* Bcrypt hash
* LanManager hash
* MD4 hash
* MD5 hash
* NTLM hash
* Generate random bytes of size provided
* RIPEMD160 hash
* SHA-1 hash
* SHA2-224 hash
* SHA2-256 hash
* SHA2-512 hash
* SHA3-224 hash
* SHA3-256 hash
* SHA3-512 hash
* Tiger hash
* UUID-1 random string
* UUID-4 random string
* Whirlpool hash
