# Benchmark

This file shows the extraction quality for pages from various websites.
You should take these results with a grain of salt. They are more or less used
for regression testing only.

The results were obtained using newscat built from `758f2d3` commit.

|News                          |Website                       |          |Pages     |TP        |TN        |FP        |FN        |F-Score   |
|:-----------------------------|:-----------------------------|:---------|---------:|---------:|---------:|---------:|---------:|---------:|
|ABC News                      |abcnews.go.com                |US        |2         |40        |94        |0         |1         |0.99      |
|Ars Technica                  |arstechnica.com               |US        |2         |62        |79        |0         |1         |0.99      |
|Associated Press              |ap.org                        |US        |2         |54        |332       |4         |0         |0.96      |
|Baltimore Sun                 |baltimoresun.com              |US        |2         |25        |398       |0         |1         |0.98      |
|Boston Globe                  |bostonglobe.com               |US        |2         |46        |64        |0         |1         |0.99      |
|Businessweek                  |businessweek.com              |US        |2         |24        |132       |0         |0         |1.00      |
|CNN                           |cnn.com                       |US        |2         |83        |389       |2         |0         |0.99      |
|Chicago Sun-Times             |suntimes.com                  |US        |2         |37        |454       |4         |0         |0.95      |
|Chicago Tribune               |chicagotribune.com            |US        |2         |28        |325       |0         |0         |1.00      |
|Detroit Free Press            |freep.com                     |US        |3         |44        |1029      |6         |0         |0.94      |
|Fox News                      |foxnews.com                   |US        |2         |38        |71        |1         |0         |0.99      |
|Los Angeles Daily News        |dailynews.com                 |US        |2         |46        |83        |1         |0         |0.99      |
|Los Angeles Times             |latimes.com                   |US        |2         |52        |349       |1         |1         |0.98      |
|National Journal              |nationaljournal.com           |US        |2         |34        |130       |1         |0         |0.99      |
|New Yok Post                  |nypost.com                    |US        |3         |50        |194       |0         |0         |1.00      |
|New York Daily News           |nydailynews.com               |US        |2         |45        |33        |0         |0         |1.00      |
|New York Times                |nytimes.com                   |US        |2         |47        |33        |0         |0         |1.00      |
|Newsday                       |newsday.com                   |US        |2         |23        |273       |0         |0         |1.00      |
|Orange County Register        |ocregister.com                |US        |2         |36        |122       |2         |1         |0.96      |
|Philadelphia Inquirer         |inquirer.com                  |US        |3         |90        |405       |5         |0         |0.97      |
|Politico                      |politico.com                  |US        |2         |55        |454       |0         |1         |0.99      |
|Reuters                       |reuters.com                   |US        |2         |51        |312       |2         |0         |0.98      |
|SFGate                        |sfgate.com                    |US        |2         |96        |27        |0         |0         |1.00      |
|San Diego Union-Tribune       |utsandiego.com                |US        |2         |37        |48        |0         |0         |1.00      |
|Seattle Times                 |seattletimes.com              |US        |2         |61        |609       |0         |2         |0.98      |
|Star Tribune                  |startribune.com               |US        |2         |34        |311       |0         |0         |1.00      |
|TIME                          |time.com                      |US        |2         |70        |95        |0         |0         |1.00      |
|Tampa Bay Times               |tampabay.com                  |US        |2         |52        |867       |0         |0         |1.00      |
|TechCrunch                    |techcrunch.com                |US        |2         |74        |180       |0         |0         |1.00      |
|The Arizona Republic          |azcentral.com                 |US        |2         |43        |70        |0         |0         |1.00      |
|The Dallas Morning News       |dallasnews.com                |US        |2         |54        |102       |1         |1         |0.98      |
|The Hill                      |thehill.com                   |US        |2         |41        |388       |0         |0         |1.00      |
|The Plain Dealer              |cleveland.com                 |US        |2         |26        |442       |2         |0         |0.96      |
|USA Today                     |usatoday.com                  |US        |2         |51        |133       |1         |0         |0.99      |
|Vanity Fair                   |vanityfair.com                |US        |2         |105       |128       |0         |6         |0.97      |
|Wall Street Journal           |wsj.com                       |US        |2         |70        |404       |0         |2         |0.99      |
|Washington Examiner           |washingtonexaminer.com        |US        |2         |61        |162       |0         |8         |0.94      |
|Washington Post               |washingtonpost.com            |US        |2         |91        |686       |1         |0         |0.99      |
|Washington Times              |washingtontimes.com           |US        |2         |98        |395       |0         |0         |1.00      |
|BBC                           |bbc.co.uk                     |UK        |2         |41        |209       |0         |2         |0.98      |
|Daily Mail                    |dailymail.co.uk               |UK        |3         |105       |1252      |4         |29        |0.86      |
|Telegraph                     |telegraph.co.uk               |UK        |2         |88        |153       |3         |2         |0.97      |
|The Guardian                  |theguardian.com               |UK        |2         |39        |311       |0         |0         |1.00      |
|The Independent               |independent.co.uk             |UK        |2         |35        |718       |0         |0         |1.00      |
|The Register                  |theregister.co.uk             |UK        |2         |51        |168       |0         |2         |0.98      |
|Calgary Herald                |calgaryherald.com             |CA        |2         |42        |645       |0         |0         |1.00      |
|Edmonton Journal              |edmontonjournal.com           |CA        |2         |39        |580       |3         |0         |0.96      |
|National Post                 |nationalpost.com              |CA        |2         |35        |241       |1         |1         |0.97      |
|The Chronicle Herald          |thechronicleherald.ca         |CA        |2         |34        |245       |0         |0         |1.00      |
|The Globe and Mail            |theglobeandmail.com           |CA        |2         |37        |253       |2         |2         |0.95      |
|The Province                  |theprovince.com               |CA        |2         |42        |813       |0         |1         |0.99      |
|Toronto Star                  |thestar.com                   |CA        |2         |36        |78        |0         |7         |0.91      |
|Vancouver Sun                 |vancouversun.com              |CA        |2         |52        |519       |0         |2         |0.98      |
|Winnipeg Free Press           |winnipegfreepress.com         |CA        |2         |42        |490       |0         |1         |0.99      |
|Bild                          |bild.de                       |DE        |3         |59        |218       |0         |6         |0.95      |
|Der Freitag                   |freitag.de                    |DE        |2         |26        |111       |0         |1         |0.98      |
|Der Tagesspiegel              |tagesspiegel.de               |DE        |2         |33        |215       |2         |0         |0.97      |
|Die Zeit                      |zeit.de                       |DE        |2         |45        |309       |2         |1         |0.97      |
|Focus                         |focus.de                      |DE        |3         |41        |331       |2         |5         |0.92      |
|Frankfurter Allgemeine        |faz.net                       |DE        |2         |16        |499       |0         |0         |1.00      |
|Handelsblatt                  |handelsblatt.com              |DE        |2         |35        |250       |1         |3         |0.95      |
|Heise                         |heise.de                      |DE        |2         |67        |418       |5         |5         |0.93      |
|Spiegel                       |spiegel.de                    |DE        |3         |71        |694       |0         |5         |0.97      |
|Tagesschau                    |tagesschau.de                 |DE        |2         |30        |158       |0         |0         |1.00      |
|Der Standard                  |derstandard.at                |AT        |2         |29        |411       |0         |1         |0.98      |
|Neue Zürcher Zeitung          |ncc.ch                        |CH        |2         |30        |134       |1         |3         |0.94      |
|                              |                              |          |*139*     |*3274*    |*21225*   |*60*      |*105*     |**0.98**  |

