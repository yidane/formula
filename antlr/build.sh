#!/usr/bin/env bash

rm ../internal/parser/*
java -jar antlr-4.9-complete.jar -Dlanguage=Go -o ../internal/parser Formula.g4