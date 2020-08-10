# golaunch
a simple .desktop file launcher

can be used with any other application like fzf which gives the user a choice from a list the returns what the user chooses
with this:
  `./golaunch | fzf | ./golaunch`

so with instantmenu from instantos:
  `./golaunch | instantmenu -i -f -q "search apps" -l 10 -c -w -1 -h -1 -bw 4 | ./golaunch`

![Example](golaunch.png)
