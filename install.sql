-- You need a create function call per exported function from
-- the go code.

create or replace function add_one(integer)
  returns integer as :MOD,'AddOne'
  LANGUAGE C STRICT;

  create or replace function inc(int8)
    returns int8 as :MOD,'Inc'
    LANGUAGE C STRICT;
