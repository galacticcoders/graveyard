# Contributing to The Graveyard

The Graveyard is a cemetery for dead businesses, events, and organizations. If you know of one that belongs here and isn't already buried, open a pull request adding it to `graves.json`.

## What belongs here

- **Dead.** The entity no longer operates under its original form. Acquisitions count only if the brand was fully retired (Credit Suisse → UBS: yes; Instagram → Meta: no). Reincarnations count and can be listed as separate entries with separate date ranges (e.g., see XFL 1.0 / XFL 2.0).
- **Notable.** A general audience should recognize the name, or the story should be interesting enough that they'll want to learn about it. Your uncle's failed dry cleaner doesn't make the cut.
- **Verifiable.** Dates and cause of death must be defensible from a reliable source (news article, Wikipedia, SEC filing, obituary).

Not accepted: people, bands, TV shows, currently-operating zombies, meme deaths ("RIP my motivation").

## How to contribute

1. Fork the repo.
2. Add a new entry to `graves.json`. Keep entries grouped by category. Find the last entry of the matching category and add yours immediately after it.
3. Open a pull request. Fill out the template.
4. A maintainer reviews, merges, and GitHub Pages redeploys within a minute.

Local preview is optional but easy:

```bash
python3 -m http.server 8000
# open http://localhost:8000
```

## Data model

```json
{
  "name": "Quibi",
  "category": "Tech",
  "born": 2020,
  "died": 2020,
  "cause": "Nobody wanted it",
  "epitaph": "$1.75 billion raised. Six months alive. A premium mobile-only streaming service during a pandemic when everyone was home."
}
```

- **`name`** : how the entity was commonly known. Use the most recognizable form ("FTX", not "FTX Trading Ltd.").
- **`category`** : exactly one of: `Gaming`, `Tech`, `Retail`, `Finance`, `Sports`. New categories require a separate discussion; don't invent one in a PR.
- **`born`** : four-digit year of founding / first public form.
- **`died`** : four-digit year the entity ceased operating. Must be `>= born`.
- **`cause`** : 2-5 word tag. The headline reason it died. Examples: `Fraud`, `Disrupted by Netflix`, `COVID + bankruptcy`, `Twitter bank run`.
- **`epitaph`** : 1-3 sentences. See voice below.

## Voice

Epitaphs are **dry, factual, slightly wry**. The facts do the work. Calibrate against existing entries - Quibi, Lehman Brothers, and Theranos are the targets.

- Do: state what happened, let the absurdity land on its own.
- Don't: crack jokes, moralize, or punch down.
- Don't: include cruelty about real human loss (bankruptcies affect real people).
- Don't: write like Wikipedia. Short sentences. Specific numbers when they're telling.

Two good examples:

> "$9 billion valuation built on a machine that didn't work. Elizabeth Holmes became a cautionary tale." — Theranos

> "158 years old. The bankruptcy that broke the global economy. Nobody was allowed to bail them out." — Lehman Brothers

## Before you open the PR

- [ ] `graves.json` still parses (`python3 -m json.tool graves.json > /dev/null`).
- [ ] No duplicate `name` — search the file first.
- [ ] `died >= born`.
- [ ] Epitaph is 1–3 sentences.
- [ ] Category is one of the five allowed values.
- [ ] You can cite a source for the `died` year.

That's it. Welcome to the cemetery.
