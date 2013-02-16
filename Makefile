smtp.go: smtp.ragel
	/home/tom/src/ragel/src/ragel -G2 -Z smtp.ragel

smtp.dot: smtp.ragel
	/home/tom/src/ragel/src/ragel -V smtp.ragel > smtp.dot

