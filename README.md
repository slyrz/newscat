![Logo](https://raw.github.com/slyrz/newscat/master/img/newscat_logo.png)

newscat is a news content / article extractor written in just 1000
lines of Go code.

### Overview

TODO...

### Training and Evaluation

A data set of 300 news articles was created by crawling top submissions from
various news related Reddit communities. The HTML pages of the news articles
were downloaded and labeled by adding a special data attribute to the
HTML elements containing relevant content.

The data set was split into training and test data.
50 randomly chosen news articles were used to train newscat. The remaining
250 news articles were used to evaluate the quality of newscat's content
extraction.
For each news article, we compared the element-level predictions with the real
labels and calculated precision, recall and the balanced F-score.

![Extraction Quality](https://raw.github.com/slyrz/newscat/master/img/newscat_graph.png)

The above figure shows the resulting F-scores ordered in increasing magnitude.
Thus, the X axis shows the percentage of articles whos F-scores fall below the
value indicated by the Y axis. In other words: the percentiles.

TODO...

### License

newscat is released under MIT license.
You can find a copy of the MIT License in the [LICENSE](./LICENSE) file.

