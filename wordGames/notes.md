# SQL For Dbase Prep

## Create letter counts
`update sandbox set acount = (select REGEXP_COUNT(word, 'a') from dual);`