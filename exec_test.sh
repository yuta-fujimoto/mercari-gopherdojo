#!/usr/bin/bash

#define
RED="\033[31m"
GREEN="\033[32m"
RESET="\033[m"

#reset
find test/muldirs -name "*.png" | xargs rm;
find test/pngToJpg -name "*.jpg" | xargs rm
find test/pngToPgm -name "*.pgm" | xargs rm;
find test/jpgToGif -name "*.gif" | xargs rm;
find test/gifToJpg -name "*.jpg" | xargs rm;
find test/jpgToPpm -name "*.ppm" | xargs rm;
rm test/42tokyo_logo.png


if [ "$1" = "r" ]; then
	exit 0
fi

mkdir -p test/NoPermFile
mkdir -p test/NoPermFile/sub
cp images/42tokyo_logo.jpg test/NoPermFile/sub/
chmod 000 test/NoPermFile/sub/42tokyo_logo.jpg

mkdir -p test/NoPermDir
chmod 000 test/NoPermDir

go build

# print result
function print_result () {
	printf "GOT:\n$1\n"
	printf "SRC:\n$2\n"

	if [ $(echo "$1" | wc -l) == $(echo "$2" | wc -l) ]; then
		printf "${GREEN}OK :) ${RESET}\n"
	else
		printf "${RED}KO :(${RESET}\n"
	fi
}

DIR=muldirs
printf "${GREEN}[NESTED DIRECTORY STRUCTURE(${DIR})]${RESET}\n"
./convert test/${DIR}
GOT=$(find test/${DIR} -name "*.png" | xargs file);
SRC=$(find test/${DIR} -name "*.jpg" | xargs file);

print_result "${GOT}" "${SRC}"

DIR=pngToJpg
printf "${GREEN}[PNG=>JPG(${DIR})]${RESET}\n"
./convert -i=png -o=jpg test/${DIR}
GOT=`find test/${DIR} -name "*.png" | xargs file`;
SRC=`find test/${DIR} -name "*.jpg" | xargs file`;

print_result "${GOT}" "${SRC}"

DIR=pngToPgm
printf "${GREEN}[PNG=>PGM(${DIR})]${RESET}\n"
./convert -i=png -o=pgm test/${DIR}
GOT=`find test/${DIR} -name "*.png" | xargs file`
SRC=`find test/${DIR} -name "*.pgm" | xargs file`

print_result "${GOT}" "${SRC}"

DIR=jpgToGif
printf "${GREEN}[JPG=>GIF(${DIR})]${RESET}\n"
./convert -i=jpg -o=gif test/${DIR}
GOT=`find test/${DIR} -name "*.gif" | xargs file`
SRC=`find test/${DIR} -name "*.jpg" | xargs file`

print_result "${GOT}" "${SRC}"

DIR=gifToJpg
printf "${GREEN}[GIF=>JPG(${DIR})]${RESET}\n"
./convert -i=gif -o=jpg test/${DIR}
GOT=`find test/${DIR} -name "*.jpg" | xargs file`
SRC=`find test/${DIR} -name "*.gif" | xargs file`

print_result "${GOT}" "${SRC}"

DIR=jpgToPpm
printf "${GREEN}[JPG=>PPM(${DIR})]${RESET}\n"
./convert -i=jpg -o=ppm test/${DIR}
GOT=`find test/${DIR} -name "*.jpg" | xargs file`
SRC=`find test/${DIR} -name "*.ppm" | xargs file`

print_result "${GOT}" "${SRC}"

DIR=42tokyo_logo.jpg
printf "${GREEN}[SPECIFY FILE NAME(${DIR})]${RESET}\n"
./convert -i=jpg -o=png test/${DIR};
GOT=`find test/ -maxdepth 1 -name "*.png" | xargs file`
SRC=`find test/ -maxdepth 1 -name "*.jpg" | xargs file`

print_result "${GOT}" "${SRC}"

printf "${RED}[NO COMMAND ARGS]${RESET}\n"
./convert

printf "${RED}[NO SUCH DIRECTORY]${RESET}\n"
./convert nosuchdirectory

printf "${RED}[NO SUCH FILE]${RESET}\n"
./convert nosuchfile.jpg

printf "${RED}[NO PERM(FILE)]${RESET}\n"
./convert test/NoPermFile

printf "${RED}[NO PERM(DIR)]${RESET}\n"
./convert test/NoPermDir

printf "${RED}[INVALID FILE FORMAT(test/NotImageErr)]${RESET}\n"
./convert test/NotImageErr

printf "${RED}[INVALID FILE FORMAT(test/NotImageErr/notimage)]${RESET}\n"
./convert test/NotImageErr/notimage

printf "${RED}[INVALID OPTION FORMAT(-i)]${RESET}\n"
./convert -i=pgm test/42tokyo_logo.jpg

printf "${RED}[INVALID OPTION FORMAT(-o)]${RESET}\n"
./convert -o=noop test/42tokyo_logo.jpg

printf "${RED}[I/O FILE FORMATS ARE SAME]${RESET}\n"
./convert -i=jpg -o=jpg test/NotImageErr/42tokyo_logo.jpg

chmod 755 test/NoPermFile/sub/42tokyo_logo.jpg
rm -rf test/NoPermFile

chmod 755 test/NoPermDir
rm -rf test/NoPermDir
