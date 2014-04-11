# Benchmark

This file shows the extraction quality for pages from various websites.
You should take these results with a grain of salt. They are more or less used
for regression testing only.

The results were obtained using newscat built from `401fed8` commit.

|News                          |Website             |        |Pages   |F-Score |
|:-----------------------------|:-------------------|:-------|-------:|-------:|
|ABC News                      |abcnews.go.com      |US      |2       |0.99    |
|Ars Technica                  |arstechnica.com     |US      |2       |0.99    |
|Associated Press              |ap.org              |US      |2       |0.96    |
|Baltimore Sun                 |baltimoresun.com    |US      |2       |0.98    |
|Boston Globe                  |bostonglobe.com     |US      |2       |0.99    |
|Businessweek                  |buisnessweek.com    |US      |1       |1.00    |
|CNN                           |cnn.com             |US      |2       |0.98    |
|Chicago Sun-Times             |suntimes.com        |US      |2       |0.90    |
|Chicago Tribune               |chicagotribune.com  |US      |2       |0.97    |
|Detroit Free Press            |freep.com           |US      |3       |0.96    |
|Fox News                      |foxnews.com         |US      |2       |0.97    |
|Los Angeles Daily News        |dailynews.com       |US      |2       |0.99    |
|Los Angeles Times             |latimes.com         |US      |1       |1.00    |
|National Journal              |nationaljournal.com |US      |2       |0.97    |
|New York Times                |nytimes.com         |US      |1       |1.00    |
|Politico                      |politico.com        |US      |1       |0.98    |
|Reuters                       |reuters.com         |US      |2       |0.99    |
|SFGate                        |sfgate.com          |US      |2       |0.99    |
|Star Tribune                  |startribune.com     |US      |2       |1.00    |
|TIME                          |time.com            |US      |2       |1.00    |
|TechCrunch                    |techcrunch.com      |US      |2       |1.00    |
|The Dallas Morning News       |dallasnews.com      |US      |2       |0.97    |
|The Plain Dealer              |cleveland.com       |US      |2       |0.98    |
|USA Today                     |usatoday.com        |US      |2       |0.99    |
|Vanity Fair                   |vanityfair.com      |US      |2       |0.97    |
|Wall Street Journal           |wsj.com             |US      |1       |0.96    |
|Washington Post               |washingtonpost.com  |US      |2       |0.99    |
|BBC                           |bbc.co.uk           |UK      |2       |0.98    |
|Daily Mail                    |dailymail.co.uk     |UK      |1       |1.00    |
|Telegraph                     |telegraph.co.uk     |UK      |1       |0.96    |
|The Guardian                  |theguardian.com     |UK      |2       |1.00    |
|The Independent               |independent.co.uk   |UK      |1       |1.00    |
|The Register                  |theregister.co.uk   |UK      |2       |0.98    |
|Bild                          |bild.de             |DE      |3       |0.96    |
|Der Freitag                   |freitag.de          |DE      |1       |1.00    |
|Der Tagesspiegel              |tagesspiegel.de     |DE      |1       |0.93    |
|Die Zeit                      |zeit.de             |DE      |2       |1.00    |
|Focus                         |focus.de            |DE      |3       |0.90    |
|Frankfurter Allgemeine        |faz.net             |DE      |1       |0.96    |
|Handelsblatt                  |handelsblatt.com    |DE      |2       |0.93    |
|Heise                         |heise.de            |DE      |2       |0.96    |
|Spiegel                       |spiegel.de          |DE      |3       |0.96    |
|Tagesschau                    |tagesschau.de       |DE      |2       |1.00    |
|Der Standard                  |derstandard.at      |AT      |2       |0.97    |
|Neue Zürcher Zeitung          |ncc.ch              |CH      |2       |0.94    |
|                              |                    |        |**84**  |**0.98**|

