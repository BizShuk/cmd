#!/bin/bash
cat tmp | awk '{print "test/"$0}' | tr '\n' '\0'| xargs -0 tee

