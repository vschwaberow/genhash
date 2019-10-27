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

Supported hash algorithms are:

* Lanmanager
* NTLM / NTHash (UTF16 Little Endian MD4)
* md4
* md5 
* bcrypt
* ripemd160
* sha1
* sha2-224, sha2-256, sha2-512
* sha3-224, sha3-256, sha3-512
* UUID Version 1
* UUID Version 4
