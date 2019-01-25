# WPCLI

This is currently a proof of concept for re-creating [WPCLI](https://github.com/wp-cli/wp-cli) in Go.

## Commands

Initial plans are to include `wp core download`, `wp core install` and `wp core is-installed`.

### Core

Contains WP Core Commands

#### Download

Downloads and extracts WordPress core files to the specified path.

#### Install

Installs WordPress using the provided details.

#### Is-Installed

Checks for a valid WordPress install at the specified path.

## License

This is MIT licensed.