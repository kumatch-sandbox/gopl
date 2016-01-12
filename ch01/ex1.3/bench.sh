#!/bin/bash

go test -bench Bench -count 1 -v | grep -e "ns/op" -e "Echo"
