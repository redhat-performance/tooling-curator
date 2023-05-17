# Tool Scrapper

## What it does

This scraper tools goes through github Organizations and gets the repositories and their topics and outputs them in a formathe that the [UI](../README.md) can display.

> NOTE: This whole process is manual and not interative wiht the UI, to be able to see the changes you want you need to run this scrapper to update wht the UI shows.

### Step by Step

0. The scraper requires a GitHub [read-only oauth token](https://github.com/settings/tokens) as environment variable `GITHUB_AUTH_TOKEN`
1. The scraper will read the [organizations.json](../public/organizations.json) file in the public folder
2. It will also configure what to ignore bassed on [ignored-repositories.json](../public/ignored-repositories.json) and [ignored-topics.json](../public/ignored-topics.json) files in public folder
3. It will loop through all `organizations` and all `repositories` to generate the structure the UI needs to display the data.

JSON example:
```json
{
  "repos": [
    {
      "org": "redhat-performance",
      "name": "aimlperf_reg_tests",
      "description": "",
      "url": "https://github.com/redhat-performance/aimlperf_reg_tests",
      "labels": []
    }
  ]
}
```

## Building and Running

First

`go mod download`

To build:

`go build main.go`

To run:

`go run main.go`

## Configuration

To configure the tool add the dessired Organizations and Filters you want to apply.

### Organizations

The [organizations.json](../public/organizations.json) file is a simple JSON string array. Add or remove organizations as you need.

Example

```json
["cloud-bulldozer", "redhat-performance"]
```

### Ignore Repositories

The [ignored-repositories.json](../public/ignored-repositories.json) file has a format of a json struct with three fields:

- global: An array of strings. All names are ignored in all organizations.
- ignoreArchived: Boolean if set to true, all archived repos are ignored.
- orgs: A map of the organizations each with a Struct containing 3 fields:

  - skip-global-ignore: Will ignore the global ignored repo list
  - skip-global-archived: Will ignore the global `ignoreArchived` value
  - repos: An array of strings. All names are ignored

Example

```json
{
  "global": [
    ".github"
  ],
  "orgs": {
    "cloud-bulldozer": {
      "skip-global-ignore": true,
      "skip-global-archived": true,
      "repos": []
    },
    "redhat-performance": {
      "repos": [
        "some-repo"
      ]
    }
  },
  "ignoreArchived": true
}
```

> Hierarchy
>
> The more granular the rule biggest the priority.
>
> - Repos at the level of orgs["org"].repos always get ignored
> - `skip-global-ignore` and `skip-global-archived` take priority over the global configs
> - `global` repos and `ignoreArchived` are the lower priority
>

### Ignore Topics

The [ignored-topics.json](../public/ignored-topics.json) file is the same format as the `Organizations` file, a JSON string array, add topics that you want to ignore from your results.

Example

```json
["topic-one", "topic-two"]
```

## Applying changes

To persist your changes for the moment the config files are being commited.
If you add new ignore rules or new organizations make sure to run the `scraper` to get the latest repository list in this repo, and then commit all files.
