# RECORDONFIG Version 3

DELETE THIS BEFORE MERGING TO MAIN

## Release Notes

DNSControl v5 will make many major changes to the internals. There should be zero user-facing
changes. Because of the refactoring, I'm releasing a few "test releases" to get feedback.

* MAJOR REFACTOR: RecordConfig now uses (a reference to) codeberg.org/miekg/dns "RDATA" struct instead of storing individual fields.
  - This enhances RFC compliance, automatically supports all RFC record types.
  - Eventually this will save memory in a future release when the old fields are removed.
  - Currently the RDATA fields are stored in the old and new location, with automatic, bi-directional syncing.
* Standardized factory for creating DomainConfig and RecordConfig struct. The old "manual" way still works, but is being deprecated.
* Replaced github.com/miekg/dns (v1) with codeberg.org/miekg/dns (v2) in various places.

## Extra testing needed!

* BIND: SOA handling has been rewritten to be easier to debug and more reliable. Shouldn't have any user-visible changes but please be on the lookout for problems.
* CLOUDFLAREAPI: CF_WORKER_ROUTES() need extra testing. Internally they were represented sometimes as `WORKER_ROUTE` and sometimes as `CF_WORKER_ROUTE`. It's amazing such complex code ever worked.  Now we use `CF_WORKER_ROUTE` exclusively. The changes were core to the worker feature. Pleaes give extra attending and testing.
* IMPORT_TRANSFORM() hasn't changed but related parts have.
* Any provider that implements pseudo-types such as CF_WORKER_ROUTES(), BUNNY_DNS_PZ(), etc.

If you maintain a DNS provider, please give this release extra testing!

I only have automated testing for the following providers. All others are at risk of being broken:
AXFRDDNS_DNSSEC AZURE_DNS AZURE_PRIVATE_DNS BIND CLOUDFLAREAPI CNR
DIGITALOCEAN GANDI_V5 GCLOUD HEDNS MYTHICBEASTS NETNOD NS1 ROUTE53
SAKURACLOUD TRANSIP VERCEL

## Developer notes

* See documentation/developer-info/cookbook.md for developer notes
* It is now significantly easier to write providers. The translation from native to models.RecordConfig is much easier thanks to new factory functions.
* Supporting new DNS record types, including custom record types, is significantly easier.

## Future

* "PTR magic" and "REV()" continue to work, basically unchanged. There is now an opportunity to reimplement them in a cleaner way.

## Workarounds required for codeberg.org/miekg/dns

The migration to the new DNS package was very smooth. However there were some things I couldn't figure out how to do:

* To implement pkg/txtutil/miekg.go ZoneifyQuoted(), we needed access to https://codeberg.org/miekg/dns/src/branch/main/internal/ddd/ddd.go. However, that is "internal" and therefore could not be imported.  I copied the files needed to `pkg/txtutil/{miekg.go miekg_test.go ddd/ddd.go}`. It would be nice to not have to copy those.
* dnsutil.Trim() returns "" when z is longer than s. That's rather unintuitive considering that strings.TrimPrefix() returns the original string in that situation. The "" result is ambiguous because it can also mean "s == z".  I wrote pkg/txtutil.StripZone() which adopts a more useful behavior (though my implementation is not as performant).
* There's no way to parse the SVCB params ("port=80 ech=1234") into a svcb.Pairs.  You can run NewData("1 port=80 ech=1234") and extract the .Value.  Not a big deal, but it would be nice to be able to construct the Pairs directly.
