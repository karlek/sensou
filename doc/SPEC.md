Specification
=============

Hex Dump
--------

02 00 00 00 00 00 00 00 03 00 00 00 68 70 65 61    ............hpea
68 30 30 30 01 00 00 00 75 6E 61 6D 03 00 00 00    h000....unam....
4E 75 6D 62 65 72 31 00 00 00 00 00 68 66 6F 6F    Number1.....hfoo
68 30 30 31 01 00 00 00 75 6E 61 6D 03 00 00 00    h001....unam....
4E 75 6D 62 65 72 32 00 00 00 00 00 68 6B 6E 69    Number2.....hkni
68 30 30 32 01 00 00 00 75 6E 61 6D 03 00 00 00    h002....unam....
4E 75 6D 62 65 72 33 00 00 00 00 00                Number3.....

Explanation
-----------

02 00 00 00 00 00 00 00       Header 
                     03       Number of units to import
               00 00 00       Value seperator 
            68 70 65 61       Raw code of base unit (hpea)
            68 30 30 30       Raw code of custom unit (h000)
                     01       Number of changed fields
               00 00 00       Value separator 
            75 6E 61 6D       Raw code of changed field (unam)
                     03       Field type*
               00 00 00       Value separator
   4E 75 6D 62 65 72 31       Field value (Number1)
         00 00 00 00 00       Field or Unit seperator


00      Integer/Bool (little-endian)
01      ??? (Probably fixed point)
02      Float  (little-endian)
03      String

24 bit integers
24 bit floats
Booleans are integer values with 0 for false and all other values for true

You can change values that have a max cap in the Warcraft 3 editor. 
500,000 is the hp limit cap in the editor, but by using a hex editing the max hp is 16,777,215
