# gate-nss-cache
Gate NSS Cache is a "minimalistic implementation on google-nss-cache" implemented in GoLang, couple it with lib-nsscache, read below for more context and goto original [google-nss-cache](https://github.com/google/nsscache) for more understanding.

Change variables in test.yml
Set GATE_CONFIG_FILE to location of test.yml by default it should be /etc/nss/nss_http.yml, so that gate-nss-cache can know from where it can read the location of rest based gate server. 

gate-nss-cache is amazingly fast and has very low overhead compared to earlier [nss_gate](https://github.com/gate-sso/nss_gate), in fact it replaces nss_gate entirely. 


nsscache - Asynchronously synchronise local NSS databases with remote restful services
========================================================================================


*nsscache* is a commandline tool and Python library that synchronises a local NSS cache from a remote http enabled rest service, such as GATE-SSO.

As soon as you have more than one machine in your network, you want to share usernames between those systems. Linux administrators have been brought up on the convention of LDAP or NIS as a directory service, and `/etc/nsswitch.conf`, `nss_ldap.so`, and `nscd` to manage their nameservice lookups.

Even small networks experience intermittent name lookup failures, such as a mail receiver sometimes returning "User not found" on a mailbox destination because of a slow socket connection over a congested network, or erratic cache behaviour by `nscd`. To combat this problem, we have separated the network from the NSS lookup codepath, by using an asynchronous cron job and a glorified script, to improve the speed and reliability of NSS lookups.


    
Read the [Google Code blog announcement](http://www.anchor.com.au/blog/2009/02/nsscache-and-ldap-reliability/) for nsscache, or more about the [motivation behind nsscache tool](https://github.com/google/nsscache/wiki/MotivationBehindNssCache).

Here's a [testimonial from Anchor Systems](http://www.anchor.com.au/blog/2009/02/nsscache-and-ldap-reliability/) on their deployment of nsscache.


Pair *gate-nss-cache* with https://github.com/gate-sso/libnss-cache to integrate the local cache with your name service switch.

Make your nssswitch.conf look like below:

```
# /etc/nsswitch.conf
#
# Example configuration of GNU Name Service Switch functionality.
# If you have the `glibc-doc-reference' and `info' packages installed, try:
# `info libc "Name Service Switch"' for information about this file.

passwd:         compat cache
group:          compat cache
shadow:         compat cache
gshadow:        files

hosts:          files dns
networks:       files

protocols:      db files
services:       db files
ethers:         db files
rpc:            db files

netgroup:       nis
```
