# Real Estate Scraping

The project for gathering actual data from https://www.hongkonghomes.com/en

`psql -h localhost -p 5432 -d real_estate -U admin --password`

# Structure
## Url generator
Service which create task to get urls of each item from special search url.
Create special task for scraper to get item urls, paginate search.

## Scraper
It reads urls from Redis

## Redis
Proxy, session, url list storing
