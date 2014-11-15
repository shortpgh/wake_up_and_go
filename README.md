I have several goals for this project. The primary is to learn golang. Secondarily, I plan on using this time to learn my way around the acme editor. I also hope that this may turn into a resource that other developers can use to learn a new language -- go or anything else.

###PLAN

With the goal to learn golang, the plan is to write some small go program in the morning right after I wake up. These programs should be small -- to be completed in roughly 30 minutes. That's usually the time I have in the morning.

####JOURNAL

Every day should have a journal entry associated with it. Each entry should be named based on the date. Each entry should start with the date along with the day that you are on in the process. While I'd like to do 30 days, each day doesn't have to be in succession. Missed days don't count -- they happen. Jot anything you think is worthwhile while you dig into the day's lesson.

The journal folder includes a shell script that will create the file and add the appropriate header to the file.

#### COMMENT!

Add top-level comments to each programming file. Include the Day # + Date line from your journal entry to help you associate the files together. This will be helpful later when you have multiple files in a single day.

Include information about how to run the program. This could be the command line, or what services you may need to start (IE., databases that you don't typically run). If a file is a sub-file for another program, make sure to point which file is the parent -- also for easy future reference.

#### COMMIT!

Don't forget to commit each day's work. It doesn't hurt to include multiple commits throughout a day.

### PROGRAM IDEAS

#### hello world

Yeah, it's a lame way to start if you are an experienced programmer. But when you are new to a language, it is still the best starting point. Use this program on day 1 to make sure you can run the language. You can also use it as a test program to make sure any tooling that you need to setup is working.

#### golang tutorial

Follow the golang tutorials. This is the best place to start.

http://golang.org/doc/code.html
https://tour.golang.org/

#### Write a Unit Test

Figure out how to write a unit test. (This is part of the tutorial).

####

#### Code Kata's Binary Search
Write a binary search in 5 different ways (each a different day). This is a good way to learn different structural ways to create a program. It can be recursive, functional, class based, and any other method that i can think up. This also provides an introduction to a list-based data type.

For binary search, first sort the elements in the array. Then split it in half. if the value matches, return the position. if it doesn't, then search in the slice before or after that based on if the value is less than or greater than.

#### Using Standard Libraries

Build a program that uses 3 standard libraries that haven't been used before. This forces me to learn some of the built-in libraries provided by the language. Repeat this several times, always using a set of 3 different libraries. The program doesn't have to be usable, so it could be some interesting output.

The library reference is here: http://golang.org/pkg/

Here are some packages that have caught my interest so far:

archive
compress
container
crypto
encoding.csv
encoding.json
io
log
math.rand
net.mail
net.smtp
net.http
path
regexp
text.tabwriter
text.template
time

and here are some potential combinations:
archive - compress - crypto		// shrink and encrypt file(s)
log - net.smtp - regexp			// email matched entries
path - encoding.csv - container	//read csv into container
net.http - encoding.json -		// get json data and 
text.template - time - io			// create a template date file



#### Use a third-party library

Pick a 3rd party library and use it. For my example, I chose Google's Analytics library. I'm familiar with the API from previous experience in other languages. This eliminates learning about the API and focus on using it in the language.

#### Use a database

Write a program that inserts and retrieves data from a database. Repeat for the different types of databases that you might use. I'm going to give both MongoDB and Postgres a try.

#### Work with files

Write a program that opens, edits, and saves text in a file. Try appending or removing specific data in the file.

#### How about CLI?

Write a program that accepts command line parameters.

