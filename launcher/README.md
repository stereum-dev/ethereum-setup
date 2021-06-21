# Run from source

### Install python3
Install python3 https://www.python.org/downloads/

### Prepare
```
python3 -m venv venv
. ./venv/bin/activate
pip install -r requirements.txt
```
This will create a virtual environment and load all the necessary python dependencies.

### Run
Example:
```
python3 launcher.py --host 1.2.3.4 --keyfile /Users/stefa2k/.ssh/id_rsa
```

Show all available parameters:
```
python3 launcher.py --help
```
