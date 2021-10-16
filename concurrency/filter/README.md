# install
	run make
		#making similink to /usr/local/bin/filter
# to delete
	make clean
		#deleting similink from /usr/local/bin/filter
		#deleting filter in current dir
# usage
	filter -d=10 -m=10 -y=2021 -s=.txt,.jpg pathtofiles/*
		#output all files with modified date
	filter -s=.jpg,.txt pathtofiles/*
		#output all files with modified date where extensions .jpg and .txt
	filter -d=15 -s=.jpg,.txt pathtofiles/*
		#output all files with modified date where extensions .jpg and .txt and date == 15
	* u can use -d 15 it's the same as -d=15 for all flagset
