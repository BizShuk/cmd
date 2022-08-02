#!/bin/bash

versions=($(git tag --points-at HEAD))
versions+=($(git log --pretty=format:'%h' -n 1))
echo -n "${versions[*]}" > version