# Panda GitHub Tooling Curator

## Updating the gh-pages

Merges to main will trigger a GitHub action to update the gh-pages branch. If you update the scraper or its configuration files,
follow the Backend instructions to run the scraper, which will update the assets for the gh-pages, and include those changes
in your commit to ensure the gh-pages reflect the changes.

## Frontend

To run the frontend locally:

```
npm install
npm run start
```

## Backend

To run backend see [here](scraper/README.md)
