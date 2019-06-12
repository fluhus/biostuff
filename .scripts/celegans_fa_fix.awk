{
	if (substr($1, 1, 1) == ">") {
		print $1;
	} else {
		print;
	}
}
