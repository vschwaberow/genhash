# genhash

genhash is a tool to generate hashes on the commandline from stdio.

It can be used to generate single or multiple hashes for usage in password databases or even in penetration testing scenarios where you want to test password cracking tools.

## Install

genhash is written in Go and has no dependencies. You can install the tool with your go installation using following command:

'
' go get -u github.com/vschwaberow/genhash
'

Or you've got the choice to download a binary distribution package under Releases.

## Usage

You can either provide a text string as argument, to be hashed by one of the program supported hash algorithms.

'
' genhash -a <HASH_ALGORITHM> <text string>
'

Or you can provide a list of text strings to be hashed over the stdin

'
' cat <LIST_OF_STRINGS> | genhash -a <HASH_ALGORITHM>
'

Supported hash algorithms are:

* md4
* md5 
* bcrypt
* sha1
* sha2-224, sha2-256, sha2-512
* sha3-224, sha3-256, sha3-512