#!/bin/bash

clear
go test -v -count 1 github.com/Switcheo/Zion/consensus/hotstuff/basic/core -run TestNewRound
