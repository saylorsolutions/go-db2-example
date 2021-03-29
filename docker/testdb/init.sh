#!/bin/bash

# Specify all scripts that need to run in this array
scripts=(00_schema.sql 01_data.sql)

for s in ${scripts[@]}; do
  su - db2inst1 -c "db2 -tvmf /custom_sql/$s | tee $s.log"
done
