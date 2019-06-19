package userMsg

func NotFound() string {
	s := `
   You must be in a git repo project directory
   for this command to run. 

   git clone <your github project>
   cd <your github project>

   Once you're in the directory you can run ci create

   ci create

`
	return s

}
