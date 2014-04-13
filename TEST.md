# Benchmark

This file shows the extraction quality for pages from various websites.
You should take these results with a grain of salt. They are more or less used
for regression testing only.

The results were obtained using newscat built from `38c7abf` commit.

|News                          |Website                       |          |Pages     |TP        |TN        |FP        |FN        |F-Score   |
|:-----------------------------|:-----------------------------|:---------|---------:|---------:|---------:|---------:|---------:|---------:|
|ABC News                      |abcnews.go.com                |US        |2         |40        |102       |1         |0         |0.99      |
|Ars Technica                  |arstechnica.com               |US        |2         |62        |79        |0         |1         |0.99      |
|Associated Press              |ap.org                        |US        |2         |54        |340       |4         |0         |0.96      |
|Baltimore Sun                 |baltimoresun.com              |US        |2         |25        |409       |0         |1         |0.98      |
|Boston Globe                  |bostonglobe.com               |US        |2         |46        |104       |0         |1         |0.99      |
|Businessweek                  |buisnessweek.com              |US        |1         |11        |71        |0         |0         |1.00      |
|CNN                           |cnn.com                       |US        |2         |83        |401       |2         |1         |0.98      |
|Chicago Sun-Times             |suntimes.com                  |US        |2         |37        |460       |8         |0         |0.90      |
|Chicago Tribune               |chicagotribune.com            |US        |2         |28        |340       |1         |1         |0.97      |
|Detroit Free Press            |freep.com                     |US        |3         |44        |1029      |4         |0         |0.96      |
|Fox News                      |foxnews.com                   |US        |2         |38        |71        |2         |0         |0.97      |
|Los Angeles Daily News        |dailynews.com                 |US        |2         |46        |83        |1         |0         |0.99      |
|Los Angeles Times             |latimes.com                   |US        |1         |16        |179       |0         |0         |1.00      |
|National Journal              |nationaljournal.com           |US        |2         |34        |132       |0         |2         |0.97      |
|New Yok Post                  |nypost.com                    |US        |3         |50        |302       |0         |3         |0.97      |
|New York Daily News           |nydailynews.com               |US        |2         |45        |33        |0         |0         |1.00      |
|New York Times                |nytimes.com                   |US        |1         |19        |15        |0         |0         |1.00      |
|Newsday                       |newsday.com                   |US        |2         |23        |315       |1         |0         |0.98      |
|Orange County Register        |ocregister.com                |US        |2         |36        |122       |1         |3         |0.95      |
|Philadelphia Inquirer         |inquirer.com                  |US        |3         |90        |409       |5         |0         |0.97      |
|Politico                      |politico.com                  |US        |1         |25        |233       |0         |1         |0.98      |
|Reuters                       |reuters.com                   |US        |2         |51        |316       |1         |0         |0.99      |
|SFGate                        |sfgate.com                    |US        |2         |96        |27        |0         |1         |0.99      |
|San Diego Union-Tribune       |utsandiego.com                |US        |2         |37        |48        |0         |0         |1.00      |
|Seattle Times                 |seattletimes.com              |US        |2         |61        |664       |1         |2         |0.98      |
|Star Tribune                  |startribune.com               |US        |2         |34        |332       |0         |0         |1.00      |
|TIME                          |time.com                      |US        |2         |70        |95        |0         |0         |1.00      |
|Tampa Bay Times               |tampabay.com                  |US        |2         |52        |871       |1         |0         |0.99      |
|TechCrunch                    |techcrunch.com                |US        |2         |74        |184       |0         |0         |1.00      |
|The Arizona Republic          |azcentral.com                 |US        |2         |43        |70        |0         |0         |1.00      |
|The Dallas Morning News       |dallasnews.com                |US        |2         |54        |131       |2         |1         |0.97      |
|The Hill                      |thehill.com                   |US        |2         |41        |402       |2         |1         |0.96      |
|The Plain Dealer              |cleveland.com                 |US        |2         |26        |444       |1         |0         |0.98      |
|USA Today                     |usatoday.com                  |US        |2         |51        |133       |1         |0         |0.99      |
|Vanity Fair                   |vanityfair.com                |US        |2         |105       |134       |0         |7         |0.97      |
|Wall Street Journal           |wsj.com                       |US        |1         |13        |209       |0         |1         |0.96      |
|Washington Examiner           |washingtonexaminer.com        |US        |2         |61        |234       |0         |13        |0.90      |
|Washington Post               |washingtonpost.com            |US        |2         |91        |688       |1         |1         |0.99      |
|Washington Times              |washingtontimes.com           |US        |2         |98        |441       |0         |0         |1.00      |
|BBC                           |bbc.co.uk                     |UK        |2         |41        |209       |0         |2         |0.98      |
|Daily Mail                    |dailymail.co.uk               |UK        |1         |18        |409       |0         |0         |1.00      |
|Telegraph                     |telegraph.co.uk               |UK        |1         |53        |87        |1         |3         |0.96      |
|The Guardian                  |theguardian.com               |UK        |2         |39        |313       |0         |0         |1.00      |
|The Independent               |independent.co.uk             |UK        |1         |22        |415       |0         |0         |1.00      |
|The Register                  |theregister.co.uk             |UK        |2         |51        |168       |0         |2         |0.98      |
|Bild                          |bild.de                       |DE        |3         |59        |245       |0         |5         |0.96      |
|Der Freitag                   |freitag.de                    |DE        |1         |10        |61        |0         |0         |1.00      |
|Der Tagesspiegel              |tagesspiegel.de               |DE        |1         |13        |94        |2         |0         |0.93      |
|Die Zeit                      |zeit.de                       |DE        |2         |44        |317       |0         |0         |1.00      |
|Focus                         |focus.de                      |DE        |3         |41        |391       |2         |7         |0.90      |
|Frankfurter Allgemeine        |faz.net                       |DE        |1         |13        |381       |0         |1         |0.96      |
|Handelsblatt                  |handelsblatt.com              |DE        |2         |35        |250       |0         |5         |0.93      |
|Heise                         |heise.de                      |DE        |2         |67        |436       |1         |4         |0.96      |
|Spiegel                       |spiegel.de                    |DE        |3         |71        |721       |0         |6         |0.96      |
|Tagesschau                    |tagesschau.de                 |DE        |2         |28        |378       |0         |0         |1.00      |
|Der Standard                  |derstandard.at                |AT        |2         |29        |427       |0         |2         |0.97      |
|Neue Zürcher Zeitung          |ncc.ch                        |CH        |2         |29        |141       |2         |2         |0.94      |
|                              |                              |          |*109*     |*2573*    |*16095*   |*48*      |*80*      |**0.98**  |

