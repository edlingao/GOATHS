#!/bin/sh

if [ ! -f models/main.db ];
  then
    sqlite3 migration/main.db < migration/schema.sql;
fi
