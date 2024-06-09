#!/bin/bash
PROJECT_PATH=$HOME/alice
if [ -e $PROJECT_PATH ]; then
    cd $PROJECT_PATH
fi
VENV_PATH=.venv/bin/activate
if [ -e $VENV_PATH ]; then
    source $VENV_PATH
fi
python3.11 -m main.py