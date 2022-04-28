#!/bin/bash
awk '{ print $2 }' | xargs kill
