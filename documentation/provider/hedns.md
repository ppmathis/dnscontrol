## Important Note
Hurricane Electric does not currently expose an official JSON or XML API, and as such, this provider interacts directly
with the web interface. Because there is no officially supported API, this provider may cease to function if Hurricane
Electric changes their interface, and you should be willing to accept this possibility before relying on this provider.

## Configuration

To use this provider, add an entry to `creds.json` with `TYPE` set to `HEDNS`
along with
your `dns.he.net` account username and password. These are the same username
and password used to log in to the [web interface](https://dns.he.net).

{% code title="creds.json" %}
```json
{
  "hedns": {
    "TYPE": "HEDNS",
    "username": "yourUsername",
    "password": "yourPassword"
  }
}
```
{% endcode %}

### Two factor authentication

If two-factor authentication has been enabled on your account you will also need to provide a valid TOTP code.
This can also be done via an environment variable:

{% code title="creds.json" %}
```json
{
  "hedns": {
    "TYPE": "HEDNS",
    "username": "yourUsername",
    "password": "yourPassword",
    "totp": "$HEDNS_TOTP"
  }
}
```
{% endcode %}

and then you can run

```shell
HEDNS_TOTP=12345 dnscontrol preview
```

It is also possible to directly provide the shared TOTP secret using the key "totp-key" in `creds.json`. This secret is
only available when first enabling two-factor authentication.

**Security Warning**:
* Anyone with access to this `creds.json` file will have *full* access to your Hurricane Electric account and will be
  able to modify and delete your DNS entries
* Storing the shared secret together with the password weakens two factor authentication because both factors are stored
  in a single place.

{% code title="creds.json" %}
```json
{
  "hedns": {
    "TYPE": "HEDNS",
    "username": "yourUsername",
    "password": "yourPassword",
    "totp-key": "yourTOTPSharedSecret"
  }
}
```
{% endcode %}

### Persistent Sessions

Normally this provider will refresh authentication with each run of dnscontrol. This can lead to issues when using
two-factor authentication if two runs occur within the time period of a single TOTP token (30 seconds), as reusing the
same token is explicitly disallowed by RFC 6238 (TOTP).

To work around this limitation, if multiple requests need to be made, the option `"session-file-path"` can be set in
`creds.json`, which is the directory where a `.hedns-session` file will be created. This can be used to allow reuse of an
existing session between runs, without the need to re-authenticate.

This option is disabled by default when this key is not present,

**Security Warning**:
* Anyone with access to this `.hedns-session` file will be able to use the existing session (until it expires) and have
  *full* access to your Hurrican Electric account and will be able to modify and delete your DNS entries.
* It should be stored in a location only trusted users can access.

{% code title="creds.json" %}
```json
{
  "hedns": {
    "TYPE": "HEDNS",
    "username": "yourUsername",
    "password": "yourPassword",
    "totp-key": "yourTOTPSharedSecret",
    "session-file-path": "."
  }
}
```
{% endcode %}

## Metadata
This provider does not recognize any special metadata fields unique to Hurricane Electric DNS.

## Usage
An example configuration:

{% code title="dnsconfig.js" %}
```javascript
var REG_NONE = NewRegistrar("none");
var DSP_HEDNS = NewDnsProvider("hedns");

D("example.com", REG_NONE, DnsProvider(DSP_HEDNS),
    A("test", "1.2.3.4"),
);
```
{% endcode %}
