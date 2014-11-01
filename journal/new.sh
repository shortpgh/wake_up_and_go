#!/bin/bash
putpath=`dirname "${BASH_SOURCE[0]}"`/
filename=$putpath`date "+%Y%m%d"`.md
echo "=Day "`ls | wc -w`": " `date "+%Y-%m-%d\n"` >> $filename
git add $filename