# gopgfuncs
Example project for creating postgresql user defined functions in go (instead of C)


## Instructions

Modify the top 2 lines of the makefile and set:

    PGINC=/usr/include/postgresql/<YOURVERSION>/server/
    DBNAME=<YOURDBNAME>

Once you've done that, just run:

    make install

That should install two new functions in to your postgres server.

To test, run psql and do:

```
psql (9.5.2, server 9.5.0)
Type "help" for help.

postgres=# select inc(15);
 inc
-----
  15
(1 row)

postgres=# select inc(15);
 inc
-----
  30
(1 row)

postgres=# select add_one(5);
 add_one
---------
       6
(1 row)
```

If you look at your /var/log/postgresql/postgresql-9.5-main.log (or whatever your version is)
you will see it print a log line for each function call.
