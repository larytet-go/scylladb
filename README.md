# scylladb

Scylla Pair Coding Task A
-------------------------

The task is to implement a command that is going to be poor-man-clone of the `wget` command. It should take 1 argument - the url of a file to download, and save it to local disk.

Each 1 second it should log current progress so far in bytes.

Example usage:

```
$ gowget http://releases.ubuntu.com/18.04.3/ubuntu-18.04.3-desktop-amd64.iso
Downloaded 123 bytes ...
Downloaded 456 bytes ...
Downloaded 789 bytes ...
...
```
