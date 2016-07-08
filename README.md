# rancher-server-secure-properties

[![Build Status][travis-image]][travis-url]

Create `cattle-local.properties` from encrypted envrionment variables

This project is **ALPHA** status.  It could change or be removed at
anytime.

## Summary

This project will look for any environment variables that start with
`ENC_CATTLE_` base64 decode the value, decrypt and put the result in
`/usr/share/cattle/war/cattle-local.properties`.

This ensures that configuration like `CATTLE_DB_CATTLE_MYSQL_PASSWORD`
cannot be gathered by any docker inspection and out of any general
docker logging that may be doing.

At this time the only AWS KMS is supported.  All AWS configuration must
be passed in through the envrionment or ec2 instance profile.

## Quick Start

*Hack Start*

Run `rancher-server` with the following command:

    sh -c 'curl -sOL
https://github.com/SungardAS/rancher-server-secure-properties/releases/download/0.1.0-alpha.4/rancher-server-secure-properties;
chmod 755 rancher-server-secure-properties;
./rancher-server-secure-properties; /service/cattle/run'

Now instead of declaring `CATTLE_DB_CATTLE_MYSQL_PASSWORD` use
`ENC_CATTLE_DB_CATTLE_MYSQL_PASSWORD` and set the value as the base64
encoded encrypted string of the password.

##TODO

**A LOT**

Once this is matures we would like to make a generic script that could
be run for any project that wants to encrypt environment variables for
docker.  Possibly running before excuting another program like `confd`.


## Sungard Availability Services | Labs
[![Sungard Availability Services | Labs][labs-logo]][labs-github-url]

This project is maintained by the Labs team at [Sungard Availability
Services](http://sungardas.com)

GitHub: [https://sungardas.github.io](https://sungardas.github.io)

Blog:
[http://blog.sungardas.com/CTOLabs/](http://blog.sungardas.com/CTOLabs/)


[labs-github-url]: https://sungardas.github.io
[labs-logo]: https://raw.githubusercontent.com/SungardAS/repo-assets/master/images/logos/sungardas-labs-logo-small.png
[travis-image]: https://travis-ci.org/SungardAS/rancher-server-secure-properties.svg?branch=master
[travis-url]: https://travis-ci.org/SungardAS/rancher-server-secure-properties
