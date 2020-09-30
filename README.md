# What is Miller?

**Miller is like awk, sed, cut, join, and sort for name-indexed data such as CSV, TSV, and tabular JSON.**

# Build status

[![Linux build status](https://travis-ci.org/johnkerl/miller.svg?branch=master)](https://travis-ci.org/johnkerl/miller)
[![Windows build status](https://ci.appveyor.com/api/projects/status/github/johnkerl/miller?branch=master&svg=true)](https://ci.appveyor.com/project/johnkerl/miller)
[![License](http://img.shields.io/badge/license-BSD2-blue.svg)](https://github.com/johnkerl/miller/blob/master/LICENSE.txt)
[![Docs](https://readthedocs.org/projects/miller/badge/?version=latest)](https://miller.readthedocs.io/en/latest/?badge=latest)

# Distributions

There's a good chance you can get Miller pre-built for your system:

[![Ubuntu](https://img.shields.io/badge/distros-ubuntu-db4923.svg)](https://launchpad.net/ubuntu/+source/miller)
[![Ubuntu 16.04 LTS](https://img.shields.io/badge/distros-ubuntu1604lts-db4923.svg)](https://launchpad.net/ubuntu/xenial/+package/miller)
[![Fedora](https://img.shields.io/badge/distros-fedora-173b70.svg)](https://apps.fedoraproject.org/packages/miller)
[![Debian](https://img.shields.io/badge/distros-debian-c70036.svg)](https://packages.debian.org/stable/miller)
[![Gentoo](https://img.shields.io/badge/distros-gentoo-4e4371.svg)](https://packages.gentoo.org/packages/sys-apps/miller)

[![Pro-Linux](https://img.shields.io/badge/distros-prolinux-3a679d.svg)](http://www.pro-linux.de/cgi-bin/DBApp/check.cgi?ShowApp..20427.100)
[![Arch Linux](https://img.shields.io/badge/distros-archlinux-1792d0.svg)](https://aur.archlinux.org/packages/miller-git)

[![NetBSD](https://img.shields.io/badge/distros-netbsd-f26711.svg)](http://pkgsrc.se/textproc/miller)
[![FreeBSD](https://img.shields.io/badge/distros-freebsd-8c0707.svg)](https://www.freshports.org/textproc/miller/)

[![Homebrew/MacOSX](https://img.shields.io/badge/distros-macosxbrew-ba832b.svg)](https://github.com/Homebrew/homebrew-core/search?utf8=%E2%9C%93&q=miller)
[![MacPorts/MacOSX](https://img.shields.io/badge/distros-macports-1376ec.svg)](https://www.macports.org/ports.php?by=name&substr=miller)
[![Chocolatey](https://img.shields.io/badge/distros-chocolatey-red.svg)](https://chocolatey.org/packages/miller)

|OS|Installation command|
|---|---|
|Linux|`yum install miller`<br/> `apt-get install miller`|
|Mac|`brew install miller`<br/>`port install miller`|
|Windows|`choco install miller`|

See also [building from source](http://johnkerl.org/miller/doc/build.html#Building_from_source).

# What can Miller do for me?

With Miller, you get to use named fields without needing to count positional
indices, using familiar formats such as CSV, TSV, JSON, and positionally-indexed.

For example, suppose you have a CSV data file like this:

```
county,tiv_2011,tiv_2012,line,construction
SEMINOLE,22890.55,20848.71,Residential,Wood
MIAMI DADE,1158674.85,1076001.08,Residential,Masonry
PALM BEACH,1174081.5,1856589.17,Residential,Masonry
MIAMI DADE,2850980.31,2650932.72,Commercial,Reinforced Masonry
HIGHLANDS,23006.41,19757.91,Residential,Wood
HIGHLANDS,49155.16,47362.96,Residential,Wood
DUVAL,1731888.18,2785551.63,Residential,Masonry
ST. JOHNS,29589.12,35207.53,Residential,Wood
```

Then, on the fly, you can add new fields which are functions of existing fields, drop fields, sort, aggregate statistically, pretty-print, and more:
```
$ mlr --icsv --opprint --barred \
  put '$tiv_delta = $tiv_2012 - $tiv_2011; unset $tiv_2011, $tiv_2012' \
  then sort -nr tiv_delta flins.csv 
+------------+-------------+----------------+
| county     | line        | tiv_delta      |
+------------+-------------+----------------+
| Duval      | Residential | 1053663.450000 |
| Palm Beach | Residential | 682507.670000  |
| St. Johns  | Residential | 5618.410000    |
| Highlands  | Residential | -1792.200000   |
| Seminole   | Residential | -2041.840000   |
| Highlands  | Residential | -3248.500000   |
| Miami Dade | Residential | -82673.770000  |
| Miami Dade | Commercial  | -200047.590000 |
+------------+-------------+----------------+
```

This is something the Unix toolkit always could have done, and arguably always
should have done.

* Miller operates on **key-value-pair data** while the familiar
Unix tools operate on integer-indexed fields: if the natural data structure for
the latter is the array, then Miller's natural data structure is the
insertion-ordered hash map.

* Miller handles a **variety of data formats**,
including but not limited to the familiar **CSV**, **TSV**, and **JSON**.
(Miller can handle **positionally-indexed data** too!)

For a few more examples please see [Miller in 10 minutes](http://johnkerl.org/miller/doc/10-min.html).

# Features

* Miller is **multi-purpose**: it's useful for **data cleaning**,
**data reduction**, **statistical reporting**, **devops**, **system
administration**, **log-file processing**, **format conversion**, and
**database-query post-processing**.

* You can use Miller to snarf and munge **log-file data**, including selecting
out relevant substreams, then produce CSV format and load that into
all-in-memory/data-frame utilities for further statistical and/or graphical
processing.

* Miller complements **data-analysis tools** such as **R**, **pandas**, etc.:
you can use Miller to **clean** and **prepare** your data. While you can do
**basic statistics** entirely in Miller, its streaming-data feature and
single-pass algorithms enable you to **reduce very large data sets**.

* Miller complements SQL **databases**: you can slice, dice, and reformat data
on the client side on its way into or out of a database. You can also reap some
of the benefits of databases for quick, setup-free one-off tasks when you just
need to query some data in disk files in a hurry.

* Miller also goes beyond the classic Unix tools by stepping fully into our
modern, **no-SQL** world: its essential record-heterogeneity property allows
Miller to operate on data where records with different schema (field names) are
interleaved.

* Miller is **streaming**: most operations need only a single record in
memory at a time, rather than ingesting all input before producing any output.
For those operations which require deeper retention (`sort`, `tac`, `stats1`),
Miller retains only as much data as needed. This means that whenever
functionally possible, you can operate on files which are larger than your
system&rsquo;s available RAM, and you can use Miller in **tail -f** contexts.

* Miller is **pipe-friendly** and interoperates with the Unix toolkit

* Miller's I/O formats include **tabular pretty-printing**, **positionally
  indexed** (Unix-toolkit style), CSV, JSON, and others

* Miller does **conversion** between formats

* Miller's **processing is format-aware**: e.g. CSV `sort` and `tac` keep header lines first

* Miller has high-throughput **performance** on par with the Unix toolkit

* Not unlike `jq` (http://stedolan.github.io/jq/) for JSON, Miller is written
in portable, modern C, with **zero runtime dependencies**. You can download or
compile a single binary, `scp` it to a faraway machine, and expect it to work.

# Documentation links

* [**Full documentation for latest release**](http://johnkerl.org/miller/doc)
* [Head docs](https://miller.readthedocs.io/en/latest) match [head code](https://github.com/johnkerl/miller); [release-specific docs](https://miller.readthedocs.io) match [release-specific code](https://github.com/johnkerl/miller/tags).
* [Miller's license is two-clause BSD](https://github.com/johnkerl/miller/blob/master/LICENSE.txt).
* [Build information including dependencies](http://johnkerl.org/miller/doc/build.html)
* [Notes about issue-labeling in the Github repo](https://github.com/johnkerl/miller/wiki/Issue-labeling)
* [Active issues](https://github.com/johnkerl/miller/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc)
