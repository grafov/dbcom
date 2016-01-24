# dbcom
Database Commander with oldschool ncurses-like UI. In memories of NC and clones :)

I planned features and UI of this app several years ago but each time can't find motivations to start it. 
Hackaton give me such motivation. Well after starting I found what this project requires significantly more time 
than couple of days during the hackaton. But I happy for spent time because work really has started.
I will continue development in my repo http://github.com/grafov/bcom

Happy hacking to all Gophergala participants! :)


## Notes about development

In case you want code some apps with text based UI in Go I recommend you excellent library: https://github.com/nsf/termbox-go

But app like dbcom requires a bit more for easy working with text forms and widgets. There are couple of libs based on termbox offer such frameworks:

* termui is a terminal dashboard.
* gocui is a minimalist Go library aimed at creating console user interfaces.

I used gocui with some patches but finally decide to rewrite package for satisfy needs of dbcom. This work in progress.

