# Benchmark

This file shows the extraction quality for pages from various websites.
You should take these results with a grain of salt. They are more or less used
for regression testing only.

The results were obtained using newscat built from `e7ba99c` commit.

|News                          |Website                       |          |Pages     |TP        |TN        |FP        |FN        |F-Score   |
|:-----------------------------|:-----------------------------|:---------|---------:|---------:|---------:|---------:|---------:|---------:|
|ABC News                      |abcnews.go.com                |US        |2         |40        |80        |0         |1         |0.99      |
|Ars Technica                  |arstechnica.com               |US        |2         |62        |46        |0         |0         |1.00      |
|Associated Press              |ap.org                        |US        |2         |54        |312       |4         |0         |0.96      |
|Baltimore Sun                 |baltimoresun.com              |US        |2         |25        |375       |0         |1         |0.98      |
|Bloomberg                     |bloomberg.com                 |US        |2         |95        |576       |13        |0         |0.94      |
|Boston Globe                  |bostonglobe.com               |US        |2         |46        |50        |0         |1         |0.99      |
|Businessweek                  |businessweek.com              |US        |2         |24        |126       |0         |0         |1.00      |
|CNBC                          |cnbc.com                      |US        |2         |65        |166       |0         |0         |1.00      |
|CNN                           |cnn.com                       |US        |2         |83        |383       |2         |0         |0.99      |
|Chicago Sun-Times             |suntimes.com                  |US        |2         |37        |454       |4         |0         |0.95      |
|Chicago Tribune               |chicagotribune.com            |US        |2         |28        |314       |0         |0         |1.00      |
|Detroit Free Press            |freep.com                     |US        |3         |44        |937       |0         |0         |1.00      |
|Forbes                        |forbes.com                    |US        |3         |90        |90        |2         |0         |0.99      |
|Fox News                      |foxnews.com                   |US        |2         |38        |59        |1         |0         |0.99      |
|Huffington Post               |huffingtonpost.com            |US        |3         |137       |320       |3         |0         |0.99      |
|International Business Times  |ibtimes.com                   |US        |2         |60        |425       |0         |1         |0.99      |
|Los Angeles Daily News        |dailynews.com                 |US        |2         |46        |69        |1         |0         |0.99      |
|Los Angeles Times             |latimes.com                   |US        |2         |52        |339       |1         |1         |0.98      |
|National Geographic           |nationalgeographic.com        |US        |2         |89        |318       |1         |0         |0.99      |
|National Journal              |nationaljournal.com           |US        |2         |34        |86        |1         |0         |0.99      |
|National Public Radio         |npr.org                       |US        |2         |62        |71        |0         |0         |1.00      |
|Nature                        |nature.com                    |US        |2         |52        |316       |0         |2         |0.98      |
|New Scientist                 |newscientist.com              |US        |2         |62        |249       |0         |0         |1.00      |
|New York Daily News           |nydailynews.com               |US        |2         |45        |29        |0         |0         |1.00      |
|New York Post                 |nypost.com                    |US        |3         |50        |157       |0         |1         |0.99      |
|New York Times                |nytimes.com                   |US        |2         |47        |33        |0         |0         |1.00      |
|Newsday                       |newsday.com                   |US        |2         |23        |258       |0         |0         |1.00      |
|Orange County Register        |ocregister.com                |US        |2         |36        |118       |2         |1         |0.96      |
|Philadelphia Inquirer         |inquirer.com                  |US        |3         |90        |402       |2         |1         |0.98      |
|Politico                      |politico.com                  |US        |2         |55        |391       |0         |1         |0.99      |
|Reuters                       |reuters.com                   |US        |2         |51        |298       |2         |0         |0.98      |
|SFGate                        |sfgate.com                    |US        |2         |96        |25        |0         |0         |1.00      |
|San Diego Union-Tribune       |utsandiego.com                |US        |2         |37        |18        |0         |0         |1.00      |
|Science Daily                 |sciencedaily.com              |US        |2         |31        |1022      |0         |0         |1.00      |
|Scientific American           |scientificamerican.com        |US        |2         |54        |233       |0         |1         |0.99      |
|Seattle Times                 |seattletimes.com              |US        |2         |61        |551       |0         |1         |0.99      |
|Star Tribune                  |startribune.com               |US        |2         |34        |311       |0         |0         |1.00      |
|TIME                          |time.com                      |US        |2         |70        |92        |0         |0         |1.00      |
|Tampa Bay Times               |tampabay.com                  |US        |2         |52        |863       |0         |0         |1.00      |
|TechCrunch                    |techcrunch.com                |US        |2         |74        |121       |0         |0         |1.00      |
|The Arizona Republic          |azcentral.com                 |US        |2         |43        |58        |0         |0         |1.00      |
|The Atlantic                  |theatlantic.com               |US        |2         |131       |566       |0         |2         |0.99      |
|The Dallas Morning News       |dallasnews.com                |US        |2         |54        |90        |1         |1         |0.98      |
|The Hill                      |thehill.com                   |US        |2         |41        |380       |0         |0         |1.00      |
|The New Yorker                |newyorker.com                 |US        |2         |48        |174       |1         |1         |0.98      |
|The Plain Dealer              |cleveland.com                 |US        |2         |26        |414       |1         |0         |0.98      |
|The Verge                     |theverge.com                  |US        |3         |42        |447       |2         |0         |0.98      |
|USA Today                     |usatoday.com                  |US        |2         |51        |121       |1         |0         |0.99      |
|Vanity Fair                   |vanityfair.com                |US        |2         |105       |100       |0         |6         |0.97      |
|Wall Street Journal           |wsj.com                       |US        |2         |70        |368       |0         |2         |0.99      |
|Washington Examiner           |washingtonexaminer.com        |US        |2         |61        |119       |0         |5         |0.96      |
|Washington Post               |washingtonpost.com            |US        |2         |91        |502       |1         |0         |0.99      |
|Washington Times              |washingtontimes.com           |US        |2         |98        |359       |0         |0         |1.00      |
|BBC                           |bbc.co.uk                     |UK        |2         |41        |207       |0         |2         |0.98      |
|Daily Mail                    |dailymail.co.uk               |UK        |3         |105       |918       |4         |21        |0.89      |
|Telegraph                     |telegraph.co.uk               |UK        |2         |88        |150       |3         |2         |0.97      |
|The Guardian                  |theguardian.com               |UK        |2         |39        |303       |0         |0         |1.00      |
|The Independent               |independent.co.uk             |UK        |2         |35        |704       |0         |0         |1.00      |
|The Register                  |theregister.co.uk             |UK        |2         |51        |160       |0         |0         |1.00      |
|CBC                           |cbc.ca                        |CA        |3         |57        |207       |0         |3         |0.97      |
|Calgary Herald                |calgaryherald.com             |CA        |2         |42        |553       |0         |0         |1.00      |
|Edmonton Journal              |edmontonjournal.com           |CA        |2         |39        |536       |3         |0         |0.96      |
|National Post                 |nationalpost.com              |CA        |2         |35        |237       |1         |1         |0.97      |
|The Chronicle Herald          |thechronicleherald.ca         |CA        |2         |34        |241       |0         |0         |1.00      |
|The Globe and Mail            |theglobeandmail.com           |CA        |2         |37        |249       |2         |2         |0.95      |
|The Province                  |theprovince.com               |CA        |2         |42        |745       |0         |1         |0.99      |
|Toronto Star                  |thestar.com                   |CA        |2         |36        |54        |0         |7         |0.91      |
|Vancouver Sun                 |vancouversun.com              |CA        |2         |52        |477       |0         |2         |0.98      |
|Winnipeg Free Press           |winnipegfreepress.com         |CA        |2         |42        |427       |0         |1         |0.99      |
|Bild                          |bild.de                       |DE        |3         |59        |200       |0         |6         |0.95      |
|CICERO                        |cicero.de                     |DE        |2         |44        |180       |1         |2         |0.97      |
|Der Freitag                   |freitag.de                    |DE        |2         |26        |109       |0         |1         |0.98      |
|Der Tagesspiegel              |tagesspiegel.de               |DE        |2         |33        |215       |2         |0         |0.97      |
|Die Zeit                      |zeit.de                       |DE        |2         |45        |309       |2         |1         |0.97      |
|Focus                         |focus.de                      |DE        |3         |41        |331       |2         |5         |0.92      |
|Frankfurter Allgemeine        |faz.net                       |DE        |2         |16        |476       |0         |0         |1.00      |
|Frankfurter Rundschau         |fr-online.de                  |DE        |2         |29        |499       |0         |0         |1.00      |
|Handelsblatt                  |handelsblatt.com              |DE        |2         |35        |248       |1         |3         |0.95      |
|Heise                         |heise.de                      |DE        |2         |67        |396       |5         |5         |0.93      |
|NEON                          |neon.de                       |DE        |2         |20        |103       |0         |0         |1.00      |
|Spiegel                       |spiegel.de                    |DE        |3         |71        |688       |0         |5         |0.97      |
|Stern                         |stern.de                      |DE        |2         |27        |290       |0         |1         |0.98      |
|Süddeutsche Zeitung Magazin   |sz-magazin.de                 |DE        |2         |42        |139       |1         |4         |0.94      |
|Tagesschau                    |tagesschau.de                 |DE        |2         |30        |136       |0         |0         |1.00      |
|The European                  |theeuropean.de                |DE        |2         |125       |225       |0         |6         |0.98      |
|Der Standard                  |derstandard.at                |AT        |2         |29        |411       |0         |1         |0.98      |
|Neue Zürcher Zeitung          |nzz.ch                        |CH        |2         |30        |128       |1         |3         |0.94      |
|                              |                              |          |*185*     |*4636*    |*26032*   |*74*      |*115*     |**0.98**  |

