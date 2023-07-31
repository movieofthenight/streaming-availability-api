## Streaming Availability API

Streaming Availability API allows getting streaming availability information of movies and series; and querying the list of available shows on streaming services such as Netflix, Disney+, Apple TV, Max and Hulu across 58 countries!

### API Key

To get an instant free subscription to start using the API, you can visit
[the RapidAPI page of the API](https://rapidapi.com/movie-of-the-night-movie-of-the-night-default/api/streaming-availability/pricing).

With a free subscription, you can send 100 requests per day.
To send more requests, you can upgrade to paid plans whenever you like.

### Useful Links

- [Official Webpage of the API](https://www.movieofthenight.com/about/api)

- [Subscription Page on RapidAPI](https://rapidapi.com/movie-of-the-night-movie-of-the-night-default/api/streaming-availability/pricing)

- [Full Documentation of All Available Endpoints](https://www.movieofthenight.com/about/api/documentation)

- [Home Page of the API on RapidAPI](https://rapidapi.com/movie-of-the-night-movie-of-the-night-default/api/streaming-availability/)

- [Main GitHub Repository of the API](https://github.com/movieofthenight/streaming-availability-api)

### Main Features

- Query streaming availability info of the movies and series via their TMDb or IMDd ids.
- Search for movies and series via their titles, genres, keywords, release years on
specific streaming services (e.g.: Get all the zombie action movies available
on Netflix and Disney+)
- Returned streaming availability info includes:
  - Deep links into the streaming services for
movies, series, seasons and episodes.
  - Available video qualities (eg. SD, HD, UHD).
  - Available subtitles and audios.
  - First detection time of the shows on the streaming services.
  - Expiry date of the shows/seasons/episodes on the streaming services.
  - All the available options to stream a show
(e.g. via subscription, to buy/rent, for free, available via an addons)
  - Price and currency information for buyable/rentable shows
- Channel and addon support (e.g. Apple TV Channels, Hulu Addons, Prime Video Channels)
- Output also includes TMDB and IMDb ids for every show.

### Terms & Conditions and Attribution

While the client libraries have MIT licenses,
the Streaming Availability API itself has further
[Terms & Conditions](https://github.com/movieofthenight/streaming-availability-api/blob/main/TERMS.md).
Make sure to read it before using the API.

Notably, the API requires an attribution to itself, if the data acquired through
is made public. You can read further about it on the
[Terms & Conditions](https://github.com/movieofthenight/streaming-availability-api/blob/main/TERMS.md)
page.

## FAQ

- **I run into an issue. How can I get help?**
  - If the issue is related to the API itself, please create a post
[here](https://rapidapi.com/movie-of-the-night-movie-of-the-night-default/api/streaming-availability/discussions),
and we will help with the issue.
  - If the issue is specific to a client library, then you can create a new issue
on the respective repository of the library.

- **API returned me some wrong data. What can I do?**
  - Send us a message with details of your findings.
You can click on the "Contact" button on this
[page](https://rapidapi.com/movie-of-the-night-movie-of-the-night-default/api/streaming-availability/details).
Once we receive the message we will take a look into the problems and fix the data.

- **I have a request to get a new streaming service supported by the API.**
  - Send us a message by clicking on the "Contact" button on this
  [page](https://rapidapi.com/movie-of-the-night-movie-of-the-night-default/api/streaming-availability/details),
  and we will get back to you.

- **I need a client library in another language.**
  - Send us a message by clicking on the "Contact" button on this
  [page](https://rapidapi.com/movie-of-the-night-movie-of-the-night-default/api/streaming-availability/details),
  and we will get back to you.

- **I need further details (e.g. posters, summaries, cast) about the shows. What can I do?**
  - Streaming Availability API works very well with
  [TMDB API](https://developer.themoviedb.org/docs), which can provide all the other
  info (that is not related to streaming availability) you need.
  You can pass the returned `tmdbId` field to the TMDB API's
  [movie-details](https://developer.themoviedb.org/reference/movie-details)
  or
  [tv-series-details](https://developer.themoviedb.org/reference/tv-series-details)
  endpoints
  and get the other details of the shows.

- **What is the difference between Streaming Availability API and TMDB API?**
  - The Movie Database (TMDB) is a community built movie and TV database, similar to IMDb.
  Their free API provides access to their database which includes posters, summaries, cast,
  and many other details about the shows. TMDB API's own streaming availability information
  data does not include deep links or any of the other features of this API such as
  available subtitles, audios, video qualities etc. Thus Streaming Availability API
  and TMDB API work hand in hand to get you all the details of the shows.

- **What is RapidAPI?**
  - RapidAPI is the world's largest API marketplace. We use RapidAPI to handle the
API subscriptions for us. You can instantly subscribe to Streaming Availability on
RapidAPI and start using the Streaming Availability API through RapidAPI right away.

## Client Libraries

1. [Go](https://github.com/movieofthenight/go-streaming-availability)
2. [TypeScript/JavaScript](https://github.com/movieofthenight/ts-streaming-availability)


## Services Supported

| Service Id | Service Name | Supported Countries |
| ---------- | ------------ | ------------------- |
| `netflix` | Netflix | 57 Countries |
| `prime` | Amazon Prime Video | 56 Countries |
| `disney` | Disney+ | 38 Countries |
| `hbo` | HBO Max | 24 Countries |
| `hulu` | Hulu | United States |
| `peacock` | Peacock | United States |
| `paramount` | Paramount+ | 18 Countries |
| `starz` | Starz | United States |
| `showtime` | Showtime | United States |
| `apple` | Apple TV+ | 52 Countries |
| `mubi` | Mubi | 55 Countries |
| `stan` | Stan | Australia |
| `now` | Now | United Kingdom, Ireland, Italy |
| `crave` | Crave | Canada |
| `all4` | All 4 | United Kingdom, Ireland |
| `iplayer` | BBC iPlayer | United Kingdom |
| `britbox` | BritBox | United Kingdom, United States, Canada, Australia, South Africa |
| `hotstar` | Hotstar | India, Canada, United Kingdom, Indonesia, Singapore |
| `zee5` | Zee5 | 58 Countries |
| `curiosity` | Curiosity Stream | 57 Countries |
| `wow` | Wow | Germany |


## Countries Supported

| Country Code | Country Name |
| ------------ | ------------ |
| `ae` | United Emirates |
| `ar` | Argentina |
| `at` | Austria |
| `au` | Australia |
| `az` | Azerbaijan |
| `be` | Belgium |
| `bg` | Bulgaria |
| `br` | Brazil |
| `ca` | Canada |
| `ch` | Switzerland |
| `cl` | Chile |
| `co` | Colombia |
| `cy` | Cyprus |
| `cz` | Czech Republic |
| `de` | Germany |
| `dk` | Denmark |
| `ec` | Ecuador |
| `ee` | Estonia |
| `es` | Spain |
| `fi` | Finland |
| `fr` | France |
| `gb` | United Kingdom |
| `gr` | Greece |
| `hk` | Hong Kong |
| `hr` | Croatia |
| `hu` | Hungary |
| `id` | Indonesia |
| `ie` | Ireland |
| `il` | Israel |
| `in` | India |
| `is` | Iceland |
| `it` | Italy |
| `jp` | Japan |
| `kr` | South Korea |
| `lt` | Lithuania |
| `md` | Moldova |
| `mk` | North Macedonia |
| `mx` | Mexico |
| `my` | Malaysia |
| `nl` | Netherlands |
| `no` | Norway |
| `nz` | New Zealand |
| `pa` | Panama |
| `pe` | Peru |
| `ph` | Philippines |
| `pl` | Poland |
| `pt` | Portugal |
| `ro` | Romania |
| `rs` | Serbia |
| `ru` | Russia |
| `se` | Sweden |
| `sg` | Singapore |
| `th` | Thailand |
| `tr` | Turkey |
| `ua` | Ukraine |
| `us` | United States |
| `vn` | Vietnam |
| `za` | South Africa |


