package userMsg

func NotFound() string {
	s := `
   You must be in a git repo project directory
   for this command to run. The files created from
   "ci create" will need to be checked into the repo.

   git clone <your github project>
   cd <your github project>

   Once you're in the directory you can run ci create

   ci create

   Now add these files.

   git add ci
   git commit 'adding ci'
   git push

`
	return s

}
