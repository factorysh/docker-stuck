Docker stuck
============

Find stuck containers and kill them.

Compile
-------

    make

Usage
-----

    docker-stuck

Returns all containers, split in good ones (wich honor `inspect` command) and bad ones.

    docker-stuck --kill

Displays stuffs, and kill bad containers


Licence
-------

3 terms BSD licence, ©2018 Mathieu Lecarme