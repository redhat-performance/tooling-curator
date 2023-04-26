# Panda GitHub Tooling Curator

## Updating the gh-pages

Merges to main will trigger a GitHub action to update the gh-pages branch. Another scheduled action runs daily to update the 
repo scrape and the gh-pages website. If you update the scraper or its configuration files, you can follow the Backend 
instructions to run the scraper locally for testing.

## Frontend

To run the frontend locally:

```
npm install
npm run start
```

## Backend

To run backend see [here](scraper/README.md)
