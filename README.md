# wcwhen

Practise go app that takes a country as an argument and displays results from previous matches and upcoming games during the current (2018) world cup.

Grabs data from [here](http://worldcup.sfg.io/matches)

``` wcwhen -g=Germany ```

```
============================================
Germany Vs Mexico

Score: 0-1

Last updated 	2018-06-21T02:43:20Z
============================================
35' goal Hirving  LOZANO Mexico
40' yellow-card Hector MORENO Mexico
58' substitution-in Edson ALVAREZ Mexico
58' substitution-out Carlos VELA Mexico
60' substitution-in Marco REUS Germany
60' substitution-out Sami KHEDIRA Germany
66' substitution-in Raul JIMENEZ Mexico
66' substitution-out Hirving  LOZANO Mexico
74' substitution-in Rafael MARQUEZ Mexico
74' substitution-out Andres GUARDADO Mexico
79' substitution-in Mario GOMEZ Germany
79' substitution-out Marvin PLATTENHARDT Germany
83' yellow-card Thomas MUELLER Germany
84' yellow-card Mats HUMMELS Germany
86' substitution-in Julian BRANDT Germany
86' substitution-out Timo WERNER Germany
90' yellow-card Hector HERRERA Mexico
```

** Todo **
========
* Refactor
    * Have broken out into packages
    * Now some parts can be simplified with smoother logic, less nested _ifs_
* Figure out cli arguents to give smarter output
    * Have added basic cli arguments using _flags_ but need to add something more robust
* TDD?




