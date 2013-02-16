smtp.go: smtp.ragel
	/home/tom/src/ragel/src/ragel -G2 -Z $<

smtp_message.go: smtp_message.ragel
	/home/tom/src/ragel/src/ragel -G2 -Z $<

smtp.dot: smtp.ragel
	/home/tom/src/ragel/src/ragel -V smtp.ragel > smtp.dot

