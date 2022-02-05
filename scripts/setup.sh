#!/bin/bash
current=$(pwd)
# echo "$current"

# Check the current machine installed python or not
pyv=0
is_3_8=0
python3.8 --version >/dev/null 2>/dev/null

if [ "$?" -eq 0 ];then
	pyv="$(python3.6 --version)"
else
	python3 --version >/dev/null
	if [ "$?" -eq 0 ]; then
		pyv="$(python3 --version)"
		is_3_8=1
	else
		echo "Python is required!"
		exit 1
	fi
fi

read -ra arr <<<"$pyv"
pyv=${arr[1]}
OLD_IFS="$IFS"
IFS='.'
read -ra arr <<<"$pyv"
IFS="$OLD_IFS"

py_major="${arr[0]}"
py_minor="${arr[1]}"
py_patch="${arr[2]}"

echo $py_major $py_minor $py_patch

if [ $py_major -eq 2 ];then
	echo "Python 3 is required!"
	exit 1
fi
if [ $py_major -eq 3 ] && [ $py_minor -lt 8 ];then
	echo "Python 3.6.5 is minimum version supported"
	echo "Your current python version is: $py_major.$py_minor.$py_patch"
	exit 1
else
	if [ $py_major -eq 3 ] && [ $py_minor -eq 6 ] && [ $py_patch -lt 5 ];then
		echo "Python 3.6.5 is minimum version supported"
		echo "Your current python version is: $py_major.$py_minor.$py_patch"
		exit 1
	fi
fi


virtualenv --version >/dev/null 2>/dev/null

if [ "$?" -eq 0 ];then
	if [ $is_3_8 -eq 0 ];then
		virtualenv -p "$(which python3.6)" venv >/dev/null
	else
		virtualenv -p "$(which python3)" venv >/dev/null
	fi
	echo "Virtual environment created"
else
	if [ $is_3_8 -eq 0 ];then
		python3.8 -m venv venv >/dev/null
	else
		python3 -m venv venv >/dev/null
	fi
	echo "Virtual environment created"
fi

echo "Installling package..."

"$(pwd)/venv/bin/pip3.6" install -r requirements.txt

if [ ! -d ./logs/ ]; then
	echo "Create logs folder..."
	mkdir "$current/logs"
	touch "$current/logs/error.log"
	touch "$current/logs/info.log"
	touch "$current/logs"

else
	if [ ! -f ./logs/logs/error.log ];then
		touch "$current"/logs/error.log
	fi

	if [ ! -f ./logs/logs/info.log ];then
		touch "$current"/logs/info.log
	fi
fi

echo "Finish!"
echo 'Run `make start` to start application'
